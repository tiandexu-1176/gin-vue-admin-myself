package global



import (
	"github.com/tiandexu-1176/gin-vue-admin-myself/server/config"
	"github.com/redis/go-redis/v9"
)

var (
	GVA_CONFIG 	config.Server
	GVA_REDIS    redis.UniversalClient
)