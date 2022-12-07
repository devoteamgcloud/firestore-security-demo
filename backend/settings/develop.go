package settings

import (
	"os"
	"time"

	"github.com/kaan-devoteam/firestore-security-demo/log"
)

var (
	Version = "dev"
	ApiKey  = os.Getenv("ApiKey")
)

const (
	ServerPort        = 8080
	DatabaseProjectID = "project-id"
	Name              = "Articles-Backend"
	ServerTimeout     = 60 * time.Second

	ServerAccessLog = true
	ServerLogLevel  = log.LevelDebug

	FirestoreRestUrl = "https://firestore.googleapis.com/v1/"
)
