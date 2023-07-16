### Error-Handling - Exception Handling

Go dilinde error handling için `error` tipi kullanılır.
    
```go
 func main() {
        f, err := os.Open("filename.ext")
        if err != nil {
            log.Fatal(err)
        }
        fmt.Println(f)
    }
   ```

### Error-Handling - Custom Errors

Go dilinde custom error oluşturmak için `error` interface'ini implemente etmek yeterlidir.

```go
    type customError struct {
        arg int
        errMsg string
    }

    func (e *customError) Error() string {
        return fmt.Sprintf("%d - %s", e.arg, e.errMsg)
    }

    func main() {
        err := test(1)
        if err != nil {
            log.Fatal(err)
        }
    }

    func test(arg int) error {
        if arg < 0 {
            return &customError{arg, "Bad Arguments"}
        }
        return nil
    }
```

### Exception Handling - Panic

Go dilinde `panic` fonksiyonu ile programın çalışması durdurulabilir. `panic` fonksiyonu `defer` ile birlikte kullanıldığında `panic` fonksiyonu çalıştırıldıktan sonra `defer` fonksiyonları çalıştırılır.

```go
    func main() {
        fmt.Println("start")
        defer func() {
            if err := recover(); err != nil {
                log.Fatal(err)
            }
        }()
        panic("something bad happened")
        fmt.Println("end")
    }
```

### Exception Handling - Defer

Go dilinde `defer` fonksiyonu ile fonksiyonun sonunda çalıştırılacak fonksiyonlar tanımlanabilir. `defer` fonksiyonları stack mantığı ile çalışır. Yani en son tanımlanan fonksiyon en önce çalışır.

```go
    func main() {
        defer fmt.Println("start")
        defer fmt.Println("middle")
        defer fmt.Println("end")
    }
```

### Exception Handling - Recover

Go dilinde `recover` fonksiyonu ile `panic` fonksiyonu ile durdurulan programın çalışması devam ettirilebilir.

```go
    func main() {
        fmt.Println("start")
        defer func() {
            if err := recover(); err != nil {
                log.Fatal(err)
            }
        }()
        panic("something bad happened")
        fmt.Println("end")
    }
```

### Exception Handling - Panic vs. OS Exit

Go dilinde `panic` fonksiyonu ile programın çalışması durdurulabilir. `panic` fonksiyonu `defer` ile birlikte kullanıldığında `panic` fonksiyonu çalıştırıldıktan sonra `defer` fonksiyonları çalıştırılır. `panic` fonksiyonu ile programın çalışması durdurulduğunda `defer` fonksiyonları çalıştırılırken `os.Exit` fonksiyonu ile programın çalışması durdurulduğunda `defer` fonksiyonları çalıştırılmaz.

```go
    func main() {
        fmt.Println("start")
        defer fmt.Println("this was deferred")
        os.Exit(-1)
    }
```