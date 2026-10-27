# Slice Internals — Header, Backing Array, Growth, Pitfalls

Slice در Go یک abstraction روی یک array است، نه خود array.

---

## 1. مدل داخلی Slice

مفهومی:

```go
type slice struct {
    ptr *T
    len int
    cap int
}
```

- `ptr` → به اولین عنصر backing array اشاره می‌کند.  
- `len` → تعداد elementهای قابل‌دسترسی.  
- `cap` → حداکثر تعداد عناصری که قبل از reallocation می‌توان اضافه کرد.

کپی‌کردن یک slice → فقط همین header کپی می‌شود.

---

## 2. Backing Array & Aliasing

```go
a := []int{1, 2, 3}
b := a
b[0] = 99
```

- هر دو به یک backing array اشاره می‌کنند.  
- تغییر روی `b` روی `a` هم دیده می‌شود.

---

## 3. append و reallocation

اگر:

```go
s := make([]int, 0, 4)
s = append(s, 1, 2, 3, 4) // cap هنوز 4 → روی همان array
s = append(s, 5)          // cap پر شده → احتمال reallocation
```

وقتی cap کافی نباشد:

- runtime یک array بزرگتر می‌سازد  
- data را کپی می‌کند  
- header به array جدید اشاره می‌کند

بنابراین aliasing ممکن است شکسته شود.

---

## 4. Growth Strategy

تقریباً:

- برای cap کوچک → دو برابر کردن  
- برای cap بزرگ → حدود 1.25x رشد

این رفتار:

- append را amortized O(1) می‌کند.  
- تعداد reallocations را محدود می‌کند.

---

## 5. Slicing = View

```go
x := []int{10, 20, 30, 40, 50}
y := x[1:4] // [20,30,40]
```

- `y` header جدید است  
- اما به همان backing array `x` اشاره می‌کند  
- تغییر روی عناصر مشترک روی هر دو اثر می‌گذارد.

---

## 6. کپی‌کردن data

برای deep copy:

```go
dst := make([]int, len(src))
copy(dst, src)
```

حالا backing array جداست.

---

## 7. nil vs empty

```go
var s1 []int      // nil
s2 := []int{}     // empty
s3 := make([]int,0)
```

- `len` هر سه 0  
- `s1 == nil` → true  
- `s2 == nil` و `s3 == nil` → false  
- append روی هر سه OK است.

---

## 8. Concurrency

- فقط read همزمان روی slice بدون write → OK.  
- هر write همزمان روی slice مشترک → data race.  
- append از چند goroutine روی یک slice بدون sync → کاملاً unsafe.

---

## 9. Preallocation

برای performance:

```go
s := make([]T, 0, N)
for ... {
    s = append(s, v)
}
```

این کار:

- تعداد reallocations را کاهش می‌دهد  
- فشار روی GC را کم می‌کند.

---

## 10. Summary

- slice header (ptr,len,cap) است.  
- copy(slice) فقط header را کپی می‌کند، نه data.  
- append می‌تواند روی همان array یا array جدید باشند.  
- slicing view جدید می‌سازد، نه کپی.  
- aliasing می‌تواند باعث باگ شود اگر حواست نباشد.
