package failover

import (
	"fmt"
	"log"
	"orch/myflag"
	"os"

	"golang.org/x/crypto/ssh"
)

var (
	logger      *log.Logger
	AllFlagArgs *myflag.OrchCfg
)

//WriteLogFile is open file write log file
func WriteLogFile() {
	file, err := os.OpenFile("./orch_failover.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		fmt.Printf("open file is failed err message=======>", err)
	}
	defer file.Close()
	logger = log.New(file, "", log.LstdFlags|log.Lshortfile)
	logger.SetFlags(log.LstdFlags | log.Lshortfile)
	CheckDeadType(logger)
}

func CheckDeadType(logger *log.Logger) {
	AllFlagArgs, err := myflag.MyFlagArgs()
	if err != nil {
		logger.Printf("get flag error =======> %v", err)
	}
	
	if AllFlagArgs.DeadStatus == "DeadMaster" {
		logger.Printf("Revocering from: %v", AllFlagArgs.DeadStatus)
		logger.Printf("New master is: %v", AllFlagArgs.NewMaster)
		ExecVipDel(AllFlagArgs)
		ExecVipAdd(AllFlagArgs)
	}
}

func ExecVipAdd(AllFlagArgs *myflag.OrchCfg) {
	client, err := ssh.Dial("tcp", AllFlagArgs.NewMaster, &ssh.ClientConfig{
		User:            AllFlagArgs.SSHUser,
		Auth:            []ssh.AuthMethod{ssh.Password(AllFlagArgs.SSHPasswd)},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	})
	if err != nil {
		logger.Printf("connect error=====> %v", err)
	}
	session, err1 := client.NewSession()
	if err1 != nil {
		logger.Printf("create session failed =======> %v", err1)
	}
	defer session.Close()
	comb, err2 := session.CombinedOutput(AllFlagArgs.CmdVipAdd)
	if err2 != nil {
		logger.Printf("command exec failed =======> %v", err2)
	}
	logger.Printf("exec command is : %v", string(comb))
}

func ExecVipDel(AllFlagArgs *myflag.OrchCfg) {
	client, err := ssh.Dial("tcp", AllFlagArgs.OldMaster, &ssh.ClientConfig{
		User:            AllFlagArgs.SSHUser,
		Auth:            []ssh.AuthMethod{ssh.Password(AllFlagArgs.SSHPasswd)},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	})
	if err != nil {
		logger.Printf("connect error =====> %v", err)
	}
	session, err1 := client.NewSession()
	if err1 != nil {
		logger.Printf("create session failed=======> %v", err1)
	}
	defer session.Close()
	comb, err2 := session.CombinedOutput(AllFlagArgs.CmdVipDel)
	if err2 != nil {
		logger.Printf("command exec failed=======> %v", err2)
	}
	logger.Printf("exec command is:%s", string(comb))
}
