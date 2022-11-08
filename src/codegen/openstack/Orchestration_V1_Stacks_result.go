package openstack

// Code generated by cloud manager.

import (
	"github.com/gophercloud/gophercloud/openstack/orchestration/v1/stacks"
)

// call result's extract function
func ExtractCreateOrchestrationV1StacksResponse(response *CreateOrchestrationV1StacksResponse) (interface{}, error) {
	return response.CreateResult.Body, response.CreateResult.Err
}

// call result's extract function
func ExtractAdoptOrchestrationV1StacksResponse(response *AdoptOrchestrationV1StacksResponse) (interface{}, error) {
	return response.AdoptResult.Body, response.AdoptResult.Err
}

// extract response info from pager for ListOrchestrationV1Stacks
func ExtractListOrchestrationV1StacksResponse(response *ListOrchestrationV1StacksResponse) ([]stacks.ListedStack, error) {
	page, err := response.Pager.AllPages()
	if err != nil {
		return nil, err
	}
	return stacks.ExtractStacks(page)
}

// call result's extract function
func ExtractGetOrchestrationV1StacksResponse(response *GetOrchestrationV1StacksResponse) (interface{}, error) {
	return response.GetResult.Body, response.GetResult.Err
}

// call result's extract function
func ExtractFindOrchestrationV1StacksResponse(response *FindOrchestrationV1StacksResponse) (interface{}, error) {
	return response.GetResult.Body, response.GetResult.Err
}

// call result's extract function
func ExtractUpdateOrchestrationV1StacksResponse(response *UpdateOrchestrationV1StacksResponse) (interface{}, error) {
	return response.UpdateResult.Body, response.UpdateResult.Err
}

// call result's extract function
func ExtractUpdatePatchOrchestrationV1StacksResponse(response *UpdatePatchOrchestrationV1StacksResponse) (interface{}, error) {
	return response.UpdateResult.Body, response.UpdateResult.Err
}

// call result's extract function
func ExtractDeleteOrchestrationV1StacksResponse(response *DeleteOrchestrationV1StacksResponse) (interface{}, error) {
	return response.DeleteResult.Body, response.DeleteResult.Err
}

// call result's extract function
func ExtractPreviewOrchestrationV1StacksResponse(response *PreviewOrchestrationV1StacksResponse) (interface{}, error) {
	return response.PreviewResult.Body, response.PreviewResult.Err
}

// call result's extract function
func ExtractAbandonOrchestrationV1StacksResponse(response *AbandonOrchestrationV1StacksResponse) (interface{}, error) {
	return response.AbandonResult.Body, response.AbandonResult.Err
}
