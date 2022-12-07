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
	DatabaseProjectID = "project-id"
	ApiKey            = "key=AIzaSyC2gQdpHk-rSNgRfvMtNIUccJQ8dy5kMGs"
	Name              = "Articles-Backend"
	ServerTimeout     = 60 * time.Second

	ServerAccessLog = true
	ServerLogLevel  = log.LevelDebug

	FirestoreRestUrl = "https://firestore.googleapis.com/v1/"
)
