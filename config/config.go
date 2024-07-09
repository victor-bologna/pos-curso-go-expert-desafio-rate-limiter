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

func LoadConfig(path *string) {
	pwd, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	defaultPath := ".env"
	if path == nil {
		path = &defaultPath
	}
	println(pwd, *path)
	err = godotenv.Load(filepath.Join(pwd, *path))
	if err != nil {
		log.Fatal("Error loading .env file", err)
	}

	redisDB, _ := strconv.Atoi(os.Getenv("REDIS_DB"))
	ipMaxReq, _ := strconv.Atoi(os.Getenv("MAX_IP_REQ"))
	tokenMaxReq, _ := strconv.Atoi(os.Getenv("MAX_TOKEN_REQ"))
	timeout, _ := strconv.Atoi(os.Getenv("TIMEOUT"))

	AppConfig = Config{
		RedisAddr:       os.Getenv("REDIS_ADDR"),
		RedisPass:       os.Getenv("REDIS_PASS"),
		RedisDB:         redisDB,
		IPMaximumReq:    ipMaxReq,
		TokenMaximumReq: tokenMaxReq,
		Timeout:         timeout,
	}
}
