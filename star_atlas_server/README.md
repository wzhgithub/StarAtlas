# remote deployment
```shell
host: http://nj.kyledong.com:9999/ 
ssh drake@nj.kyledong.com -p 8027
pwd: MEOWMEOWMEOW!@#
sudo su
```

# build & run 
```shell
./build.sh
docker compose up -d 
```
---
# go into service docker container and run udp script
```shell
docker exec -it star_atlas_server /bin/bash
sh gen_udp.sh
```
---
# go into mongodb container
```shell
docker ps -a 
docker exec -it [container_id] /bin/bash
# connet mongo
```shell 
mongo mongodb://docker:mongopw@127.0.0.1:27017
> show dbs
admin   0.000GB
config  0.000GB
db      0.000GB
local   0.000GB
test    0.000GB
> use db
switched to db db
> show collections
myCollection
> db.myCollection.insert({"x" :  7})
WriteResult({ "nInserted" : 1 })
> db.myCollection.find()
{ "_id" : ObjectId("62c948b5137dfb20119b768e"), "x" : 1 }
{ "_id" : ObjectId("62c951b45bede65ea6726e7a"), "x" : 3 }
{ "_id" : ObjectId("62c95279f99036948849e06b"), "x" : 7 }
> db.myCollection.find({"x" : 3})
{ "_id" : ObjectId("62c951b45bede65ea6726e7a"), "x" : 3 }
```
