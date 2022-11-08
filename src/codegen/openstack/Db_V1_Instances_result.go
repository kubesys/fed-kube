package openstack

// Code generated by cloud manager.

import (
	"github.com/gophercloud/gophercloud/openstack/db/v1/instances"
)

// call result's extract function
func ExtractCreateDbV1InstancesResponse(response *CreateDbV1InstancesResponse) (interface{}, error) {
	return response.CreateResult.Body, response.CreateResult.Err
}

// extract response info from pager for ListDbV1Instances
func ExtractListDbV1InstancesResponse(response *ListDbV1InstancesResponse) ([]instances.Instance, error) {
	page, err := response.Pager.AllPages()
	if err != nil {
		return nil, err
	}
	return instances.ExtractInstances(page)
}

// call result's extract function
func ExtractGetDbV1InstancesResponse(response *GetDbV1InstancesResponse) (interface{}, error) {
	return response.GetResult.Body, response.GetResult.Err
}

// call result's extract function
func ExtractDeleteDbV1InstancesResponse(response *DeleteDbV1InstancesResponse) (interface{}, error) {
	return response.DeleteResult.Body, response.DeleteResult.Err
}

// call result's extract function
func ExtractEnableRootUserDbV1InstancesResponse(response *EnableRootUserDbV1InstancesResponse) (interface{}, error) {
	return response.EnableRootUserResult.Body, response.EnableRootUserResult.Err
}

// call result's extract function
func ExtractIsRootEnabledDbV1InstancesResponse(response *IsRootEnabledDbV1InstancesResponse) (interface{}, error) {
	return response.IsRootEnabledResult.Body, response.IsRootEnabledResult.Err
}

// call result's extract function
func ExtractRestartDbV1InstancesResponse(response *RestartDbV1InstancesResponse) (interface{}, error) {
	return response.ActionResult.Body, response.ActionResult.Err
}

// call result's extract function
func ExtractResizeDbV1InstancesResponse(response *ResizeDbV1InstancesResponse) (interface{}, error) {
	return response.ActionResult.Body, response.ActionResult.Err
}

// call result's extract function
func ExtractResizeVolumeDbV1InstancesResponse(response *ResizeVolumeDbV1InstancesResponse) (interface{}, error) {
	return response.ActionResult.Body, response.ActionResult.Err
}

// call result's extract function
func ExtractAttachConfigurationGroupDbV1InstancesResponse(response *AttachConfigurationGroupDbV1InstancesResponse) (interface{}, error) {
	return response.ConfigurationResult.Body, response.ConfigurationResult.Err
}

// call result's extract function
func ExtractDetachConfigurationGroupDbV1InstancesResponse(response *DetachConfigurationGroupDbV1InstancesResponse) (interface{}, error) {
	return response.ConfigurationResult.Body, response.ConfigurationResult.Err
}
