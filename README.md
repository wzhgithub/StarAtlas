# StarAtlas

## star_atlas_server base gin service
```
cd star_atlas_server
# if not exists go.mod file just execute this step
touch go.mod & go mod edit --module=star_atlas_server
# local run & debug
sh debug.sh
# release & pushlish
sh build.sh
# test
cd test && python udp_send.py
```
