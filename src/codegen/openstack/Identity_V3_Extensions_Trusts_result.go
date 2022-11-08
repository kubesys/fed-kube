package openstack

// Code generated by cloud manager.

import (
    "github.com/gophercloud/gophercloud/openstack/identity/v3/extensions/trusts"
)


// call result's extract function
func ExtractCreateIdentityV3ExtensionsTrustsResponse(response *CreateIdentityV3ExtensionsTrustsResponse)(interface{}, error){
    return response.CreateResult.Body, response.CreateResult.Err
}



// call result's extract function
func ExtractDeleteIdentityV3ExtensionsTrustsResponse(response *DeleteIdentityV3ExtensionsTrustsResponse)(interface{}, error){
    return response.DeleteResult.Body, response.DeleteResult.Err
}



//extract response info from pager for ListIdentityV3ExtensionsTrusts
func ExtractListIdentityV3ExtensionsTrustsResponse(response *ListIdentityV3ExtensionsTrustsResponse)([]trusts.Trust,error){
	page, err := response.Pager.AllPages()
	if err != nil {
		return nil, err
	}
	return trusts.ExtractTrusts(page)
}


// call result's extract function
func ExtractGetIdentityV3ExtensionsTrustsResponse(response *GetIdentityV3ExtensionsTrustsResponse)(interface{}, error){
    return response.GetResult.Body, response.GetResult.Err
}



//extract response info from pager for ListRolesIdentityV3ExtensionsTrusts
func ExtractListRolesIdentityV3ExtensionsTrustsResponse(response *ListRolesIdentityV3ExtensionsTrustsResponse)([]trusts.Role,error){
	page, err := response.Pager.AllPages()
	if err != nil {
		return nil, err
	}
	return trusts.ExtractRoles(page)
}


// call result's extract function
func ExtractGetRoleIdentityV3ExtensionsTrustsResponse(response *GetRoleIdentityV3ExtensionsTrustsResponse)(interface{}, error){
    return response.GetRoleResult.Body, response.GetRoleResult.Err
}



// call result's extract function
func ExtractCheckRoleIdentityV3ExtensionsTrustsResponse(response *CheckRoleIdentityV3ExtensionsTrustsResponse)(interface{}, error){
    return response.CheckRoleResult.Body, response.CheckRoleResult.Err
}
