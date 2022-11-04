package openstack

// Code generated by cloud manager.





import (
    "github.com/gophercloud/gophercloud/openstack/networking/v2/extensions/lbaas_v2/loadbalancers"
    "github.com/gophercloud/gophercloud/pagination"
)
//request struct for the ListNetworkingV2ExtensionsLbaas_v2Loadbalancers
type ListNetworkingV2ExtensionsLbaas_v2LoadbalancersRequest struct{
    Opts loadbalancers.ListOpts
}

func NewListNetworkingV2ExtensionsLbaas_v2LoadbalancersRequest()*ListNetworkingV2ExtensionsLbaas_v2LoadbalancersRequest{
    return &ListNetworkingV2ExtensionsLbaas_v2LoadbalancersRequest{}
}

//response struct for the ListNetworkingV2ExtensionsLbaas_v2Loadbalancers
type ListNetworkingV2ExtensionsLbaas_v2LoadbalancersResponse struct{
    Pager pagination.Pager
}

func NewListNetworkingV2ExtensionsLbaas_v2LoadbalancersResponse(pager pagination.Pager,)*ListNetworkingV2ExtensionsLbaas_v2LoadbalancersResponse {
    return &ListNetworkingV2ExtensionsLbaas_v2LoadbalancersResponse{
            Pager:pager,
    }
}

// action function
func (oc *OpenstackClient) ListNetworkingV2ExtensionsLbaas_v2Loadbalancers(req *ListNetworkingV2ExtensionsLbaas_v2LoadbalancersRequest)(*ListNetworkingV2ExtensionsLbaas_v2LoadbalancersResponse){
    return NewListNetworkingV2ExtensionsLbaas_v2LoadbalancersResponse(loadbalancers.List(oc.Client,req.Opts, ))

}
//request struct for the CreateNetworkingV2ExtensionsLbaas_v2Loadbalancers
type CreateNetworkingV2ExtensionsLbaas_v2LoadbalancersRequest struct{
    Opts loadbalancers.CreateOpts
}

func NewCreateNetworkingV2ExtensionsLbaas_v2LoadbalancersRequest()*CreateNetworkingV2ExtensionsLbaas_v2LoadbalancersRequest{
    return &CreateNetworkingV2ExtensionsLbaas_v2LoadbalancersRequest{}
}

//response struct for the CreateNetworkingV2ExtensionsLbaas_v2Loadbalancers
type CreateNetworkingV2ExtensionsLbaas_v2LoadbalancersResponse struct{
    CreateResult loadbalancers.CreateResult
}

func NewCreateNetworkingV2ExtensionsLbaas_v2LoadbalancersResponse(createResult loadbalancers.CreateResult,)*CreateNetworkingV2ExtensionsLbaas_v2LoadbalancersResponse {
    return &CreateNetworkingV2ExtensionsLbaas_v2LoadbalancersResponse{
            CreateResult:createResult,
    }
}

// action function
func (oc *OpenstackClient) CreateNetworkingV2ExtensionsLbaas_v2Loadbalancers(req *CreateNetworkingV2ExtensionsLbaas_v2LoadbalancersRequest)(*CreateNetworkingV2ExtensionsLbaas_v2LoadbalancersResponse){
    return NewCreateNetworkingV2ExtensionsLbaas_v2LoadbalancersResponse(loadbalancers.Create(oc.Client,req.Opts, ))

}
//request struct for the GetNetworkingV2ExtensionsLbaas_v2Loadbalancers
type GetNetworkingV2ExtensionsLbaas_v2LoadbalancersRequest struct{
    Id string
}

func NewGetNetworkingV2ExtensionsLbaas_v2LoadbalancersRequest()*GetNetworkingV2ExtensionsLbaas_v2LoadbalancersRequest{
    return &GetNetworkingV2ExtensionsLbaas_v2LoadbalancersRequest{}
}

//response struct for the GetNetworkingV2ExtensionsLbaas_v2Loadbalancers
type GetNetworkingV2ExtensionsLbaas_v2LoadbalancersResponse struct{
    GetResult loadbalancers.GetResult
}

func NewGetNetworkingV2ExtensionsLbaas_v2LoadbalancersResponse(getResult loadbalancers.GetResult,)*GetNetworkingV2ExtensionsLbaas_v2LoadbalancersResponse {
    return &GetNetworkingV2ExtensionsLbaas_v2LoadbalancersResponse{
            GetResult:getResult,
    }
}

