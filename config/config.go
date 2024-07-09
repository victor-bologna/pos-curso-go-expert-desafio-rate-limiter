package config

import (
	"log"
	"os"
	"path/filepath"
	"strconv"

	"github.com/joho/godotenv"
)

type Config struct {
	RedisAddr       string
	RedisPass       string
	RedisDB         int
	IPMaximumReq    int
	IPTTL           int
	TokenMaximumReq int
	TokenTTL        int
	Timeout         int
}

var AppConfig Config

func LoadConfig() {
	pwd, err := os.Getwd()
	if err != nil {
		log.Print("Error getting current directory: %v", err)
	}
	err = godotenv.Load(filepath.Join(pwd, ".env"))
	if err != nil {
		log.Printf("Error loading .env file: %v, using default values...", err)
	}

	redisAddr := os.Getenv("REDIS_ADDR")
	if redisAddr == "" {
		log.Printf("Redis Addr is empty, converting to default value: redis:6379")
		redisAddr = "redis:6379"
	}
	redisDB, err := strconv.Atoi(os.Getenv("REDIS_DB"))
	if err != nil {
		log.Printf("Error converting REDIS_DB to integer: %v, converting to default value: 0", err)
		redisDB = 0
	}
	ipMaxReq, err := strconv.Atoi(os.Getenv("IP_MAX_REQ"))
	if err != nil {
		log.Printf("Error converting IP_MAX_REQ to integer: %v, converting to default value: 5", err)
		ipMaxReq = 5
	}
	tokenMaxReq, err := strconv.Atoi(os.Getenv("TOKEN_MAX_REQ"))
	if err != nil {
		log.Printf("Error converting TOKEN_MAX_REQ to integer: %v, converting to default value: 5", err)
		tokenMaxReq = 5
	}
	timeout, err := strconv.Atoi(os.Getenv("TIMEOUT"))
	if err != nil {
		log.Printf("Error converting TIMEOUT to integer: %v, converting to default value: 10", err)
		timeout = 10
	}

	AppConfig = Config{
		RedisAddr:       redisAddr,
		RedisPass:       os.Getenv("REDIS_PASS"),
		RedisDB:         redisDB,
		IPMaximumReq:    ipMaxReq,
		TokenMaximumReq: tokenMaxReq,
		Timeout:         timeout,
	}
}
