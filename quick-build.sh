#!/bin/bash
echo Build start `date`
export GOPATH=`pwd`
go install h_temp_mon_app
mv bin/h_temp_mon_app ./
echo Build end `date`