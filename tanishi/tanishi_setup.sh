#! /bin/bash

uname -a > osinfo.txt
lscpu >> osinfo.txt
free -h >> osinfo.txt
ls $1 *.go | xargs sed -r -n '/(Getenv)/p' > /tmp/env.txt

systemctl list-unit-files > systemd_list.txt
echo "------------" >> systemd_list.txt
systemctl list-unit-files | grep enabled >> systemd_list.txt

bash <(curl -Ss https://my-netdata.io/kickstart.sh)

apt-get install htop iotop
