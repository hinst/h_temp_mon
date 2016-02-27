#!/bin/bash
bash deploy-unpack.sh
export GOPATH=`pwd`
go test h_temp_mon