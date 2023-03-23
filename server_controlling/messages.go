package server_controlling

import (
	"context"
	"github.com/redis/go-redis/v9"
	"github.com/senyehor/go_server/utils"
	log "github.com/sirupsen/logrus"
	"time"
)

func setCurrentStatus(status string, r redis.Client) {
	result := r.Set(context.Background(), utils.ServerControllingConfig.CurrentStatusKey(), status, 0*time.Second)
	if result.Err() != nil {
		log.Error(result.Err())
		panic("failed to set current status")
	}
}
