@echo off

cd .\cmd\inverter_exporter

set GOOS=linux
set GOARCH=arm

go build .

scp .\inverter_exporter root@10.11.10.194:/root/addons/inverter_exporter

ssh root@10.11.10.194 "chmod +x /root/addons/inverter_exporter/inverter_exporter"

del inverter_exporter

cd ..\..\