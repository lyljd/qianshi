package svc

import (
	"github.com/go-redis/redis/v8"
	"os"
	"qianshi/service/vcode/api/internal/config"
	"qianshi/service/vcode/api/internal/svc/email"
	"strconv"
)

type ServiceContext struct {
	Config       config.Config
	RedisCli     *redis.Client
	EmailCli     *email.Client
	EmailTitle   string
	EmailContent string
	VcodeTTL     int
	VcodeCD      int
}

func NewServiceContext(c config.Config) *ServiceContext {
	ttl, _ := strconv.Atoi(os.Getenv("VCODE_TTL"))
	cd, _ := strconv.Atoi(os.Getenv("VCODE_CD"))

	return &ServiceContext{
		Config: c,
		RedisCli: redis.NewClient(&redis.Options{
			Addr:     os.Getenv("REDIS_ADDR"),
			Username: os.Getenv("REDIS_USERNAME"),
			Password: os.Getenv("REDIS_PASSWORD"),
		}),
		EmailCli: email.NewClient(&email.Options{
			Hostname: os.Getenv("EMAIL_HOSTNAME"),
			Username: os.Getenv("EMAIL_USERNAME"),
			Password: os.Getenv("EMAIL_PASSWORD"),
		}),
		EmailTitle:   os.Getenv("EMAIL_TITLE"),
		EmailContent: os.Getenv("EMAIL_CONTENT"),
		VcodeTTL:     ttl,
		VcodeCD:      cd,
	}
}
