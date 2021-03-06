package settings

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

var environments = map[string]string{
	"production":    "/Users/artemchuzhmarov/Desktop/IdeaWorkspace/chat/src/settings/prod.json",
	"preproduction": "/app/settings/pre.json",
	"mac":           "/Users/artemchuzhmarov/Desktop/IdeaWorkspace/chat/src/settings/mac.json",
	"windows":       "E:\\work\\gopath\\src\\gochat\\settings\\windows.json",
}

type Settings struct {
	PrivateKeyPath     string
	PublicKeyPath      string
	JWTExpirationDelta int
}

var settings *Settings = nil
var env = "preproduction"

func Init() {
	env = os.Getenv("GO_ENV")
	if env == "" {
		fmt.Println("Warning: Setting preproduction environment due to lack of GO_ENV value")
		env = "preproduction"
	}
	LoadSettingsByEnv(env)
}

func LoadSettingsByEnv(env string) {
	content, err := ioutil.ReadFile(environments[env])
	if err != nil {
		fmt.Println("Error while reading config file", err)
	}
	settings = &Settings{}

	jsonErr := json.Unmarshal(content, settings)
	if jsonErr != nil {
		fmt.Println("Error while parsing config file", jsonErr)
	}
}

func GetEnvironment() string {
	return env
}

func Get() *Settings {
	if settings == nil {
		Init()
	}
	return settings
}

func IsTestEnvironment() bool {
	return env == "tests"
}
