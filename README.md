# URL shortener

URL shortener JSON API.

### Usage
```sh
$ git clone https://github.com/markusvaikjarv/url_shortener.git
$ cd url_shortener
$ go run main.go responseStructs.go
```

```
$ curl localhost:4321/shorten?url=facebook.com

{
  "Success": true,
  "ShortenedURL": "localhost:4321/get/0",
  "Error": ""
}
```
```
$ curl localhost:4321/shorten?url=http://sitethatdoesnotexist.ee

{
  "Success": false,
  "ShortenedURL": "",
  "Error": "URL seems unreachable"
}
```
```
$ curl localhost:4321/show

{
  "0":"http://facebook.com",
  "1":"http://facebook.com",
  "2":"https://github.com"
}
```
```
$ curl localhost:4321/get/2
{
  "Exists": true,
  "URL": "https://github.com"
}
```
