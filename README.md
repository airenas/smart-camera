https://techship.com/faq/how-to-guide-control-and-set-up-a-data-connection-in-linux-using-modemmanager-as-connection-manager/

https://techship.com/faq/how-to-step-by-step-set-up-a-data-connection-over-qmi-interface-using-qmicli-and-in-kernel-driver-qmi-wwan-in-linux/

https://lists.freedesktop.org/archives/libqmi-devel/2017-August/002410.html


mmcli --scan-modems 
mmcli --list-modems
mmcli --modem=0
mmcli --modem=0 --enable
mmcli -m 0 --simple-connect='apn=telenorbg,allowed-auth=none,ip-type=ipv4'


qmicli --device=/dev/cdc-wdm0 --device-open-proxy --wds-start-network="ip-type=4,apn=telenorbg" --client-no-release-cid


busybox udhcpc -q -f -n -i wwan0


qmicli -p -d /dev/cdc-wdm0 --device-open-net='net-raw-ip|net-no-qos-header' --wds-start-network="apn='telenorbg',ip-type=4" --client-no-release-cid



qmicli --device=/dev/cdc-wdm0 --device-open-proxy --wds-stop-network= --client-cid=


qmicli --device=/dev/cdc-wdm0 --device-open-proxy --wds-start-network="ip-type=4,apn=telenorbg" --client-no-release-cid

# check status
qmicli -d /dev/cdc-wdm0 --dms-get-operating-mode
qmicli -p -d /dev/cdc-wdm0 --wds-get-packet-service-status
qmicli -p -d /dev/cdc-wdm0 --wds-get-current-settings

qmicli -p -d /dev/cdc-wdm0 --nas-get-signal-info

ip addr add 10.165.216.237 dev wwan0

ip route add default dev wwan0

ip addr

ip route



ping -I wwan0 8.8.8.8

speedtest --source=10.165.216.237
