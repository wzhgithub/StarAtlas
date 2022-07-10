#!/bin/bash
set -ex
rm -f star_atlas_server
cd ../client
sh run.sh
cp bin/utest tpinit.sh ../star_atlas_server
cd -
chmod a+x utest && mv tpinit.sh gen_udp.sh
docker build --platform=linux/amd64 -t star_atlas_server .
