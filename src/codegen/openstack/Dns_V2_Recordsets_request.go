package openstack

// Code generated by cloud manager.





import (
    "github.com/gophercloud/gophercloud/openstack/dns/v2/recordsets"
    "github.com/gophercloud/gophercloud/pagination"
)
//request struct for the ListByZoneDnsV2Recordsets
type ListByZoneDnsV2RecordsetsRequest struct{
    ZoneID string
    Opts recordsets.ListOpts
}

func NewListByZoneDnsV2RecordsetsRequest()*ListByZoneDnsV2RecordsetsRequest{
    return &ListByZoneDnsV2RecordsetsRequest{}
}

//response struct for the ListByZoneDnsV2Recordsets
type ListByZoneDnsV2RecordsetsResponse struct{
    Pager pagination.Pager
}

func NewListByZoneDnsV2RecordsetsResponse(pager pagination.Pager,)*ListByZoneDnsV2RecordsetsResponse {
    return &ListByZoneDnsV2RecordsetsResponse{
            Pager:pager,
    }
}

// action function
func (oc *OpenstackClient) ListByZoneDnsV2Recordsets(req *ListByZoneDnsV2RecordsetsRequest)(*ListByZoneDnsV2RecordsetsResponse){
    return NewListByZoneDnsV2RecordsetsResponse(recordsets.ListByZone(oc.Client,req.ZoneID,req.Opts, ))

}
//request struct for the GetDnsV2Recordsets
type GetDnsV2RecordsetsRequest struct{
    ZoneID string
    RrsetID string
}

func NewGetDnsV2RecordsetsRequest()*GetDnsV2RecordsetsRequest{
    return &GetDnsV2RecordsetsRequest{}
}

//response struct for the GetDnsV2Recordsets
type GetDnsV2RecordsetsResponse struct{
    GetResult recordsets.GetResult
}

func NewGetDnsV2RecordsetsResponse(getResult recordsets.GetResult,)*GetDnsV2RecordsetsResponse {
    return &GetDnsV2RecordsetsResponse{
            GetResult:getResult,
    }
}

// action function
func (oc *OpenstackClient) GetDnsV2Recordsets(req *GetDnsV2RecordsetsRequest)(*GetDnsV2RecordsetsResponse){
    return NewGetDnsV2RecordsetsResponse(recordsets.Get(oc.Client,req.ZoneID,req.RrsetID, ))

}
//request struct for the CreateDnsV2Recordsets
type CreateDnsV2RecordsetsRequest struct{
    ZoneID string
    Opts recordsets.CreateOpts
}

func NewCreateDnsV2RecordsetsRequest()*CreateDnsV2RecordsetsRequest{
    return &CreateDnsV2RecordsetsRequest{}
}

//response struct for the CreateDnsV2Recordsets
type CreateDnsV2RecordsetsResponse struct{
    CreateResult recordsets.CreateResult
}

func NewCreateDnsV2RecordsetsResponse(createResult recordsets.CreateResult,)*CreateDnsV2RecordsetsResponse {
    return &CreateDnsV2RecordsetsResponse{
            CreateResult:createResult,
    }
}

// action function
func (oc *OpenstackClient) CreateDnsV2Recordsets(req *CreateDnsV2RecordsetsRequest)(*CreateDnsV2RecordsetsResponse){
    return NewCreateDnsV2RecordsetsResponse(recordsets.Create(oc.Client,req.ZoneID,req.Opts, ))

}
//request struct for the UpdateDnsV2Recordsets
type UpdateDnsV2RecordsetsRequest struct{
    ZoneID string
    RrsetID string
    Opts recordsets.UpdateOpts
}

func NewUpdateDnsV2RecordsetsRequest()*UpdateDnsV2RecordsetsRequest{
    return &UpdateDnsV2RecordsetsRequest{}
}

//response struct for the UpdateDnsV2Recordsets
type UpdateDnsV2RecordsetsResponse struct{
    UpdateResult recordsets.UpdateResult
}

func NewUpdateDnsV2RecordsetsResponse(updateResult recordsets.UpdateResult,)*UpdateDnsV2RecordsetsResponse {
    return &UpdateDnsV2RecordsetsResponse{
            UpdateResult:updateResult,
    }
}

// action function
func (oc *OpenstackClient) UpdateDnsV2Recordsets(req *UpdateDnsV2RecordsetsRequest)(*UpdateDnsV2RecordsetsResponse){
    return NewUpdateDnsV2RecordsetsResponse(recordsets.Update(oc.Client,req.ZoneID,req.RrsetID,req.Opts, ))

}
//request struct for the DeleteDnsV2Recordsets
type DeleteDnsV2RecordsetsRequest struct{
    ZoneID string
    RrsetID string
}

func NewDeleteDnsV2RecordsetsRequest()*DeleteDnsV2RecordsetsRequest{
    return &DeleteDnsV2RecordsetsRequest{}
}

//response struct for the DeleteDnsV2Recordsets
type DeleteDnsV2RecordsetsResponse struct{
    DeleteResult recordsets.DeleteResult
}

func NewDeleteDnsV2RecordsetsResponse(deleteResult recordsets.DeleteResult,)*DeleteDnsV2RecordsetsResponse {
    return &DeleteDnsV2RecordsetsResponse{
            DeleteResult:deleteResult,
    }
}

// action function
func (oc *OpenstackClient) DeleteDnsV2Recordsets(req *DeleteDnsV2RecordsetsRequest)(*DeleteDnsV2RecordsetsResponse){
    return NewDeleteDnsV2RecordsetsResponse(recordsets.Delete(oc.Client,req.ZoneID,req.RrsetID, ))

}