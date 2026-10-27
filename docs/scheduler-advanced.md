# Advanced Scheduler Internals — sysmon, Parking, Run Queues, Preemption

این داک جزئیات پیشرفته‌تر scheduler Go را پوشش می‌دهد.

---

## 1. sysmon

Thread اختصاصی runtime:

- monitor goroutineها  
- trigger preemption  
- مدیریت timers  
- کمک به netpoller  
- گاهی trigger GC

---

## 2. Parking / Unparking

- هنگام بلاک شدن روی channel/mutex/cond → goroutine با `gopark` park می‌شود.  
- وقتی شرط آزاد شد → با `ready`/`goready` unpark می‌شود و به run queue بازمی‌گردد.

---

## 3. Run Queues

- per-P local runqueue  
- یک global runqueue  
- work-stealing برای توازن بین Pها

---

## 4. Syscall Handling

هنگام syscall blocking:

- G به حالت syscall می‌رود.  
- M در syscall می‌ماند.  
- P به M دیگر داده می‌شود.  

این مانع از توقف scheduler به‌خاطر syscall می‌شود.

---

## 5. Async Preemption

- Go می‌تواند goroutineهایی که مدت زیادی CPU می‌گیرند را قطع کند.  
- این کار در safe pointها انجام می‌شود.  
- برای GC و fairness ضروری است.

---

## 6. Summary

Scheduler Go:

- از M/P/G استفاده می‌کند.  
- sysmon برای کنترل و preemption دارد.  
- runqueueهای محلی + global دارد.  
- از work-stealing و async preemption برای scaling و fairness استفاده می‌کند.
