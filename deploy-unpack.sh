#!/bin/bash
mv content.tar ../
rm -R *
mv ../content.tar ./
tar xvf content.tar
