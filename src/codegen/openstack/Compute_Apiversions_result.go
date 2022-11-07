package openstack

// Code generated by cloud manager.

import (
    "github.com/gophercloud/gophercloud/openstack/compute/apiversions"
)


//extract response info from pager for ListComputeApiversions
func ExtractListComputeApiversionsResponse(response *ListComputeApiversionsResponse)([]apiversions.APIVersion,error){
	page, err := response.Pager.AllPages()
	if err != nil {
		return nil, err
	}
	return apiversions.ExtractAPIVersions(page)
}


// call result's extract function
func ExtractGetComputeApiversionsResponse(response *GetComputeApiversionsResponse)(interface{}, error){
    return response.GetResult.Body, response.GetResult.Err
}
