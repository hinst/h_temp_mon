#!/bin/bash
mv content.tar ../
rm -R *
mv ../content.tar ./
tar xf content.tar
