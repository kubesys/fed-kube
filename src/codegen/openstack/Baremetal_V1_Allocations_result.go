package openstack

// Code generated by cloud manager.

import (
	"github.com/gophercloud/gophercloud/openstack/baremetal/v1/allocations"
)

// call result's extract function
func ExtractCreateBaremetalV1AllocationsResponse(response *CreateBaremetalV1AllocationsResponse) (interface{}, error) {
	return response.CreateResult.Body, response.CreateResult.Err
}

// extract response info from pager for ListBaremetalV1Allocations
func ExtractListBaremetalV1AllocationsResponse(response *ListBaremetalV1AllocationsResponse) ([]allocations.Allocation, error) {
	page, err := response.Pager.AllPages()
	if err != nil {
		return nil, err
	}
	return allocations.ExtractAllocations(page)
}

// call result's extract function
func ExtractGetBaremetalV1AllocationsResponse(response *GetBaremetalV1AllocationsResponse) (interface{}, error) {
	return response.GetResult.Body, response.GetResult.Err
}

// call result's extract function
func ExtractDeleteBaremetalV1AllocationsResponse(response *DeleteBaremetalV1AllocationsResponse) (interface{}, error) {
	return response.DeleteResult.Body, response.DeleteResult.Err
}
