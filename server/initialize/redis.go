package initialize

import(
	"context"
	"github.com/redis/go-redis/v9"
	"github.com/tiandexu-1176/gin-vue-admin-myself/server/global"
)



func initRedisClient(redisCfg config.Redis)(redis.UniversalClient, error) {

	var client redis.UniversalClient
	if redisCfg.useCluster {
		client = redis.NewClusterClient(&redis.ClusterOptions {
			Addrs: redisCfg.ClusterAddrs,
			Password: redisCfg.Password,
		})
	}else {
		client = redis.NewClient(&redis.Options{
			Addr: redisCfg.Addr,
			Password: redisCfg.Password,
			DB: redisCfg.DB,
		})
	}
	pong, err: = client.Ping(context.Background()).Result()
	
	return client, nil
}

funct Redis() {
	redisClient, err := initRedisClient(global.GVA_CONFIG.Server)
	if nil != err {
		
	}
	global.GVA_REDIS = redisClient
}