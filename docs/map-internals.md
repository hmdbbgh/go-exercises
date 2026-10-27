# Map Internals — Buckets, Hashing, Overflow, Growth, Iteration

Map در Go یک hash-table پیچیده با bucketهای ثابت و overflow chainهاست.

---

## 1. Map Value → Pointer به `hmap`

خود `map[K]V` در Go یک value است که pointer به `hmap` را نگه می‌دارد:

```go
type hmap struct {
    count      int
    flags      uint8
    B          uint8          // log2(bucket count)
    hash0      uint32         // hash seed
    buckets    unsafe.Pointer // [2^B]bmap
    oldbuckets unsafe.Pointer // برای grow
    nevacuate  uintptr
}
```

---

## 2. Bucket Layout

هر bucket (`bmap`) چیزی شبیه این است:

```text
+----------+-------------------+-------------------+------------+
| tophash  |       keys[8]     |     values[8]     |  overflow  |
+----------+-------------------+-------------------+------------+
```

- `tophash[8]` → بخش بالایی hash برای هر slot  
- `keys[8]` و `values[8]` → تا ۸ key/value  
- `overflow` → pointer به bucket اضافی در صورت پر شدن

---

## 3. Hashing

برای key:

1. hash(key) حساب می‌شود.  
2. پایین‌ترین B بیت → index bucket.  
3. داخل bucket با استفاده از `tophash` و سپس مقایسه‌ی key، slot یافت می‌شود.  
4. در صورت نیاز، overflow buckets هم بررسی می‌شوند.

---

## 4. Load Factor & Growing

Map وقتی grow می‌شود که:

- تعداد عناصر زیاد شود (load factor بالا برود)، یا  
- طول chain overflow زیاد شود.

در این حالت:

- `B` یک واحد زیاد می‌شود (دو برابر شدن bucketها).  
- bucketهای جدید تخصیص داده می‌شوند.  
- `oldbuckets` نگه‌دارنده‌ی قدیمی‌هاست.  
- در طول دسترسی‌های بعدی، انتقال تدریجی (evacuation) انجام می‌شود.

این grow incremental است و از pauseهای بزرگ جلوگیری می‌کند.

---

## 5. Iteration Randomness

ترتیب iteration روی map:

```go
for k, v := range m {}
```

عمداً random است:

- به دلیل hash seed (`hash0`) تصادفی برای هر map.  
- semantics مشخصاً order را تضمین نمی‌کند.  

بنابراین:

> هرگز روی ترتیب map تکیه نکن.

---

## 6. nil Map

```go
var m map[string]int // nil
```

- `m == nil` → true  
- read → `m["x"]` → 0 (zero-value)  
- write → panic (`assignment to entry in nil map`)

`make(map[K]V)` یک map non-nil قابل نوشتن می‌سازد.

---

## 7. Copying Maps

```go
m2 := m1
```

- فقط pointer به `hmap` کپی می‌شود.  
- هر دو به همان ساختار داخلی اشاره می‌کنند.  
- تغییر در یکی روی دیگری اثر دارد.  
- deep copy باید دستی انجام شود:

```go
m2 := make(map[K]V, len(m1))
for k, v := range m1 {
    m2[k] = v
}
```

---

## 8. Concurrency

Map:

- برای concurrent write/read ایمن نیست.  
- دو goroutine که هم‌زمان روی یک map write کنند → behavior undefined و ممکن است panic شود (`concurrent map writes`).

راه‌حل:

- `sync.RWMutex` + map  
- یا `sync.Map` با understanding کامل trade-off.

---

## 9. Summary

- map value pointer به `hmap` را نگه می‌دارد.  
- data در bucketها و overflow bucketها ذخیره می‌شود.  
- iteration order تصادفی است.  
- copy(map) فقط reference است.  
- write concurrent روی map بدون sync → ممنوع.
