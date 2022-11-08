package openstack

// Code generated by cloud manager.

import ()

// call result's extract function
func ExtractCreateIdentityV3TokensResponse(response *CreateIdentityV3TokensResponse) (interface{}, error) {
	return response.CreateResult.Body, response.CreateResult.Err
}

// call result's extract function
func ExtractGetIdentityV3TokensResponse(response *GetIdentityV3TokensResponse) (interface{}, error) {
	return response.GetResult.Body, response.GetResult.Err
}

// call result's extract function
func ExtractRevokeIdentityV3TokensResponse(response *RevokeIdentityV3TokensResponse) (interface{}, error) {
	return response.RevokeResult.Body, response.RevokeResult.Err
}