// action function
func (oc *OpenstackClient) GetNetworkingV2ExtensionsLbaas_v2Loadbalancers(req *GetNetworkingV2ExtensionsLbaas_v2LoadbalancersRequest)(*GetNetworkingV2ExtensionsLbaas_v2LoadbalancersResponse){
    return NewGetNetworkingV2ExtensionsLbaas_v2LoadbalancersResponse(loadbalancers.Get(oc.Client,req.Id, ))

}
//request struct for the UpdateNetworkingV2ExtensionsLbaas_v2Loadbalancers
type UpdateNetworkingV2ExtensionsLbaas_v2LoadbalancersRequest struct{
    Id string
    Opts loadbalancers.UpdateOpts
}

func NewUpdateNetworkingV2ExtensionsLbaas_v2LoadbalancersRequest()*UpdateNetworkingV2ExtensionsLbaas_v2LoadbalancersRequest{
    return &UpdateNetworkingV2ExtensionsLbaas_v2LoadbalancersRequest{}
}

//response struct for the UpdateNetworkingV2ExtensionsLbaas_v2Loadbalancers
type UpdateNetworkingV2ExtensionsLbaas_v2LoadbalancersResponse struct{
    UpdateResult loadbalancers.UpdateResult
}

func NewUpdateNetworkingV2ExtensionsLbaas_v2LoadbalancersResponse(updateResult loadbalancers.UpdateResult,)*UpdateNetworkingV2ExtensionsLbaas_v2LoadbalancersResponse {
    return &UpdateNetworkingV2ExtensionsLbaas_v2LoadbalancersResponse{
            UpdateResult:updateResult,
    }
}

// action function
func (oc *OpenstackClient) UpdateNetworkingV2ExtensionsLbaas_v2Loadbalancers(req *UpdateNetworkingV2ExtensionsLbaas_v2LoadbalancersRequest)(*UpdateNetworkingV2ExtensionsLbaas_v2LoadbalancersResponse){
    return NewUpdateNetworkingV2ExtensionsLbaas_v2LoadbalancersResponse(loadbalancers.Update(oc.Client,req.Id,req.Opts, ))

}
//request struct for the DeleteNetworkingV2ExtensionsLbaas_v2Loadbalancers
type DeleteNetworkingV2ExtensionsLbaas_v2LoadbalancersRequest struct{
    Id string
}

func NewDeleteNetworkingV2ExtensionsLbaas_v2LoadbalancersRequest()*DeleteNetworkingV2ExtensionsLbaas_v2LoadbalancersRequest{
    return &DeleteNetworkingV2ExtensionsLbaas_v2LoadbalancersRequest{}
}

//response struct for the DeleteNetworkingV2ExtensionsLbaas_v2Loadbalancers
type DeleteNetworkingV2ExtensionsLbaas_v2LoadbalancersResponse struct{
    DeleteResult loadbalancers.DeleteResult
}

func NewDeleteNetworkingV2ExtensionsLbaas_v2LoadbalancersResponse(deleteResult loadbalancers.DeleteResult,)*DeleteNetworkingV2ExtensionsLbaas_v2LoadbalancersResponse {
    return &DeleteNetworkingV2ExtensionsLbaas_v2LoadbalancersResponse{
            DeleteResult:deleteResult,
    }
}

// action function
func (oc *OpenstackClient) DeleteNetworkingV2ExtensionsLbaas_v2Loadbalancers(req *DeleteNetworkingV2ExtensionsLbaas_v2LoadbalancersRequest)(*DeleteNetworkingV2ExtensionsLbaas_v2LoadbalancersResponse){
    return NewDeleteNetworkingV2ExtensionsLbaas_v2LoadbalancersResponse(loadbalancers.Delete(oc.Client,req.Id, ))

}
//request struct for the CascadingDeleteNetworkingV2ExtensionsLbaas_v2Loadbalancers
type CascadingDeleteNetworkingV2ExtensionsLbaas_v2LoadbalancersRequest struct{
    Id string
}

func NewCascadingDeleteNetworkingV2ExtensionsLbaas_v2LoadbalancersRequest()*CascadingDeleteNetworkingV2ExtensionsLbaas_v2LoadbalancersRequest{
    return &CascadingDeleteNetworkingV2ExtensionsLbaas_v2LoadbalancersRequest{}
}

//response struct for the CascadingDeleteNetworkingV2ExtensionsLbaas_v2Loadbalancers
type CascadingDeleteNetworkingV2ExtensionsLbaas_v2LoadbalancersResponse struct{
    DeleteResult loadbalancers.DeleteResult
}

