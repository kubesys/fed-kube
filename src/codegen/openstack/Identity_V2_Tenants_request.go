package openstack

// Code generated by cloud manager.

import (
	"github.com/gophercloud/gophercloud/openstack/identity/v2/tenants"
	"github.com/gophercloud/gophercloud/pagination"
)

// request struct for the ListIdentityV2Tenants
type ListIdentityV2TenantsRequest struct {
	Opts *tenants.ListOpts
}

func NewListIdentityV2TenantsRequest() *ListIdentityV2TenantsRequest {
	return &ListIdentityV2TenantsRequest{}
}

// response struct for the ListIdentityV2Tenants
type ListIdentityV2TenantsResponse struct {
	Pager pagination.Pager
}

func NewListIdentityV2TenantsResponse(pager pagination.Pager) *ListIdentityV2TenantsResponse {
	return &ListIdentityV2TenantsResponse{
		Pager: pager,
	}
}

// action function
func (oc *OpenstackClient) ListIdentityV2Tenants(req *ListIdentityV2TenantsRequest) *ListIdentityV2TenantsResponse {
	return NewListIdentityV2TenantsResponse(tenants.List(oc.Client, req.Opts))

}

// request struct for the CreateIdentityV2Tenants
type CreateIdentityV2TenantsRequest struct {
	Opts tenants.CreateOpts
}

func NewCreateIdentityV2TenantsRequest() *CreateIdentityV2TenantsRequest {
	return &CreateIdentityV2TenantsRequest{}
}

// response struct for the CreateIdentityV2Tenants
type CreateIdentityV2TenantsResponse struct {
	CreateResult tenants.CreateResult
}

func NewCreateIdentityV2TenantsResponse(createResult tenants.CreateResult) *CreateIdentityV2TenantsResponse {
	return &CreateIdentityV2TenantsResponse{
		CreateResult: createResult,
	}
}

// action function
func (oc *OpenstackClient) CreateIdentityV2Tenants(req *CreateIdentityV2TenantsRequest) *CreateIdentityV2TenantsResponse {
	return NewCreateIdentityV2TenantsResponse(tenants.Create(oc.Client, req.Opts))

}

// request struct for the GetIdentityV2Tenants
type GetIdentityV2TenantsRequest struct {
	Id string
}

func NewGetIdentityV2TenantsRequest() *GetIdentityV2TenantsRequest {
	return &GetIdentityV2TenantsRequest{}
}

// response struct for the GetIdentityV2Tenants
type GetIdentityV2TenantsResponse struct {
	GetResult tenants.GetResult
}

func NewGetIdentityV2TenantsResponse(getResult tenants.GetResult) *GetIdentityV2TenantsResponse {
	return &GetIdentityV2TenantsResponse{
		GetResult: getResult,
	}
}

// action function
func (oc *OpenstackClient) GetIdentityV2Tenants(req *GetIdentityV2TenantsRequest) *GetIdentityV2TenantsResponse {
	return NewGetIdentityV2TenantsResponse(tenants.Get(oc.Client, req.Id))

}

// request struct for the UpdateIdentityV2Tenants
type UpdateIdentityV2TenantsRequest struct {
	Id   string
	Opts tenants.UpdateOpts
}

func NewUpdateIdentityV2TenantsRequest() *UpdateIdentityV2TenantsRequest {
	return &UpdateIdentityV2TenantsRequest{}
}

// response struct for the UpdateIdentityV2Tenants
type UpdateIdentityV2TenantsResponse struct {
	UpdateResult tenants.UpdateResult
}

func NewUpdateIdentityV2TenantsResponse(updateResult tenants.UpdateResult) *UpdateIdentityV2TenantsResponse {
	return &UpdateIdentityV2TenantsResponse{
		UpdateResult: updateResult,
	}
}

// action function
func (oc *OpenstackClient) UpdateIdentityV2Tenants(req *UpdateIdentityV2TenantsRequest) *UpdateIdentityV2TenantsResponse {
	return NewUpdateIdentityV2TenantsResponse(tenants.Update(oc.Client, req.Id, req.Opts))

}

// request struct for the DeleteIdentityV2Tenants
type DeleteIdentityV2TenantsRequest struct {
	Id string
}

func NewDeleteIdentityV2TenantsRequest() *DeleteIdentityV2TenantsRequest {
	return &DeleteIdentityV2TenantsRequest{}
}

// response struct for the DeleteIdentityV2Tenants
type DeleteIdentityV2TenantsResponse struct {
	DeleteResult tenants.DeleteResult
}

func NewDeleteIdentityV2TenantsResponse(deleteResult tenants.DeleteResult) *DeleteIdentityV2TenantsResponse {
	return &DeleteIdentityV2TenantsResponse{
		DeleteResult: deleteResult,
	}
}

// action function
func (oc *OpenstackClient) DeleteIdentityV2Tenants(req *DeleteIdentityV2TenantsRequest) *DeleteIdentityV2TenantsResponse {
	return NewDeleteIdentityV2TenantsResponse(tenants.Delete(oc.Client, req.Id))

}
