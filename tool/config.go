package tool

import (
	"os"
	"strconv"

	"github.com/joho/godotenv"
	"go.uber.org/zap"
)

// EnvConfig Config read from environment variables
var EnvConfig = new(Config)

// Config Gshop Config struct
type Config struct {
	// Gshop
	Gshop struct {
		HOST  string
		Debug int
	}
	// DB
	DB struct {
		Address string
	}
	// Redis
	Redis struct {
		Address string
	}
	// JWT
	JWT struct {
		Key        string
		Expiration int
	}
}

// DoEnv Read the configuration from the specified file to the environment variable, the default is .env under gshop
func DoEnv(fileNames ...string) {
	err := godotenv.Load(fileNames...)
	if err != nil {
		if fileNames == nil {
			Log.Panic("No configuration file is found, check if there is a .env file under the gshop project and readable", zap.Error(err))
		} else {
			Log.Panic("No configuration file found, check if the file path is correct and readable", zap.Error(err))
		}
	}
}

// InitConfig Read environment variables to Config structure
func InitConfig() {
	EnvConfig.Gshop.HOST = func() string {
		h := os.Getenv("GSHOP_HOST")
		if h != "" {
			return h
		}
		Log.Warn("GSHOP_HOST is not read, replace it with the default localhost:8848")
		return "localhost:8848"
	}()

	EnvConfig.Gshop.Debug = func() int {
		ed := os.Getenv("GSHOP_DEBUG")
		if ed != "" {
			d, err := strconv.Atoi(ed)
			if err != nil {
				Log.Warn("GSHOP_DEBUG cannot be parsed, replaced with the default 1")
				return 1
			}
			return d
		}
		Log.Warn("GSHOP_DEBUG is not read, replaced with the default 1")
		return 1
	}()

	EnvConfig.DB.Address = func() string {
		a := os.Getenv("GSHOP_DB_ADDRESS")
		if a != "" {
			return a
		}
		Log.Panic("GSHOP_DB_ADDRESS is not read!")
		return ""
	}()

	EnvConfig.Redis.Address = func() string {
		a := os.Getenv("GSHOP_REDIS_ADDRESS")
		if a != "" {
			return a
		}
		Log.Panic("GSHOP_REDIS_ADDRESS is not read!")
		return ""
	}()

	EnvConfig.JWT.Key = func() string {
		k := os.Getenv("JWT_KEY")
		if k != "" {
			return k
		}
		Log.Panic("JWT_KEY is not read!")
		return ""
	}()

	EnvConfig.JWT.Expiration = func() int {
		e := os.Getenv("JWT_Expiration")
		if e != "" {
			d, err := strconv.Atoi(e)
			if err != nil {
				Log.Warn("JWT_Expiration cannot be parsed, replaced with the default 604800")
				return 604800
			}
			return d
		}
		Log.Warn("JWT_Expiration cannot be parsed, replaced with the default 604800")
		return 604800
	}()

}
