package openstack

// Code generated by cloud manager.





import (
    "github.com/gophercloud/gophercloud/openstack/compute/v2/extensions/volumeattach"
    "github.com/gophercloud/gophercloud/pagination"
)
//request struct for the ListComputeV2ExtensionsVolumeattach
type ListComputeV2ExtensionsVolumeattachRequest struct{
    ServerID string
}

func NewListComputeV2ExtensionsVolumeattachRequest()*ListComputeV2ExtensionsVolumeattachRequest{
    return &ListComputeV2ExtensionsVolumeattachRequest{}
}

//response struct for the ListComputeV2ExtensionsVolumeattach
type ListComputeV2ExtensionsVolumeattachResponse struct{
    Pager pagination.Pager
}

func NewListComputeV2ExtensionsVolumeattachResponse(pager pagination.Pager,)*ListComputeV2ExtensionsVolumeattachResponse {
    return &ListComputeV2ExtensionsVolumeattachResponse{
            Pager:pager,
    }
}

// action function
func (oc *OpenstackClient) ListComputeV2ExtensionsVolumeattach(req *ListComputeV2ExtensionsVolumeattachRequest)(*ListComputeV2ExtensionsVolumeattachResponse){
    return NewListComputeV2ExtensionsVolumeattachResponse(volumeattach.List(oc.Client,req.ServerID, ))

}
//request struct for the CreateComputeV2ExtensionsVolumeattach
type CreateComputeV2ExtensionsVolumeattachRequest struct{
    ServerID string
    Opts volumeattach.CreateOpts
}

func NewCreateComputeV2ExtensionsVolumeattachRequest()*CreateComputeV2ExtensionsVolumeattachRequest{
    return &CreateComputeV2ExtensionsVolumeattachRequest{}
}

//response struct for the CreateComputeV2ExtensionsVolumeattach
type CreateComputeV2ExtensionsVolumeattachResponse struct{
    CreateResult volumeattach.CreateResult
}

func NewCreateComputeV2ExtensionsVolumeattachResponse(createResult volumeattach.CreateResult,)*CreateComputeV2ExtensionsVolumeattachResponse {
    return &CreateComputeV2ExtensionsVolumeattachResponse{
            CreateResult:createResult,
    }
}

// action function
func (oc *OpenstackClient) CreateComputeV2ExtensionsVolumeattach(req *CreateComputeV2ExtensionsVolumeattachRequest)(*CreateComputeV2ExtensionsVolumeattachResponse){
    return NewCreateComputeV2ExtensionsVolumeattachResponse(volumeattach.Create(oc.Client,req.ServerID,req.Opts, ))

}
//request struct for the GetComputeV2ExtensionsVolumeattach
type GetComputeV2ExtensionsVolumeattachRequest struct{
    ServerID string
    AttachmentID string
}

func NewGetComputeV2ExtensionsVolumeattachRequest()*GetComputeV2ExtensionsVolumeattachRequest{
    return &GetComputeV2ExtensionsVolumeattachRequest{}
}

//response struct for the GetComputeV2ExtensionsVolumeattach
type GetComputeV2ExtensionsVolumeattachResponse struct{
    GetResult volumeattach.GetResult
}

func NewGetComputeV2ExtensionsVolumeattachResponse(getResult volumeattach.GetResult,)*GetComputeV2ExtensionsVolumeattachResponse {
    return &GetComputeV2ExtensionsVolumeattachResponse{
            GetResult:getResult,
    }
}

// action function
func (oc *OpenstackClient) GetComputeV2ExtensionsVolumeattach(req *GetComputeV2ExtensionsVolumeattachRequest)(*GetComputeV2ExtensionsVolumeattachResponse){
    return NewGetComputeV2ExtensionsVolumeattachResponse(volumeattach.Get(oc.Client,req.ServerID,req.AttachmentID, ))

}
//request struct for the DeleteComputeV2ExtensionsVolumeattach
type DeleteComputeV2ExtensionsVolumeattachRequest struct{
    ServerID string
    AttachmentID string
}

func NewDeleteComputeV2ExtensionsVolumeattachRequest()*DeleteComputeV2ExtensionsVolumeattachRequest{
    return &DeleteComputeV2ExtensionsVolumeattachRequest{}
}

//response struct for the DeleteComputeV2ExtensionsVolumeattach
type DeleteComputeV2ExtensionsVolumeattachResponse struct{
    DeleteResult volumeattach.DeleteResult
}

func NewDeleteComputeV2ExtensionsVolumeattachResponse(deleteResult volumeattach.DeleteResult,)*DeleteComputeV2ExtensionsVolumeattachResponse {
    return &DeleteComputeV2ExtensionsVolumeattachResponse{
            DeleteResult:deleteResult,
    }
}

// action function
func (oc *OpenstackClient) DeleteComputeV2ExtensionsVolumeattach(req *DeleteComputeV2ExtensionsVolumeattachRequest)(*DeleteComputeV2ExtensionsVolumeattachResponse){
    return NewDeleteComputeV2ExtensionsVolumeattachResponse(volumeattach.Delete(oc.Client,req.ServerID,req.AttachmentID, ))

}