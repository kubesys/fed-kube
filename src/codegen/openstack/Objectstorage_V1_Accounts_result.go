package openstack

// Code generated by cloud manager.

import (
)


// call result's extract function
func ExtractGetObjectstorageV1AccountsResponse(response *GetObjectstorageV1AccountsResponse)(interface{}, error){
    return response.GetResult.Body, response.GetResult.Err
}



// call result's extract function
func ExtractUpdateObjectstorageV1AccountsResponse(response *UpdateObjectstorageV1AccountsResponse)(interface{}, error){
    return response.UpdateResult.Body, response.UpdateResult.Err
}
