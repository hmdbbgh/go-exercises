# Benchmarking, Profiling & Tracing in Go

ابزارهای داخلی Go برای performance خیلی قدرتمند هستند.

---

## 1. Benchmarks

ساختار:

```go
func BenchmarkX(b *testing.B) {
    for i := 0; i < b.N; i++ {
        X()
    }
}
```

اجرا:

```bash
go test -bench=. -benchmem
```

به `ns/op`, `B/op`, `allocs/op` توجه کن.

---

## 2. Escape Analysis

```bash
go build -gcflags="-m"
```

برای فهمیدن این‌که چه چیزهایی روی heap می‌روند، و چرا.

---

## 3. CPU Profiling

```bash
go test -bench=. -cpuprofile cpu.out
go tool pprof cpu.out
```

داخل pprof:

- `top`  
- `list Func`  
- `web`

---

## 4. Memory Profiling

```bash
go test -bench=. -memprofile mem.out
go tool pprof mem.out
```

برای دیدن:

- چه توابعی زیاد allocate می‌کنند.  
- چه objectهایی حجم زیادی می‌گیرند.

---

## 5. Block & Mutex Profiling

```bash
go test -blockprofile block.out
go tool pprof block.out

go test -mutexprofile mutex.out
go tool pprof mutex.out
```

برای پیدا کردن:

- نقاطی که goroutineها زیاد block شده‌اند.  
- contention روی mutexها.

---

## 6. Trace

```bash
go test -trace trace.out
go tool trace trace.out
```

نمایش:

- timeline goroutineها  
- channel events  
- syscalls  
- GC events  
- زمان‌بندی scheduler

---

## 7. Summary

- همیشه performance را اندازه بگیر، حدس نزن.  
- از benchmark برای APIهای مهم استفاده کن.  
- از pprof و trace برای یافتن bottleneck استفاده کن.  
- escape analysis را جدی بگیر.
