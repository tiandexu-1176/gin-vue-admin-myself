package global



import (
	"server/config"
	"github.com/redis/go-redis/v9"
)

var (
	GVA_CONFIG 	config.Server
	GVA_REDIS    redis.UniversalClient
)