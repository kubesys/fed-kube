package openstack

// Code generated by cloud manager.

import (
	"github.com/gophercloud/gophercloud/openstack/networking/v2/extensions/agents"
	"github.com/gophercloud/gophercloud/pagination"
)

// request struct for the ListNetworkingV2ExtensionsAgents
type ListNetworkingV2ExtensionsAgentsRequest struct {
	Opts agents.ListOpts
}

func NewListNetworkingV2ExtensionsAgentsRequest() *ListNetworkingV2ExtensionsAgentsRequest {
	return &ListNetworkingV2ExtensionsAgentsRequest{}
}

// response struct for the ListNetworkingV2ExtensionsAgents
type ListNetworkingV2ExtensionsAgentsResponse struct {
	Pager pagination.Pager
}

func NewListNetworkingV2ExtensionsAgentsResponse(pager pagination.Pager) *ListNetworkingV2ExtensionsAgentsResponse {
	return &ListNetworkingV2ExtensionsAgentsResponse{
		Pager: pager,
	}
}

// action function
func (oc *OpenstackClient) ListNetworkingV2ExtensionsAgents(req *ListNetworkingV2ExtensionsAgentsRequest) *ListNetworkingV2ExtensionsAgentsResponse {
	return NewListNetworkingV2ExtensionsAgentsResponse(agents.List(oc.Client, req.Opts))

}

// request struct for the GetNetworkingV2ExtensionsAgents
type GetNetworkingV2ExtensionsAgentsRequest struct {
	Id string
}

func NewGetNetworkingV2ExtensionsAgentsRequest() *GetNetworkingV2ExtensionsAgentsRequest {
	return &GetNetworkingV2ExtensionsAgentsRequest{}
}

// response struct for the GetNetworkingV2ExtensionsAgents
type GetNetworkingV2ExtensionsAgentsResponse struct {
	GetResult agents.GetResult
}

func NewGetNetworkingV2ExtensionsAgentsResponse(getResult agents.GetResult) *GetNetworkingV2ExtensionsAgentsResponse {
	return &GetNetworkingV2ExtensionsAgentsResponse{
		GetResult: getResult,
	}
}

// action function
func (oc *OpenstackClient) GetNetworkingV2ExtensionsAgents(req *GetNetworkingV2ExtensionsAgentsRequest) *GetNetworkingV2ExtensionsAgentsResponse {
	return NewGetNetworkingV2ExtensionsAgentsResponse(agents.Get(oc.Client, req.Id))

}

// request struct for the UpdateNetworkingV2ExtensionsAgents
type UpdateNetworkingV2ExtensionsAgentsRequest struct {
	Id   string
	Opts agents.UpdateOpts
}

func NewUpdateNetworkingV2ExtensionsAgentsRequest() *UpdateNetworkingV2ExtensionsAgentsRequest {
	return &UpdateNetworkingV2ExtensionsAgentsRequest{}
}

// response struct for the UpdateNetworkingV2ExtensionsAgents
type UpdateNetworkingV2ExtensionsAgentsResponse struct {
	UpdateResult agents.UpdateResult
}

func NewUpdateNetworkingV2ExtensionsAgentsResponse(updateResult agents.UpdateResult) *UpdateNetworkingV2ExtensionsAgentsResponse {
	return &UpdateNetworkingV2ExtensionsAgentsResponse{
		UpdateResult: updateResult,
	}
}

// action function
func (oc *OpenstackClient) UpdateNetworkingV2ExtensionsAgents(req *UpdateNetworkingV2ExtensionsAgentsRequest) *UpdateNetworkingV2ExtensionsAgentsResponse {
	return NewUpdateNetworkingV2ExtensionsAgentsResponse(agents.Update(oc.Client, req.Id, req.Opts))

}

// request struct for the DeleteNetworkingV2ExtensionsAgents
type DeleteNetworkingV2ExtensionsAgentsRequest struct {
	Id string
}

func NewDeleteNetworkingV2ExtensionsAgentsRequest() *DeleteNetworkingV2ExtensionsAgentsRequest {
	return &DeleteNetworkingV2ExtensionsAgentsRequest{}
}

// response struct for the DeleteNetworkingV2ExtensionsAgents
type DeleteNetworkingV2ExtensionsAgentsResponse struct {
	DeleteResult agents.DeleteResult
}

func NewDeleteNetworkingV2ExtensionsAgentsResponse(deleteResult agents.DeleteResult) *DeleteNetworkingV2ExtensionsAgentsResponse {
	return &DeleteNetworkingV2ExtensionsAgentsResponse{
		DeleteResult: deleteResult,
	}
}

// action function
func (oc *OpenstackClient) DeleteNetworkingV2ExtensionsAgents(req *DeleteNetworkingV2ExtensionsAgentsRequest) *DeleteNetworkingV2ExtensionsAgentsResponse {
	return NewDeleteNetworkingV2ExtensionsAgentsResponse(agents.Delete(oc.Client, req.Id))

}

// request struct for the ListDHCPNetworksNetworkingV2ExtensionsAgents
type ListDHCPNetworksNetworkingV2ExtensionsAgentsRequest struct {
	Id string
}

func NewListDHCPNetworksNetworkingV2ExtensionsAgentsRequest() *ListDHCPNetworksNetworkingV2ExtensionsAgentsRequest {
	return &ListDHCPNetworksNetworkingV2ExtensionsAgentsRequest{}
}

// response struct for the ListDHCPNetworksNetworkingV2ExtensionsAgents
type ListDHCPNetworksNetworkingV2ExtensionsAgentsResponse struct {
	ListDHCPNetworksResult agents.ListDHCPNetworksResult
}

