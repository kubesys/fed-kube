package openstack

// Code generated by cloud manager.

import (
	"github.com/gophercloud/gophercloud/openstack/identity/v3/projects"
	"github.com/gophercloud/gophercloud/pagination"
)

// request struct for the ListIdentityV3Projects
type ListIdentityV3ProjectsRequest struct {
	Opts projects.ListOpts
}

func NewListIdentityV3ProjectsRequest() *ListIdentityV3ProjectsRequest {
	return &ListIdentityV3ProjectsRequest{}
}

// response struct for the ListIdentityV3Projects
type ListIdentityV3ProjectsResponse struct {
	Pager pagination.Pager
}

func NewListIdentityV3ProjectsResponse(pager pagination.Pager) *ListIdentityV3ProjectsResponse {
	return &ListIdentityV3ProjectsResponse{
		Pager: pager,
	}
}

// action function
func (oc *OpenstackClient) ListIdentityV3Projects(req *ListIdentityV3ProjectsRequest) *ListIdentityV3ProjectsResponse {
	return NewListIdentityV3ProjectsResponse(projects.List(oc.Client, req.Opts))

}

// request struct for the ListAvailableIdentityV3Projects
type ListAvailableIdentityV3ProjectsRequest struct {
}

func NewListAvailableIdentityV3ProjectsRequest() *ListAvailableIdentityV3ProjectsRequest {
	return &ListAvailableIdentityV3ProjectsRequest{}
}

// response struct for the ListAvailableIdentityV3Projects
type ListAvailableIdentityV3ProjectsResponse struct {
	Pager pagination.Pager
}

func NewListAvailableIdentityV3ProjectsResponse(pager pagination.Pager) *ListAvailableIdentityV3ProjectsResponse {
	return &ListAvailableIdentityV3ProjectsResponse{
		Pager: pager,
	}
}

// action function
func (oc *OpenstackClient) ListAvailableIdentityV3Projects(req *ListAvailableIdentityV3ProjectsRequest) *ListAvailableIdentityV3ProjectsResponse {
	return NewListAvailableIdentityV3ProjectsResponse(projects.ListAvailable(oc.Client))

}

// request struct for the GetIdentityV3Projects
type GetIdentityV3ProjectsRequest struct {
	Id string
}

func NewGetIdentityV3ProjectsRequest() *GetIdentityV3ProjectsRequest {
	return &GetIdentityV3ProjectsRequest{}
}

// response struct for the GetIdentityV3Projects
type GetIdentityV3ProjectsResponse struct {
	GetResult projects.GetResult
}

func NewGetIdentityV3ProjectsResponse(getResult projects.GetResult) *GetIdentityV3ProjectsResponse {
	return &GetIdentityV3ProjectsResponse{
		GetResult: getResult,
	}
}

// action function
func (oc *OpenstackClient) GetIdentityV3Projects(req *GetIdentityV3ProjectsRequest) *GetIdentityV3ProjectsResponse {
	return NewGetIdentityV3ProjectsResponse(projects.Get(oc.Client, req.Id))

}

// request struct for the CreateIdentityV3Projects
type CreateIdentityV3ProjectsRequest struct {
	Opts projects.CreateOpts
}

func NewCreateIdentityV3ProjectsRequest() *CreateIdentityV3ProjectsRequest {
	return &CreateIdentityV3ProjectsRequest{}
}

// response struct for the CreateIdentityV3Projects
type CreateIdentityV3ProjectsResponse struct {
	CreateResult projects.CreateResult
}

func NewCreateIdentityV3ProjectsResponse(createResult projects.CreateResult) *CreateIdentityV3ProjectsResponse {
	return &CreateIdentityV3ProjectsResponse{
		CreateResult: createResult,
	}
}

// action function
func (oc *OpenstackClient) CreateIdentityV3Projects(req *CreateIdentityV3ProjectsRequest) *CreateIdentityV3ProjectsResponse {
	return NewCreateIdentityV3ProjectsResponse(projects.Create(oc.Client, req.Opts))

}

// request struct for the DeleteIdentityV3Projects
type DeleteIdentityV3ProjectsRequest struct {
	ProjectID string
}

func NewDeleteIdentityV3ProjectsRequest() *DeleteIdentityV3ProjectsRequest {
	return &DeleteIdentityV3ProjectsRequest{}
}

// response struct for the DeleteIdentityV3Projects
type DeleteIdentityV3ProjectsResponse struct {
	DeleteResult projects.DeleteResult
}

func NewDeleteIdentityV3ProjectsResponse(deleteResult projects.DeleteResult) *DeleteIdentityV3ProjectsResponse {
	return &DeleteIdentityV3ProjectsResponse{
		DeleteResult: deleteResult,
	}
}

// action function
func (oc *OpenstackClient) DeleteIdentityV3Projects(req *DeleteIdentityV3ProjectsRequest) *DeleteIdentityV3ProjectsResponse {
	return NewDeleteIdentityV3ProjectsResponse(projects.Delete(oc.Client, req.ProjectID))

}

// request struct for the UpdateIdentityV3Projects
type UpdateIdentityV3ProjectsRequest struct {
	Id   string
	Opts projects.UpdateOpts
}

func NewUpdateIdentityV3ProjectsRequest() *UpdateIdentityV3ProjectsRequest {
	return &UpdateIdentityV3ProjectsRequest{}
}

// response struct for the UpdateIdentityV3Projects
type UpdateIdentityV3ProjectsResponse struct {
	UpdateResult projects.UpdateResult
}

func NewUpdateIdentityV3ProjectsResponse(updateResult projects.UpdateResult) *UpdateIdentityV3ProjectsResponse {
	return &UpdateIdentityV3ProjectsResponse{
		UpdateResult: updateResult,
	}
}

// action function
func (oc *OpenstackClient) UpdateIdentityV3Projects(req *UpdateIdentityV3ProjectsRequest) *UpdateIdentityV3ProjectsResponse {
	return NewUpdateIdentityV3ProjectsResponse(projects.Update(oc.Client, req.Id, req.Opts))

}
