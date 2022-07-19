#!/bin/bash
set -ex
rm -f star_atlas_server
cd ../client
sh run.sh
sh run_trans.sh
cp bin/utest bin/trans tpinit.sh ../star_atlas_server
cd -
chmod a+x utest && chmod a+x trans && mv tpinit.sh gen_udp.sh
docker build --platform=linux/amd64 -t star_atlas_server .
