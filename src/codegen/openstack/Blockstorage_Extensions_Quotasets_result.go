package openstack

// Code generated by cloud manager.

import ()

// call result's extract function
func ExtractGetBlockstorageExtensionsQuotasetsResponse(response *GetBlockstorageExtensionsQuotasetsResponse) (interface{}, error) {
	return response.GetResult.Body, response.GetResult.Err
}

// call result's extract function
func ExtractGetDefaultsBlockstorageExtensionsQuotasetsResponse(response *GetDefaultsBlockstorageExtensionsQuotasetsResponse) (interface{}, error) {
	return response.GetResult.Body, response.GetResult.Err
}

// call result's extract function
func ExtractGetUsageBlockstorageExtensionsQuotasetsResponse(response *GetUsageBlockstorageExtensionsQuotasetsResponse) (interface{}, error) {
	return response.GetUsageResult.Body, response.GetUsageResult.Err
}

// call result's extract function
func ExtractUpdateBlockstorageExtensionsQuotasetsResponse(response *UpdateBlockstorageExtensionsQuotasetsResponse) (interface{}, error) {
	return response.UpdateResult.Body, response.UpdateResult.Err
}

// call result's extract function
func ExtractDeleteBlockstorageExtensionsQuotasetsResponse(response *DeleteBlockstorageExtensionsQuotasetsResponse) (interface{}, error) {
	return response.DeleteResult.Body, response.DeleteResult.Err
}
