#!/bin/bash
echo Unpack start `date`
mv content.tar ../
rm -R *
mv ../content.tar ./
tar xf content.tar
echo Unpack end `date`
