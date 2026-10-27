# Context Internals — Cancellation, Deadlines, Propagation, Values

`context` پکیج central برای کنترل lifetime و metadata در Go است.

---

## 1. Context Tree

هر context یک node در درخت است:

- `context.Background()` → ریشه  
- `WithCancel`, `WithTimeout`, `WithDeadline`, `WithValue` → nodeهای جدید

Cancel روی یک parent → همه‌ی childها را cancel می‌کند.

---

## 2. Cancellation

```go
ctx, cancel := context.WithCancel(parent)
defer cancel()
```

- `ctx.Done()` یک channel است که هنگام cancel شدن بسته می‌شود.  
- `ctx.Err()` دلیل cancel را نگه می‌دارد.

goroutineها باید به `ctx.Done()` گوش دهند و خودشان terminate شوند.

---

## 3. Timeout & Deadline

```go
ctx, cancel := context.WithTimeout(parent, 2*time.Second)
```

یا:

```go
ctx, cancel := context.WithDeadline(parent, t)
```

- وقتی زمان برسد → context به صورت خودکار cancel می‌شود.  
- `ctx.Err()` در این حالت `context.DeadlineExceeded` است.

---

## 4. Value

```go
ctx = context.WithValue(ctx, key, value)
v := ctx.Value(key)
```

فقط برای:

- metadata کوچک  
- request ID  
- auth token  
- tracing info

مناسب است. context storage برای data بزرگ یا dependency injection نیست.

---

## 5. Rules of Thumb

- `context.Context` باید اولین پارامتر توابع public باشد.  
- از `context.Background` فقط در entrypointها استفاده کن (main, tests).  
- در handlerها و لایه‌های میانی، context را forward کن.  
- حتماً `cancel()` را در استفاده از `WithCancel`/`WithTimeout` فراخوانی کن تا leak نشود.

---

## 6. Summary

- context یک درخت immutable از cancelable nodeهاست.  
- cancelation و deadline از parent به child propagate می‌شود.  
- goroutineها باید `Done()` را چک کنند و خودشان terminate شوند.  
- context برای مدیریت lifecycle در distributed systems حیاتی است.
