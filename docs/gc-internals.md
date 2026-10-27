# Go GC Internals — Concurrent Mark-Sweep, Tricolor, Pacer

GC Go:

- concurrent  
- mark-sweep  
- non-moving  
- tricolor

طراحی شده برای latency پایین و عملکرد مناسب در سرویس‌های بزرگ.

---

## 1. Tricolor Marking

objectها:

- White → candidate برای جمع‌آوری  
- Gray → reachable ولی scan نشده  
- Black → reachable و scan شده

روند:

1. تمام heap white.  
2. rootها خاکستری.  
3. تا وقتی gray وجود دارد → scan و رنگ کردن.  
4. در پایان: whiteها آزاد می‌شوند.

---

## 2. Phases

1. Mark setup (STW کوتاه)  
2. Concurrent mark (با write barrier)  
3. Mark termination (STW کوتاه)  
4. Sweep (incremental، پس‌زمینه)

---

## 3. GOGC و Heap Growth

`GOGC` درصد رشد مجاز heap بین دو GC است:

- 100 → heap می‌تواند ~2x شود.  
- کم‌تر → GC بیشتر، heap کوچک‌تر.  
- بیشتر → GC کمتر، heap بزرگ‌تر.

---

## 4. Pacer & Mutator Assist

- pacer سرعت GC را کنترل می‌کند.  
- اگر GC عقب باشد، allocationها موظف‌اند کمی کار GC انجام دهند (mutator assist).

---

## 5. Summary

- GC Go tricolor mark-sweep concurrent است.  
- pauseها معمولاً کوتاه و قابل‌قبول‌اند.  
- behavior قابل تنظیم با GOGC است.  
- طراحی شده برای سرویس‌های cloud-native.
