package openstack

// Code generated by cloud manager.

import (
    "github.com/gophercloud/gophercloud/openstack/networking/v2/ports"
)


//extract response info from pager for ListNetworkingV2Ports
func ExtractListNetworkingV2PortsResponse(response *ListNetworkingV2PortsResponse)([]ports.Port,error){
	page, err := response.Pager.AllPages()
	if err != nil {
		return nil, err
	}
	return ports.ExtractPorts(page)
}


// call result's extract function
func ExtractGetNetworkingV2PortsResponse(response *GetNetworkingV2PortsResponse)(interface{}, error){
    return response.GetResult.Body, response.GetResult.Err
}



// call result's extract function
func ExtractCreateNetworkingV2PortsResponse(response *CreateNetworkingV2PortsResponse)(interface{}, error){
    return response.CreateResult.Body, response.CreateResult.Err
}



// call result's extract function
func ExtractUpdateNetworkingV2PortsResponse(response *UpdateNetworkingV2PortsResponse)(interface{}, error){
    return response.UpdateResult.Body, response.UpdateResult.Err
}



// call result's extract function
func ExtractDeleteNetworkingV2PortsResponse(response *DeleteNetworkingV2PortsResponse)(interface{}, error){
    return response.DeleteResult.Body, response.DeleteResult.Err
}
