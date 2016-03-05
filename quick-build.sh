#!/bin/bash
echo Build start `date`
export GOPATH=`pwd`
go build h_temp_mon_app
echo Build end `date`