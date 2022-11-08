package openstack

// Code generated by cloud manager.

import (
	"github.com/gophercloud/gophercloud/openstack/compute/v2/flavors"
)

// extract response info from pager for ListDetailComputeV2Flavors
func ExtractListDetailComputeV2FlavorsResponse(response *ListDetailComputeV2FlavorsResponse) ([]flavors.Flavor, error) {
	page, err := response.Pager.AllPages()
	if err != nil {
		return nil, err
	}
	return flavors.ExtractFlavors(page)
}

// call result's extract function
func ExtractCreateComputeV2FlavorsResponse(response *CreateComputeV2FlavorsResponse) (interface{}, error) {
	return response.CreateResult.Body, response.CreateResult.Err
}

// call result's extract function
func ExtractGetComputeV2FlavorsResponse(response *GetComputeV2FlavorsResponse) (interface{}, error) {
	return response.GetResult.Body, response.GetResult.Err
}

// call result's extract function
func ExtractDeleteComputeV2FlavorsResponse(response *DeleteComputeV2FlavorsResponse) (interface{}, error) {
	return response.DeleteResult.Body, response.DeleteResult.Err
}

// extract response info from pager for ListAccessesComputeV2Flavors
func ExtractListAccessesComputeV2FlavorsResponse(response *ListAccessesComputeV2FlavorsResponse) ([]flavors.FlavorAccess, error) {
	page, err := response.Pager.AllPages()
	if err != nil {
		return nil, err
	}
	return flavors.ExtractAccesses(page)
}

// call result's extract function
func ExtractAddAccessComputeV2FlavorsResponse(response *AddAccessComputeV2FlavorsResponse) (interface{}, error) {
	return response.AddAccessResult.Body, response.AddAccessResult.Err
}

// call result's extract function
func ExtractRemoveAccessComputeV2FlavorsResponse(response *RemoveAccessComputeV2FlavorsResponse) (interface{}, error) {
	return response.RemoveAccessResult.Body, response.RemoveAccessResult.Err
}

// call result's extract function
func ExtractListExtraSpecsComputeV2FlavorsResponse(response *ListExtraSpecsComputeV2FlavorsResponse) (interface{}, error) {
	return response.ListExtraSpecsResult.Body, response.ListExtraSpecsResult.Err
}

// call result's extract function
func ExtractGetExtraSpecComputeV2FlavorsResponse(response *GetExtraSpecComputeV2FlavorsResponse) (interface{}, error) {
	return response.GetExtraSpecResult.Body, response.GetExtraSpecResult.Err
}

// call result's extract function
func ExtractCreateExtraSpecsComputeV2FlavorsResponse(response *CreateExtraSpecsComputeV2FlavorsResponse) (interface{}, error) {
	return response.CreateExtraSpecsResult.Body, response.CreateExtraSpecsResult.Err
}

// call result's extract function
func ExtractUpdateExtraSpecComputeV2FlavorsResponse(response *UpdateExtraSpecComputeV2FlavorsResponse) (interface{}, error) {
	return response.UpdateExtraSpecResult.Body, response.UpdateExtraSpecResult.Err
}

// call result's extract function
func ExtractDeleteExtraSpecComputeV2FlavorsResponse(response *DeleteExtraSpecComputeV2FlavorsResponse) (interface{}, error) {
	return response.DeleteExtraSpecResult.Body, response.DeleteExtraSpecResult.Err
}
