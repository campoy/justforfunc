# justforfunc 19: mastering io.Pipes

We've all used a bit of the io package ... right?
Implementations of `io.Writer` and `io.Reader` can be literally found everywhere ... but do you really know the io package well?

<div style="text-align:center">
    <a href="https://www.youtube.com/watch?v=LHZ2CAZE6Gs&feature=youtu.be&list=PL6">
        <img src="https://img.youtube.com/vi/LHZ2CAZE6Gs/0.jpg" alt="justforfunc 19: mastering io.Pipes">
        <p>justforfunc 19: mastering io.Pipes</p>
    </a>
</div>

In this episode I write a program to do cat images into iTerm2, and to do that I use the best pieces of the io package.

<div style="text-align:center">
    <a href="https://www.youtube.com/watch?v=LHZ2CAZE6Gs&feature=youtu.be&list=PL6">
        <img src="https://img.youtube.com/vi/LHZ2CAZE6Gs/0.jpg" alt="justforfunc 19: mastering io.Pipes">
        <p>justforfunc 19: mastering io.Pipes</p>
    </a>
</div>


I use four types in the io package:
- [PipeReader](https://golang.org/pkg/io/#PipeReader)
- [PipeWriter](https://golang.org/pkg/io/#PipeWriter)
- [MultiReader](https://golang.org/pkg/io/#MultiReader)
- [MultiWriter](https://golang.org/pkg/io/#MultiWriter)

For a more complete version of this code, check [github.com/campoy/tools](https://github.com/campoy/tools/tree/master/imgcat).
