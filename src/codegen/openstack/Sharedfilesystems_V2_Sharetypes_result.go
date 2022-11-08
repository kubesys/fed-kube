package openstack

// Code generated by cloud manager.

import (
	"github.com/gophercloud/gophercloud/openstack/sharedfilesystems/v2/sharetypes"
)

// call result's extract function
func ExtractCreateSharedfilesystemsV2SharetypesResponse(response *CreateSharedfilesystemsV2SharetypesResponse) (interface{}, error) {
	return response.CreateResult.Body, response.CreateResult.Err
}

// call result's extract function
func ExtractDeleteSharedfilesystemsV2SharetypesResponse(response *DeleteSharedfilesystemsV2SharetypesResponse) (interface{}, error) {
	return response.DeleteResult.Body, response.DeleteResult.Err
}

// extract response info from pager for ListSharedfilesystemsV2Sharetypes
func ExtractListSharedfilesystemsV2SharetypesResponse(response *ListSharedfilesystemsV2SharetypesResponse) ([]sharetypes.ShareType, error) {
	page, err := response.Pager.AllPages()
	if err != nil {
		return nil, err
	}
	return sharetypes.ExtractShareTypes(page)
}

// call result's extract function
func ExtractGetDefaultSharedfilesystemsV2SharetypesResponse(response *GetDefaultSharedfilesystemsV2SharetypesResponse) (interface{}, error) {
	return response.GetDefaultResult.Body, response.GetDefaultResult.Err
}

// call result's extract function
func ExtractGetExtraSpecsSharedfilesystemsV2SharetypesResponse(response *GetExtraSpecsSharedfilesystemsV2SharetypesResponse) (interface{}, error) {
	return response.GetExtraSpecsResult.Body, response.GetExtraSpecsResult.Err
}

// call result's extract function
func ExtractSetExtraSpecsSharedfilesystemsV2SharetypesResponse(response *SetExtraSpecsSharedfilesystemsV2SharetypesResponse) (interface{}, error) {
	return response.SetExtraSpecsResult.Body, response.SetExtraSpecsResult.Err
}

// call result's extract function
func ExtractUnsetExtraSpecsSharedfilesystemsV2SharetypesResponse(response *UnsetExtraSpecsSharedfilesystemsV2SharetypesResponse) (interface{}, error) {
	return response.UnsetExtraSpecsResult.Body, response.UnsetExtraSpecsResult.Err
}

// call result's extract function
func ExtractShowAccessSharedfilesystemsV2SharetypesResponse(response *ShowAccessSharedfilesystemsV2SharetypesResponse) (interface{}, error) {
	return response.ShowAccessResult.Body, response.ShowAccessResult.Err
}

// call result's extract function
func ExtractAddAccessSharedfilesystemsV2SharetypesResponse(response *AddAccessSharedfilesystemsV2SharetypesResponse) (interface{}, error) {
	return response.AddAccessResult.Body, response.AddAccessResult.Err
}

// call result's extract function
func ExtractRemoveAccessSharedfilesystemsV2SharetypesResponse(response *RemoveAccessSharedfilesystemsV2SharetypesResponse) (interface{}, error) {
	return response.RemoveAccessResult.Body, response.RemoveAccessResult.Err
}
