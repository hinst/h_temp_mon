#!/bin/bash
echo Unpack start `date`
mv content.tar ../
mv pkg ../
rm -R *
mv ../content.tar ./
mv ../pkg ./
tar xf content.tar
echo Unpack end `date`
