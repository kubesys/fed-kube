package openstack

// Code generated by cloud manager.

import (
	"github.com/gophercloud/gophercloud/openstack/identity/v3/roles"
	"github.com/gophercloud/gophercloud/pagination"
)

// request struct for the ListIdentityV3Roles
type ListIdentityV3RolesRequest struct {
	Opts roles.ListOpts
}

func NewListIdentityV3RolesRequest() *ListIdentityV3RolesRequest {
	return &ListIdentityV3RolesRequest{}
}

// response struct for the ListIdentityV3Roles
type ListIdentityV3RolesResponse struct {
	Pager pagination.Pager
}

func NewListIdentityV3RolesResponse(pager pagination.Pager) *ListIdentityV3RolesResponse {
	return &ListIdentityV3RolesResponse{
		Pager: pager,
	}
}

// action function
func (oc *OpenstackClient) ListIdentityV3Roles(req *ListIdentityV3RolesRequest) *ListIdentityV3RolesResponse {
	return NewListIdentityV3RolesResponse(roles.List(oc.Client, req.Opts))

}

// request struct for the GetIdentityV3Roles
type GetIdentityV3RolesRequest struct {
	Id string
}

func NewGetIdentityV3RolesRequest() *GetIdentityV3RolesRequest {
	return &GetIdentityV3RolesRequest{}
}

// response struct for the GetIdentityV3Roles
type GetIdentityV3RolesResponse struct {
	GetResult roles.GetResult
}

func NewGetIdentityV3RolesResponse(getResult roles.GetResult) *GetIdentityV3RolesResponse {
	return &GetIdentityV3RolesResponse{
		GetResult: getResult,
	}
}

// action function
func (oc *OpenstackClient) GetIdentityV3Roles(req *GetIdentityV3RolesRequest) *GetIdentityV3RolesResponse {
	return NewGetIdentityV3RolesResponse(roles.Get(oc.Client, req.Id))

}

// request struct for the CreateIdentityV3Roles
type CreateIdentityV3RolesRequest struct {
	Opts roles.CreateOpts
}

func NewCreateIdentityV3RolesRequest() *CreateIdentityV3RolesRequest {
	return &CreateIdentityV3RolesRequest{}
}

// response struct for the CreateIdentityV3Roles
type CreateIdentityV3RolesResponse struct {
	CreateResult roles.CreateResult
}

func NewCreateIdentityV3RolesResponse(createResult roles.CreateResult) *CreateIdentityV3RolesResponse {
	return &CreateIdentityV3RolesResponse{
		CreateResult: createResult,
	}
}

// action function
func (oc *OpenstackClient) CreateIdentityV3Roles(req *CreateIdentityV3RolesRequest) *CreateIdentityV3RolesResponse {
	return NewCreateIdentityV3RolesResponse(roles.Create(oc.Client, req.Opts))

}

// request struct for the UpdateIdentityV3Roles
type UpdateIdentityV3RolesRequest struct {
	RoleID string
	Opts   roles.UpdateOpts
}

func NewUpdateIdentityV3RolesRequest() *UpdateIdentityV3RolesRequest {
	return &UpdateIdentityV3RolesRequest{}
}

// response struct for the UpdateIdentityV3Roles
type UpdateIdentityV3RolesResponse struct {
	UpdateResult roles.UpdateResult
}

func NewUpdateIdentityV3RolesResponse(updateResult roles.UpdateResult) *UpdateIdentityV3RolesResponse {
	return &UpdateIdentityV3RolesResponse{
		UpdateResult: updateResult,
	}
}

// action function
func (oc *OpenstackClient) UpdateIdentityV3Roles(req *UpdateIdentityV3RolesRequest) *UpdateIdentityV3RolesResponse {
	return NewUpdateIdentityV3RolesResponse(roles.Update(oc.Client, req.RoleID, req.Opts))

}

// request struct for the DeleteIdentityV3Roles
type DeleteIdentityV3RolesRequest struct {
	RoleID string
}

func NewDeleteIdentityV3RolesRequest() *DeleteIdentityV3RolesRequest {
	return &DeleteIdentityV3RolesRequest{}
}

// response struct for the DeleteIdentityV3Roles
type DeleteIdentityV3RolesResponse struct {
	DeleteResult roles.DeleteResult
}

func NewDeleteIdentityV3RolesResponse(deleteResult roles.DeleteResult) *DeleteIdentityV3RolesResponse {
	return &DeleteIdentityV3RolesResponse{
		DeleteResult: deleteResult,
	}
}

