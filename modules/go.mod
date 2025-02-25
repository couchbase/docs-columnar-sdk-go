module docs

require (
	github.com/couchbaselabs/gocbcolumnar v0.0.1
	github.com/sirupsen/logrus v1.9.3
)

require (
	github.com/couchbase/gocbcore/v10 v10.5.4 // indirect
	github.com/couchbaselabs/gocbconnstr v1.0.5 // indirect
	github.com/golang/snappy v0.0.4 // indirect
	github.com/google/uuid v1.6.0 // indirect
	github.com/konsorten/go-windows-terminal-sequences v1.0.3 // indirect
	golang.org/x/sys v0.30.0 // indirect
)

go 1.23

replace github.com/couchbaselabs/gocbcolumnar => ../../gocb-columnar
