package openstack

// Code generated by cloud manager.

import (
    "github.com/gophercloud/gophercloud/openstack/identity/v3/regions"
)


//extract response info from pager for ListIdentityV3Regions
func ExtractListIdentityV3RegionsResponse(response *ListIdentityV3RegionsResponse)([]regions.Region,error){
	page, err := response.Pager.AllPages()
	if err != nil {
		return nil, err
	}
	return regions.ExtractRegions(page)
}


// call result's extract function
func ExtractGetIdentityV3RegionsResponse(response *GetIdentityV3RegionsResponse)(interface{}, error){
    return response.GetResult.Body, response.GetResult.Err
}



// call result's extract function
func ExtractCreateIdentityV3RegionsResponse(response *CreateIdentityV3RegionsResponse)(interface{}, error){
    return response.CreateResult.Body, response.CreateResult.Err
}



// call result's extract function
func ExtractUpdateIdentityV3RegionsResponse(response *UpdateIdentityV3RegionsResponse)(interface{}, error){
    return response.UpdateResult.Body, response.UpdateResult.Err
}



// call result's extract function
func ExtractDeleteIdentityV3RegionsResponse(response *DeleteIdentityV3RegionsResponse)(interface{}, error){
    return response.DeleteResult.Body, response.DeleteResult.Err
}
