package openstack

// Code generated by cloud manager.

import (
	"github.com/gophercloud/gophercloud/openstack/networking/v2/extensions/fwaas_v2/groups"
)

// extract response info from pager for ListNetworkingV2ExtensionsFwaas_v2Groups
func ExtractListNetworkingV2ExtensionsFwaas_v2GroupsResponse(response *ListNetworkingV2ExtensionsFwaas_v2GroupsResponse) ([]groups.Group, error) {
	page, err := response.Pager.AllPages()
	if err != nil {
		return nil, err
	}
	return groups.ExtractGroups(page)
}

// call result's extract function
func ExtractGetNetworkingV2ExtensionsFwaas_v2GroupsResponse(response *GetNetworkingV2ExtensionsFwaas_v2GroupsResponse) (interface{}, error) {
	return response.GetResult.Body, response.GetResult.Err
}

// call result's extract function
func ExtractCreateNetworkingV2ExtensionsFwaas_v2GroupsResponse(response *CreateNetworkingV2ExtensionsFwaas_v2GroupsResponse) (interface{}, error) {
	return response.CreateResult.Body, response.CreateResult.Err
}

// call result's extract function
func ExtractUpdateNetworkingV2ExtensionsFwaas_v2GroupsResponse(response *UpdateNetworkingV2ExtensionsFwaas_v2GroupsResponse) (interface{}, error) {
	return response.UpdateResult.Body, response.UpdateResult.Err
}

// call result's extract function
func ExtractDeleteNetworkingV2ExtensionsFwaas_v2GroupsResponse(response *DeleteNetworkingV2ExtensionsFwaas_v2GroupsResponse) (interface{}, error) {
	return response.DeleteResult.Body, response.DeleteResult.Err
}
