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
	TokenMaximumReq int
	Timeout         int
}

var AppConfig Config

func LoadConfig() {
	pwd, err := os.Getwd()
	if err != nil {
		log.Printf("Error getting current directory: %v", err)
	}
	err = godotenv.Load(filepath.Join(pwd, ".env"))
	if err != nil {
		log.Printf("Error loading .env file: %v, using default values...", err)
	}

	redisAddr := os.Getenv("REDIS_ADDR")
	if redisAddr == "" {
		redisAddr = "localhost:6379"
		log.Printf("Redis Addr is empty, converting to default value: %v", redisAddr)
	}
	redisDB, err := strconv.Atoi(os.Getenv("REDIS_DB"))
	if err != nil {
		redisDB = 0
		log.Printf("Error converting REDIS_DB to integer: %v, converting to default value: %v", err, redisDB)
	}
	ipMaxReq, err := strconv.Atoi(os.Getenv("IP_MAX_REQ"))
	if err != nil {
		ipMaxReq = 5
		log.Printf("Error converting IP_MAX_REQ to integer: %v, converting to default value: %v", err, ipMaxReq)
	}
	tokenMaxReq, err := strconv.Atoi(os.Getenv("TOKEN_MAX_REQ"))
	if err != nil {
		tokenMaxReq = 5
		log.Printf("Error converting TOKEN_MAX_REQ to integer: %v, converting to default value: %v", err, tokenMaxReq)
	}
	timeout, err := strconv.Atoi(os.Getenv("TIMEOUT"))
	if err != nil {
		timeout = 10
		log.Printf("Error converting TIMEOUT to integer: %v, converting to default value: %v", err, timeout)
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
