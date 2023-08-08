# 集群搭建

[Reference](https://juejin.cn/post/6974325733812355085)

tar -czh . | docker build --progress=plain -t cache-server -

docker network create -d bridge net --subnet=172.18.0.0/24 --gateway=172.18.0.1

docker run -it --net net --ip 172.18.0.11 -p 22345:12345 -p 22346:12346 cache-server -type rocksdb -node 172.18.0.11
docker run -it --net net --ip 172.18.0.12 -p 32345:12345 -p 32346:12346 cache-server -type rocksdb -node 172.18.0.12 -cluster 172.18.0.11  