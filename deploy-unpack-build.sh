#!/bin/bash
bash deploy-unpack.sh
export GOPATH=`pwd`
go build h_temp_mon_app