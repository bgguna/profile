# Server

The server for this web application is built using [Go](https://golang.org).

* [gin-gonic](https://gin-gonic.com) is used for the HTTP endpoints
* [sqlite](https://sqlite.org/index.html) is used for storing contact messages

A `makefile` is available to setup/handle the server easily.

### Compiles and run in development mode

```
make run-dev
```

### Builds the binary file

```
make build
```

### Runs the binary file

```
make run
```

### Lists available `make` commands

```
make help
```
