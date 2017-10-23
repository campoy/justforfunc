# Using the Go tracer

Watch the episode here:

<div style="text-align:center">
    <a href="https://www.youtube.com/watch?v=ySy3sR1LFCQ&feature=youtu.be&list=PL6">
        <img src="https://img.youtube.com/vi/ySy3sR1LFCQ/0.jpg" alt="justforfunc 22: using the Go tracer">
        <p>justforfunc 22: using the Go tracer</p>
    </a>
</div>

What is the fastest way to compute a Mandelbrot set like the one below?

![mandelbrot](out.png)

The final times are

| mode                  | seconds |
|-----------------------|---------|
| sequential            |   4.669 |
| pixels                |   3.004 |
| rows                  |   0.689 |
| workers               |   2.967 |
| workers + buffer      |   1.226 |
| row workers           |   0.714 |
| row workers + buffer  |   0.698 |

You can run the benchmarks yourself by running:

```bash
$ go test -bench=.

```

Also, modify the constant named `complexity` in the `pixel` function to compute simpler or more complex fractals.
