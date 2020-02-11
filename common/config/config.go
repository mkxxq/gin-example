package config

import (
	"fmt"
	"os"
	"reflect"
	"strconv"
	"strings"

	"ginexample/utils/log"
)

type GlobalConfig struct {
	CloudSql *CloudSqlConfig
	Port     string
}
type CloudSqlConfig struct {
	ConnectionAddr string
	User           string
	Password       string
	Database       string
	// Maximum number of connections.
	MaxConn int
	// Maximum timeout.
	MaxIdle int
	// Connection maximum survival time.
	MaxLifeTime int64
	// Database source.
	Source string
}

// Conf is global config
var Conf *GlobalConfig

// Setup method
func Setup() {
	Conf = new(GlobalConfig)
	Conf.CloudSql = initCloudSQLConfig()
	Conf.Port = mayGetEnv("PORT")
}
func initCloudSQLConfig() *CloudSqlConfig {
	connectionAddr := mustGetEnv("SQL_CONNECTION_ADDR")
	user := mustGetEnv("SQL_USER")
	database := mustGetEnv("SQL_DATABASE")
	// NOTE: password may be empty
	password := os.Getenv("SQL_PASSWORD")
	maxCoon, err := strconv.Atoi(mustGetEnv("SQL_MAX_CONN"))
	if err != nil {
		log.Fatalln(err.Error())
	}
	maxIdle, err := strconv.Atoi(mustGetEnv("SQL_MAX_IDLE"))
	if err != nil {
		log.Fatalln(err.Error())
	}
	maxLifeTime, err := strconv.ParseInt(mustGetEnv("SQL_MAX_LIFE_TIME"), 10, 64)
	if err != nil {
		log.Fatalln(err.Error())
	}
	if err != nil {
		log.Fatalln(err.Error())
	}
	var format string
	var source string
	// connect mysql by tcp
	format = "%s:%s@tcp(%s)/%s?charset=utf8&parseTime=true&loc=UTC"
	source = fmt.Sprintf(format, user, password, connectionAddr, database)
	// connect mysql by unix socket
	// format = "%s:%s@unix(%s)/%s?charset=utf8&parseTime=true&loc=UTC"
	// source = fmt.Sprintf(format, user, password, connectionName, database)
	return &CloudSqlConfig{
		ConnectionAddr: connectionAddr,
		User:           user,
		Password:       password,
		Database:       database,
		MaxConn:        maxCoon,
		MaxIdle:        maxIdle,
		MaxLifeTime:    maxLifeTime,
		Source:         source,
	}
}
func mustGetEnv(k string) string {
	v := os.Getenv(k)
	if isEmpty(v) {
		log.Fatalf("%s environment variable not set.", k)
	}
	return v
}
func mayGetEnv(k string) string {
	return os.Getenv(k)
}

func isEmpty(v interface{}) bool {
	if v == nil {
		return true
	}
	return isValueEmpty(reflect.ValueOf(v))
}
func isValueEmpty(value reflect.Value) bool {
	switch value.Kind() {
	case reflect.String:
		return len(strings.TrimSpace(value.String())) == 0
	case reflect.Bool:
		return !value.Bool()
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return value.Int() == 0
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		return value.Uint() == 0
	case reflect.Float32, reflect.Float64:
		return value.Float() == 0
	case reflect.Complex64, reflect.Complex128:
		return value.Complex() == 0
	case reflect.Interface, reflect.Ptr:
		if value.IsNil() {
			return true
		}
		return isEmpty(value.Elem())
	case reflect.Array, reflect.Slice, reflect.Chan:
		return value.Len() == 0
	case reflect.Map:
		return len(value.MapKeys()) == 0
	case reflect.Func:
		return value.IsNil()
	default:
		return reflect.DeepEqual(value.Interface(), reflect.Zero(value.Type()).Interface())
	}
}
