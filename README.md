# logfmt
lite logrus formatter

## install
```
% go get -u -v github.com/lostelk/logfmt
```

## usage
```go
package main

import (
    "github.com/Sirupsen/logrus"
    "github.com/lostelk/logfmt"
)

func main() {
    logrus.SetFormatter(logfmt.DefaultFormatter)
    logrus.Infof("hello world")
}
```
