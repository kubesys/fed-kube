package openstack

// Code generated by cloud manager.

import (
    "github.com/gophercloud/gophercloud/openstack/containerinfra/apiversions"
    "github.com/gophercloud/gophercloud/pagination"
)
//request struct for the ListContainerinfraApiversions
type ListContainerinfraApiversionsRequest struct{
}

func NewListContainerinfraApiversionsRequest()*ListContainerinfraApiversionsRequest{
    return &ListContainerinfraApiversionsRequest{}
}

//response struct for the ListContainerinfraApiversions
type ListContainerinfraApiversionsResponse struct{
    Pager pagination.Pager
}

func NewListContainerinfraApiversionsResponse(pager pagination.Pager,)*ListContainerinfraApiversionsResponse {
    return &ListContainerinfraApiversionsResponse{
            Pager:pager,
    }
}

// action function
func (oc *OpenstackClient) ListContainerinfraApiversions(req *ListContainerinfraApiversionsRequest)(*ListContainerinfraApiversionsResponse){
    return NewListContainerinfraApiversionsResponse(apiversions.List(oc.Client, ))

}
//request struct for the GetContainerinfraApiversions
type GetContainerinfraApiversionsRequest struct{
    V string
}

func NewGetContainerinfraApiversionsRequest()*GetContainerinfraApiversionsRequest{
    return &GetContainerinfraApiversionsRequest{}
}

//response struct for the GetContainerinfraApiversions
type GetContainerinfraApiversionsResponse struct{
    GetResult apiversions.GetResult
}

func NewGetContainerinfraApiversionsResponse(getResult apiversions.GetResult,)*GetContainerinfraApiversionsResponse {
    return &GetContainerinfraApiversionsResponse{
            GetResult:getResult,
    }
}

// action function
func (oc *OpenstackClient) GetContainerinfraApiversions(req *GetContainerinfraApiversionsRequest)(*GetContainerinfraApiversionsResponse){
    return NewGetContainerinfraApiversionsResponse(apiversions.Get(oc.Client,req.V, ))

}