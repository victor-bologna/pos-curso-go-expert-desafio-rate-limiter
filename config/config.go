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
	IPTimeout       int
	TokenMaximumReq int
	TokenTimeout    int
}

var AppConfig Config

func LoadConfig(path string) {
	pwd, err := os.Getwd()
	if err != nil {
		log.Printf("Error getting current directory: %v", err)
	}
	rootPath := filepath.Join(pwd, path)
	err = godotenv.Load(filepath.Join(rootPath, ".env"))
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
	ipTimeout, err := strconv.Atoi(os.Getenv("IP_TIMEOUT"))
	if err != nil {
		ipTimeout = 10
		log.Printf("Error converting IP_TIMEOUT to integer: %v, converting to default value: %v", err, ipTimeout)
	}
	tokenTimeout, err := strconv.Atoi(os.Getenv("TOKEN_TIMEOUT"))
	if err != nil {
		tokenTimeout = 10
		log.Printf("Error converting TOKEN_TIMEOUT to integer: %v, converting to default value: %v", err, tokenTimeout)
	}

	AppConfig = Config{
		RedisAddr:       redisAddr,
		RedisPass:       os.Getenv("REDIS_PASS"),
		RedisDB:         redisDB,
		IPMaximumReq:    ipMaxReq,
		IPTimeout:       ipTimeout,
		TokenMaximumReq: tokenMaxReq,
		TokenTimeout:    tokenTimeout,
	}
}
