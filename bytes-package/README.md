### Bytes Paketi

Go dilinde bytes paketi, bayt dizileri üzerinde işlem yapmak için kullanılır. Bayt dizileri, değiştirilemez dizilerdir. Bu nedenle bayt dizileri üzerindeki işlemler, yeni bir bayt dizisi oluşturularak yapılır.

#### Bayt Dizisi Oluşturma

```go
package main
import "fmt"

func main() {
    var b []byte
    b = []byte{'G', 'o', 'l', 'a', 'n', 'g'}
    fmt.Println(b)
}
```
Yukarıdaki örnekte, bayt dizisi oluşturulmuştur. Bayt dizisi oluşturulurken, bayt dizisinin uzunluğu belirtilmemiştir. Bu nedenle bayt dizisinin uzunluğu, bayt dizisi içerisindeki eleman sayısı kadardır.
