package openstack

// Code generated by cloud manager.

import (
	"github.com/gophercloud/gophercloud/openstack/baremetal/v1/allocations"
	"github.com/gophercloud/gophercloud/pagination"
)

// request struct for the CreateBaremetalV1Allocations
type CreateBaremetalV1AllocationsRequest struct {
	Opts allocations.CreateOpts
}

func NewCreateBaremetalV1AllocationsRequest() *CreateBaremetalV1AllocationsRequest {
	return &CreateBaremetalV1AllocationsRequest{}
}

// response struct for the CreateBaremetalV1Allocations
type CreateBaremetalV1AllocationsResponse struct {
	CreateResult allocations.CreateResult
}

func NewCreateBaremetalV1AllocationsResponse(createResult allocations.CreateResult) *CreateBaremetalV1AllocationsResponse {
	return &CreateBaremetalV1AllocationsResponse{
		CreateResult: createResult,
	}
}

// action function
func (oc *OpenstackClient) CreateBaremetalV1Allocations(req *CreateBaremetalV1AllocationsRequest) *CreateBaremetalV1AllocationsResponse {
	return NewCreateBaremetalV1AllocationsResponse(allocations.Create(oc.Client, req.Opts))

}

// request struct for the ListBaremetalV1Allocations
type ListBaremetalV1AllocationsRequest struct {
	Opts allocations.ListOpts
}

func NewListBaremetalV1AllocationsRequest() *ListBaremetalV1AllocationsRequest {
	return &ListBaremetalV1AllocationsRequest{}
}

// response struct for the ListBaremetalV1Allocations
type ListBaremetalV1AllocationsResponse struct {
	Pager pagination.Pager
}

func NewListBaremetalV1AllocationsResponse(pager pagination.Pager) *ListBaremetalV1AllocationsResponse {
	return &ListBaremetalV1AllocationsResponse{
		Pager: pager,
	}
}

// action function
func (oc *OpenstackClient) ListBaremetalV1Allocations(req *ListBaremetalV1AllocationsRequest) *ListBaremetalV1AllocationsResponse {
	return NewListBaremetalV1AllocationsResponse(allocations.List(oc.Client, req.Opts))

}

// request struct for the GetBaremetalV1Allocations
type GetBaremetalV1AllocationsRequest struct {
	Id string
}

func NewGetBaremetalV1AllocationsRequest() *GetBaremetalV1AllocationsRequest {
	return &GetBaremetalV1AllocationsRequest{}
}

// response struct for the GetBaremetalV1Allocations
type GetBaremetalV1AllocationsResponse struct {
	GetResult allocations.GetResult
}

func NewGetBaremetalV1AllocationsResponse(getResult allocations.GetResult) *GetBaremetalV1AllocationsResponse {
	return &GetBaremetalV1AllocationsResponse{
		GetResult: getResult,
	}
}

// action function
func (oc *OpenstackClient) GetBaremetalV1Allocations(req *GetBaremetalV1AllocationsRequest) *GetBaremetalV1AllocationsResponse {
	return NewGetBaremetalV1AllocationsResponse(allocations.Get(oc.Client, req.Id))

}

// request struct for the DeleteBaremetalV1Allocations
type DeleteBaremetalV1AllocationsRequest struct {
	Id string
}

func NewDeleteBaremetalV1AllocationsRequest() *DeleteBaremetalV1AllocationsRequest {
	return &DeleteBaremetalV1AllocationsRequest{}
}

// response struct for the DeleteBaremetalV1Allocations
type DeleteBaremetalV1AllocationsResponse struct {
	DeleteResult allocations.DeleteResult
}

func NewDeleteBaremetalV1AllocationsResponse(deleteResult allocations.DeleteResult) *DeleteBaremetalV1AllocationsResponse {
	return &DeleteBaremetalV1AllocationsResponse{
		DeleteResult: deleteResult,
	}
}

// action function
func (oc *OpenstackClient) DeleteBaremetalV1Allocations(req *DeleteBaremetalV1AllocationsRequest) *DeleteBaremetalV1AllocationsResponse {
	return NewDeleteBaremetalV1AllocationsResponse(allocations.Delete(oc.Client, req.Id))

}
