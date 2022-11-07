package openstack

// Code generated by cloud manager.

import (
    "github.com/gophercloud/gophercloud/openstack/networking/v2/extensions/fwaas/rules"
)


//extract response info from pager for ListNetworkingV2ExtensionsFwaasRules
func ExtractListNetworkingV2ExtensionsFwaasRulesResponse(response *ListNetworkingV2ExtensionsFwaasRulesResponse)([]rules.Rule,error){
	page, err := response.Pager.AllPages()
	if err != nil {
		return nil, err
	}
	return rules.ExtractRules(page)
}


// call result's extract function
func ExtractCreateNetworkingV2ExtensionsFwaasRulesResponse(response *CreateNetworkingV2ExtensionsFwaasRulesResponse)(interface{}, error){
    return response.CreateResult.Body, response.CreateResult.Err
}



// call result's extract function
func ExtractGetNetworkingV2ExtensionsFwaasRulesResponse(response *GetNetworkingV2ExtensionsFwaasRulesResponse)(interface{}, error){
    return response.GetResult.Body, response.GetResult.Err
}



// call result's extract function
func ExtractUpdateNetworkingV2ExtensionsFwaasRulesResponse(response *UpdateNetworkingV2ExtensionsFwaasRulesResponse)(interface{}, error){
    return response.UpdateResult.Body, response.UpdateResult.Err
}



// call result's extract function
func ExtractDeleteNetworkingV2ExtensionsFwaasRulesResponse(response *DeleteNetworkingV2ExtensionsFwaasRulesResponse)(interface{}, error){
    return response.DeleteResult.Body, response.DeleteResult.Err
}
