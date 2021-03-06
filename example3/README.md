# example3

## Obtaining the Dependencies

This example uses [dep](https://github.com/golang/dep), which is a prototype
dependency management tool for Go.

To install dep, follow the instructions in dep's [documention](https://golang.github.io/dep/docs/installation.html).

Installing the correct versions of the depencies used in example1 can be done through
the use of dep's `ensure` command. More documentation on the `ensure` command can
be found [here](https://golang.github.io/dep/docs/daily-dep.html).

For this example, run the following to catch up on installing the dependencies.

```
$ dep ensure -v
```

## Build Static Binary

To build the single, static, binary for example3, run the following in `example3/`:

```
$ go build -o bin/example3
```

This should create an executable binary in the local `bin` directory.

## Running the Executable Binary

To run the built binary, use the following:

```
$ ./bin/example3
```

This should produce something similar to the following output:

```
2018/02/11 21:06:30 started server on localhost:8080
starting goroutine for listening for history
```

Then, open your browser and visit `localhost:8080`
