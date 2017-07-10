# justforfunc 15: a code review with logging, errors, and signals

Back to code reviews!
This time I review a piece of code by [Sandeep](https://twitter.com/sandeepdinesh),
from his repo [logpipe](https://github.com/thesandlord/logpipe).

The program simply reads from standard input and sends each line both to standard
output and a logging service.

What could go wrong? Many things!

<div style="text-align:center">
    <a href="https://www.youtube.com/watch?v=c5ufcpTGIJM&feature=youtu.be&list=PL6">
        <img src="https://img.youtube.com/vi/c5ufcpTGIJM/0.jpg" alt="justforfunc 15: a code review with logging, errors, and signals">
        <p>justforfunc 15: a code review with logging, errors, and signals</p>
    </a>
</div>

References:

- cloud logging docs: https://cloud.google.com/go/logging
- original repo: https://github.com/thesandlord/logpipe
- source code: https://github.com/campoy/justforfunc/tree/master/15-logpipe