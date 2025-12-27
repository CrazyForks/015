package utils

import (
	"io"
	"strings"
	"sync"

	"github.com/spf13/viper"
)

var (
	v       *viper.Viper
	envOnce sync.Once
)

func InitEnv(props EnvOption) {
	if v != nil {
		return
	}
	envOnce.Do(func() {
		v = viper.New()
		if props.ConfigData != nil {
			v.ReadConfig(props.ConfigData)
			return
		}
		for _, name := range props.ConfigName {
			v.SetConfigName(name)
		}
		for _, viperConfigType := range props.ConfigType {
			v.SetConfigType(viperConfigType)
		}
		v.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
		for _, path := range props.ConfigPath {
			v.AddConfigPath(path)
		}
		v.AutomaticEnv()
		v.WatchConfig()
		err := v.ReadInConfig()
		if err != nil {
			panic(err)
		}
	})
}

type Option interface {
	applyTo(*EnvOption)
}

type EnvOption struct {
	DefaultValue string
	ConfigPath   []string
	ConfigName   []string
	ConfigType   []string
	ConfigData   io.Reader // 测试环境使用
}

type WithDefaultValue string

func (o WithDefaultValue) applyTo(props *EnvOption) {
	props.DefaultValue = string(o)
}

func getEnvOptions(options ...Option) EnvOption {
	props := EnvOption{
		DefaultValue: "",
		ConfigPath:   []string{".", "../"},
		ConfigName:   []string{"config"},
		ConfigType:   []string{"yaml"},
	}
	for _, option := range options {
		option.applyTo(&props)
	}
	return props
}

func GetEnv(key string, options ...Option) string {
	props := getEnvOptions(options...)
	InitEnv(props)
	value := v.GetString(key)

	if value == "" && props.DefaultValue != "" {
		return props.DefaultValue
	}
	return value
}

func GetEnvWithDefault(key string, defaultValue string) string {
	return GetEnv(key, WithDefaultValue(defaultValue))
}

func GetEnvMapString(key string) map[string]string {
	props := getEnvOptions()
	InitEnv(props)
	return v.GetStringMapString(key)
}

func SetEnv(key string, value string) {
	v.Set(key, value)
}
