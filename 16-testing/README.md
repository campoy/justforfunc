# justforfunc 16: unit testing HTTP servers

Let's cover the basics of unit testing in Go and then show how you can test
[http.HandlerFunc](https://golang.org/pkg/net/http#HandlerFunc) and
[http.Handler](https://golang.org/pkg/net/http#Handler).

You'll learn how to use "testing", including subtests and examples, and "net/http/httptest" with ResponseRecorder and Server.

<div style="text-align:center">
    <a href="https://www.youtube.com/watch?v=hVFEV-ieeew&feature=youtu.be&list=PL6">
        <img src="https://img.youtube.com/vi/hVFEV-ieeew/0.jpg" alt="justforfunc 16: unit testing HTTP servers">
        <p>justforfunc 16: unit testing HTTP servers</p>
    </a>
</div>

References:

- testing: https://golang.org/pkg/testing/
- net/http/httptest: https://golang.org/pkg/net/http/httptest
- Testable Examples in Go: https://blog.golang.org/examples
- source code: https://github.com/campoy/justforfunc/tree/master/16-testing
