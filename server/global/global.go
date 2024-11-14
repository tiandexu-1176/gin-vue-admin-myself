package global



import (
	"server/config"
	"github.com/tiandexu-1176/gin-vue-admin-myself/github.com/redis/go-redis/v9"
)

var (
	GVA_CONFIG 	config.Server
	GVA_REDIS    redis.UniversalClient
)