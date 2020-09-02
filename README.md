## orch_hook

### orchestrator hook
```bash
[root@slave ~]# go build main.go
[root@slave ~]# mv main /usr/local/bin/
[root@slave ~]# ./main -h
Usage of ./main:
  -cluster_name string
        cluster name (default "one")
  -dead_state string
        master status judge (default "DeadMaster")
  -interface string
        interface name (default "ens33")
  -new_master string
        New master host:port (default "slave")
  -old_master string
        Old master host:port (default "master:22")
  -ssh_addr string
        ssh address (default "192.168.0.1")
  -ssh_passwd string
        input ssh password
  -ssh_user string
        input ssh user (default "root")
  -vip_addr string
        vip address (default "192.168.0.120")

cluster_name:集群名称，对应meta.cluster表的cluster_name字段
dead_state:故障转移类型，通过外部Orchestrator的内置环境变量获取，环境变量名称：{failureType}
interface:VIP地址所在的网卡名称
new_master:通过选举得出的新的master及其ssh的端口，使用内置环境变量：{successorHost}获取hostname，后面跟ssh的端口。格式：<hostname>:<sshport>
old_master:同上，故障的老master及其ssh端口
vip_addr:VIP地址
ssh_passwd:ssh连接时的密码，使用root账户，--ssh_user参数无效，如果想使其有效，那么修改代码failover.go

[root@slave ~]# cat failover.sh
bash exec /usr/local/bin/main --dead_state={failureType} --vip_addr='your vip' --ssh_passwd='your passwd' --old_master={failedHost}:22 --interface='eth0' --new_master={successorHost}:22

```
