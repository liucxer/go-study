module awesomeProject

go 1.16

replace (
	github.com/coreos/bbolt => github.com/coreos/bbolt v1.3.3
	github.com/go-cmd/cmd => github.com/liucxer/cmd v1.0.2
	google.golang.org/grpc => google.golang.org/grpc v1.26.0
)

require (
	github.com/davecgh/go-spew v1.1.1
	github.com/shirou/gopsutil/v3 v3.23.4
	github.com/sirupsen/logrus v1.9.0
)
