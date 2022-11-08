package openstack

// Code generated by cloud manager.

import (
    "github.com/gophercloud/gophercloud/openstack/compute/v2/extensions/servergroups"
    "github.com/gophercloud/gophercloud/pagination"
)
//request struct for the ListComputeV2ExtensionsServergroups
type ListComputeV2ExtensionsServergroupsRequest struct{
    Opts servergroups.ListOpts
}

func NewListComputeV2ExtensionsServergroupsRequest()*ListComputeV2ExtensionsServergroupsRequest{
    return &ListComputeV2ExtensionsServergroupsRequest{}
}

//response struct for the ListComputeV2ExtensionsServergroups
type ListComputeV2ExtensionsServergroupsResponse struct{
    Pager pagination.Pager
}

func NewListComputeV2ExtensionsServergroupsResponse(pager pagination.Pager,)*ListComputeV2ExtensionsServergroupsResponse {
    return &ListComputeV2ExtensionsServergroupsResponse{
            Pager:pager,
    }
}

// action function
func (oc *OpenstackClient) ListComputeV2ExtensionsServergroups(req *ListComputeV2ExtensionsServergroupsRequest)(*ListComputeV2ExtensionsServergroupsResponse){
    return NewListComputeV2ExtensionsServergroupsResponse(servergroups.List(oc.Client,req.Opts, ))

}
//request struct for the CreateComputeV2ExtensionsServergroups
type CreateComputeV2ExtensionsServergroupsRequest struct{
    Opts servergroups.CreateOpts
}

func NewCreateComputeV2ExtensionsServergroupsRequest()*CreateComputeV2ExtensionsServergroupsRequest{
    return &CreateComputeV2ExtensionsServergroupsRequest{}
}

//response struct for the CreateComputeV2ExtensionsServergroups
type CreateComputeV2ExtensionsServergroupsResponse struct{
    CreateResult servergroups.CreateResult
}

func NewCreateComputeV2ExtensionsServergroupsResponse(createResult servergroups.CreateResult,)*CreateComputeV2ExtensionsServergroupsResponse {
    return &CreateComputeV2ExtensionsServergroupsResponse{
            CreateResult:createResult,
    }
}

// action function
func (oc *OpenstackClient) CreateComputeV2ExtensionsServergroups(req *CreateComputeV2ExtensionsServergroupsRequest)(*CreateComputeV2ExtensionsServergroupsResponse){
    return NewCreateComputeV2ExtensionsServergroupsResponse(servergroups.Create(oc.Client,req.Opts, ))

}
//request struct for the GetComputeV2ExtensionsServergroups
type GetComputeV2ExtensionsServergroupsRequest struct{
    Id string
}

func NewGetComputeV2ExtensionsServergroupsRequest()*GetComputeV2ExtensionsServergroupsRequest{
    return &GetComputeV2ExtensionsServergroupsRequest{}
}

//response struct for the GetComputeV2ExtensionsServergroups
type GetComputeV2ExtensionsServergroupsResponse struct{
    GetResult servergroups.GetResult
}

func NewGetComputeV2ExtensionsServergroupsResponse(getResult servergroups.GetResult,)*GetComputeV2ExtensionsServergroupsResponse {
    return &GetComputeV2ExtensionsServergroupsResponse{
            GetResult:getResult,
    }
}

// action function
func (oc *OpenstackClient) GetComputeV2ExtensionsServergroups(req *GetComputeV2ExtensionsServergroupsRequest)(*GetComputeV2ExtensionsServergroupsResponse){
    return NewGetComputeV2ExtensionsServergroupsResponse(servergroups.Get(oc.Client,req.Id, ))

}
//request struct for the DeleteComputeV2ExtensionsServergroups
type DeleteComputeV2ExtensionsServergroupsRequest struct{
    Id string
}

func NewDeleteComputeV2ExtensionsServergroupsRequest()*DeleteComputeV2ExtensionsServergroupsRequest{
    return &DeleteComputeV2ExtensionsServergroupsRequest{}
}

//response struct for the DeleteComputeV2ExtensionsServergroups
type DeleteComputeV2ExtensionsServergroupsResponse struct{
    DeleteResult servergroups.DeleteResult
}

func NewDeleteComputeV2ExtensionsServergroupsResponse(deleteResult servergroups.DeleteResult,)*DeleteComputeV2ExtensionsServergroupsResponse {
    return &DeleteComputeV2ExtensionsServergroupsResponse{
            DeleteResult:deleteResult,
    }
}

// action function
func (oc *OpenstackClient) DeleteComputeV2ExtensionsServergroups(req *DeleteComputeV2ExtensionsServergroupsRequest)(*DeleteComputeV2ExtensionsServergroupsResponse){
    return NewDeleteComputeV2ExtensionsServergroupsResponse(servergroups.Delete(oc.Client,req.Id, ))

}