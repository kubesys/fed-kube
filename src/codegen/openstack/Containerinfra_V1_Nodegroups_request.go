package openstack

// Code generated by cloud manager.

import (
	"github.com/gophercloud/gophercloud/openstack/containerinfra/v1/nodegroups"
	"github.com/gophercloud/gophercloud/pagination"
)

// request struct for the GetContainerinfraV1Nodegroups
type GetContainerinfraV1NodegroupsRequest struct {
	ClusterID   string
	NodeGroupID string
}

func NewGetContainerinfraV1NodegroupsRequest() *GetContainerinfraV1NodegroupsRequest {
	return &GetContainerinfraV1NodegroupsRequest{}
}

// response struct for the GetContainerinfraV1Nodegroups
type GetContainerinfraV1NodegroupsResponse struct {
	GetResult nodegroups.GetResult
}

func NewGetContainerinfraV1NodegroupsResponse(getResult nodegroups.GetResult) *GetContainerinfraV1NodegroupsResponse {
	return &GetContainerinfraV1NodegroupsResponse{
		GetResult: getResult,
	}
}

// action function
func (oc *OpenstackClient) GetContainerinfraV1Nodegroups(req *GetContainerinfraV1NodegroupsRequest) *GetContainerinfraV1NodegroupsResponse {
	return NewGetContainerinfraV1NodegroupsResponse(nodegroups.Get(oc.Client, req.ClusterID, req.NodeGroupID))

}

// request struct for the ListContainerinfraV1Nodegroups
type ListContainerinfraV1NodegroupsRequest struct {
	ClusterID string
	Opts      nodegroups.ListOpts
}

func NewListContainerinfraV1NodegroupsRequest() *ListContainerinfraV1NodegroupsRequest {
	return &ListContainerinfraV1NodegroupsRequest{}
}

// response struct for the ListContainerinfraV1Nodegroups
type ListContainerinfraV1NodegroupsResponse struct {
	Pager pagination.Pager
}

func NewListContainerinfraV1NodegroupsResponse(pager pagination.Pager) *ListContainerinfraV1NodegroupsResponse {
	return &ListContainerinfraV1NodegroupsResponse{
		Pager: pager,
	}
}

// action function
func (oc *OpenstackClient) ListContainerinfraV1Nodegroups(req *ListContainerinfraV1NodegroupsRequest) *ListContainerinfraV1NodegroupsResponse {
	return NewListContainerinfraV1NodegroupsResponse(nodegroups.List(oc.Client, req.ClusterID, req.Opts))

}

// request struct for the CreateContainerinfraV1Nodegroups
type CreateContainerinfraV1NodegroupsRequest struct {
	ClusterID string
	Opts      nodegroups.CreateOpts
}

func NewCreateContainerinfraV1NodegroupsRequest() *CreateContainerinfraV1NodegroupsRequest {
	return &CreateContainerinfraV1NodegroupsRequest{}
}

// response struct for the CreateContainerinfraV1Nodegroups
type CreateContainerinfraV1NodegroupsResponse struct {
	CreateResult nodegroups.CreateResult
}

func NewCreateContainerinfraV1NodegroupsResponse(createResult nodegroups.CreateResult) *CreateContainerinfraV1NodegroupsResponse {
	return &CreateContainerinfraV1NodegroupsResponse{
		CreateResult: createResult,
	}
}

// action function
func (oc *OpenstackClient) CreateContainerinfraV1Nodegroups(req *CreateContainerinfraV1NodegroupsRequest) *CreateContainerinfraV1NodegroupsResponse {
	return NewCreateContainerinfraV1NodegroupsResponse(nodegroups.Create(oc.Client, req.ClusterID, req.Opts))

}

// request struct for the DeleteContainerinfraV1Nodegroups
type DeleteContainerinfraV1NodegroupsRequest struct {
	ClusterID   string
	NodeGroupID string
}

func NewDeleteContainerinfraV1NodegroupsRequest() *DeleteContainerinfraV1NodegroupsRequest {
	return &DeleteContainerinfraV1NodegroupsRequest{}
}

// response struct for the DeleteContainerinfraV1Nodegroups
type DeleteContainerinfraV1NodegroupsResponse struct {
	DeleteResult nodegroups.DeleteResult
}

func NewDeleteContainerinfraV1NodegroupsResponse(deleteResult nodegroups.DeleteResult) *DeleteContainerinfraV1NodegroupsResponse {
	return &DeleteContainerinfraV1NodegroupsResponse{
		DeleteResult: deleteResult,
	}
}

// action function
func (oc *OpenstackClient) DeleteContainerinfraV1Nodegroups(req *DeleteContainerinfraV1NodegroupsRequest) *DeleteContainerinfraV1NodegroupsResponse {
	return NewDeleteContainerinfraV1NodegroupsResponse(nodegroups.Delete(oc.Client, req.ClusterID, req.NodeGroupID))

}
