package openstack

// Code generated by cloud manager.

import (
)


// call result's extract function
func ExtractResetNetworkComputeV2ExtensionsResetnetworkResponse(response *ResetNetworkComputeV2ExtensionsResetnetworkResponse)(interface{}, error){
    return response.ResetResult.Body, response.ResetResult.Err
}
