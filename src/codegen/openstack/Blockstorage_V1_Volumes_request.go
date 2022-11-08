package openstack

// Code generated by cloud manager.

import (
	"github.com/gophercloud/gophercloud/openstack/blockstorage/v1/volumes"
	"github.com/gophercloud/gophercloud/pagination"
)

// request struct for the CreateBlockstorageV1Volumes
type CreateBlockstorageV1VolumesRequest struct {
	Opts volumes.CreateOpts
}

func NewCreateBlockstorageV1VolumesRequest() *CreateBlockstorageV1VolumesRequest {
	return &CreateBlockstorageV1VolumesRequest{}
}

// response struct for the CreateBlockstorageV1Volumes
type CreateBlockstorageV1VolumesResponse struct {
	CreateResult volumes.CreateResult
}

func NewCreateBlockstorageV1VolumesResponse(createResult volumes.CreateResult) *CreateBlockstorageV1VolumesResponse {
	return &CreateBlockstorageV1VolumesResponse{
		CreateResult: createResult,
	}
}

// action function
func (oc *OpenstackClient) CreateBlockstorageV1Volumes(req *CreateBlockstorageV1VolumesRequest) *CreateBlockstorageV1VolumesResponse {
	return NewCreateBlockstorageV1VolumesResponse(volumes.Create(oc.Client, req.Opts))

}

// request struct for the DeleteBlockstorageV1Volumes
type DeleteBlockstorageV1VolumesRequest struct {
	Id string
}

func NewDeleteBlockstorageV1VolumesRequest() *DeleteBlockstorageV1VolumesRequest {
	return &DeleteBlockstorageV1VolumesRequest{}
}

// response struct for the DeleteBlockstorageV1Volumes
type DeleteBlockstorageV1VolumesResponse struct {
	DeleteResult volumes.DeleteResult
}

func NewDeleteBlockstorageV1VolumesResponse(deleteResult volumes.DeleteResult) *DeleteBlockstorageV1VolumesResponse {
	return &DeleteBlockstorageV1VolumesResponse{
		DeleteResult: deleteResult,
	}
}

// action function
func (oc *OpenstackClient) DeleteBlockstorageV1Volumes(req *DeleteBlockstorageV1VolumesRequest) *DeleteBlockstorageV1VolumesResponse {
	return NewDeleteBlockstorageV1VolumesResponse(volumes.Delete(oc.Client, req.Id))

}

// request struct for the GetBlockstorageV1Volumes
type GetBlockstorageV1VolumesRequest struct {
	Id string
}

func NewGetBlockstorageV1VolumesRequest() *GetBlockstorageV1VolumesRequest {
	return &GetBlockstorageV1VolumesRequest{}
}

// response struct for the GetBlockstorageV1Volumes
type GetBlockstorageV1VolumesResponse struct {
	GetResult volumes.GetResult
}

func NewGetBlockstorageV1VolumesResponse(getResult volumes.GetResult) *GetBlockstorageV1VolumesResponse {
	return &GetBlockstorageV1VolumesResponse{
		GetResult: getResult,
	}
}

// action function
func (oc *OpenstackClient) GetBlockstorageV1Volumes(req *GetBlockstorageV1VolumesRequest) *GetBlockstorageV1VolumesResponse {
	return NewGetBlockstorageV1VolumesResponse(volumes.Get(oc.Client, req.Id))

}

// request struct for the ListBlockstorageV1Volumes
type ListBlockstorageV1VolumesRequest struct {
	Opts volumes.ListOpts
}

func NewListBlockstorageV1VolumesRequest() *ListBlockstorageV1VolumesRequest {
	return &ListBlockstorageV1VolumesRequest{}
}

// response struct for the ListBlockstorageV1Volumes
type ListBlockstorageV1VolumesResponse struct {
	Pager pagination.Pager
}

func NewListBlockstorageV1VolumesResponse(pager pagination.Pager) *ListBlockstorageV1VolumesResponse {
	return &ListBlockstorageV1VolumesResponse{
		Pager: pager,
	}
}

// action function
func (oc *OpenstackClient) ListBlockstorageV1Volumes(req *ListBlockstorageV1VolumesRequest) *ListBlockstorageV1VolumesResponse {
	return NewListBlockstorageV1VolumesResponse(volumes.List(oc.Client, req.Opts))

}

// request struct for the UpdateBlockstorageV1Volumes
type UpdateBlockstorageV1VolumesRequest struct {
	Id   string
	Opts volumes.UpdateOpts
}

func NewUpdateBlockstorageV1VolumesRequest() *UpdateBlockstorageV1VolumesRequest {
	return &UpdateBlockstorageV1VolumesRequest{}
}

// response struct for the UpdateBlockstorageV1Volumes
type UpdateBlockstorageV1VolumesResponse struct {
	UpdateResult volumes.UpdateResult
}

func NewUpdateBlockstorageV1VolumesResponse(updateResult volumes.UpdateResult) *UpdateBlockstorageV1VolumesResponse {
	return &UpdateBlockstorageV1VolumesResponse{
		UpdateResult: updateResult,
	}
}

// action function
func (oc *OpenstackClient) UpdateBlockstorageV1Volumes(req *UpdateBlockstorageV1VolumesRequest) *UpdateBlockstorageV1VolumesResponse {
	return NewUpdateBlockstorageV1VolumesResponse(volumes.Update(oc.Client, req.Id, req.Opts))

}
