# orch_hook

orchestrator hook
go build main.go
bash exec /usr/local/bin/main --dead_state=$1 --vip_addr='your vip' --ssh_passwd='your passwd' --old_master=$2:22 --interface='eth0' --new_master=$3:22
