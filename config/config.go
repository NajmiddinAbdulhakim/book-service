package config

import (
	"os"

	"github.com/spf13/cast"
)

type Config struct {
	PostgresHost     string
	PostgresPort     int
	PostgresDB       string
	PostgresUser     string
	PostgresPassword string
	RPCPort         string
}

func Load() Config {
	c := Config{}


	c.PostgresHost = cast.ToString(look("POSTGRES_HOST", "localhost"))
	c.PostgresPort = cast.ToInt(look("POSTGRES_HOST", 5432))
	c.PostgresDB = cast.ToString(look("POSTGRES_DATABASE", "mybook"))
	c.PostgresUser = cast.ToString(look("POSTGRES_USER", "najmiddin"))
	c.PostgresPassword = cast.ToString(look("POSTGRES_PASSWORD", "1234"))

	c.RPCPort = cast.ToString(look(`RPC_PORT`,))
	
	return c
}

func look(key string, defVal interface{}) interface{} {
	_, exists := os.LookupEnv(key)
	if exists {
		return os.Getenv(key)
	}
	return defVal
}
