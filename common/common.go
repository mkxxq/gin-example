package common

import (
	"os"
	"path/filepath"

	"github.com/joho/godotenv"

	"ginexample/common/config"
	"ginexample/common/mysql"
	"ginexample/utils"
	"ginexample/utils/log"
)

func Setup() {
	config.Setup()
	if err := mysql.Setup(); err != nil {
		panic(err)
	}

}

func LoadEnv(envFile string) error {
	err := godotenv.Load(envFile)
	if err != nil {
		notEnvFile := os.Getenv("NOT_ENV_FILE")
		if notEnvFile != "true" {
			return err
		}
	} else {
		log.Warnf("load env from %s success\n", envFile)
	}
	return nil
}
func LoadNormalEnv() error {
	rootPath, err := utils.GetRootPath()
	if err != nil {
		return err
	}
	normalEnvFile := filepath.Join(rootPath, ".env")
	return LoadEnv(normalEnvFile)
}
func LoadTestEnv() error {
	rootPath, err := utils.GetRootPath()
	if err != nil {
		return err
	}
	testEnvFile := filepath.Join(rootPath, "test.env")
	return LoadEnv(testEnvFile)
}
