package openstack

// Code generated by cloud manager.





import (
    "github.com/gophercloud/gophercloud/openstack/networking/v2/extensions/lbaas/monitors"
    "github.com/gophercloud/gophercloud/pagination"
)
//request struct for the ListNetworkingV2ExtensionsLbaasMonitors
type ListNetworkingV2ExtensionsLbaasMonitorsRequest struct{
    Opts monitors.ListOpts
}

func NewListNetworkingV2ExtensionsLbaasMonitorsRequest()*ListNetworkingV2ExtensionsLbaasMonitorsRequest{
    return &ListNetworkingV2ExtensionsLbaasMonitorsRequest{}
}

//response struct for the ListNetworkingV2ExtensionsLbaasMonitors
type ListNetworkingV2ExtensionsLbaasMonitorsResponse struct{
    Pager pagination.Pager
}

func NewListNetworkingV2ExtensionsLbaasMonitorsResponse(pager pagination.Pager,)*ListNetworkingV2ExtensionsLbaasMonitorsResponse {
    return &ListNetworkingV2ExtensionsLbaasMonitorsResponse{
            Pager:pager,
    }
}

// action function
func (oc *OpenstackClient) ListNetworkingV2ExtensionsLbaasMonitors(req *ListNetworkingV2ExtensionsLbaasMonitorsRequest)(*ListNetworkingV2ExtensionsLbaasMonitorsResponse){
    return NewListNetworkingV2ExtensionsLbaasMonitorsResponse(monitors.List(oc.Client,req.Opts, ))

}
//request struct for the CreateNetworkingV2ExtensionsLbaasMonitors
type CreateNetworkingV2ExtensionsLbaasMonitorsRequest struct{
    Opts monitors.CreateOpts
}

func NewCreateNetworkingV2ExtensionsLbaasMonitorsRequest()*CreateNetworkingV2ExtensionsLbaasMonitorsRequest{
    return &CreateNetworkingV2ExtensionsLbaasMonitorsRequest{}
}

//response struct for the CreateNetworkingV2ExtensionsLbaasMonitors
type CreateNetworkingV2ExtensionsLbaasMonitorsResponse struct{
    CreateResult monitors.CreateResult
}

func NewCreateNetworkingV2ExtensionsLbaasMonitorsResponse(createResult monitors.CreateResult,)*CreateNetworkingV2ExtensionsLbaasMonitorsResponse {
    return &CreateNetworkingV2ExtensionsLbaasMonitorsResponse{
            CreateResult:createResult,
    }
}

// action function
func (oc *OpenstackClient) CreateNetworkingV2ExtensionsLbaasMonitors(req *CreateNetworkingV2ExtensionsLbaasMonitorsRequest)(*CreateNetworkingV2ExtensionsLbaasMonitorsResponse){
    return NewCreateNetworkingV2ExtensionsLbaasMonitorsResponse(monitors.Create(oc.Client,req.Opts, ))

}
//request struct for the GetNetworkingV2ExtensionsLbaasMonitors
type GetNetworkingV2ExtensionsLbaasMonitorsRequest struct{
    Id string
}

func NewGetNetworkingV2ExtensionsLbaasMonitorsRequest()*GetNetworkingV2ExtensionsLbaasMonitorsRequest{
    return &GetNetworkingV2ExtensionsLbaasMonitorsRequest{}
}

//response struct for the GetNetworkingV2ExtensionsLbaasMonitors
type GetNetworkingV2ExtensionsLbaasMonitorsResponse struct{
    GetResult monitors.GetResult
}

func NewGetNetworkingV2ExtensionsLbaasMonitorsResponse(getResult monitors.GetResult,)*GetNetworkingV2ExtensionsLbaasMonitorsResponse {
    return &GetNetworkingV2ExtensionsLbaasMonitorsResponse{
            GetResult:getResult,
    }
}

// action function
func (oc *OpenstackClient) GetNetworkingV2ExtensionsLbaasMonitors(req *GetNetworkingV2ExtensionsLbaasMonitorsRequest)(*GetNetworkingV2ExtensionsLbaasMonitorsResponse){
    return NewGetNetworkingV2ExtensionsLbaasMonitorsResponse(monitors.Get(oc.Client,req.Id, ))

}
//request struct for the UpdateNetworkingV2ExtensionsLbaasMonitors
type UpdateNetworkingV2ExtensionsLbaasMonitorsRequest struct{
    Id string
    Opts monitors.UpdateOpts
}

func NewUpdateNetworkingV2ExtensionsLbaasMonitorsRequest()*UpdateNetworkingV2ExtensionsLbaasMonitorsRequest{
    return &UpdateNetworkingV2ExtensionsLbaasMonitorsRequest{}
}

//response struct for the UpdateNetworkingV2ExtensionsLbaasMonitors
type UpdateNetworkingV2ExtensionsLbaasMonitorsResponse struct{
    UpdateResult monitors.UpdateResult
}

func NewUpdateNetworkingV2ExtensionsLbaasMonitorsResponse(updateResult monitors.UpdateResult,)*UpdateNetworkingV2ExtensionsLbaasMonitorsResponse {
    return &UpdateNetworkingV2ExtensionsLbaasMonitorsResponse{
            UpdateResult:updateResult,
    }
}

// action function
func (oc *OpenstackClient) UpdateNetworkingV2ExtensionsLbaasMonitors(req *UpdateNetworkingV2ExtensionsLbaasMonitorsRequest)(*UpdateNetworkingV2ExtensionsLbaasMonitorsResponse){
    return NewUpdateNetworkingV2ExtensionsLbaasMonitorsResponse(monitors.Update(oc.Client,req.Id,req.Opts, ))

}
//request struct for the DeleteNetworkingV2ExtensionsLbaasMonitors
type DeleteNetworkingV2ExtensionsLbaasMonitorsRequest struct{
    Id string
}

func NewDeleteNetworkingV2ExtensionsLbaasMonitorsRequest()*DeleteNetworkingV2ExtensionsLbaasMonitorsRequest{
    return &DeleteNetworkingV2ExtensionsLbaasMonitorsRequest{}
}

//response struct for the DeleteNetworkingV2ExtensionsLbaasMonitors
type DeleteNetworkingV2ExtensionsLbaasMonitorsResponse struct{
    DeleteResult monitors.DeleteResult
}

func NewDeleteNetworkingV2ExtensionsLbaasMonitorsResponse(deleteResult monitors.DeleteResult,)*DeleteNetworkingV2ExtensionsLbaasMonitorsResponse {
    return &DeleteNetworkingV2ExtensionsLbaasMonitorsResponse{
            DeleteResult:deleteResult,
    }
}

// action function
func (oc *OpenstackClient) DeleteNetworkingV2ExtensionsLbaasMonitors(req *DeleteNetworkingV2ExtensionsLbaasMonitorsRequest)(*DeleteNetworkingV2ExtensionsLbaasMonitorsResponse){
    return NewDeleteNetworkingV2ExtensionsLbaasMonitorsResponse(monitors.Delete(oc.Client,req.Id, ))

}