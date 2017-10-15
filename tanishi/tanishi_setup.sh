#! /bin/bash

uname -a >> env.txt
lscpu >> env.txt
free -h >> env.txt

bash <(curl -Ss https://my-netdata.io/kickstart.sh)

apt-get install htop iotop
