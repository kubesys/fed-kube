package openstack

// Code generated by cloud manager.

import (
	"github.com/gophercloud/gophercloud/openstack/db/v1/flavors"
	"github.com/gophercloud/gophercloud/pagination"
)

// request struct for the ListDbV1Flavors
type ListDbV1FlavorsRequest struct {
}

func NewListDbV1FlavorsRequest() *ListDbV1FlavorsRequest {
	return &ListDbV1FlavorsRequest{}
}

// response struct for the ListDbV1Flavors
type ListDbV1FlavorsResponse struct {
	Pager pagination.Pager
}

func NewListDbV1FlavorsResponse(pager pagination.Pager) *ListDbV1FlavorsResponse {
	return &ListDbV1FlavorsResponse{
		Pager: pager,
	}
}

// action function
func (oc *OpenstackClient) ListDbV1Flavors(req *ListDbV1FlavorsRequest) *ListDbV1FlavorsResponse {
	return NewListDbV1FlavorsResponse(flavors.List(oc.Client))

}

// request struct for the GetDbV1Flavors
type GetDbV1FlavorsRequest struct {
	Id string
}

func NewGetDbV1FlavorsRequest() *GetDbV1FlavorsRequest {
	return &GetDbV1FlavorsRequest{}
}

// response struct for the GetDbV1Flavors
type GetDbV1FlavorsResponse struct {
	GetResult flavors.GetResult
}

func NewGetDbV1FlavorsResponse(getResult flavors.GetResult) *GetDbV1FlavorsResponse {
	return &GetDbV1FlavorsResponse{
		GetResult: getResult,
	}
}

// action function
func (oc *OpenstackClient) GetDbV1Flavors(req *GetDbV1FlavorsRequest) *GetDbV1FlavorsResponse {
	return NewGetDbV1FlavorsResponse(flavors.Get(oc.Client, req.Id))

}
