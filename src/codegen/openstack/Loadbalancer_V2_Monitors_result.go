package openstack

// Code generated by cloud manager.

import (
    "github.com/gophercloud/gophercloud/openstack/loadbalancer/v2/monitors"
)


//extract response info from pager for ListLoadbalancerV2Monitors
func ExtractListLoadbalancerV2MonitorsResponse(response *ListLoadbalancerV2MonitorsResponse)([]monitors.Monitor,error){
	page, err := response.Pager.AllPages()
	if err != nil {
		return nil, err
	}
	return monitors.ExtractMonitors(page)
}


// call result's extract function
func ExtractCreateLoadbalancerV2MonitorsResponse(response *CreateLoadbalancerV2MonitorsResponse)(interface{}, error){
    return response.CreateResult.Body, response.CreateResult.Err
}



// call result's extract function
func ExtractGetLoadbalancerV2MonitorsResponse(response *GetLoadbalancerV2MonitorsResponse)(interface{}, error){
    return response.GetResult.Body, response.GetResult.Err
}



// call result's extract function
func ExtractUpdateLoadbalancerV2MonitorsResponse(response *UpdateLoadbalancerV2MonitorsResponse)(interface{}, error){
    return response.UpdateResult.Body, response.UpdateResult.Err
}



// call result's extract function
func ExtractDeleteLoadbalancerV2MonitorsResponse(response *DeleteLoadbalancerV2MonitorsResponse)(interface{}, error){
    return response.DeleteResult.Body, response.DeleteResult.Err
}
