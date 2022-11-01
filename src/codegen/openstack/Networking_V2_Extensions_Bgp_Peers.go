package openstack

// Code generated by cloud manager.





import (
    "github.com/gophercloud/gophercloud/openstack/networking/v2/extensions/bgp/peers"
    "github.com/gophercloud/gophercloud/pagination"
)
//request struct for the ListNetworkingV2ExtensionsBgpPeers
type ListNetworkingV2ExtensionsBgpPeersRequest struct{
}

func NewListNetworkingV2ExtensionsBgpPeersRequest()*ListNetworkingV2ExtensionsBgpPeersRequest{
    return &ListNetworkingV2ExtensionsBgpPeersRequest{}
}

//response struct for the ListNetworkingV2ExtensionsBgpPeers
type ListNetworkingV2ExtensionsBgpPeersResponse struct{
    Pager pagination.Pager
}

func NewListNetworkingV2ExtensionsBgpPeersResponse(pager pagination.Pager,)*ListNetworkingV2ExtensionsBgpPeersResponse {
    return &ListNetworkingV2ExtensionsBgpPeersResponse{
            Pager:pager,
    }
}

// action function
func (oc *OpenstackClient) ListNetworkingV2ExtensionsBgpPeers(req *ListNetworkingV2ExtensionsBgpPeersRequest)(*ListNetworkingV2ExtensionsBgpPeersResponse){
    return NewListNetworkingV2ExtensionsBgpPeersResponse(peers.List(oc.Client, ))

}
//request struct for the GetNetworkingV2ExtensionsBgpPeers
type GetNetworkingV2ExtensionsBgpPeersRequest struct{
    Id string
}

func NewGetNetworkingV2ExtensionsBgpPeersRequest()*GetNetworkingV2ExtensionsBgpPeersRequest{
    return &GetNetworkingV2ExtensionsBgpPeersRequest{}
}

//response struct for the GetNetworkingV2ExtensionsBgpPeers
type GetNetworkingV2ExtensionsBgpPeersResponse struct{
    GetResult peers.GetResult
}

func NewGetNetworkingV2ExtensionsBgpPeersResponse(getResult peers.GetResult,)*GetNetworkingV2ExtensionsBgpPeersResponse {
    return &GetNetworkingV2ExtensionsBgpPeersResponse{
            GetResult:getResult,
    }
}

// action function
func (oc *OpenstackClient) GetNetworkingV2ExtensionsBgpPeers(req *GetNetworkingV2ExtensionsBgpPeersRequest)(*GetNetworkingV2ExtensionsBgpPeersResponse){
    return NewGetNetworkingV2ExtensionsBgpPeersResponse(peers.Get(oc.Client,req.Id, ))

}
//request struct for the CreateNetworkingV2ExtensionsBgpPeers
type CreateNetworkingV2ExtensionsBgpPeersRequest struct{
    Opts peers.CreateOpts
}

func NewCreateNetworkingV2ExtensionsBgpPeersRequest()*CreateNetworkingV2ExtensionsBgpPeersRequest{
    return &CreateNetworkingV2ExtensionsBgpPeersRequest{}
}

//response struct for the CreateNetworkingV2ExtensionsBgpPeers
type CreateNetworkingV2ExtensionsBgpPeersResponse struct{
    CreateResult peers.CreateResult
}

func NewCreateNetworkingV2ExtensionsBgpPeersResponse(createResult peers.CreateResult,)*CreateNetworkingV2ExtensionsBgpPeersResponse {
    return &CreateNetworkingV2ExtensionsBgpPeersResponse{
            CreateResult:createResult,
    }
}

// action function
func (oc *OpenstackClient) CreateNetworkingV2ExtensionsBgpPeers(req *CreateNetworkingV2ExtensionsBgpPeersRequest)(*CreateNetworkingV2ExtensionsBgpPeersResponse){
    return NewCreateNetworkingV2ExtensionsBgpPeersResponse(peers.Create(oc.Client,req.Opts, ))

}
//request struct for the DeleteNetworkingV2ExtensionsBgpPeers
type DeleteNetworkingV2ExtensionsBgpPeersRequest struct{
    BgpPeerID string
}

func NewDeleteNetworkingV2ExtensionsBgpPeersRequest()*DeleteNetworkingV2ExtensionsBgpPeersRequest{
    return &DeleteNetworkingV2ExtensionsBgpPeersRequest{}
}

//response struct for the DeleteNetworkingV2ExtensionsBgpPeers
type DeleteNetworkingV2ExtensionsBgpPeersResponse struct{
    DeleteResult peers.DeleteResult
}

func NewDeleteNetworkingV2ExtensionsBgpPeersResponse(deleteResult peers.DeleteResult,)*DeleteNetworkingV2ExtensionsBgpPeersResponse {
    return &DeleteNetworkingV2ExtensionsBgpPeersResponse{
            DeleteResult:deleteResult,
    }
}

// action function
func (oc *OpenstackClient) DeleteNetworkingV2ExtensionsBgpPeers(req *DeleteNetworkingV2ExtensionsBgpPeersRequest)(*DeleteNetworkingV2ExtensionsBgpPeersResponse){
    return NewDeleteNetworkingV2ExtensionsBgpPeersResponse(peers.Delete(oc.Client,req.BgpPeerID, ))

}
//request struct for the UpdateNetworkingV2ExtensionsBgpPeers
type UpdateNetworkingV2ExtensionsBgpPeersRequest struct{
    BgpPeerID string
    Opts peers.UpdateOpts
}

func NewUpdateNetworkingV2ExtensionsBgpPeersRequest()*UpdateNetworkingV2ExtensionsBgpPeersRequest{
    return &UpdateNetworkingV2ExtensionsBgpPeersRequest{}
}

//response struct for the UpdateNetworkingV2ExtensionsBgpPeers
type UpdateNetworkingV2ExtensionsBgpPeersResponse struct{
    UpdateResult peers.UpdateResult
}

func NewUpdateNetworkingV2ExtensionsBgpPeersResponse(updateResult peers.UpdateResult,)*UpdateNetworkingV2ExtensionsBgpPeersResponse {
    return &UpdateNetworkingV2ExtensionsBgpPeersResponse{
            UpdateResult:updateResult,
    }
}

// action function
func (oc *OpenstackClient) UpdateNetworkingV2ExtensionsBgpPeers(req *UpdateNetworkingV2ExtensionsBgpPeersRequest)(*UpdateNetworkingV2ExtensionsBgpPeersResponse){
    return NewUpdateNetworkingV2ExtensionsBgpPeersResponse(peers.Update(oc.Client,req.BgpPeerID,req.Opts, ))

}