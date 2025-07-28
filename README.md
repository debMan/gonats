# Simulate multi NATS pods crash

## Commands

SSH into NC-Teh-<X>

```shell
. /etc/kolla/admin-openrc.sh && cd ~/new-cluster/ && pipenv shell
REGION="ts-2"
PROJECT_ID=$(openstack project list | grep "${REGION}" | cut -f2 -d\  )
openstack  server list --project "${PROJECT_ID}" --long -c ID -c Name -c Host  | grep worker-ls

# | 53430ed3-7da5-4093-95c9-4b4f4b5dad16 | okd4-worker-ls-4   | compute21 |
# | 15b0b9c1-101b-478d-a8e8-7743495d5b0e | okd4-worker-ls-3   | compute09 |
# | 79c955d2-3193-412b-a649-f4481fd1ea3e | okd4-worker-ls-2   | compute17 |
# | 8bc6ea0c-b6ca-44cc-9582-19338f5c514c | okd4-worker-ls-1   | compute03 |
# | a3ddc5ab-2c95-4356-b153-db214a4e89ec | okd4-worker-ls-0   | compute18 |

## Note the hosts
## Then ssh into hosts and kill the process of the VM

ssh compute<X>

VMID=<NODE_UUID>
VM_PS_NUMBER=$(ps aux | grep $VMID | head -n1  | awk '{ print $2 }')
sudo kill -9 ${VM_PS_NUMBER}
