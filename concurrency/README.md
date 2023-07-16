### Concurrency

Concurrency, birbirinden bağımsız çalışmaların bir araya gelmesi, birleşmesi.
Birbirinden bağımsız çalışan işlemler, birbirlerini etkilemeden çalışırlar.

Bir bilgisayarda, birden fazla iş aynı anda çalışabilir. 
Bu işlerin birbirlerini etkilemeden çalışmaları için, işlemler arasında bir koordinasyon sağlanması gerekir. 
Bu koordinasyonu sağlayan yapıya, **thread** denir.

Her uygulama, en az bir thread içerir. Bu thread, uygulamanın main thread'idir.
Bir uygulama, birden fazla thread içerebilir. Bu thread'ler, birbirlerinden bağımsız çalışırlar.
Her thread sadece kendi işini yapar. Bir thread, diğer thread'lerin işini yapamaz.

**Context Switching**: Bir thread'in işini yaparken, başka bir thread'in işini yapması için, 
ilk thread'in işini yarım bırakıp, diğer thread'e geçmesi işlemine denir. 
Bu işlem, çok hızlı bir şekilde gerçekleşir ve bu sayede bir bilgisayar, birden fazla işi aynı anda yapmış olur.

**Parelelism**: Bir bilgisayarın birden fazla işi aynı anda yapmasıdır. 
Bir bilgisayarın birden fazla işi aynı anda yapabilmesi için, birden fazla çekirdeğe sahip olması gerekir.
Aynı anda birden fazla işi yapabilen bilgisayarlara, **multi-core** denir. 
Yeni bir thread oluşturulduğunda, bu thread, farklı bir çekirdekte çalışabilir.

**Gorooutines**: Go dilinde, thread'lere(lightweight thread) **goroutine** denir. 
Go dilinde, bir thread oluşturmak için, **go** anahtar kelimesi kullanılır. Bir thread oluşturulduğunda, 
bu thread, main thread'den bağımsız olarak çalışır, bu thread'in işi bittiğinde, main thread'e geri döner.

Normal bir fonksiyonu çalıştırdığımızda main thread'de çalışır. 
Goroutine oluşturduğumuzda, bu fonksiyon, main thread'den bağımsız olarak çalışır.

**Channels**: Thread'ler arasında haberleşme sağlamak için kullanılır.
Bir thread, bir channel'a mesaj gönderir, diğer thread, bu mesajı alır.

Channellar '<-' operatörü ile oluşturulur. 
Bir channel'a mesaj göndermek için, channel'ın isminin soluna '<-' operatörü ile mesaj gönderilir.
Bir channel'dan mesaj almak için, channel'ın isminin sağına '<-' operatörü ile mesaj alınır.
Varsayılan olarak, veri gönderme ve alma işlemleri, karşı taraf hazır olana kadar bekler.

**Örnek**: Bir channel oluşturup, bu channel'a mesaj gönderelim ve bu mesajı alalım.

```go
package main
import "fmt"

func main() {
    ch := make(chan string)
    go func() {
        ch <- "Merhaba"
    }()
    msg := <-ch
    fmt.Println(msg)
}
```

**Buffered Channels**: Bir channel'a mesaj göndermek için, channel'ın isminin soluna '<-' operatörü ile mesaj gönderilir.
Bir channel'dan mesaj almak için, channel'ın isminin sağına '<-' operatörü ile mesaj alınır.

Normal channel'dan farklı olarak, buffered channel'lar, mesajları biriktirirler.
**Örnek**: Bir buffered channel oluşturup, bu channel'a mesaj gönderelim ve bu mesajı alalım.

```go
package main
import "fmt"

func main() {
    ch := make(chan string, 1)
    ch <- "Merhaba"
    msg := <-ch
    fmt.Println(msg)
}
```
