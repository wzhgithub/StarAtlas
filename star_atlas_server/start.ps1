cd  E:\yiyun\StarAtlas-main
cd client
#编译遥控和mock数据客户端
docker run  -v ${PWD}:/app -w /app --platform linux/amd64 buildessential/debian:latest bash build.sh
docker run  -v ${PWD}:/app -w /app --platform linux/amd64 buildessential/debian:latest bash build_trans.sh
cp -Force bin/utest ../star_atlas_server
cp -Force bin/trans ../star_atlas_server
cp -Force tpinit.sh ../star_atlas_server
cp -Recurse -Force bin/conf ../star_atlas_server

#初始化go编译环境
cd ../star_atlas_server
cp -Force go.mod.bk go.mod


#编译服务端代码
docker build --platform=linux/amd64 -t star_atlas_server .

docker rm -f $(docker ps -a -q)
docker compose up -d