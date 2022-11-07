package openstack

// Code generated by cloud manager.

import (
    "github.com/gophercloud/gophercloud/openstack/identity/v3/endpoints"
)


// call result's extract function
func ExtractCreateIdentityV3EndpointsResponse(response *CreateIdentityV3EndpointsResponse)(interface{}, error){
    return response.CreateResult.Body, response.CreateResult.Err
}



//extract response info from pager for ListIdentityV3Endpoints
func ExtractListIdentityV3EndpointsResponse(response *ListIdentityV3EndpointsResponse)([]endpoints.Endpoint,error){
	page, err := response.Pager.AllPages()
	if err != nil {
		return nil, err
	}
	return endpoints.ExtractEndpoints(page)
}


// call result's extract function
func ExtractUpdateIdentityV3EndpointsResponse(response *UpdateIdentityV3EndpointsResponse)(interface{}, error){
    return response.UpdateResult.Body, response.UpdateResult.Err
}



// call result's extract function
func ExtractDeleteIdentityV3EndpointsResponse(response *DeleteIdentityV3EndpointsResponse)(interface{}, error){
    return response.DeleteResult.Body, response.DeleteResult.Err
}