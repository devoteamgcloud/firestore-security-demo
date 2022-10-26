package settings

import (
	"time"

	"github.com/kaan-devoteam/one-click-deploy-demo/log"
)

var (
	Version = "dev"
)

const (
	ServerPort        = 8080
	DatabaseProjectID = "kaan-sandbox"
	Name              = "Gossip-Backend"
	ServerTimeout     = 60 * time.Second

	ServerAccessLog = true
	ServerLogLevel  = log.LevelDebug
)
