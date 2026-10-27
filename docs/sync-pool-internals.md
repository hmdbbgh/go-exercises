# sync.Pool Internals — Object Pooling & GC Interaction

`sync.Pool` برای reuse objectهای موقتی طراحی شده است.

---

## 1. Use Case

برای:

- buffers (`[]byte`)  
- encoder/decoderها  
- structهای موقتی، per-request

نه برای:

- connection pooling  
- caching بلندمدت  
- objectهای stateful

---

## 2. مدل داخلی

به‌صورت ساده:

- per-P local pool (`poolLocal`)  
- یک اسلات private بسیار سریع  
- یک صف shared برای موارد بیشتر

Get:

- اگر private غیر nil → برمی‌گرداند.  
- وگرنه از shared/local/سایر Pها → و در نهایت New().

Put:

- اگر private خالی → می‌گذارد آنجا.  
- وگرنه در shared queue.

---

## 3. GC Behavior

GC می‌تواند محتوای pool را flush کند:

- objectهایی که فقط در pool referenced هستند جمع‌آوری می‌شوند.  
- به این معنی است که pool **cache بلندمدت** نیست.

---

## 4. Summary

- sync.Pool برای کاهش allocations در hot path عالی است.  
- GC محتوا را در طول زمان پاک می‌کند.  
- برای objectهای short-lived و resettable استفاده شود.