func NewCascadingDeleteNetworkingV2ExtensionsLbaas_v2LoadbalancersResponse(deleteResult loadbalancers.DeleteResult,)*CascadingDeleteNetworkingV2ExtensionsLbaas_v2LoadbalancersResponse {
    return &CascadingDeleteNetworkingV2ExtensionsLbaas_v2LoadbalancersResponse{
            DeleteResult:deleteResult,
    }
}

// action function
func (oc *OpenstackClient) CascadingDeleteNetworkingV2ExtensionsLbaas_v2Loadbalancers(req *CascadingDeleteNetworkingV2ExtensionsLbaas_v2LoadbalancersRequest)(*CascadingDeleteNetworkingV2ExtensionsLbaas_v2LoadbalancersResponse){
    return NewCascadingDeleteNetworkingV2ExtensionsLbaas_v2LoadbalancersResponse(loadbalancers.CascadingDelete(oc.Client,req.Id, ))

}
//request struct for the GetStatusesNetworkingV2ExtensionsLbaas_v2Loadbalancers
type GetStatusesNetworkingV2ExtensionsLbaas_v2LoadbalancersRequest struct{
    Id string
}

func NewGetStatusesNetworkingV2ExtensionsLbaas_v2LoadbalancersRequest()*GetStatusesNetworkingV2ExtensionsLbaas_v2LoadbalancersRequest{
    return &GetStatusesNetworkingV2ExtensionsLbaas_v2LoadbalancersRequest{}
}

//response struct for the GetStatusesNetworkingV2ExtensionsLbaas_v2Loadbalancers
type GetStatusesNetworkingV2ExtensionsLbaas_v2LoadbalancersResponse struct{
    GetStatusesResult loadbalancers.GetStatusesResult
}

func NewGetStatusesNetworkingV2ExtensionsLbaas_v2LoadbalancersResponse(getStatusesResult loadbalancers.GetStatusesResult,)*GetStatusesNetworkingV2ExtensionsLbaas_v2LoadbalancersResponse {
    return &GetStatusesNetworkingV2ExtensionsLbaas_v2LoadbalancersResponse{
            GetStatusesResult:getStatusesResult,
    }
}

// action function
func (oc *OpenstackClient) GetStatusesNetworkingV2ExtensionsLbaas_v2Loadbalancers(req *GetStatusesNetworkingV2ExtensionsLbaas_v2LoadbalancersRequest)(*GetStatusesNetworkingV2ExtensionsLbaas_v2LoadbalancersResponse){
    return NewGetStatusesNetworkingV2ExtensionsLbaas_v2LoadbalancersResponse(loadbalancers.GetStatuses(oc.Client,req.Id, ))

}
//request struct for the GetStatsNetworkingV2ExtensionsLbaas_v2Loadbalancers
type GetStatsNetworkingV2ExtensionsLbaas_v2LoadbalancersRequest struct{
    Id string
}

func NewGetStatsNetworkingV2ExtensionsLbaas_v2LoadbalancersRequest()*GetStatsNetworkingV2ExtensionsLbaas_v2LoadbalancersRequest{
    return &GetStatsNetworkingV2ExtensionsLbaas_v2LoadbalancersRequest{}
}

//response struct for the GetStatsNetworkingV2ExtensionsLbaas_v2Loadbalancers
type GetStatsNetworkingV2ExtensionsLbaas_v2LoadbalancersResponse struct{
    StatsResult loadbalancers.StatsResult
}

func NewGetStatsNetworkingV2ExtensionsLbaas_v2LoadbalancersResponse(statsResult loadbalancers.StatsResult,)*GetStatsNetworkingV2ExtensionsLbaas_v2LoadbalancersResponse {
    return &GetStatsNetworkingV2ExtensionsLbaas_v2LoadbalancersResponse{
            StatsResult:statsResult,
    }
}

// action function
func (oc *OpenstackClient) GetStatsNetworkingV2ExtensionsLbaas_v2Loadbalancers(req *GetStatsNetworkingV2ExtensionsLbaas_v2LoadbalancersRequest)(*GetStatsNetworkingV2ExtensionsLbaas_v2LoadbalancersResponse){
    return NewGetStatsNetworkingV2ExtensionsLbaas_v2LoadbalancersResponse(loadbalancers.GetStats(oc.Client,req.Id, ))

}