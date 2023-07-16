Go programlama dilinde spesifik bir ortam değişkenini almak için `os` paketinin `Getenv()` fonksiyonu kullanılır.

```go
package main
import (
    "fmt"
    "os"
)

func main() {
    fmt.Println("PATH:", os.Getenv("PATH"))
    fmt.Println("HOME:", os.Getenv("HOME"))
    fmt.Println("USER:", os.Getenv("USER"))
}
```

Sistem ortam değişkenlerini almak için `os.Environ()` fonksiyonu kullanılır. Bu fonksiyon bir dizi döndürür. Dizinin her bir elemanı bir ortam değişkenini temsil eder.

```go
package main
import (
    "fmt"
    "os"
)

func main() {
    env := os.Environ()
    for _, v := range env {
        fmt.Println(v)
    }
}
```