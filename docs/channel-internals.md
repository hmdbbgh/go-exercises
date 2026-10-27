# Channel Internals — hchan, Blocking Semantics, Close

Channel primitive اصلی Go برای communication + synchronization است.

---

## 1. ساختار `hchan`

Channel value در Go pointer به `hchan` است:

```go
type hchan struct {
    qcount   uint
    dataqsiz uint
    buf      unsafe.Pointer
    elemsize uint16
    closed   uint32
    elemtype *_type

    sendx    uint
    recvx    uint

    recvq    waitq
    sendq    waitq

    lock     mutex
}
```

---

## 2. Unbuffered vs Buffered

- Unbuffered: `dataqsiz = 0`, `buf = nil`  
  - send و receive باید هم‌زمان اتفاق بیفتند.  
- Buffered: `dataqsiz > 0`, `buf` به ring buffer اشاره می‌کند.  
  - send تا پرشدن buffer بدون بلاک‌شدن ممکن است.  
  - receive تا empty شدن buffer بدون بلاک ممکن است.

---

## 3. Send (`ch <- x`)

- اگر receiver منتظر باشد → value مستقیماً به او داده می‌شود.  
- در غیر این صورت، اگر buffer جا داشته باشد → x در buffer قرار می‌گیرد.  
- اگر نه → goroutine sender در `sendq` park می‌شود.

---

## 4. Receive (`x := <-ch`)

- اگر buffer خالی نباشد → یک value از آن برداشته می‌شود.  
- اگر sender منتظر باشد → value را مستقیم از او می‌گیرد.  
- اگر channel بسته باشد → zero-value برمی‌گردد، با ok=false.  
- در غیر این صورت → goroutine در `recvq` park می‌شود.

---

## 5. Close

```go
close(ch)
```

- flag `closed` ست می‌شود.  
- همه goroutineهای منتظر در `recvq` بیدار می‌شوند.  
- receive بعدی از channel بسته: zero-value + ok=false.  
- send روی channel بسته → panic.

---

## 6. nil Channel

```go
var ch chan int // nil
```

- send و receive روی nil channel → بلاک دائمی (deadlock اگر تنها goroutine باشد).  
- close روی nil channel → panic.

این ویژگی در select برای disable کردن caseها استفاده می‌شود.

---

## 7. Synchronization & Memory Ordering

Send/receive روی channel علاوه بر synchronization:

- memory barrier هم هستند.  
- تمام writes قبل از send، برای goroutine receiver بعد از receive قابل مشاهده‌اند.

---

## 8. Summary

- channel با `hchan` پیاده‌سازی شده است.  
- buffered/unbuffered semantics دارد.  
- close رفتار کاملاً مشخص و deterministic دارد.  
- send روی channel بسته panic می‌دهد.  
- nil channel همیشه بلاک می‌کند.