func NewListDHCPNetworksNetworkingV2ExtensionsAgentsResponse(listDHCPNetworksResult agents.ListDHCPNetworksResult) *ListDHCPNetworksNetworkingV2ExtensionsAgentsResponse {
	return &ListDHCPNetworksNetworkingV2ExtensionsAgentsResponse{
		ListDHCPNetworksResult: listDHCPNetworksResult,
	}
}

// action function
func (oc *OpenstackClient) ListDHCPNetworksNetworkingV2ExtensionsAgents(req *ListDHCPNetworksNetworkingV2ExtensionsAgentsRequest) *ListDHCPNetworksNetworkingV2ExtensionsAgentsResponse {
	return NewListDHCPNetworksNetworkingV2ExtensionsAgentsResponse(agents.ListDHCPNetworks(oc.Client, req.Id))

}

// request struct for the ScheduleDHCPNetworkNetworkingV2ExtensionsAgents
type ScheduleDHCPNetworkNetworkingV2ExtensionsAgentsRequest struct {
	Id   string
	Opts agents.ScheduleDHCPNetworkOpts
}

func NewScheduleDHCPNetworkNetworkingV2ExtensionsAgentsRequest() *ScheduleDHCPNetworkNetworkingV2ExtensionsAgentsRequest {
	return &ScheduleDHCPNetworkNetworkingV2ExtensionsAgentsRequest{}
}

// response struct for the ScheduleDHCPNetworkNetworkingV2ExtensionsAgents
type ScheduleDHCPNetworkNetworkingV2ExtensionsAgentsResponse struct {
	ScheduleDHCPNetworkResult agents.ScheduleDHCPNetworkResult
}

func NewScheduleDHCPNetworkNetworkingV2ExtensionsAgentsResponse(scheduleDHCPNetworkResult agents.ScheduleDHCPNetworkResult) *ScheduleDHCPNetworkNetworkingV2ExtensionsAgentsResponse {
	return &ScheduleDHCPNetworkNetworkingV2ExtensionsAgentsResponse{
		ScheduleDHCPNetworkResult: scheduleDHCPNetworkResult,
	}
}

// action function
func (oc *OpenstackClient) ScheduleDHCPNetworkNetworkingV2ExtensionsAgents(req *ScheduleDHCPNetworkNetworkingV2ExtensionsAgentsRequest) *ScheduleDHCPNetworkNetworkingV2ExtensionsAgentsResponse {
	return NewScheduleDHCPNetworkNetworkingV2ExtensionsAgentsResponse(agents.ScheduleDHCPNetwork(oc.Client, req.Id, req.Opts))

}

// request struct for the RemoveDHCPNetworkNetworkingV2ExtensionsAgents
type RemoveDHCPNetworkNetworkingV2ExtensionsAgentsRequest struct {
	Id        string
	NetworkID string
}

func NewRemoveDHCPNetworkNetworkingV2ExtensionsAgentsRequest() *RemoveDHCPNetworkNetworkingV2ExtensionsAgentsRequest {
	return &RemoveDHCPNetworkNetworkingV2ExtensionsAgentsRequest{}
}

// response struct for the RemoveDHCPNetworkNetworkingV2ExtensionsAgents
type RemoveDHCPNetworkNetworkingV2ExtensionsAgentsResponse struct {
	RemoveDHCPNetworkResult agents.RemoveDHCPNetworkResult
}

func NewRemoveDHCPNetworkNetworkingV2ExtensionsAgentsResponse(removeDHCPNetworkResult agents.RemoveDHCPNetworkResult) *RemoveDHCPNetworkNetworkingV2ExtensionsAgentsResponse {
	return &RemoveDHCPNetworkNetworkingV2ExtensionsAgentsResponse{
		RemoveDHCPNetworkResult: removeDHCPNetworkResult,
	}
}

// action function
func (oc *OpenstackClient) RemoveDHCPNetworkNetworkingV2ExtensionsAgents(req *RemoveDHCPNetworkNetworkingV2ExtensionsAgentsRequest) *RemoveDHCPNetworkNetworkingV2ExtensionsAgentsResponse {
	return NewRemoveDHCPNetworkNetworkingV2ExtensionsAgentsResponse(agents.RemoveDHCPNetwork(oc.Client, req.Id, req.NetworkID))

}

// request struct for the ListBGPSpeakersNetworkingV2ExtensionsAgents
type ListBGPSpeakersNetworkingV2ExtensionsAgentsRequest struct {
	AgentID string
}

func NewListBGPSpeakersNetworkingV2ExtensionsAgentsRequest() *ListBGPSpeakersNetworkingV2ExtensionsAgentsRequest {
	return &ListBGPSpeakersNetworkingV2ExtensionsAgentsRequest{}
}

// response struct for the ListBGPSpeakersNetworkingV2ExtensionsAgents
type ListBGPSpeakersNetworkingV2ExtensionsAgentsResponse struct {
	Pager pagination.Pager
}

func NewListBGPSpeakersNetworkingV2ExtensionsAgentsResponse(pager pagination.Pager) *ListBGPSpeakersNetworkingV2ExtensionsAgentsResponse {
	return &ListBGPSpeakersNetworkingV2ExtensionsAgentsResponse{
		Pager: pager,
	}
}

// action function
func (oc *OpenstackClient) ListBGPSpeakersNetworkingV2ExtensionsAgents(req *ListBGPSpeakersNetworkingV2ExtensionsAgentsRequest) *ListBGPSpeakersNetworkingV2ExtensionsAgentsResponse {
	return NewListBGPSpeakersNetworkingV2ExtensionsAgentsResponse(agents.ListBGPSpeakers(oc.Client, req.AgentID))

}
