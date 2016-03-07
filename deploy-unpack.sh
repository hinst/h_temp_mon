#!/bin/bash
echo Unpack start `date`
mv content.tar ../
mv pkg ../
rm -R *
mv ../content.tar ./
mv ../pkg ./
tar xf content.tar
chmod a+rx css -R
chmod a+rx js -R
echo Unpack end `date`
