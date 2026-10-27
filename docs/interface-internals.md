# Interface Internals — type/value Pair, itab, Nil Traps

Interface در Go تأمین‌کننده‌ی polymorphism بر اساس رفتار (method set) است.

---

## 1. Interface Representation

دو مدل اصلی:

### Empty interface (`interface{}`)

```go
type eface struct {
    _type *_type
    data  unsafe.Pointer
}
```

### Non-empty interface

```go
type iface struct {
    tab  *itab
    data unsafe.Pointer
}

type itab struct {
    inter *interfacetype // interface type
    _type *_type         // concrete type
    fun   [1]uintptr     // method function pointers
}
```

بنابراین interface همیشه:

- نوع واقعی (dynamic type)  
- pointer به داده

را نگه می‌دارد.

---

## 2. Method Sets و سازگاری با Interface

Method set:

- برای `T`: متدهایی با receiver نوع `T`.  
- برای `*T`: متدهایی با receiver `T` و `*T`.

برای اینکه یک type یک interface را پیاده‌سازی کند، method set آن باید همه متدهای interface را داشته باشد.

---

## 3. Nil Interface vs Typed Nil

دام مهم:

```go
var r io.Reader
fmt.Println(r == nil) // true
```

اما:

```go
var f *os.File = nil
var r io.Reader = f
fmt.Println(r == nil) // false
```

علت:

- interface == nil فقط وقتی که هم type و هم data nil باشند.  
- در مثال دوم: type = `*os.File` (غیر nil)، data = nil.  
- پس interface value non-nil است.

---

## 4. Equality

برای دو interface:

- اگر dynamic type متفاوت باشد → `==` false.  
- اگر dynamic type یکسان باشد → valueها با قوانین همان type مقایسه می‌شوند.  
- اگر type غیر قابل مقایسه باشد (slice, map, func) → مقایسه panic می‌دهد.

---

## 5. Type Assertion & Type Switch

```go
v, ok := i.(T)
```

- اگر dynamic type == T → ok=true  
- در غیر این صورت → ok=false (و بدون ok، panic)

type switch:

```go
switch v := i.(type) {
case string:
case int:
case fmt.Stringer:
default:
}
```

با dynamic type dispatch می‌کند.

---

## 6. Boxing & Performance

قرار دادن value در interface معمولاً:

- آن را روی heap می‌برد (boxing/escape).  
- نیازمند نگهداری type info است.  

در مسیرهای performance-critical، استفادهٔ زیاد از interface می‌تواند overhead ایجاد کند.

---

## 7. Summary

- interface value همیشه (type, data) است.  
- empty interface فقط type/data دارد، non-empty از itab استفاده می‌کند.  
- typed nil ≠ interface nil.  
- equality وابسته به نوع و قوانین مقایسه type است.  
- boxing می‌تواند روی performance اثر بگذارد.
