# sync Package Internals — Mutex, RWMutex, WaitGroup, Once, Cond, Atomic

پکیج `sync` ابزارهای سطح‌پایین concurrency را فراهم می‌کند.

---

## 1. Mutex

```go
type Mutex struct {
    state int32
    sema  uint32
}
```

- `Lock` با CAS تلاش می‌کند سریع lock بگیرد؛ در صورت contention وارد مسیر پیچیده‌تر می‌شود.  
- `Unlock` lock را آزاد کرده و در صورت وجود waiter، یکی را بیدار می‌کند.  
- starvation mode برای جلوگیری از گرسنگی طولانی وجود دارد.

هرگز struct حاوی Mutex را کپی نکن.

---

## 2. RWMutex

- `RLock`/`RUnlock` برای readerها  
- `Lock`/`Unlock` برای writer

Readers همزمان می‌توانند وارد شوند، اما:

- اگر readers مداوم وارد شوند، writer ممکن است دیر اجرا شود.  
- برای workloadی که writeها کم‌اند مفید است.

---

## 3. WaitGroup

برای منتظر ماندن روی چند goroutine:

```go
wg.Add(1)
go func() {
    defer wg.Done()
    ...
}()
wg.Wait()
```

- `Add` شمارنده را تغییر می‌دهد.  
- `Done` معادل `Add(-1)` است.  
- اگر شمارنده زیر صفر برود → panic.  
- استفاده نادرست/مجدد از WaitGroup می‌تواند race ایجاد کند.

---

## 4. Once

برای اجرای یک initialization فقط یک بار:

```go
var once sync.Once
once.Do(func() { ... })
```

- از atomic load/store + Mutex استفاده می‌کند.  
- fast-path: اگر done شده باشد، هزینه ناچیز دارد.

---

## 5. Cond

Condition variable:

```go
c := sync.NewCond(&mu)
```

الگو:

```go
mu.Lock()
for !condition {
    c.Wait()
}
// condition برقرار است
mu.Unlock()
```

- Wait lock را آزاد می‌کند، goroutine را park می‌کند، سپس بعد از بیدار شدن دوباره lock می‌گیرد.  
- باید همیشه در یک loop استفاده شود.

---

## 6. sync/atomic

عملیات‌های atomic برای typeهای عددی و pointer:

- `Load`, `Store`, `Add`, `Swap`, `CompareAndSwap`

این‌ها memory barrier دارند و از reordering جلوگیری می‌کنند.

مثال:

```go
atomic.StoreInt32(&flag, 1) // release
if atomic.LoadInt32(&flag) == 1 { // acquire
    ...
}
```

---

## 7. Summary

- Mutex/RWMutex برای محافظت از shared state.  
- WaitGroup برای مدیریت lifecycle goroutineها.  
- Once برای init امن.  
- Cond برای سناریوهای هماهنگ‌سازی پیچیده.  
- atomic برای موارد کم‌حجم و performance-critical.

استفاده درست از این primitives برای ساخت سیستم concurrent سالم ضروری است.
