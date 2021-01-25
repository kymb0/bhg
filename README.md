# bhg
Code samples for the No Starch Press Black Hat Go

srs biz.

# Chapter Overview
* Chapter 1: Go Fundamentals and Concepts
* Chapter 2: TCP and Go: Scanners and Proxies
* Chapter 3: HTTP Clients: Remote Interaction with Tools
* Chapter 4: HTTP Servers: Routing and Middleware
* Chapter 5: Exploiting DNS: Recon and More
* Chapter 6: SMB and NTLM: A Peek Down the Rabbit Hole
* Chapter 7: Databases and Filesystems: Pilfering and Abusing
* Chapter 8: Packet Processing: Living on the Wire
* Chapter 9: Exploit Code: Writing and Porting
* Chapter 10: Extending Tools: Using Go Plugins and Lua
* Chapter 11: Cryptography: Implementing and Attacking
* Chapter 12: Windows: System Interaction and Analysis
* Chapter 13: Steganography: Hiding Data
* Chapter 14: Command and Control: Building a RAT


# What This Book Isn’t
```
This book is not an introduction to Go programming in general but an
introduction to using Go for developing security tools. We are hackers and
then coders—in that order. None of us have ever been software engineers.
This means that, as hackers, we put a premium on function over elegance.
In many instances, we’ve opted to code as hackers do, disregarding some
of the idioms or best practices of software design. As consultants, time is
always scarce; developing simpler code is often faster and, therefore, preferable over elegance. When you need to quickly create a solution to a problem,
style concerns come secondary.
This is bound to anger Go purists, who will likely tweet at us that we
don’t gracefully handle all error conditions, that our examples could be
optimized, or that better constructs or methods are available to produce
the desired results. We’re not, in most cases, concerned with teaching you
the best, the most elegant, or 100 percent idiomatic solutions, unless doing
so will concretely benefit the end result. Although we’ll briefly cover the
language syntax, we do so purely to establish a baseline foundation upon
which we can build. After all, this isn’t Learning to Program Elegantly with
Go—this is Black Hat Go.
```

# Why Use Go for Hacking?
```
Prior to Go, you could prioritize ease of use by using dynamically typed
languages—such as Python, Ruby, or PHP—at the expense of performance
and safety. Alternatively, you could choose a statically typed language,
like C or C++, that offers high performance and safety but isn’t very userfriendly. Go is stripped of much of the ugliness of C, its primary ancestor,
making development more user-friendly. At the same time, it’s a statically
typed language that produces syntax errors at compile time, increasing
your assurance that your code will actually run safely. As it’s compiled, it
performs more optimally than interpreted languages and was designed
with multicore computing considerations, making concurrent programming a breeze.
These reasons for using Go don’t concern security practitioners specifically. However, many of the language’s features are particularly useful for
hackers and adversaries:
Clean package management system Go’s package management solution is elegant and integrated directly with Go’s tooling. Through the
use of the go binary, you can easily download, compile, and install packages and dependencies, which makes consuming third-party libraries
simple and generally free from conflict.
Cross-compilation One of the best features in Go is its ability to
cross-compile executables. So long as your code doesn’t interact with
raw C, you can easily write code on your Linux or Mac system but compile the code in a Windows-friendly, Portable Executable format.
Rich standard library Time spent developing in other languages has
helped us appreciate the extent of Go’s standard library. Many modern
languages lack the standard libraries required to perform many common
tasks such as crypto, network communications, database connectivity,
and data encoding (JSON, XML, Base64, hex). Go includes many of
these critical functions and libraries as part of the language’s standard
packaging, reducing the effort necessary to correctly set up your development environment or to call the functions.
Concurrency Unlike languages that have been around longer, Go
was released around the same time as the initial mainstream multicore
processors became available. For this reason, Go’s concurrency patterns
and performance optimizations are tuned specifically to this model.
```
