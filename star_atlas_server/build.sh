#!/bin/bash
set -ex
rm -f star_atlas_server
cd ../client
sh run.sh
cp bin/utest ../star_atlas_server
cd -
chmod a+x utest
docker build --platform=linux/amd64 -t star_atlas_server .
