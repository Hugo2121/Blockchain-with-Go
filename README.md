A simple concurrent Blockchain app written in Go language.

There are two types of models: 
Work queue: capable of managing n workers executing tasks simultaneously and queuing maxJob tasks.
Worker routine: processes tasks from jobs unless StopRequests has a message requesting the worker to halt.

How to run:

under go folder

$ export GOPATH=`pwd`
$ go run src/main/main.go
