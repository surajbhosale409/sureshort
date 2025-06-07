# sureshort
A URL Shortner Service

This is a simple URL shortener service that accepts a URL as an argument over a REST API and
returns a shortened URL as a result. (Similar to bitly.com or tinyurl.com)

### Supported Features
<p>

- To provide `/app/create` API that accepts
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
- To provide `/<shortened-url>` API which supports a HTTP GET call and redirects the caller to original URL associated with short version.

</p>

### Upcoming Features
<p>


- To provide `/app/metrics` API which returns top 3 domain names that have been shortened the most
number of times. For eg. if the user hasshortened 4 YouTube video links, 1 StackOverflow link,2
Wikipedia links and 6 Udemy tutorial links. Then the output would be:<br>
    ```
    Udemy: 6,
    YouTube: 4,
    Wikipedia: 2
    ```

</p>