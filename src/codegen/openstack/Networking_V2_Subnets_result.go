package openstack

// Code generated by cloud manager.

import (
	"github.com/gophercloud/gophercloud/openstack/networking/v2/subnets"
)

// extract response info from pager for ListNetworkingV2Subnets
func ExtractListNetworkingV2SubnetsResponse(response *ListNetworkingV2SubnetsResponse) ([]subnets.Subnet, error) {
	page, err := response.Pager.AllPages()
	if err != nil {
		return nil, err
	}
	return subnets.ExtractSubnets(page)
}

// call result's extract function
func ExtractGetNetworkingV2SubnetsResponse(response *GetNetworkingV2SubnetsResponse) (interface{}, error) {
	return response.GetResult.Body, response.GetResult.Err
}

// call result's extract function
func ExtractCreateNetworkingV2SubnetsResponse(response *CreateNetworkingV2SubnetsResponse) (interface{}, error) {
	return response.CreateResult.Body, response.CreateResult.Err
}

// call result's extract function
func ExtractUpdateNetworkingV2SubnetsResponse(response *UpdateNetworkingV2SubnetsResponse) (interface{}, error) {
	return response.UpdateResult.Body, response.UpdateResult.Err
}

// call result's extract function
func ExtractDeleteNetworkingV2SubnetsResponse(response *DeleteNetworkingV2SubnetsResponse) (interface{}, error) {
	return response.DeleteResult.Body, response.DeleteResult.Err
}
