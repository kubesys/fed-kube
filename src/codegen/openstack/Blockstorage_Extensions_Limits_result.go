package openstack

// Code generated by cloud manager.

import (
)


// call result's extract function
func ExtractGetBlockstorageExtensionsLimitsResponse(response *GetBlockstorageExtensionsLimitsResponse)(interface{}, error){
    return response.GetResult.Body, response.GetResult.Err
}
