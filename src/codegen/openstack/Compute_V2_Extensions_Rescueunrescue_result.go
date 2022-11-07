package openstack

// Code generated by cloud manager.

import (
)


// call result's extract function
func ExtractRescueComputeV2ExtensionsRescueunrescueResponse(response *RescueComputeV2ExtensionsRescueunrescueResponse)(interface{}, error){
    return response.RescueResult.Body, response.RescueResult.Err
}



// call result's extract function
func ExtractUnrescueComputeV2ExtensionsRescueunrescueResponse(response *UnrescueComputeV2ExtensionsRescueunrescueResponse)(interface{}, error){
    return response.UnrescueResult.Body, response.UnrescueResult.Err
}
