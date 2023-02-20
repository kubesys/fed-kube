package gen

import (
	"encoding/json"
	"log"
	"os"
)

type CloudConfig struct {
	CloudType       string            `json:"CloudType"`
	RegistryConfigs []*RegistryConfig `json:"RegistryConfigs"`
	APICodeConfigs  []*APICodeConfig  `json:"APICodeConfigs"`
	JavaCodeConfig  *JavaCodeConfig   `json:"JavaCodeConfig"`
}

type RegistryConfig struct {
	CodeGenConfig       CodeGenConfig `json:"CodeGenConfig"`
	RegistryImportPaths []string      `json:"RegistryImportPaths"`
	CreateFuncPre       string        `json:"CreateFuncPre"`
	FuncAction          string        `json:"FuncAction"`
	CodeType            string        `json:"CodeType"` //Request or Response
}

type APICodeConfig struct {
	CodeGenConfig  CodeGenConfig `json:"CodeGenConfig"`
	SourceCodePath string        `json:"SourceCodePath"`
	CodeType       string        `json:"CodeType"` //Request or Result
}
type CodeGenConfig struct {
	TemplatePath string `json:"TemplatePath"`
	CodePath     string `json:"CodePath"`
}

type JavaCodeConfig struct {
	Class     CodeGenConfig `json:"Class"`
	Domain    CodeGenConfig `json:"Domain"`
	Impl      CodeGenConfig `json:"Impl"`
	Lifecycle CodeGenConfig `json:"Lifecycle"`
	Spec      CodeGenConfig `json:"Spec"`
}

func LoadCloudConfig(configPath string) *CloudConfig {
	config := &CloudConfig{}
	bytes, err := os.ReadFile(configPath)
	if err != nil {
		log.Fatal("read config error, err=", err)
	}
	err = json.Unmarshal(bytes, config)
	if err != nil {
		log.Fatal("json unmarshal error, ", err)
	}
	return config
}
