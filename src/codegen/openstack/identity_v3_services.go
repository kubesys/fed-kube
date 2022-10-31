package openstack

// Code generated by cloud manager.





import (
    "github.com/gophercloud/gophercloud/openstack/identity/v3/services"
    "github.com/gophercloud/gophercloud/pagination"
)
//request struct for the CreateIdentityV3Services
type CreateIdentityV3ServicesRequest struct{
    Opts services.CreateOpts
}

func NewCreateIdentityV3ServicesRequest()*CreateIdentityV3ServicesRequest{
    return &CreateIdentityV3ServicesRequest{}
}

//response struct for the CreateIdentityV3Services
type CreateIdentityV3ServicesResponse struct{
    CreateResult services.CreateResult
}

func NewCreateIdentityV3ServicesResponse(createResult services.CreateResult,)*CreateIdentityV3ServicesResponse {
    return &CreateIdentityV3ServicesResponse{
            CreateResult:createResult,
    }
}

// action function
func (oc *OpenstackClient) CreateIdentityV3Services(req *CreateIdentityV3ServicesRequest)(*CreateIdentityV3ServicesResponse){
    return NewCreateIdentityV3ServicesResponse(services.Create(oc.Client,req.Opts, ))

}
//request struct for the ListIdentityV3Services
type ListIdentityV3ServicesRequest struct{
    Opts services.ListOpts
}

func NewListIdentityV3ServicesRequest()*ListIdentityV3ServicesRequest{
    return &ListIdentityV3ServicesRequest{}
}

//response struct for the ListIdentityV3Services
type ListIdentityV3ServicesResponse struct{
    Pager pagination.Pager
}

func NewListIdentityV3ServicesResponse(pager pagination.Pager,)*ListIdentityV3ServicesResponse {
    return &ListIdentityV3ServicesResponse{
            Pager:pager,
    }
}

// action function
func (oc *OpenstackClient) ListIdentityV3Services(req *ListIdentityV3ServicesRequest)(*ListIdentityV3ServicesResponse){
    return NewListIdentityV3ServicesResponse(services.List(oc.Client,req.Opts, ))

}
//request struct for the GetIdentityV3Services
type GetIdentityV3ServicesRequest struct{
    ServiceID string
}

func NewGetIdentityV3ServicesRequest()*GetIdentityV3ServicesRequest{
    return &GetIdentityV3ServicesRequest{}
}

//response struct for the GetIdentityV3Services
type GetIdentityV3ServicesResponse struct{
    GetResult services.GetResult
}

func NewGetIdentityV3ServicesResponse(getResult services.GetResult,)*GetIdentityV3ServicesResponse {
    return &GetIdentityV3ServicesResponse{
            GetResult:getResult,
    }
}

// action function
func (oc *OpenstackClient) GetIdentityV3Services(req *GetIdentityV3ServicesRequest)(*GetIdentityV3ServicesResponse){
    return NewGetIdentityV3ServicesResponse(services.Get(oc.Client,req.ServiceID, ))

}
//request struct for the UpdateIdentityV3Services
type UpdateIdentityV3ServicesRequest struct{
    ServiceID string
    Opts services.UpdateOpts
}

func NewUpdateIdentityV3ServicesRequest()*UpdateIdentityV3ServicesRequest{
    return &UpdateIdentityV3ServicesRequest{}
}

//response struct for the UpdateIdentityV3Services
type UpdateIdentityV3ServicesResponse struct{
    UpdateResult services.UpdateResult
}

func NewUpdateIdentityV3ServicesResponse(updateResult services.UpdateResult,)*UpdateIdentityV3ServicesResponse {
    return &UpdateIdentityV3ServicesResponse{
            UpdateResult:updateResult,
    }
}

// action function
func (oc *OpenstackClient) UpdateIdentityV3Services(req *UpdateIdentityV3ServicesRequest)(*UpdateIdentityV3ServicesResponse){
    return NewUpdateIdentityV3ServicesResponse(services.Update(oc.Client,req.ServiceID,req.Opts, ))

}
//request struct for the DeleteIdentityV3Services
type DeleteIdentityV3ServicesRequest struct{
    ServiceID string
}

func NewDeleteIdentityV3ServicesRequest()*DeleteIdentityV3ServicesRequest{
    return &DeleteIdentityV3ServicesRequest{}
}

//response struct for the DeleteIdentityV3Services
type DeleteIdentityV3ServicesResponse struct{
    DeleteResult services.DeleteResult
}

func NewDeleteIdentityV3ServicesResponse(deleteResult services.DeleteResult,)*DeleteIdentityV3ServicesResponse {
    return &DeleteIdentityV3ServicesResponse{
            DeleteResult:deleteResult,
    }
}

// action function
func (oc *OpenstackClient) DeleteIdentityV3Services(req *DeleteIdentityV3ServicesRequest)(*DeleteIdentityV3ServicesResponse){
    return NewDeleteIdentityV3ServicesResponse(services.Delete(oc.Client,req.ServiceID, ))

}