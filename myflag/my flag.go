package myflag

import (
	"flag"
	"fmt"
	_ "os/exec"
)

var (
	o OrchCfg
)

type OrchCfg struct {
	DeadStatus string
	OldMaster  string
	NewMaster  string
	Cluster    string
	SSHAddr    string
	SSHUser    string
	SSHPasswd  string
	VipAddr    string
	Interface  string
	CmdVipAdd  string
	CmdVipDel  string
	CmdVipStat string
}

func MyFlagArgs() (*OrchCfg, error) {
	flag.StringVar(&o.SSHUser, "ssh_user", "root", "input ssh user")
	flag.StringVar(&o.SSHPasswd, "ssh_passwd", "", "input ssh password")
	flag.StringVar(&o.DeadStatus, "dead_state", "DeadMaster", "master status judge")
	flag.StringVar(&o.OldMaster, "old_master", "master:22", "Old master host:port")
	flag.StringVar(&o.NewMaster, "new_master", "slave", "New master host:port")
	flag.StringVar(&o.Cluster, "cluster_name", "one", "cluster name")
	flag.StringVar(&o.VipAddr, "vip_addr", "192.168.0.120", "vip address")
	flag.StringVar(&o.Interface, "interface", "ens33", "interface name")
	flag.StringVar(&o.SSHAddr, "ssh_addr", "192.168.0.1", "ssh address")
	flag.Parse()
	o.CmdVipAdd = fmt.Sprintf("ip addr add %v dev %v", o.VipAddr, o.Interface)
	o.CmdVipDel = fmt.Sprintf("ip addr del %v dev %v", o.VipAddr, o.Interface)
	o.CmdVipStat = fmt.Sprintf("ping -c 1 -W 1 %v", o.VipAddr)
	return &OrchCfg{SSHUser: o.SSHUser, SSHPasswd: o.SSHPasswd, DeadStatus: o.DeadStatus, OldMaster: o.OldMaster, NewMaster: o.NewMaster, Cluster: o.Cluster, VipAddr: o.VipAddr, Interface: o.Interface, CmdVipAdd: o.CmdVipAdd, CmdVipDel: o.CmdVipDel, CmdVipStat: o.CmdVipStat, SSHAddr: o.SSHAddr}, nil

}
