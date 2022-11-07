package openstack

// Code generated by cloud manager.

import (
    "github.com/gophercloud/gophercloud/openstack/clustering/v1/nodes"
)


// call result's extract function
func ExtractCreateClusteringV1NodesResponse(response *CreateClusteringV1NodesResponse)(interface{}, error){
    return response.CreateResult.Body, response.CreateResult.Err
}



// call result's extract function
func ExtractUpdateClusteringV1NodesResponse(response *UpdateClusteringV1NodesResponse)(interface{}, error){
    return response.UpdateResult.Body, response.UpdateResult.Err
}



//extract response info from pager for ListClusteringV1Nodes
func ExtractListClusteringV1NodesResponse(response *ListClusteringV1NodesResponse)([]nodes.Node,error){
	page, err := response.Pager.AllPages()
	if err != nil {
		return nil, err
	}
	return nodes.ExtractNodes(page)
}


// call result's extract function
func ExtractDeleteClusteringV1NodesResponse(response *DeleteClusteringV1NodesResponse)(interface{}, error){
    return response.DeleteResult.Body, response.DeleteResult.Err
}



// call result's extract function
func ExtractGetClusteringV1NodesResponse(response *GetClusteringV1NodesResponse)(interface{}, error){
    return response.GetResult.Body, response.GetResult.Err
}



// call result's extract function
func ExtractOpsClusteringV1NodesResponse(response *OpsClusteringV1NodesResponse)(interface{}, error){
    return response.ActionResult.Body, response.ActionResult.Err
}



// call result's extract function
func ExtractRecoverClusteringV1NodesResponse(response *RecoverClusteringV1NodesResponse)(interface{}, error){
    return response.ActionResult.Body, response.ActionResult.Err
}



// call result's extract function
func ExtractCheckClusteringV1NodesResponse(response *CheckClusteringV1NodesResponse)(interface{}, error){
    return response.ActionResult.Body, response.ActionResult.Err
}