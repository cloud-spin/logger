# Logger [![Build Status](https://travis-ci.com/cloud-spin/logger.svg?branch=master)](https://travis-ci.com/cloud-spin/logger) [![codecov](https://codecov.io/gh/cloud-spin/logger/branch/master/graph/badge.svg)](https://codecov.io/gh/cloud-spin/logger) [![Go Report Card](https://goreportcard.com/badge/github.com/cloud-spin/logger)](https://goreportcard.com/report/github.com/cloud-spin/logger) [![GoDoc](https://godoc.org/github.com/cloud-spin/logger?status.svg)](https://godoc.org/github.com/cloud-spin/logger)

Package logger provides standard logging methods around the standard log package.
Each logged line is prefixed with the logging level (Debug, Info, Warn, ...).
The Enabled and Level configurations are respected every time the logging methods are called.
Logger provides a Logger interface for easy mocking and also an OnLog handler to help identify logged data in unit tests.

#### Install

From a configured [Go environment](https://golang.org/doc/install#testing):
```sh
go get -u github.com/cloud-spin/logger
```

If you are using dep:
```sh
dep ensure -add github.com/cloud-spin/logger
```


#### How to Use

Below example starts a enabled logger to Info level and logs a Info message with one parameter.

```go
package main

import (
	"fmt"

	"github.com/cloud-spin/logger"
)

func main() {
	configs := &logger.Configs{
		Enabled: true,
		Level:   logger.LevelInfo,
	}
	logger, err := logger.NewLogger(configs)
	if err != nil {
		fmt.Printf("Expected: logger initialized; Got: %s", err.Error())
	}

	logger.Info("Info message with '%s' param", "string")
}
```

Output:
```
[INFO] 2018/09/10 11:20:11 Info message with 'string' param
```

Also refer to the tests at [logger_test.go](logger_test.go).


## License
MIT, see [LICENSE](LICENSE).

"Use, abuse, have fun and contribute back!"


## Contributions
See [Contributions.md](https://github.com/cloud-spin/docs/blob/master/contributing.md).
