# sureshort
A URL Shortner Service

This is a simple URL shortener service that accepts a URL as an argument over a REST API and
returns a shortened URL as a result. And then redirects you to the original URL whenever you try to open shortened URL (Similar to https://bitly.com or https://tinyurl.com)

### Supported Features
<p>

- Provides `/app/create` API that accepts
    1. a URL as query parameter named `url` over a HTTP GET request and returns shortened URL
        ```
        <server-address>/app/create?url=google.com
        ```
    2. a URL as argument/field in JSON request body over HTTP POST request and returns shortened URL
        ```
        {"url": "google.com"}
        ```
  For which server returns an response as follows:
  
  text/html:
  ```
  <server-address>/e14f0993
  ```
  application/json
  ```
    {"shortened_url": "<server-address>/e14f0993"}
  ```
- Provides `/<shortened-url>` API which supports a HTTP GET call and redirects the caller to original URL associated with short version.

- Provides `/app/metrics` API which returns top 3 domain names that have been shortened the most
number of times. For eg. if the user hasshortened 4 YouTube video links, 1 StackOverflow link,2
Wikipedia links and 6 Udemy tutorial links. Then the output would be:<br>
    ```
    Udemy: 6,
    YouTube: 4,
    Wikipedia: 2
    ```

</p>

### Steps to build and run the service

#### Build
```bash
$ go mod tidy
$ go build
```

#### Run
If you need to updated configurations like server listener address and port, you can update config/config.defaults.yaml for now
```bash
$ ./sureshort

   ____    __
  / __/___/ /  ___
 / _// __/ _ \/ _ \
/___/\__/_//_/\___/ v4.13.4
High performance, minimalist Go web framework
https://echo.labstack.com
____________________________________O/_______
                                    O\
⇨ http server started on [::]:80
```

#### Test
```bash
$ go test -coverprofile=coverage.out ./...
        github.com/surajbhosale409/sureshort            coverage: 0.0% of statements
ok      github.com/surajbhosale409/sureshort/pkg        0.330s  coverage: 100.0% of statements
ok      github.com/surajbhosale409/sureshort/service    1.459s  coverage: 91.1% of statements
```