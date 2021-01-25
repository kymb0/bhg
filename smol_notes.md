# just smol things
_reduce bloat in your binary:_
  * `$ go build -ldflags "-w -s"`

_Cross compilation:_
```
To cross-compile, you need to set a constraint. This is just a means to
pass information to the build command about the operating system and
architecture for which you’d like to compile your code. These constraints
include GOOS (for the operating system) and GOARCH (for the architecture).
You can introduce build constraints in three ways: via the command
line, code comments, or a file suffix naming convention. We’ll discuss the
command line method here and leave the other two methods for you to
research if you wish.
Let’s suppose that you want to cross-compile your previous hello.go
program residing on a macOS system so that it runs on a Linux 64-bit 
8 Chapter 1
architecture. You can accomplish this via the command line by setting the
GOOS and GOARCH constraints when running the build command:
$ GOOS="linux" GOARCH="amd64" go build hello.go
$ ls
hello hello.go
$ file hello
hello: ELF 64-bit LSB executable, x86-64, version 1 (SYSV), statically linked, not stripped
The output confirms that the resulting binary is a 64-bit ELF (Linux) file.
The cross-compilation process is much simpler in Go than in just about
any other modern programming language. The only real “gotcha” happens
when you try to cross-compile applications that use native C bindings. We’ll
stay out of the weeds and let you dig into those challenges independently.
Depending on the packages you import and the projects you develop, you
may not have to worry about that very often.
```
