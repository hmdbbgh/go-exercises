# Go Memory Model — Deep Dive (Senior-Level)

Go یک زبان **value-oriented** است.

این یعنی:

- بیشتر چیزها **value** هستند  
- assignment = **copy**  
- function calls = **copy**  
- struct = value  
- array = value  
- interface = value (اما حاوی reference)  
- slices/maps/channels/functions = reference-types اما **خود value** هستند

این مدل باعث می‌شود Go:

- قابل‌پیش‌بینی باشد  
- بدون undefined behavior سطح زبان باشد  
- برای performance بهینه باشد  

اما برای استفاده حرفه‌ای باید memory model را دقیق بشناسیم.

---

## 1. Stack vs Heap

Go دو محل اصلی برای تخصیص حافظه دارد:

### Stack

- بسیار سریع  
- با خروج از تابع آزاد می‌شود  
- بدون نیاز به GC  
- محدود به همان goroutine

### Heap

- shared بین goroutineها  
- آزادسازی توسط GC  
- برای داده‌هایی که طول عمرشان از scope فعلی فراتر می‌رود

---

## 2. Escape Analysis

کامپایلر تصمیم می‌گیرد که یک مقدار:

- روی stack بماند، یا  
- به heap **escape** کند.

مثال:

```go
func Foo() *int {
    x := 10
    return &x // x به heap می‌رود
}

func Bar() int {
    x := 10
    return x // روی stack می‌ماند
}
```

برای دیدن تصمیم‌های کامپایلر:

```bash
go build -gcflags="-m"
```

---

## 3. Value Types vs Reference-Like Values

اصطلاح «reference type» در Go رسمی نیست، اما مفید است.

- `struct`, `array`, `int`, `float`, ... → **valueهای ساده**  
- `slice`, `map`, `chan`, `func`, `interface`, `string` → valueهایی که **داخلشان pointer** نگه می‌دارند.

مثال slice:

```go
type sliceHeader struct {
    ptr *T
    len int
    cap int
}
```

یعنی:

```go
s2 := s1 // فقط header (ptr,len,cap) کپی می‌شود، نه خود data
```

---

## 4. Strings, Slices, Maps, Interfaces

به‌صورت ساده:

- `string` → (ptr, len)  
- `[]T` → (ptr, len, cap)  
- `map[K]V` → ptr به `hmap`  
- `interface{}` → (type ptr, data ptr)

کپی‌کردن این‌ها همیشه سبک است (sizeof چند word)، ولی دادهٔ واقعی جای دیگری است.

---

## 5. Pass by Value — همیشه

Go فقط **pass by value** دارد؛ اما آن value ممکن است خودش pointer باشد.

```go
func f(u User)       {} // کپی کامل struct
func g(u *User)      {} // کپی pointer، struct مشترک
func h(s []int)      {} // کپی header، backing array مشترک
func k(m map[int]int){} // کپی pointer به hmap، map مشترک
```

---

## 6. Zero-Value

هر type در Go یک zero-value معتبر دارد:

- عددی → 0  
- bool → false  
- string → ""  
- pointer/map/slice/chan → nil  
- struct → تمام فیلدها zero-value

این کار:

- کد را ساده‌تر می‌کند (نیاز کمتر به init دستی)  
- امن‌تر است (nil-safe zero values)

---

## 7. Practical Rules

- structهای کوچک را by-value pass کن.  
- structهای بزرگ یا با state مشترک را با pointer.  
- slice/map همین‌طور هم reference-like هستند، pointer نیاز ندارند.  
- escape analysis را برای codeهای حساس به performance چک کن.  
- interface و closureها غالباً باعث heap allocation می‌شوند.

---

## 8. Summary

- Go value-centric است.  
- تفاوت بین valueهای ساده و valueهای حاوی pointer را بشناس.  
- copy همیشه value را کپی می‌کند، نه الزاماً data را.  
- stack/heap را کامپایلر براساس escape analysis انتخاب می‌کند.  
- درک این مدل روی طراحی API، performance و correctness اثر مستقیم دارد.
