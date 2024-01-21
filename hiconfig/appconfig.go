package hiconfig

import (
	"fmt"
	"sync"

	"github.com/gookit/config/v2"
	"github.com/gookit/config/v2/yaml"
)

var lock = &sync.Mutex{}

type AppConfigType struct {
	boolconfigs   map[string]bool
	stringconfigs map[string]string
}

func assertConfigItemExistence(key string) {
	value := config.String(key)
	if value == "" {
		panic("Configuration error: '" + key + "' is empty")
	}
}

func NewAppConfig() AppConfigType {
	config.WithOptions(config.ParseEnv)

	stringmap := make(map[string]string)
	boolmap := make(map[string]bool)

	// load flag info
	flagkeys1 := []string{"config", "dev"}
	err := config.LoadFlags(flagkeys1)
	if err != nil {
		panic(err)
	}

	config.AddDriver(yaml.Driver)

	// TODO: check file exists
	configfile := config.String("config")
	// load config file
	err = config.LoadFiles(configfile)
	if err != nil {
		panic(err)
	}

	assertConfigItemExistence("hugoproject")

	appconfig := AppConfigType{boolmap, stringmap}
	appconfig.setBool("dev", config.Bool("dev"))
	appconfig.setString("hugoproject", config.String("hugoproject"), "/")
	appconfig.setString("backendbaseurl", config.String("backendnaseurl"), "http://localhost:1313/")
	appconfig.setString("serverbinding", config.String("serverbinding"), "127.0.0.1:3000")
	appconfig.setString("gitcommand", config.String("gitcommand"), "git")

	if appconfig.getBool("dev") {
		fmt.Println("Development modus enabled (dev)")
	}
	fmt.Println("Working with (hugoproject): " + appconfig.HugoProject())
	fmt.Println("Listening to (backendbaseUrl): " + appconfig.BackendBaseUrl())
	fmt.Println("Service connects to (serverbinding): " + appconfig.ServerBinding())

	return appconfig
}

func (appconfig AppConfigType) setBool(key string, value bool) {
	appconfig.boolconfigs[key] = value
}

func (appconfig AppConfigType) getBool(key string) bool {
	return appconfig.boolconfigs[key]
}

func (appconfig AppConfigType) setString(key string, value string, defaultvalue string) {
	if value == "" {
		value = defaultvalue
	}
	appconfig.stringconfigs[key] = value
}

func (appconfig AppConfigType) getString(key string) string {
	return appconfig.stringconfigs[key]
}

func ensureLastSlash(text string) string {
	lastchar := text[len(text)-1:]
	if lastchar != "/" {
		return text + "/"
	}
	return text
}

func (appconfig AppConfigType) BackendBaseUrl() string {
	return ensureLastSlash(appconfig.getString("backendbaseurl"))
}

func (appconfig AppConfigType) HugoProject() string {
	return ensureLastSlash(appconfig.getString("hugoproject"))
}

func (appconfig AppConfigType) ServerBinding() string {
	return appconfig.getString("serverbinding")
}

func (appconfig AppConfigType) GitCommand() string {
	return appconfig.getString("gitcommand")
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
