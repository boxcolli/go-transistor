## Run Redis
- ```$ docker run -d -p 6379:6379 --name myredis redis redis-server --notify-keyspace-events KEA```
- ```$ docker run -v ./redis.conf:/usr/local/etc/redis/redis.conf -d -p 6379:6379 --name myredis redis redis-server /usr/local/etc/redis/redis.conf```

## Redis CLI
- ```$ docker exec -it myredis redis-cli```
- 