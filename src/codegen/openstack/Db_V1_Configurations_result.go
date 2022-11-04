package openstack

// Code generated by cloud manager.





import (
    "github.com/gophercloud/gophercloud/openstack/db/v1/configurations"
)


//extract response info from pager for ListDbV1Configurations
func ExtractListDbV1ConfigurationsResponse(response *ListDbV1ConfigurationsResponse)([]configurations.Config,error){
	page, err := response.Pager.AllPages()
	if err != nil {
		return nil, err
	}
	return configurations.ExtractConfigs(page)
}


//extract response info from pager for ListDatastoreParamsDbV1Configurations
func ExtractListDatastoreParamsDbV1ConfigurationsResponse(response *ListDatastoreParamsDbV1ConfigurationsResponse)([]configurations.Param,error){
	page, err := response.Pager.AllPages()
	if err != nil {
		return nil, err
	}
	return configurations.ExtractParams(page)
}


//extract response info from pager for ListGlobalParamsDbV1Configurations
func ExtractListGlobalParamsDbV1ConfigurationsResponse(response *ListGlobalParamsDbV1ConfigurationsResponse)([]configurations.Param,error){
	page, err := response.Pager.AllPages()
	if err != nil {
		return nil, err
	}
	return configurations.ExtractParams(page)
}