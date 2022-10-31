package openstack

// Code generated by cloud manager.





import (
    "github.com/gophercloud/gophercloud/openstack/networking/v2/extensions/layer3/floatingips"
    "github.com/gophercloud/gophercloud/pagination"
)
//request struct for the ListNetworkingV2ExtensionsLayer3Floatingips
type ListNetworkingV2ExtensionsLayer3FloatingipsRequest struct{
    Opts floatingips.ListOpts
}

func NewListNetworkingV2ExtensionsLayer3FloatingipsRequest()*ListNetworkingV2ExtensionsLayer3FloatingipsRequest{
    return &ListNetworkingV2ExtensionsLayer3FloatingipsRequest{}
}

//response struct for the ListNetworkingV2ExtensionsLayer3Floatingips
type ListNetworkingV2ExtensionsLayer3FloatingipsResponse struct{
    Pager pagination.Pager
}

func NewListNetworkingV2ExtensionsLayer3FloatingipsResponse(pager pagination.Pager,)*ListNetworkingV2ExtensionsLayer3FloatingipsResponse {
    return &ListNetworkingV2ExtensionsLayer3FloatingipsResponse{
            Pager:pager,
    }
}

// action function
func (oc *OpenstackClient) ListNetworkingV2ExtensionsLayer3Floatingips(req *ListNetworkingV2ExtensionsLayer3FloatingipsRequest)(*ListNetworkingV2ExtensionsLayer3FloatingipsResponse){
    return NewListNetworkingV2ExtensionsLayer3FloatingipsResponse(floatingips.List(oc.Client,req.Opts, ))

}
//request struct for the CreateNetworkingV2ExtensionsLayer3Floatingips
type CreateNetworkingV2ExtensionsLayer3FloatingipsRequest struct{
    Opts floatingips.CreateOpts
}

func NewCreateNetworkingV2ExtensionsLayer3FloatingipsRequest()*CreateNetworkingV2ExtensionsLayer3FloatingipsRequest{
    return &CreateNetworkingV2ExtensionsLayer3FloatingipsRequest{}
}

//response struct for the CreateNetworkingV2ExtensionsLayer3Floatingips
type CreateNetworkingV2ExtensionsLayer3FloatingipsResponse struct{
    CreateResult floatingips.CreateResult
}

func NewCreateNetworkingV2ExtensionsLayer3FloatingipsResponse(createResult floatingips.CreateResult,)*CreateNetworkingV2ExtensionsLayer3FloatingipsResponse {
    return &CreateNetworkingV2ExtensionsLayer3FloatingipsResponse{
            CreateResult:createResult,
    }
}

// action function
func (oc *OpenstackClient) CreateNetworkingV2ExtensionsLayer3Floatingips(req *CreateNetworkingV2ExtensionsLayer3FloatingipsRequest)(*CreateNetworkingV2ExtensionsLayer3FloatingipsResponse){
    return NewCreateNetworkingV2ExtensionsLayer3FloatingipsResponse(floatingips.Create(oc.Client,req.Opts, ))

}
//request struct for the GetNetworkingV2ExtensionsLayer3Floatingips
type GetNetworkingV2ExtensionsLayer3FloatingipsRequest struct{
    Id string
}

func NewGetNetworkingV2ExtensionsLayer3FloatingipsRequest()*GetNetworkingV2ExtensionsLayer3FloatingipsRequest{
    return &GetNetworkingV2ExtensionsLayer3FloatingipsRequest{}
}

//response struct for the GetNetworkingV2ExtensionsLayer3Floatingips
type GetNetworkingV2ExtensionsLayer3FloatingipsResponse struct{
    GetResult floatingips.GetResult
}

func NewGetNetworkingV2ExtensionsLayer3FloatingipsResponse(getResult floatingips.GetResult,)*GetNetworkingV2ExtensionsLayer3FloatingipsResponse {
    return &GetNetworkingV2ExtensionsLayer3FloatingipsResponse{
            GetResult:getResult,
    }
}

// action function
func (oc *OpenstackClient) GetNetworkingV2ExtensionsLayer3Floatingips(req *GetNetworkingV2ExtensionsLayer3FloatingipsRequest)(*GetNetworkingV2ExtensionsLayer3FloatingipsResponse){
    return NewGetNetworkingV2ExtensionsLayer3FloatingipsResponse(floatingips.Get(oc.Client,req.Id, ))

}
//request struct for the UpdateNetworkingV2ExtensionsLayer3Floatingips
type UpdateNetworkingV2ExtensionsLayer3FloatingipsRequest struct{
    Id string
    Opts floatingips.UpdateOpts
}

func NewUpdateNetworkingV2ExtensionsLayer3FloatingipsRequest()*UpdateNetworkingV2ExtensionsLayer3FloatingipsRequest{
    return &UpdateNetworkingV2ExtensionsLayer3FloatingipsRequest{}
}

//response struct for the UpdateNetworkingV2ExtensionsLayer3Floatingips
type UpdateNetworkingV2ExtensionsLayer3FloatingipsResponse struct{
    UpdateResult floatingips.UpdateResult
}

func NewUpdateNetworkingV2ExtensionsLayer3FloatingipsResponse(updateResult floatingips.UpdateResult,)*UpdateNetworkingV2ExtensionsLayer3FloatingipsResponse {
    return &UpdateNetworkingV2ExtensionsLayer3FloatingipsResponse{
            UpdateResult:updateResult,
    }
}

// action function
func (oc *OpenstackClient) UpdateNetworkingV2ExtensionsLayer3Floatingips(req *UpdateNetworkingV2ExtensionsLayer3FloatingipsRequest)(*UpdateNetworkingV2ExtensionsLayer3FloatingipsResponse){
    return NewUpdateNetworkingV2ExtensionsLayer3FloatingipsResponse(floatingips.Update(oc.Client,req.Id,req.Opts, ))

}
//request struct for the DeleteNetworkingV2ExtensionsLayer3Floatingips
type DeleteNetworkingV2ExtensionsLayer3FloatingipsRequest struct{
    Id string
}

func NewDeleteNetworkingV2ExtensionsLayer3FloatingipsRequest()*DeleteNetworkingV2ExtensionsLayer3FloatingipsRequest{
    return &DeleteNetworkingV2ExtensionsLayer3FloatingipsRequest{}
}

//response struct for the DeleteNetworkingV2ExtensionsLayer3Floatingips
type DeleteNetworkingV2ExtensionsLayer3FloatingipsResponse struct{
    DeleteResult floatingips.DeleteResult
}

func NewDeleteNetworkingV2ExtensionsLayer3FloatingipsResponse(deleteResult floatingips.DeleteResult,)*DeleteNetworkingV2ExtensionsLayer3FloatingipsResponse {
    return &DeleteNetworkingV2ExtensionsLayer3FloatingipsResponse{
            DeleteResult:deleteResult,
    }
}

// action function
func (oc *OpenstackClient) DeleteNetworkingV2ExtensionsLayer3Floatingips(req *DeleteNetworkingV2ExtensionsLayer3FloatingipsRequest)(*DeleteNetworkingV2ExtensionsLayer3FloatingipsResponse){
    return NewDeleteNetworkingV2ExtensionsLayer3FloatingipsResponse(floatingips.Delete(oc.Client,req.Id, ))

}