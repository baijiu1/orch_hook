package main

import (
	_ "fmt"
	"orch/failover"
)

func main() {
	//写日志，里面调用了myflag包
	failover.WriteLogFile()
}
