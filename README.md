## orch_hook
该脚本为vip漂移脚本，我们是使用的cmdb记录主机的vip等信息的，所以vip是从cmdb取的
也可以用orchestrator自己的配置：DetectClusterDomainQuery或者其他探测语句把vip写到表里
### orchestrator hook
```bash
[root@slave ~]# go build main.go
[root@slave ~]# mv main /usr/local/bin/
[root@slave ~]# ./main -h
Usage of ./main:
  -ClusterName string
        cluster name (default "one")
  -DeadStatus string
        master status judge (default "DeadMaster")
  -NewMaster string
        interface name (default "ens33")
  -NewMasterPort string
        New master host:port (default "slave")
  -OldMaster string
        Old master host:port (default "master:22")
  -OldMasterPort string
        ssh address (default "192.168.0.1")

cluster_name:集群名称，对应meta.cluster表的cluster_name字段
DeadStatus:故障转移类型，通过外部Orchestrator的内置环境变量获取，环境变量名称：{failureType}
NewMaster:使用内置环境变量：{successorHost}获取hostname
NewMasterPort:新主端口
OldMaster:同上，故障的老master
OldMasterPort:老主端口

[root@slave ~]# cat failover.sh
bash exec /usr/local/bin/main --DeadStatus={failureType} --ClusterName='cluster name' --OldMaster={failedHost} --OldMasterPort=3306 --NewMaster={successorHost} --NewMasterPort=3306

```
