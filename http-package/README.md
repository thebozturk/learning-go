### HTTP PACKAGE

Go dilinde http paketi ile beraber gelen bir çok fonksiyon bulunmaktadır. Bu fonksiyonlar ile beraber http istekleri atabilir, http isteklerine cevap verebilir, http isteklerini yönlendirebilir ve http isteklerini dinleyebiliriz.

### Reverse Proxy

Go dilinde http paketi ile beraber gelen `httputil.ReverseProxy` tipi ile http isteklerini yönlendirebiliriz. Bu tipi kullanarak http isteklerini yönlendirmek için `httputil.NewSingleHostReverseProxy()` fonksiyonunu kullanabiliriz.
Mesela bizim bir uygulamamız var ve 100 tane sunucuda çalışıyor, burada tek bir sunucu üzerinden bu 100 sunucuya erişmek istiyoruz. Bu durumda reverse proxy kullanabiliriz.
O sunucuların adresini mevcut domaine yönlendirebiliriz. Bu sayede istekler tek bir sunucuya gelir ve bu sunucu isteği ilgili sunucuya yönlendirir.

```go
func main() {
    proxy := httputil.NewSingleHostReverseProxy(&url.URL{
        Scheme: "http",
        Host:   "localhost:8080",
    })
    log.Fatal(http.ListenAndServe(":8081", proxy))
}
```

### HTTP İstekleri

Go dilinde http istekleri atmak için `http.Get()` fonksiyonu kullanılır. Bu fonksiyon bize bir `http.Response` ve `error` döndürür. `http.Response` içerisinde isteğin durumu, başlıkları ve gövdesi bulunur. `error` ise isteğin başarılı olup olmadığını belirtir.

```go
resp, err := http.Get("https://www.google.com")
if err != nil {
    log.Fatal(err)
}
```

### HTTP İsteklerine Cevap Vermek

Go dilinde http isteklerine cevap vermek için `http.ResponseWriter` ve `*http.Request` parametrelerini alan bir fonksiyon yazmamız gerekmektedir. Bu fonksiyonu `http.HandleFunc()` fonksiyonuna parametre olarak verdiğimizde isteklere cevap verebiliriz.

```go
func handler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
}

func main() {
    http.HandleFunc("/", handler)
    log.Fatal(http.ListenAndServe(":8080", nil))
}
```