// action function
func (oc *OpenstackClient) DeleteIdentityV3Roles(req *DeleteIdentityV3RolesRequest) *DeleteIdentityV3RolesResponse {
	return NewDeleteIdentityV3RolesResponse(roles.Delete(oc.Client, req.RoleID))

}

// request struct for the ListAssignmentsIdentityV3Roles
type ListAssignmentsIdentityV3RolesRequest struct {
	Opts roles.ListAssignmentsOpts
}

func NewListAssignmentsIdentityV3RolesRequest() *ListAssignmentsIdentityV3RolesRequest {
	return &ListAssignmentsIdentityV3RolesRequest{}
}

// response struct for the ListAssignmentsIdentityV3Roles
type ListAssignmentsIdentityV3RolesResponse struct {
	Pager pagination.Pager
}

func NewListAssignmentsIdentityV3RolesResponse(pager pagination.Pager) *ListAssignmentsIdentityV3RolesResponse {
	return &ListAssignmentsIdentityV3RolesResponse{
		Pager: pager,
	}
}

// action function
func (oc *OpenstackClient) ListAssignmentsIdentityV3Roles(req *ListAssignmentsIdentityV3RolesRequest) *ListAssignmentsIdentityV3RolesResponse {
	return NewListAssignmentsIdentityV3RolesResponse(roles.ListAssignments(oc.Client, req.Opts))

}

// request struct for the ListAssignmentsOnResourceIdentityV3Roles
type ListAssignmentsOnResourceIdentityV3RolesRequest struct {
	Opts roles.ListAssignmentsOnResourceOpts
}

func NewListAssignmentsOnResourceIdentityV3RolesRequest() *ListAssignmentsOnResourceIdentityV3RolesRequest {
	return &ListAssignmentsOnResourceIdentityV3RolesRequest{}
}

// response struct for the ListAssignmentsOnResourceIdentityV3Roles
type ListAssignmentsOnResourceIdentityV3RolesResponse struct {
	Pager pagination.Pager
}

func NewListAssignmentsOnResourceIdentityV3RolesResponse(pager pagination.Pager) *ListAssignmentsOnResourceIdentityV3RolesResponse {
	return &ListAssignmentsOnResourceIdentityV3RolesResponse{
		Pager: pager,
	}
}

// action function
func (oc *OpenstackClient) ListAssignmentsOnResourceIdentityV3Roles(req *ListAssignmentsOnResourceIdentityV3RolesRequest) *ListAssignmentsOnResourceIdentityV3RolesResponse {
	return NewListAssignmentsOnResourceIdentityV3RolesResponse(roles.ListAssignmentsOnResource(oc.Client, req.Opts))

}

// request struct for the AssignIdentityV3Roles
type AssignIdentityV3RolesRequest struct {
	RoleID string
	Opts   roles.AssignOpts
}

func NewAssignIdentityV3RolesRequest() *AssignIdentityV3RolesRequest {
	return &AssignIdentityV3RolesRequest{}
}

// response struct for the AssignIdentityV3Roles
type AssignIdentityV3RolesResponse struct {
	AssignmentResult roles.AssignmentResult
}

func NewAssignIdentityV3RolesResponse(assignmentResult roles.AssignmentResult) *AssignIdentityV3RolesResponse {
	return &AssignIdentityV3RolesResponse{
		AssignmentResult: assignmentResult,
	}
}

// action function
func (oc *OpenstackClient) AssignIdentityV3Roles(req *AssignIdentityV3RolesRequest) *AssignIdentityV3RolesResponse {
	return NewAssignIdentityV3RolesResponse(roles.Assign(oc.Client, req.RoleID, req.Opts))

}

// request struct for the UnassignIdentityV3Roles
type UnassignIdentityV3RolesRequest struct {
	RoleID string
	Opts   roles.UnassignOpts
}

func NewUnassignIdentityV3RolesRequest() *UnassignIdentityV3RolesRequest {
	return &UnassignIdentityV3RolesRequest{}
}

// response struct for the UnassignIdentityV3Roles
type UnassignIdentityV3RolesResponse struct {
	UnassignmentResult roles.UnassignmentResult
}

func NewUnassignIdentityV3RolesResponse(unassignmentResult roles.UnassignmentResult) *UnassignIdentityV3RolesResponse {
	return &UnassignIdentityV3RolesResponse{
		UnassignmentResult: unassignmentResult,
	}
}

// action function
func (oc *OpenstackClient) UnassignIdentityV3Roles(req *UnassignIdentityV3RolesRequest) *UnassignIdentityV3RolesResponse {
	return NewUnassignIdentityV3RolesResponse(roles.Unassign(oc.Client, req.RoleID, req.Opts))

}
