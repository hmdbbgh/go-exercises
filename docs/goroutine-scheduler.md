# Goroutine Scheduler — M/P/G Model, Work Stealing, Preemption

Scheduler Go مسئول زمان‌بندی goroutineها روی OS threadها است.

---

## 1. M/P/G Model

سه جزء اصلی:

- **G**: goroutine  
- **M**: OS thread (Machine)  
- **P**: Processor منطقی برای scheduler

مدل:

- P run queue، timers و caches را نگه می‌دارد.  
- M فقط با داشتن یک P می‌تواند goroutine اجرا کند.  
- G روی M با کمک P اجرا می‌شود.

`GOMAXPROCS` تعداد Pها را تعیین می‌کند.

---

## 2. ایجاد Goroutine

```go
go f()
```

- یک G جدید ایجاد می‌کند.  
- آن را در run queue یک P قرار می‌دهد.  
- اگر M مناسب همان P فعال باشد، آن goroutine را اجرا می‌کند.

---

## 3. Run Queues

- هر P یک run queue محلی دارد.  
- یک queue global هم وجود دارد.  
- scheduler اول از queue محلی استفاده می‌کند، سپس در صورت نیاز به global مراجعه می‌کند.

---

## 4. Work Stealing

اگر P run queue خالی داشته باشد:

- یک P دیگر را به‌طور تصادفی انتخاب می‌کند.  
- نیمی از goroutineهای runnable آن را می‌دزدد.  
- این کار load را متوازن می‌کند.

---

## 5. Blocking Syscalls

اگر goroutine یک syscall blocking انجام دهد:

- G → حالت syscall  
- M → در syscall گیر می‌کند  
- P از M جدا می‌شود  
- یک M دیگر P را می‌گیرد و goroutineهای دیگر را اجرا می‌کند.  

وقتی syscall تمام شود، G دوباره runnable می‌شود و در queue قرار می‌گیرد.

---

## 6. sysmon

Thread ویژه runtime:

- monitor کردن goroutineهای طولانی-running  
- trigger کردن preemption  
- مدیریت timers  
- کمک به شروع GC  
- کنترل وضعیت کلی scheduler

---

## 7. Preemption

Go می‌تواند goroutineهایی که مدت طولانی در حال اجرا هستند را preempt کند:

- برای جلوگیری از starvation  
- برای کمک به GC  
- برای حفظ responsiveness سیستم

preemption در نقاطی انجام می‌شود که از نظر GC safe هستند (safe points).

---

## 8. Deadlock Detection

اگر:

- هیچ goroutine runnable نباشد  
- و برنامه هنوز تمام نشده باشد

runtime panic می‌دهد:

```text
fatal error: all goroutines are asleep - deadlock!
```

---

## 9. Summary

- Scheduler Go با مدل M/P/G کار می‌کند.  
- از work-stealing برای balance استفاده می‌کند.  
- sysmon برای preemption و monitoring حضور دارد.  
- syscalls blocking با الگوریتم handoff مدیریت می‌شوند.  
- طراحی شده برای scalability بالا روی multi-core.
