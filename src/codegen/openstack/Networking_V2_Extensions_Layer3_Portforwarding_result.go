package openstack

// Code generated by cloud manager.

import (
    "github.com/gophercloud/gophercloud/openstack/networking/v2/extensions/layer3/portforwarding"
)


//extract response info from pager for ListNetworkingV2ExtensionsLayer3Portforwarding
func ExtractListNetworkingV2ExtensionsLayer3PortforwardingResponse(response *ListNetworkingV2ExtensionsLayer3PortforwardingResponse)([]portforwarding.PortForwarding,error){
	page, err := response.Pager.AllPages()
	if err != nil {
		return nil, err
	}
	return portforwarding.ExtractPortForwardings(page)
}


// call result's extract function
func ExtractGetNetworkingV2ExtensionsLayer3PortforwardingResponse(response *GetNetworkingV2ExtensionsLayer3PortforwardingResponse)(interface{}, error){
    return response.GetResult.Body, response.GetResult.Err
}



// call result's extract function
func ExtractCreateNetworkingV2ExtensionsLayer3PortforwardingResponse(response *CreateNetworkingV2ExtensionsLayer3PortforwardingResponse)(interface{}, error){
    return response.CreateResult.Body, response.CreateResult.Err
}



// call result's extract function
func ExtractUpdateNetworkingV2ExtensionsLayer3PortforwardingResponse(response *UpdateNetworkingV2ExtensionsLayer3PortforwardingResponse)(interface{}, error){
    return response.UpdateResult.Body, response.UpdateResult.Err
}



// call result's extract function
func ExtractDeleteNetworkingV2ExtensionsLayer3PortforwardingResponse(response *DeleteNetworkingV2ExtensionsLayer3PortforwardingResponse)(interface{}, error){
    return response.DeleteResult.Body, response.DeleteResult.Err
}
