
# progress

Making a progress for `Reader` or `Writer`.

View the [docs][].


## Usage


### Write bytes in progress

```go
filename := "progress_test.go"
f, err := os.Open(filename)
defer f.Close()
if err != nil {
  log.Fatalln(err)
}
fs, err := os.Stat(filename)
if err != nil {
  log.Fatalln(err)
}

p := NewWriter()
p.Total = fs.Size()
p.Progress = func(current, total, expected int64) {
  log.Println(current, total, expected)
  assert.Equal(t, true, current <= total)
}

b := new(bytes.Buffer)
w := io.MultiWriter(p, b)
_, err = io.Copy(w, f)
if err != nil {
  log.Fatalln(err)
}
```


[docs]: http://godoc.org/github.com/go-libs/progress

