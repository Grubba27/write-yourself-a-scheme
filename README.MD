
This is a Repo that was build alongside [eatonphil](twitch.tv/eatonphil) from the twitch stream.

[Part 1: A lexer](https://www.youtube.com/watch?v=lZNhZI-dN9k)
[Part 2: Parsing](https://www.youtube.com/watch?v=5ttFEPQopXc)
[Part 3: AST walking interpreter](https://www.youtube.com/watch?v=YwmGcverSHI)
[Part 4: Cleanup and Fibonacci](https://www.youtube.com/watch?v=skDhTWILH8I)

Give him a star on the original [repo](https://github.com/eatonphil/livescheme) and follow him on twitch if you like the content. 

```
$ go mod tidy
$ go test
$ go build
$ cat examples/fib.scm
(fn fib (a)
      (if (< a 2)
	  a
	  (add (fib (sub a 1)) (fib (sub a 2)))))

(fib 11)
$ ./livescheme examples/fib.scm
89

```