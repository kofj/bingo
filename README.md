# bingo
Benchmark test for golang standard library.



## Integer Format

Compare format integer with `fmt` package and `strconv` package.

```bash
go test -bench "FmtSprintInt|StrconvFmtInt"
```
