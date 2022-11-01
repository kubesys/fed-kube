package openstack

// Code generated by cloud manager.





import (
    "github.com/gophercloud/gophercloud/openstack/networking/v2/extensions/subnetpools"
    "github.com/gophercloud/gophercloud/pagination"
)
//request struct for the ListNetworkingV2ExtensionsSubnetpools
type ListNetworkingV2ExtensionsSubnetpoolsRequest struct{
    Opts subnetpools.ListOpts
}

func NewListNetworkingV2ExtensionsSubnetpoolsRequest()*ListNetworkingV2ExtensionsSubnetpoolsRequest{
    return &ListNetworkingV2ExtensionsSubnetpoolsRequest{}
}

//response struct for the ListNetworkingV2ExtensionsSubnetpools
type ListNetworkingV2ExtensionsSubnetpoolsResponse struct{
    Pager pagination.Pager
}

func NewListNetworkingV2ExtensionsSubnetpoolsResponse(pager pagination.Pager,)*ListNetworkingV2ExtensionsSubnetpoolsResponse {
    return &ListNetworkingV2ExtensionsSubnetpoolsResponse{
            Pager:pager,
    }
}

// action function
func (oc *OpenstackClient) ListNetworkingV2ExtensionsSubnetpools(req *ListNetworkingV2ExtensionsSubnetpoolsRequest)(*ListNetworkingV2ExtensionsSubnetpoolsResponse){
    return NewListNetworkingV2ExtensionsSubnetpoolsResponse(subnetpools.List(oc.Client,req.Opts, ))

}
//request struct for the GetNetworkingV2ExtensionsSubnetpools
type GetNetworkingV2ExtensionsSubnetpoolsRequest struct{
    Id string
}

func NewGetNetworkingV2ExtensionsSubnetpoolsRequest()*GetNetworkingV2ExtensionsSubnetpoolsRequest{
    return &GetNetworkingV2ExtensionsSubnetpoolsRequest{}
}

//response struct for the GetNetworkingV2ExtensionsSubnetpools
type GetNetworkingV2ExtensionsSubnetpoolsResponse struct{
    GetResult subnetpools.GetResult
}

func NewGetNetworkingV2ExtensionsSubnetpoolsResponse(getResult subnetpools.GetResult,)*GetNetworkingV2ExtensionsSubnetpoolsResponse {
    return &GetNetworkingV2ExtensionsSubnetpoolsResponse{
            GetResult:getResult,
    }
}

// action function
func (oc *OpenstackClient) GetNetworkingV2ExtensionsSubnetpools(req *GetNetworkingV2ExtensionsSubnetpoolsRequest)(*GetNetworkingV2ExtensionsSubnetpoolsResponse){
    return NewGetNetworkingV2ExtensionsSubnetpoolsResponse(subnetpools.Get(oc.Client,req.Id, ))

}
//request struct for the CreateNetworkingV2ExtensionsSubnetpools
type CreateNetworkingV2ExtensionsSubnetpoolsRequest struct{
    Opts subnetpools.CreateOpts
}

func NewCreateNetworkingV2ExtensionsSubnetpoolsRequest()*CreateNetworkingV2ExtensionsSubnetpoolsRequest{
    return &CreateNetworkingV2ExtensionsSubnetpoolsRequest{}
}

//response struct for the CreateNetworkingV2ExtensionsSubnetpools
type CreateNetworkingV2ExtensionsSubnetpoolsResponse struct{
    CreateResult subnetpools.CreateResult
}

func NewCreateNetworkingV2ExtensionsSubnetpoolsResponse(createResult subnetpools.CreateResult,)*CreateNetworkingV2ExtensionsSubnetpoolsResponse {
    return &CreateNetworkingV2ExtensionsSubnetpoolsResponse{
            CreateResult:createResult,
    }
}

// action function
func (oc *OpenstackClient) CreateNetworkingV2ExtensionsSubnetpools(req *CreateNetworkingV2ExtensionsSubnetpoolsRequest)(*CreateNetworkingV2ExtensionsSubnetpoolsResponse){
    return NewCreateNetworkingV2ExtensionsSubnetpoolsResponse(subnetpools.Create(oc.Client,req.Opts, ))

}
//request struct for the UpdateNetworkingV2ExtensionsSubnetpools
type UpdateNetworkingV2ExtensionsSubnetpoolsRequest struct{
    SubnetPoolID string
    Opts subnetpools.UpdateOpts
}

func NewUpdateNetworkingV2ExtensionsSubnetpoolsRequest()*UpdateNetworkingV2ExtensionsSubnetpoolsRequest{
    return &UpdateNetworkingV2ExtensionsSubnetpoolsRequest{}
}

//response struct for the UpdateNetworkingV2ExtensionsSubnetpools
type UpdateNetworkingV2ExtensionsSubnetpoolsResponse struct{
    UpdateResult subnetpools.UpdateResult
}

func NewUpdateNetworkingV2ExtensionsSubnetpoolsResponse(updateResult subnetpools.UpdateResult,)*UpdateNetworkingV2ExtensionsSubnetpoolsResponse {
    return &UpdateNetworkingV2ExtensionsSubnetpoolsResponse{
            UpdateResult:updateResult,
    }
}

// action function
func (oc *OpenstackClient) UpdateNetworkingV2ExtensionsSubnetpools(req *UpdateNetworkingV2ExtensionsSubnetpoolsRequest)(*UpdateNetworkingV2ExtensionsSubnetpoolsResponse){
    return NewUpdateNetworkingV2ExtensionsSubnetpoolsResponse(subnetpools.Update(oc.Client,req.SubnetPoolID,req.Opts, ))

}
//request struct for the DeleteNetworkingV2ExtensionsSubnetpools
type DeleteNetworkingV2ExtensionsSubnetpoolsRequest struct{
    Id string
}

func NewDeleteNetworkingV2ExtensionsSubnetpoolsRequest()*DeleteNetworkingV2ExtensionsSubnetpoolsRequest{
    return &DeleteNetworkingV2ExtensionsSubnetpoolsRequest{}
}

//response struct for the DeleteNetworkingV2ExtensionsSubnetpools
type DeleteNetworkingV2ExtensionsSubnetpoolsResponse struct{
    DeleteResult subnetpools.DeleteResult
}

func NewDeleteNetworkingV2ExtensionsSubnetpoolsResponse(deleteResult subnetpools.DeleteResult,)*DeleteNetworkingV2ExtensionsSubnetpoolsResponse {
    return &DeleteNetworkingV2ExtensionsSubnetpoolsResponse{
            DeleteResult:deleteResult,
    }
}

// action function
func (oc *OpenstackClient) DeleteNetworkingV2ExtensionsSubnetpools(req *DeleteNetworkingV2ExtensionsSubnetpoolsRequest)(*DeleteNetworkingV2ExtensionsSubnetpoolsResponse){
    return NewDeleteNetworkingV2ExtensionsSubnetpoolsResponse(subnetpools.Delete(oc.Client,req.Id, ))

}