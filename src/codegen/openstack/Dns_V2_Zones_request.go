package openstack

// Code generated by cloud manager.

import (
    "github.com/gophercloud/gophercloud/openstack/dns/v2/zones"
    "github.com/gophercloud/gophercloud/pagination"
)
//request struct for the ListDnsV2Zones
type ListDnsV2ZonesRequest struct{
    Opts zones.ListOpts
}

func NewListDnsV2ZonesRequest()*ListDnsV2ZonesRequest{
    return &ListDnsV2ZonesRequest{}
}

//response struct for the ListDnsV2Zones
type ListDnsV2ZonesResponse struct{
    Pager pagination.Pager
}

func NewListDnsV2ZonesResponse(pager pagination.Pager,)*ListDnsV2ZonesResponse {
    return &ListDnsV2ZonesResponse{
            Pager:pager,
    }
}

// action function
func (oc *OpenstackClient) ListDnsV2Zones(req *ListDnsV2ZonesRequest)(*ListDnsV2ZonesResponse){
    return NewListDnsV2ZonesResponse(zones.List(oc.Client,req.Opts, ))

}
//request struct for the GetDnsV2Zones
type GetDnsV2ZonesRequest struct{
    ZoneID string
}

func NewGetDnsV2ZonesRequest()*GetDnsV2ZonesRequest{
    return &GetDnsV2ZonesRequest{}
}

//response struct for the GetDnsV2Zones
type GetDnsV2ZonesResponse struct{
    GetResult zones.GetResult
}

func NewGetDnsV2ZonesResponse(getResult zones.GetResult,)*GetDnsV2ZonesResponse {
    return &GetDnsV2ZonesResponse{
            GetResult:getResult,
    }
}

// action function
func (oc *OpenstackClient) GetDnsV2Zones(req *GetDnsV2ZonesRequest)(*GetDnsV2ZonesResponse){
    return NewGetDnsV2ZonesResponse(zones.Get(oc.Client,req.ZoneID, ))

}
//request struct for the CreateDnsV2Zones
type CreateDnsV2ZonesRequest struct{
    Opts zones.CreateOpts
}

func NewCreateDnsV2ZonesRequest()*CreateDnsV2ZonesRequest{
    return &CreateDnsV2ZonesRequest{}
}

//response struct for the CreateDnsV2Zones
type CreateDnsV2ZonesResponse struct{
    CreateResult zones.CreateResult
}

func NewCreateDnsV2ZonesResponse(createResult zones.CreateResult,)*CreateDnsV2ZonesResponse {
    return &CreateDnsV2ZonesResponse{
            CreateResult:createResult,
    }
}

// action function
func (oc *OpenstackClient) CreateDnsV2Zones(req *CreateDnsV2ZonesRequest)(*CreateDnsV2ZonesResponse){
    return NewCreateDnsV2ZonesResponse(zones.Create(oc.Client,req.Opts, ))

}
//request struct for the UpdateDnsV2Zones
type UpdateDnsV2ZonesRequest struct{
    ZoneID string
    Opts zones.UpdateOpts
}

func NewUpdateDnsV2ZonesRequest()*UpdateDnsV2ZonesRequest{
    return &UpdateDnsV2ZonesRequest{}
}

//response struct for the UpdateDnsV2Zones
type UpdateDnsV2ZonesResponse struct{
    UpdateResult zones.UpdateResult
}

func NewUpdateDnsV2ZonesResponse(updateResult zones.UpdateResult,)*UpdateDnsV2ZonesResponse {
    return &UpdateDnsV2ZonesResponse{
            UpdateResult:updateResult,
    }
}

// action function
func (oc *OpenstackClient) UpdateDnsV2Zones(req *UpdateDnsV2ZonesRequest)(*UpdateDnsV2ZonesResponse){
    return NewUpdateDnsV2ZonesResponse(zones.Update(oc.Client,req.ZoneID,req.Opts, ))

}
//request struct for the DeleteDnsV2Zones
type DeleteDnsV2ZonesRequest struct{
    ZoneID string
}

func NewDeleteDnsV2ZonesRequest()*DeleteDnsV2ZonesRequest{
    return &DeleteDnsV2ZonesRequest{}
}

//response struct for the DeleteDnsV2Zones
type DeleteDnsV2ZonesResponse struct{
    DeleteResult zones.DeleteResult
}

func NewDeleteDnsV2ZonesResponse(deleteResult zones.DeleteResult,)*DeleteDnsV2ZonesResponse {
    return &DeleteDnsV2ZonesResponse{
            DeleteResult:deleteResult,
    }
}

// action function
func (oc *OpenstackClient) DeleteDnsV2Zones(req *DeleteDnsV2ZonesRequest)(*DeleteDnsV2ZonesResponse){
    return NewDeleteDnsV2ZonesResponse(zones.Delete(oc.Client,req.ZoneID, ))

}