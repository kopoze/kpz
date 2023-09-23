package config

import (
	"errors"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"reflect"
	"runtime"
	"strings"

	"github.com/spf13/viper"
)

type SpecificConfig interface{}

type ConfigOptions struct {
	Kopoze   KopozeConfig `mapstructure:"kopoze"`
	Docker   DockerConfig `mapstructure:"docker"`
	Git      GitConfig    `mapstructure:"git"`
	Database DBConfig     `mapsctructure:"database"`
}

type Config struct {
	ConfigOptions //`mapstructure:",squash"`
}

// Create new Config.
func NewConfig() Config {
	return Config{
		ConfigOptions{
			Kopoze:   NewKopozeConfig(KopozeOptions{}),
			Docker:   NewDockerConfig(DockerOptions{}),
			Git:      NewGitConfig(GitOptions{}),
			Database: NewDBConfig(DBOptions{}),
		},
	}
}

func Configure() {
	log.Println("Initializing config")
	viper.SetConfigName(FILE_CONFIG)
	viper.SetConfigType(FILE_TYPE)
	viper.AddConfigPath(getConfigPath())

	var conf = NewConfig()

	SetConfig("kopoze", conf.Kopoze.KopozeOptions)
	SetConfig("docker", conf.Docker.DockerOptions)
	SetConfig("git", conf.Git.GitOptions)
	SetConfig("database", conf.Database.DBOptions)

	if err := viper.SafeWriteConfig(); err != nil {
		log.Println(err)
	}
}

// Dynamically set config from struct to viper.
func SetConfig(g string, o SpecificConfig) {
	v := reflect.ValueOf(o)
	typeOfs := v.Type()
	for i := 0; i < v.NumField(); i++ {
		viper.Set(fmt.Sprintf("%s.%s", strings.ToLower(g), typeOfs.Field(i).Name), v.Field(i).Interface())
	}
}

func LoadConfig() Config {
	log.Println("Reading existing config")
	viper.SetConfigName(FILE_CONFIG)
	viper.SetConfigType(FILE_TYPE)
	viper.AddConfigPath(getConfigPath())
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("fatal error config file: %w", err))
	}
	config := ConfigOptions{}
	viper.UnmarshalKey("database", &config.Database.DBOptions)
	viper.UnmarshalKey("docker", &config.Docker.DockerOptions)
	viper.UnmarshalKey("git", &config.Git.GitOptions)
	viper.UnmarshalKey("kopoze", &config.Kopoze.KopozeOptions)
	return Config{
		config,
	}
}

func getConfigPath() string {
	configPath := filepath.Join(getUserHomeDir(), ".kopoze")
	if err := os.MkdirAll(configPath, os.ModePerm); err != nil {
		panic(err)
	}
	return configPath
}

func getUserHomeDir() string {
	// Source: https://stackoverflow.com/a/7922977/5527968
	if runtime.GOOS == "windows" {
		home := os.Getenv("HOMEDRIVE") + os.Getenv("HOMEPATH")
		if home == "" {
			home = os.Getenv("USERPROFILE")
		}
		return home
	}
	return os.Getenv("HOME")
}

// Update `dst` struct value to match `src` struct.
func updateOpts(src, dst interface{}) (interface{}, error) {
	srcValue := reflect.ValueOf(src)
	dstValue := reflect.ValueOf(dst)

	if srcValue.Kind() != reflect.Struct {
		return nil, errors.New("src must be a struct")
	}

	if dstValue.Kind() != reflect.Struct {
		return nil, errors.New("dst must be a struct")
	}

	modifiedDst := reflect.New(dstValue.Type()).Elem()

	// Copy the original dst struct into modifiedDst
	modifiedDst.Set(dstValue)

	for i := 0; i < srcValue.NumField(); i++ {
		srcField := srcValue.Field(i)
		dstField := modifiedDst.FieldByName(srcValue.Type().Field(i).Name)

		if dstField.IsValid() && dstField.CanSet() {
			if srcField.IsValid() {
				srcFieldValue := srcField.Interface()
				if !isEmpty(srcFieldValue) {
					dstField.Set(srcField)
				}
			}
		}
	}

	return modifiedDst.Interface(), nil
}

// Function to check if a value is empty (nil or the zero value of its type)
func isEmpty(value interface{}) bool {
	if value == nil {
		return true
	}

	v := reflect.ValueOf(value)
	if v.Kind() == reflect.Ptr && v.IsNil() {
		return true
	}

	if v.Kind() == reflect.String {
		return v.Len() == 0
	}

	return v.IsZero()
}
