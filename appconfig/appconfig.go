package appconfig

import (
    "fmt"
    "sync"	
    "github.com/gookit/config/v2"
    "github.com/gookit/config/v2/yaml"
)

var lock = &sync.Mutex{}

type AppConfigType struct {
	dev bool
	backendBaseUrl string
	hugoProject string
}

func assertConfigItemExistence(key string) {
	if config.String(key) == "" {
		panic("Configuration error: '"+key+"' is empty")
	}
}

func NewAppConfig() AppConfigType {
	config.WithOptions(config.ParseEnv)

	// load flag info
	flagkeys1 := []string{"config","dev"}
	err := config.LoadFlags(flagkeys1)
	if err != nil {
		panic(err)
	}

	config.AddDriver(yaml.Driver)

	// load config file
	err = config.LoadFiles(config.String("config"))
	if err != nil {
		panic(err)
	}

	dev := config.Bool("dev")

	assertConfigItemExistence("backendBaseUrl")
	backendBaseUrl := config.String("backendBaseUrl")
	fmt.Println(backendBaseUrl)
	
	assertConfigItemExistence("hugoProject")
	hugoProject := config.String("hugoProject")
	fmt.Println(hugoProject)
	
	return AppConfigType {
		dev,
		backendBaseUrl,
		hugoProject,
	}
}

func (appconfig AppConfigType) BackendBaseUrl() string {
	return appconfig.backendBaseUrl
}

func (appconfig AppConfigType) HugoProject() string {
	return appconfig.hugoProject
}

// AppConfig singleton

var instance *AppConfigType

func AppConfig() *AppConfigType {
    if instance == nil {
        lock.Lock()
        defer lock.Unlock()
        if instance == nil {
            conf := NewAppConfig()
			instance = &conf
        }
	}
    return instance
}