package openstack

// Code generated by cloud manager.

import (
)


// call result's extract function
func ExtractListBaremetalApiversionsResponse(response *ListBaremetalApiversionsResponse)(interface{}, error){
    return response.ListResult.Body, response.ListResult.Err
}



// call result's extract function
func ExtractGetBaremetalApiversionsResponse(response *GetBaremetalApiversionsResponse)(interface{}, error){
    return response.GetResult.Body, response.GetResult.Err
}