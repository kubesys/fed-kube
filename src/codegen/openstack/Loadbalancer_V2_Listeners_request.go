package openstack

// Code generated by cloud manager.

import (
	"github.com/gophercloud/gophercloud/openstack/loadbalancer/v2/listeners"
	"github.com/gophercloud/gophercloud/pagination"
)

// request struct for the ListLoadbalancerV2Listeners
type ListLoadbalancerV2ListenersRequest struct {
	Opts listeners.ListOpts
}

func NewListLoadbalancerV2ListenersRequest() *ListLoadbalancerV2ListenersRequest {
	return &ListLoadbalancerV2ListenersRequest{}
}

// response struct for the ListLoadbalancerV2Listeners
type ListLoadbalancerV2ListenersResponse struct {
	Pager pagination.Pager
}

func NewListLoadbalancerV2ListenersResponse(pager pagination.Pager) *ListLoadbalancerV2ListenersResponse {
	return &ListLoadbalancerV2ListenersResponse{
		Pager: pager,
	}
}

// action function
func (oc *OpenstackClient) ListLoadbalancerV2Listeners(req *ListLoadbalancerV2ListenersRequest) *ListLoadbalancerV2ListenersResponse {
	return NewListLoadbalancerV2ListenersResponse(listeners.List(oc.Client, req.Opts))

}

// request struct for the CreateLoadbalancerV2Listeners
type CreateLoadbalancerV2ListenersRequest struct {
	Opts listeners.CreateOpts
}

func NewCreateLoadbalancerV2ListenersRequest() *CreateLoadbalancerV2ListenersRequest {
	return &CreateLoadbalancerV2ListenersRequest{}
}

// response struct for the CreateLoadbalancerV2Listeners
type CreateLoadbalancerV2ListenersResponse struct {
	CreateResult listeners.CreateResult
}

func NewCreateLoadbalancerV2ListenersResponse(createResult listeners.CreateResult) *CreateLoadbalancerV2ListenersResponse {
	return &CreateLoadbalancerV2ListenersResponse{
		CreateResult: createResult,
	}
}

// action function
func (oc *OpenstackClient) CreateLoadbalancerV2Listeners(req *CreateLoadbalancerV2ListenersRequest) *CreateLoadbalancerV2ListenersResponse {
	return NewCreateLoadbalancerV2ListenersResponse(listeners.Create(oc.Client, req.Opts))

}

// request struct for the GetLoadbalancerV2Listeners
type GetLoadbalancerV2ListenersRequest struct {
	Id string
}

func NewGetLoadbalancerV2ListenersRequest() *GetLoadbalancerV2ListenersRequest {
	return &GetLoadbalancerV2ListenersRequest{}
}

// response struct for the GetLoadbalancerV2Listeners
type GetLoadbalancerV2ListenersResponse struct {
	GetResult listeners.GetResult
}

func NewGetLoadbalancerV2ListenersResponse(getResult listeners.GetResult) *GetLoadbalancerV2ListenersResponse {
	return &GetLoadbalancerV2ListenersResponse{
		GetResult: getResult,
	}
}

// action function
func (oc *OpenstackClient) GetLoadbalancerV2Listeners(req *GetLoadbalancerV2ListenersRequest) *GetLoadbalancerV2ListenersResponse {
	return NewGetLoadbalancerV2ListenersResponse(listeners.Get(oc.Client, req.Id))

}

// request struct for the UpdateLoadbalancerV2Listeners
type UpdateLoadbalancerV2ListenersRequest struct {
	Id   string
	Opts listeners.UpdateOpts
}

func NewUpdateLoadbalancerV2ListenersRequest() *UpdateLoadbalancerV2ListenersRequest {
	return &UpdateLoadbalancerV2ListenersRequest{}
}

// response struct for the UpdateLoadbalancerV2Listeners
type UpdateLoadbalancerV2ListenersResponse struct {
	UpdateResult listeners.UpdateResult
}

func NewUpdateLoadbalancerV2ListenersResponse(updateResult listeners.UpdateResult) *UpdateLoadbalancerV2ListenersResponse {
	return &UpdateLoadbalancerV2ListenersResponse{
		UpdateResult: updateResult,
	}
}

// action function
func (oc *OpenstackClient) UpdateLoadbalancerV2Listeners(req *UpdateLoadbalancerV2ListenersRequest) *UpdateLoadbalancerV2ListenersResponse {
	return NewUpdateLoadbalancerV2ListenersResponse(listeners.Update(oc.Client, req.Id, req.Opts))

}

// request struct for the DeleteLoadbalancerV2Listeners
type DeleteLoadbalancerV2ListenersRequest struct {
	Id string
}

func NewDeleteLoadbalancerV2ListenersRequest() *DeleteLoadbalancerV2ListenersRequest {
	return &DeleteLoadbalancerV2ListenersRequest{}
}

// response struct for the DeleteLoadbalancerV2Listeners
type DeleteLoadbalancerV2ListenersResponse struct {
	DeleteResult listeners.DeleteResult
}

func NewDeleteLoadbalancerV2ListenersResponse(deleteResult listeners.DeleteResult) *DeleteLoadbalancerV2ListenersResponse {
	return &DeleteLoadbalancerV2ListenersResponse{
		DeleteResult: deleteResult,
	}
}

// action function
func (oc *OpenstackClient) DeleteLoadbalancerV2Listeners(req *DeleteLoadbalancerV2ListenersRequest) *DeleteLoadbalancerV2ListenersResponse {
	return NewDeleteLoadbalancerV2ListenersResponse(listeners.Delete(oc.Client, req.Id))

}

// request struct for the GetStatsLoadbalancerV2Listeners
type GetStatsLoadbalancerV2ListenersRequest struct {
	Id string
}

func NewGetStatsLoadbalancerV2ListenersRequest() *GetStatsLoadbalancerV2ListenersRequest {
	return &GetStatsLoadbalancerV2ListenersRequest{}
}

// response struct for the GetStatsLoadbalancerV2Listeners
type GetStatsLoadbalancerV2ListenersResponse struct {
	StatsResult listeners.StatsResult
}

func NewGetStatsLoadbalancerV2ListenersResponse(statsResult listeners.StatsResult) *GetStatsLoadbalancerV2ListenersResponse {
	return &GetStatsLoadbalancerV2ListenersResponse{
		StatsResult: statsResult,
	}
}

// action function
func (oc *OpenstackClient) GetStatsLoadbalancerV2Listeners(req *GetStatsLoadbalancerV2ListenersRequest) *GetStatsLoadbalancerV2ListenersResponse {
	return NewGetStatsLoadbalancerV2ListenersResponse(listeners.GetStats(oc.Client, req.Id))

}
