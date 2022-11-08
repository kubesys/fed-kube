package openstack

// Code generated by cloud manager.

import (
	"github.com/gophercloud/gophercloud/openstack/networking/v2/extensions/fwaas_v2/groups"
	"github.com/gophercloud/gophercloud/pagination"
)

// request struct for the ListNetworkingV2ExtensionsFwaas_v2Groups
type ListNetworkingV2ExtensionsFwaas_v2GroupsRequest struct {
	Opts groups.ListOpts
}

func NewListNetworkingV2ExtensionsFwaas_v2GroupsRequest() *ListNetworkingV2ExtensionsFwaas_v2GroupsRequest {
	return &ListNetworkingV2ExtensionsFwaas_v2GroupsRequest{}
}

// response struct for the ListNetworkingV2ExtensionsFwaas_v2Groups
type ListNetworkingV2ExtensionsFwaas_v2GroupsResponse struct {
	Pager pagination.Pager
}

func NewListNetworkingV2ExtensionsFwaas_v2GroupsResponse(pager pagination.Pager) *ListNetworkingV2ExtensionsFwaas_v2GroupsResponse {
	return &ListNetworkingV2ExtensionsFwaas_v2GroupsResponse{
		Pager: pager,
	}
}

// action function
func (oc *OpenstackClient) ListNetworkingV2ExtensionsFwaas_v2Groups(req *ListNetworkingV2ExtensionsFwaas_v2GroupsRequest) *ListNetworkingV2ExtensionsFwaas_v2GroupsResponse {
	return NewListNetworkingV2ExtensionsFwaas_v2GroupsResponse(groups.List(oc.Client, req.Opts))

}

// request struct for the GetNetworkingV2ExtensionsFwaas_v2Groups
type GetNetworkingV2ExtensionsFwaas_v2GroupsRequest struct {
	Id string
}

func NewGetNetworkingV2ExtensionsFwaas_v2GroupsRequest() *GetNetworkingV2ExtensionsFwaas_v2GroupsRequest {
	return &GetNetworkingV2ExtensionsFwaas_v2GroupsRequest{}
}

// response struct for the GetNetworkingV2ExtensionsFwaas_v2Groups
type GetNetworkingV2ExtensionsFwaas_v2GroupsResponse struct {
	GetResult groups.GetResult
}

func NewGetNetworkingV2ExtensionsFwaas_v2GroupsResponse(getResult groups.GetResult) *GetNetworkingV2ExtensionsFwaas_v2GroupsResponse {
	return &GetNetworkingV2ExtensionsFwaas_v2GroupsResponse{
		GetResult: getResult,
	}
}

// action function
func (oc *OpenstackClient) GetNetworkingV2ExtensionsFwaas_v2Groups(req *GetNetworkingV2ExtensionsFwaas_v2GroupsRequest) *GetNetworkingV2ExtensionsFwaas_v2GroupsResponse {
	return NewGetNetworkingV2ExtensionsFwaas_v2GroupsResponse(groups.Get(oc.Client, req.Id))

}

// request struct for the CreateNetworkingV2ExtensionsFwaas_v2Groups
type CreateNetworkingV2ExtensionsFwaas_v2GroupsRequest struct {
	Opts groups.CreateOpts
}

func NewCreateNetworkingV2ExtensionsFwaas_v2GroupsRequest() *CreateNetworkingV2ExtensionsFwaas_v2GroupsRequest {
	return &CreateNetworkingV2ExtensionsFwaas_v2GroupsRequest{}
}

// response struct for the CreateNetworkingV2ExtensionsFwaas_v2Groups
type CreateNetworkingV2ExtensionsFwaas_v2GroupsResponse struct {
	CreateResult groups.CreateResult
}

func NewCreateNetworkingV2ExtensionsFwaas_v2GroupsResponse(createResult groups.CreateResult) *CreateNetworkingV2ExtensionsFwaas_v2GroupsResponse {
	return &CreateNetworkingV2ExtensionsFwaas_v2GroupsResponse{
		CreateResult: createResult,
	}
}

// action function
func (oc *OpenstackClient) CreateNetworkingV2ExtensionsFwaas_v2Groups(req *CreateNetworkingV2ExtensionsFwaas_v2GroupsRequest) *CreateNetworkingV2ExtensionsFwaas_v2GroupsResponse {
	return NewCreateNetworkingV2ExtensionsFwaas_v2GroupsResponse(groups.Create(oc.Client, req.Opts))

}

// request struct for the UpdateNetworkingV2ExtensionsFwaas_v2Groups
type UpdateNetworkingV2ExtensionsFwaas_v2GroupsRequest struct {
	Id   string
	Opts groups.UpdateOpts
}

func NewUpdateNetworkingV2ExtensionsFwaas_v2GroupsRequest() *UpdateNetworkingV2ExtensionsFwaas_v2GroupsRequest {
	return &UpdateNetworkingV2ExtensionsFwaas_v2GroupsRequest{}
}

// response struct for the UpdateNetworkingV2ExtensionsFwaas_v2Groups
type UpdateNetworkingV2ExtensionsFwaas_v2GroupsResponse struct {
	UpdateResult groups.UpdateResult
}

func NewUpdateNetworkingV2ExtensionsFwaas_v2GroupsResponse(updateResult groups.UpdateResult) *UpdateNetworkingV2ExtensionsFwaas_v2GroupsResponse {
	return &UpdateNetworkingV2ExtensionsFwaas_v2GroupsResponse{
		UpdateResult: updateResult,
	}
}

// action function
func (oc *OpenstackClient) UpdateNetworkingV2ExtensionsFwaas_v2Groups(req *UpdateNetworkingV2ExtensionsFwaas_v2GroupsRequest) *UpdateNetworkingV2ExtensionsFwaas_v2GroupsResponse {
	return NewUpdateNetworkingV2ExtensionsFwaas_v2GroupsResponse(groups.Update(oc.Client, req.Id, req.Opts))

}

// request struct for the DeleteNetworkingV2ExtensionsFwaas_v2Groups
type DeleteNetworkingV2ExtensionsFwaas_v2GroupsRequest struct {
	Id string
}

func NewDeleteNetworkingV2ExtensionsFwaas_v2GroupsRequest() *DeleteNetworkingV2ExtensionsFwaas_v2GroupsRequest {
	return &DeleteNetworkingV2ExtensionsFwaas_v2GroupsRequest{}
}

// response struct for the DeleteNetworkingV2ExtensionsFwaas_v2Groups
type DeleteNetworkingV2ExtensionsFwaas_v2GroupsResponse struct {
	DeleteResult groups.DeleteResult
}

func NewDeleteNetworkingV2ExtensionsFwaas_v2GroupsResponse(deleteResult groups.DeleteResult) *DeleteNetworkingV2ExtensionsFwaas_v2GroupsResponse {
	return &DeleteNetworkingV2ExtensionsFwaas_v2GroupsResponse{
		DeleteResult: deleteResult,
	}
}

// action function
func (oc *OpenstackClient) DeleteNetworkingV2ExtensionsFwaas_v2Groups(req *DeleteNetworkingV2ExtensionsFwaas_v2GroupsRequest) *DeleteNetworkingV2ExtensionsFwaas_v2GroupsResponse {
	return NewDeleteNetworkingV2ExtensionsFwaas_v2GroupsResponse(groups.Delete(oc.Client, req.Id))

}
