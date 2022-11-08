package gen

import (
	cloud_manager "cloud_manager/src/analyzer"
	openstack "cloud_manager/src/codegen/openstack"
	"cloud_manager/src/utils"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/ecs"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func GenCloudCode(config *CloudConfig) error {
	analyzer := cloud_manager.NewCloudAPIAnalyzer()
	switch config.Kind {
	case "Aliyun":
		client := ecs.Client{}
		analyzer.ExtractCloudAPIs(client)
		for _, registryConfig := range config.RegistryConfigs {
			GenRegistryCode(config.Kind, analyzer, registryConfig)
		}
	case "Openstack":
		for _, apiCodeConfig := range config.APICodeConfigs {
			GenAPICode(apiCodeConfig)
		}
		client := openstack.OpenstackClient{}
		analyzer.ExtractCloudAPIs(client)
		for _, registryConfig := range config.RegistryConfigs {
			GenRegistryCode(config.Kind, analyzer, registryConfig)
		}
	default:
		log.Fatalln("unsupport cloud kind, ", config.Kind)
	}
	return nil
}

func GenRegistryCode(cloudKind string, analyzer *cloud_manager.CloudAPIAnalyzer, registryConfig *RegistryConfig) {
	requestRegistryInfo := analyzer.ExtractRequestInfos(registryConfig.CreateFuncPre)
	if cloudKind == "Openstack" && registryConfig.CodeType == "Result" {
		for _, requestInfo := range requestRegistryInfo.RequestInfos {
			requestInfo.CreateFunctionName = strings.Replace(requestInfo.CreateFunctionName, "New", "Extract", -1)
			requestInfo.CreateFunctionName = requestInfo.CreateFunctionName[0:len(requestInfo.CreateFunctionName)-len("Request")] + "Response"
		}
	}
	requestRegistryInfo.ImportPaths = registryConfig.RegistryImportPaths
	data := utils.Struct2Map(requestRegistryInfo)
	params := make(map[string]interface{})
	params["packageName"] = "registry"
	params["kind"] = cloudKind
	params["action"] = registryConfig.FuncAction
	params["type"] = registryConfig.CodeType
	codePath := filepath.Join(registryConfig.CodeGenConfig.CodePath, strings.ToLower(cloudKind)+"_"+strings.ToLower(registryConfig.FuncAction)+"_"+strings.ToLower(registryConfig.CodeType)+"_registry.go")
	GenAndSaveCode(registryConfig.CodeGenConfig.TemplatePath, codePath, data, params)
}

func GenAPICode(config *APICodeConfig) {
	ma := cloud_manager.NewModuleAnalyzer()
	resourceInfos, err := ma.DoAnalyze(config.SourceCodePath)
	if err != nil {
		log.Fatalln("analyze source code error, ", err)
	}
	for _, resourceInfo := range resourceInfos {
		if len(resourceInfo.ActionInfos) == 0 {
			continue
		}
		log.Printf("gen %s code for actions in resource %s\n", config.CodeType, resourceInfo.ResourcePackageName)
		data := utils.Struct2Map(resourceInfo)
		params := map[string]interface{}{
			"packageName": "openstack",
		}
		filenName := utils.JoinName(resourceInfo.ResourcePath, "openstack", "_") + "_" + strings.ToLower(config.CodeType) + ".go"
		codePath := filepath.Join(config.CodeGenConfig.CodePath, filenName)
		GenAndSaveCode(config.CodeGenConfig.TemplatePath, codePath, data, params)
	}
}

func GenAndSaveCode(templatePath, codePath string, data, params map[string]interface{}) {
	code, err := GenCode(templatePath, data, params)
	if err != nil {
		log.Fatalln("Gen Registry Code error, ", err)
	}
	file, err := os.Create(codePath)
	if err != nil {
		log.Fatalln("Create Code File for registry error, ", err)
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			log.Fatalln("Close Code File error, ", err)
		}
	}(file)
	_, err = file.Write(code)
	if err != nil {
		log.Fatalln("Write registry code to file error, ", err)
	}
}
