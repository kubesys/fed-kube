package openstack

// Code generated by cloud manager.

import (
	"github.com/gophercloud/gophercloud/openstack/networking/v2/extensions/trunks"
)

// call result's extract function
func ExtractCreateNetworkingV2ExtensionsTrunksResponse(response *CreateNetworkingV2ExtensionsTrunksResponse) (interface{}, error) {
	return response.CreateResult.Body, response.CreateResult.Err
}

// call result's extract function
func ExtractDeleteNetworkingV2ExtensionsTrunksResponse(response *DeleteNetworkingV2ExtensionsTrunksResponse) (interface{}, error) {
	return response.DeleteResult.Body, response.DeleteResult.Err
}

// extract response info from pager for ListNetworkingV2ExtensionsTrunks
func ExtractListNetworkingV2ExtensionsTrunksResponse(response *ListNetworkingV2ExtensionsTrunksResponse) ([]trunks.Trunk, error) {
	page, err := response.Pager.AllPages()
	if err != nil {
		return nil, err
	}
	return trunks.ExtractTrunks(page)
}

// call result's extract function
func ExtractGetNetworkingV2ExtensionsTrunksResponse(response *GetNetworkingV2ExtensionsTrunksResponse) (interface{}, error) {
	return response.GetResult.Body, response.GetResult.Err
}

// call result's extract function
func ExtractUpdateNetworkingV2ExtensionsTrunksResponse(response *UpdateNetworkingV2ExtensionsTrunksResponse) (interface{}, error) {
	return response.UpdateResult.Body, response.UpdateResult.Err
}

// call result's extract function
func ExtractGetSubportsNetworkingV2ExtensionsTrunksResponse(response *GetSubportsNetworkingV2ExtensionsTrunksResponse) (interface{}, error) {
	return response.GetSubportsResult.Body, response.GetSubportsResult.Err
}

// call result's extract function
func ExtractAddSubportsNetworkingV2ExtensionsTrunksResponse(response *AddSubportsNetworkingV2ExtensionsTrunksResponse) (interface{}, error) {
	return response.UpdateSubportsResult.Body, response.UpdateSubportsResult.Err
}

// call result's extract function
func ExtractRemoveSubportsNetworkingV2ExtensionsTrunksResponse(response *RemoveSubportsNetworkingV2ExtensionsTrunksResponse) (interface{}, error) {
	return response.UpdateSubportsResult.Body, response.UpdateSubportsResult.Err
}
