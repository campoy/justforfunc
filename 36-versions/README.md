# Versions, build constraints, and ldflags

During my [Go Tooling in Action Workshop](https://github.com/campoy/go-tooling-workshop)
many people ask how to include version information to binaries,
so I decided to make an episode out of it!

We will solve the problem using three different techniques:

- constants and build constraints ([code](consts))
- variables and build constraints ([code](vars))
- variables and ldflags ([code](ldflags))

In each directory you can build the `development` and `production`
releases by running `make dev` or `make prod`.

<div style="text-align:center">
    <a href="https://youtu.be/-XSlev-d7UY">
        <img src="https://img.youtube.com/vi/-XSlev-d7UY/0.jpg" alt="justforfunc #36: Versions, build constraints, and ldflags">
        <p>justforfunc #36: Versions, build constraints, and ldflags</p>
    </a>
</div>