package openstack

// Code generated by cloud manager.





import (
    "github.com/gophercloud/gophercloud/openstack/networking/v2/ports"
    "github.com/gophercloud/gophercloud/pagination"
)
//request struct for the ListNetworkingV2Ports
type ListNetworkingV2PortsRequest struct{
    Opts ports.ListOpts
}

func NewListNetworkingV2PortsRequest()*ListNetworkingV2PortsRequest{
    return &ListNetworkingV2PortsRequest{}
}

//response struct for the ListNetworkingV2Ports
type ListNetworkingV2PortsResponse struct{
    Pager pagination.Pager
}

func NewListNetworkingV2PortsResponse(pager pagination.Pager,)*ListNetworkingV2PortsResponse {
    return &ListNetworkingV2PortsResponse{
            Pager:pager,
    }
}

// action function
func (oc *OpenstackClient) ListNetworkingV2Ports(req *ListNetworkingV2PortsRequest)(*ListNetworkingV2PortsResponse){
    return NewListNetworkingV2PortsResponse(ports.List(oc.Client,req.Opts, ))

}
//request struct for the GetNetworkingV2Ports
type GetNetworkingV2PortsRequest struct{
    Id string
}

func NewGetNetworkingV2PortsRequest()*GetNetworkingV2PortsRequest{
    return &GetNetworkingV2PortsRequest{}
}

//response struct for the GetNetworkingV2Ports
type GetNetworkingV2PortsResponse struct{
    GetResult ports.GetResult
}

func NewGetNetworkingV2PortsResponse(getResult ports.GetResult,)*GetNetworkingV2PortsResponse {
    return &GetNetworkingV2PortsResponse{
            GetResult:getResult,
    }
}

// action function
func (oc *OpenstackClient) GetNetworkingV2Ports(req *GetNetworkingV2PortsRequest)(*GetNetworkingV2PortsResponse){
    return NewGetNetworkingV2PortsResponse(ports.Get(oc.Client,req.Id, ))

}
//request struct for the CreateNetworkingV2Ports
type CreateNetworkingV2PortsRequest struct{
    Opts ports.CreateOpts
}

func NewCreateNetworkingV2PortsRequest()*CreateNetworkingV2PortsRequest{
    return &CreateNetworkingV2PortsRequest{}
}

//response struct for the CreateNetworkingV2Ports
type CreateNetworkingV2PortsResponse struct{
    CreateResult ports.CreateResult
}

func NewCreateNetworkingV2PortsResponse(createResult ports.CreateResult,)*CreateNetworkingV2PortsResponse {
    return &CreateNetworkingV2PortsResponse{
            CreateResult:createResult,
    }
}

// action function
func (oc *OpenstackClient) CreateNetworkingV2Ports(req *CreateNetworkingV2PortsRequest)(*CreateNetworkingV2PortsResponse){
    return NewCreateNetworkingV2PortsResponse(ports.Create(oc.Client,req.Opts, ))

}
//request struct for the UpdateNetworkingV2Ports
type UpdateNetworkingV2PortsRequest struct{
    Id string
    Opts ports.UpdateOpts
}

func NewUpdateNetworkingV2PortsRequest()*UpdateNetworkingV2PortsRequest{
    return &UpdateNetworkingV2PortsRequest{}
}

//response struct for the UpdateNetworkingV2Ports
type UpdateNetworkingV2PortsResponse struct{
    UpdateResult ports.UpdateResult
}

func NewUpdateNetworkingV2PortsResponse(updateResult ports.UpdateResult,)*UpdateNetworkingV2PortsResponse {
    return &UpdateNetworkingV2PortsResponse{
            UpdateResult:updateResult,
    }
}

// action function
func (oc *OpenstackClient) UpdateNetworkingV2Ports(req *UpdateNetworkingV2PortsRequest)(*UpdateNetworkingV2PortsResponse){
    return NewUpdateNetworkingV2PortsResponse(ports.Update(oc.Client,req.Id,req.Opts, ))

}
//request struct for the DeleteNetworkingV2Ports
type DeleteNetworkingV2PortsRequest struct{
    Id string
}

func NewDeleteNetworkingV2PortsRequest()*DeleteNetworkingV2PortsRequest{
    return &DeleteNetworkingV2PortsRequest{}
}

//response struct for the DeleteNetworkingV2Ports
type DeleteNetworkingV2PortsResponse struct{
    DeleteResult ports.DeleteResult
}

func NewDeleteNetworkingV2PortsResponse(deleteResult ports.DeleteResult,)*DeleteNetworkingV2PortsResponse {
    return &DeleteNetworkingV2PortsResponse{
            DeleteResult:deleteResult,
    }
}

// action function
func (oc *OpenstackClient) DeleteNetworkingV2Ports(req *DeleteNetworkingV2PortsRequest)(*DeleteNetworkingV2PortsResponse){
    return NewDeleteNetworkingV2PortsResponse(ports.Delete(oc.Client,req.Id, ))

}