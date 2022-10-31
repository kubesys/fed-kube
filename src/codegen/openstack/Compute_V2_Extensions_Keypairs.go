package openstack

// Code generated by cloud manager.





import (
    "github.com/gophercloud/gophercloud/openstack/compute/v2/extensions/keypairs"
    "github.com/gophercloud/gophercloud/pagination"
)
//request struct for the ListComputeV2ExtensionsKeypairs
type ListComputeV2ExtensionsKeypairsRequest struct{
    Opts keypairs.ListOpts
}

func NewListComputeV2ExtensionsKeypairsRequest()*ListComputeV2ExtensionsKeypairsRequest{
    return &ListComputeV2ExtensionsKeypairsRequest{}
}

//response struct for the ListComputeV2ExtensionsKeypairs
type ListComputeV2ExtensionsKeypairsResponse struct{
    Pager pagination.Pager
}

func NewListComputeV2ExtensionsKeypairsResponse(pager pagination.Pager,)*ListComputeV2ExtensionsKeypairsResponse {
    return &ListComputeV2ExtensionsKeypairsResponse{
            Pager:pager,
    }
}

// action function
func (oc *OpenstackClient) ListComputeV2ExtensionsKeypairs(req *ListComputeV2ExtensionsKeypairsRequest)(*ListComputeV2ExtensionsKeypairsResponse){
    return NewListComputeV2ExtensionsKeypairsResponse(keypairs.List(oc.Client,req.Opts, ))

}
//request struct for the CreateComputeV2ExtensionsKeypairs
type CreateComputeV2ExtensionsKeypairsRequest struct{
    Opts keypairs.CreateOpts
}

func NewCreateComputeV2ExtensionsKeypairsRequest()*CreateComputeV2ExtensionsKeypairsRequest{
    return &CreateComputeV2ExtensionsKeypairsRequest{}
}

//response struct for the CreateComputeV2ExtensionsKeypairs
type CreateComputeV2ExtensionsKeypairsResponse struct{
    CreateResult keypairs.CreateResult
}

func NewCreateComputeV2ExtensionsKeypairsResponse(createResult keypairs.CreateResult,)*CreateComputeV2ExtensionsKeypairsResponse {
    return &CreateComputeV2ExtensionsKeypairsResponse{
            CreateResult:createResult,
    }
}

// action function
func (oc *OpenstackClient) CreateComputeV2ExtensionsKeypairs(req *CreateComputeV2ExtensionsKeypairsRequest)(*CreateComputeV2ExtensionsKeypairsResponse){
    return NewCreateComputeV2ExtensionsKeypairsResponse(keypairs.Create(oc.Client,req.Opts, ))

}
//request struct for the GetComputeV2ExtensionsKeypairs
type GetComputeV2ExtensionsKeypairsRequest struct{
    Name string
    Opts keypairs.GetOpts
}

func NewGetComputeV2ExtensionsKeypairsRequest()*GetComputeV2ExtensionsKeypairsRequest{
    return &GetComputeV2ExtensionsKeypairsRequest{}
}

//response struct for the GetComputeV2ExtensionsKeypairs
type GetComputeV2ExtensionsKeypairsResponse struct{
    GetResult keypairs.GetResult
}

func NewGetComputeV2ExtensionsKeypairsResponse(getResult keypairs.GetResult,)*GetComputeV2ExtensionsKeypairsResponse {
    return &GetComputeV2ExtensionsKeypairsResponse{
            GetResult:getResult,
    }
}

// action function
func (oc *OpenstackClient) GetComputeV2ExtensionsKeypairs(req *GetComputeV2ExtensionsKeypairsRequest)(*GetComputeV2ExtensionsKeypairsResponse){
    return NewGetComputeV2ExtensionsKeypairsResponse(keypairs.Get(oc.Client,req.Name,req.Opts, ))

}
//request struct for the DeleteComputeV2ExtensionsKeypairs
type DeleteComputeV2ExtensionsKeypairsRequest struct{
    Name string
    Opts keypairs.DeleteOpts
}

func NewDeleteComputeV2ExtensionsKeypairsRequest()*DeleteComputeV2ExtensionsKeypairsRequest{
    return &DeleteComputeV2ExtensionsKeypairsRequest{}
}

//response struct for the DeleteComputeV2ExtensionsKeypairs
type DeleteComputeV2ExtensionsKeypairsResponse struct{
    DeleteResult keypairs.DeleteResult
}

func NewDeleteComputeV2ExtensionsKeypairsResponse(deleteResult keypairs.DeleteResult,)*DeleteComputeV2ExtensionsKeypairsResponse {
    return &DeleteComputeV2ExtensionsKeypairsResponse{
            DeleteResult:deleteResult,
    }
}

// action function
func (oc *OpenstackClient) DeleteComputeV2ExtensionsKeypairs(req *DeleteComputeV2ExtensionsKeypairsRequest)(*DeleteComputeV2ExtensionsKeypairsResponse){
    return NewDeleteComputeV2ExtensionsKeypairsResponse(keypairs.Delete(oc.Client,req.Name,req.Opts, ))

}