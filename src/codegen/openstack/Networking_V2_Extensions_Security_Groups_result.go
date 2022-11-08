package openstack

// Code generated by cloud manager.

import (
	"github.com/gophercloud/gophercloud/openstack/networking/v2/extensions/security/groups"
)

// extract response info from pager for ListNetworkingV2ExtensionsSecurityGroups
func ExtractListNetworkingV2ExtensionsSecurityGroupsResponse(response *ListNetworkingV2ExtensionsSecurityGroupsResponse) ([]groups.SecGroup, error) {
	page, err := response.Pager.AllPages()
	if err != nil {
		return nil, err
	}
	return groups.ExtractGroups(page)
}

// call result's extract function
func ExtractCreateNetworkingV2ExtensionsSecurityGroupsResponse(response *CreateNetworkingV2ExtensionsSecurityGroupsResponse) (interface{}, error) {
	return response.CreateResult.Body, response.CreateResult.Err
}

// call result's extract function
func ExtractUpdateNetworkingV2ExtensionsSecurityGroupsResponse(response *UpdateNetworkingV2ExtensionsSecurityGroupsResponse) (interface{}, error) {
	return response.UpdateResult.Body, response.UpdateResult.Err
}

// call result's extract function
func ExtractGetNetworkingV2ExtensionsSecurityGroupsResponse(response *GetNetworkingV2ExtensionsSecurityGroupsResponse) (interface{}, error) {
	return response.GetResult.Body, response.GetResult.Err
}

// call result's extract function
func ExtractDeleteNetworkingV2ExtensionsSecurityGroupsResponse(response *DeleteNetworkingV2ExtensionsSecurityGroupsResponse) (interface{}, error) {
	return response.DeleteResult.Body, response.DeleteResult.Err
}
