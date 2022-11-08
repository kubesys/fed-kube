package openstack

// Code generated by cloud manager.

import (
	"github.com/gophercloud/gophercloud/openstack/identity/v3/endpoints"
	"github.com/gophercloud/gophercloud/pagination"
)

// request struct for the CreateIdentityV3Endpoints
type CreateIdentityV3EndpointsRequest struct {
	Opts endpoints.CreateOpts
}

func NewCreateIdentityV3EndpointsRequest() *CreateIdentityV3EndpointsRequest {
	return &CreateIdentityV3EndpointsRequest{}
}

// response struct for the CreateIdentityV3Endpoints
type CreateIdentityV3EndpointsResponse struct {
	CreateResult endpoints.CreateResult
}

func NewCreateIdentityV3EndpointsResponse(createResult endpoints.CreateResult) *CreateIdentityV3EndpointsResponse {
	return &CreateIdentityV3EndpointsResponse{
		CreateResult: createResult,
	}
}

// action function
func (oc *OpenstackClient) CreateIdentityV3Endpoints(req *CreateIdentityV3EndpointsRequest) *CreateIdentityV3EndpointsResponse {
	return NewCreateIdentityV3EndpointsResponse(endpoints.Create(oc.Client, req.Opts))

}

// request struct for the ListIdentityV3Endpoints
type ListIdentityV3EndpointsRequest struct {
	Opts endpoints.ListOpts
}

func NewListIdentityV3EndpointsRequest() *ListIdentityV3EndpointsRequest {
	return &ListIdentityV3EndpointsRequest{}
}

// response struct for the ListIdentityV3Endpoints
type ListIdentityV3EndpointsResponse struct {
	Pager pagination.Pager
}

func NewListIdentityV3EndpointsResponse(pager pagination.Pager) *ListIdentityV3EndpointsResponse {
	return &ListIdentityV3EndpointsResponse{
		Pager: pager,
	}
}

// action function
func (oc *OpenstackClient) ListIdentityV3Endpoints(req *ListIdentityV3EndpointsRequest) *ListIdentityV3EndpointsResponse {
	return NewListIdentityV3EndpointsResponse(endpoints.List(oc.Client, req.Opts))

}

// request struct for the UpdateIdentityV3Endpoints
type UpdateIdentityV3EndpointsRequest struct {
	EndpointID string
	Opts       endpoints.UpdateOpts
}

func NewUpdateIdentityV3EndpointsRequest() *UpdateIdentityV3EndpointsRequest {
	return &UpdateIdentityV3EndpointsRequest{}
}

// response struct for the UpdateIdentityV3Endpoints
type UpdateIdentityV3EndpointsResponse struct {
	UpdateResult endpoints.UpdateResult
}

func NewUpdateIdentityV3EndpointsResponse(updateResult endpoints.UpdateResult) *UpdateIdentityV3EndpointsResponse {
	return &UpdateIdentityV3EndpointsResponse{
		UpdateResult: updateResult,
	}
}

// action function
func (oc *OpenstackClient) UpdateIdentityV3Endpoints(req *UpdateIdentityV3EndpointsRequest) *UpdateIdentityV3EndpointsResponse {
	return NewUpdateIdentityV3EndpointsResponse(endpoints.Update(oc.Client, req.EndpointID, req.Opts))

}

// request struct for the DeleteIdentityV3Endpoints
type DeleteIdentityV3EndpointsRequest struct {
	EndpointID string
}

func NewDeleteIdentityV3EndpointsRequest() *DeleteIdentityV3EndpointsRequest {
	return &DeleteIdentityV3EndpointsRequest{}
}

// response struct for the DeleteIdentityV3Endpoints
type DeleteIdentityV3EndpointsResponse struct {
	DeleteResult endpoints.DeleteResult
}

func NewDeleteIdentityV3EndpointsResponse(deleteResult endpoints.DeleteResult) *DeleteIdentityV3EndpointsResponse {
	return &DeleteIdentityV3EndpointsResponse{
		DeleteResult: deleteResult,
	}
}

// action function
func (oc *OpenstackClient) DeleteIdentityV3Endpoints(req *DeleteIdentityV3EndpointsRequest) *DeleteIdentityV3EndpointsResponse {
	return NewDeleteIdentityV3EndpointsResponse(endpoints.Delete(oc.Client, req.EndpointID))

}
