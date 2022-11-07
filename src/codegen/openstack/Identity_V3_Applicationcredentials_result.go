package openstack

// Code generated by cloud manager.

import (
    "github.com/gophercloud/gophercloud/openstack/identity/v3/applicationcredentials"
)


//extract response info from pager for ListIdentityV3Applicationcredentials
func ExtractListIdentityV3ApplicationcredentialsResponse(response *ListIdentityV3ApplicationcredentialsResponse)([]applicationcredentials.ApplicationCredential,error){
	page, err := response.Pager.AllPages()
	if err != nil {
		return nil, err
	}
	return applicationcredentials.ExtractApplicationCredentials(page)
}


// call result's extract function
func ExtractGetIdentityV3ApplicationcredentialsResponse(response *GetIdentityV3ApplicationcredentialsResponse)(interface{}, error){
    return response.GetResult.Body, response.GetResult.Err
}



// call result's extract function
func ExtractCreateIdentityV3ApplicationcredentialsResponse(response *CreateIdentityV3ApplicationcredentialsResponse)(interface{}, error){
    return response.CreateResult.Body, response.CreateResult.Err
}



// call result's extract function
func ExtractDeleteIdentityV3ApplicationcredentialsResponse(response *DeleteIdentityV3ApplicationcredentialsResponse)(interface{}, error){
    return response.DeleteResult.Body, response.DeleteResult.Err
}



//extract response info from pager for ListAccessRulesIdentityV3Applicationcredentials
func ExtractListAccessRulesIdentityV3ApplicationcredentialsResponse(response *ListAccessRulesIdentityV3ApplicationcredentialsResponse)([]applicationcredentials.AccessRule,error){
	page, err := response.Pager.AllPages()
	if err != nil {
		return nil, err
	}
	return applicationcredentials.ExtractAccessRules(page)
}


// call result's extract function
func ExtractGetAccessRuleIdentityV3ApplicationcredentialsResponse(response *GetAccessRuleIdentityV3ApplicationcredentialsResponse)(interface{}, error){
    return response.GetAccessRuleResult.Body, response.GetAccessRuleResult.Err
}



// call result's extract function
func ExtractDeleteAccessRuleIdentityV3ApplicationcredentialsResponse(response *DeleteAccessRuleIdentityV3ApplicationcredentialsResponse)(interface{}, error){
    return response.DeleteResult.Body, response.DeleteResult.Err
}
