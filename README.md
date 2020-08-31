## orch_hook

### orchestrator hook
```bash
go build main.go
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


bash exec /usr/local/bin/main --dead_state={failureType} --vip_addr='your vip' --ssh_passwd='your passwd' --old_master={failedHost}:22 --interface='eth0' --new_master={successorHost}:22
:22是SSH的端口
```
