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


// call result's extract function
func ExtractCreateDbV1ConfigurationsResponse(response *CreateDbV1ConfigurationsResponse)(interface{}, error){
    return response.CreateResult.Body, response.CreateResult.Err
}



// call result's extract function
func ExtractGetDbV1ConfigurationsResponse(response *GetDbV1ConfigurationsResponse)(interface{}, error){
    return response.GetResult.Body, response.GetResult.Err
}



// call result's extract function
func ExtractUpdateDbV1ConfigurationsResponse(response *UpdateDbV1ConfigurationsResponse)(interface{}, error){
    return response.UpdateResult.Body, response.UpdateResult.Err
}



// call result's extract function
func ExtractReplaceDbV1ConfigurationsResponse(response *ReplaceDbV1ConfigurationsResponse)(interface{}, error){
    return response.ReplaceResult.Body, response.ReplaceResult.Err
}



// call result's extract function
func ExtractDeleteDbV1ConfigurationsResponse(response *DeleteDbV1ConfigurationsResponse)(interface{}, error){
    return response.DeleteResult.Body, response.DeleteResult.Err
}



//extract response info from pager for ListDatastoreParamsDbV1Configurations
func ExtractListDatastoreParamsDbV1ConfigurationsResponse(response *ListDatastoreParamsDbV1ConfigurationsResponse)([]configurations.Param,error){
	page, err := response.Pager.AllPages()
	if err != nil {
		return nil, err
	}
	return configurations.ExtractParams(page)
}


// call result's extract function
func ExtractGetDatastoreParamDbV1ConfigurationsResponse(response *GetDatastoreParamDbV1ConfigurationsResponse)(interface{}, error){
    return response.ParamResult.Body, response.ParamResult.Err
}



//extract response info from pager for ListGlobalParamsDbV1Configurations
func ExtractListGlobalParamsDbV1ConfigurationsResponse(response *ListGlobalParamsDbV1ConfigurationsResponse)([]configurations.Param,error){
	page, err := response.Pager.AllPages()
	if err != nil {
		return nil, err
	}
	return configurations.ExtractParams(page)
}


// call result's extract function
func ExtractGetGlobalParamDbV1ConfigurationsResponse(response *GetGlobalParamDbV1ConfigurationsResponse)(interface{}, error){
    return response.ParamResult.Body, response.ParamResult.Err
}
