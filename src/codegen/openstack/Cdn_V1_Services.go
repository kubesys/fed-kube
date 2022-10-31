package openstack

// Code generated by cloud manager.





import (
    "github.com/gophercloud/gophercloud/openstack/cdn/v1/services"
    "github.com/gophercloud/gophercloud/pagination"
)
//request struct for the ListCdnV1Services
type ListCdnV1ServicesRequest struct{
    Opts services.ListOpts
}

func NewListCdnV1ServicesRequest()*ListCdnV1ServicesRequest{
    return &ListCdnV1ServicesRequest{}
}

//response struct for the ListCdnV1Services
type ListCdnV1ServicesResponse struct{
    Pager pagination.Pager
}

func NewListCdnV1ServicesResponse(pager pagination.Pager,)*ListCdnV1ServicesResponse {
    return &ListCdnV1ServicesResponse{
            Pager:pager,
    }
}

// action function
func (oc *OpenstackClient) ListCdnV1Services(req *ListCdnV1ServicesRequest)(*ListCdnV1ServicesResponse){
    return NewListCdnV1ServicesResponse(services.List(oc.Client,req.Opts, ))

}
//request struct for the CreateCdnV1Services
type CreateCdnV1ServicesRequest struct{
    Opts services.CreateOpts
}

func NewCreateCdnV1ServicesRequest()*CreateCdnV1ServicesRequest{
    return &CreateCdnV1ServicesRequest{}
}

//response struct for the CreateCdnV1Services
type CreateCdnV1ServicesResponse struct{
    CreateResult services.CreateResult
}

func NewCreateCdnV1ServicesResponse(createResult services.CreateResult,)*CreateCdnV1ServicesResponse {
    return &CreateCdnV1ServicesResponse{
            CreateResult:createResult,
    }
}

// action function
func (oc *OpenstackClient) CreateCdnV1Services(req *CreateCdnV1ServicesRequest)(*CreateCdnV1ServicesResponse){
    return NewCreateCdnV1ServicesResponse(services.Create(oc.Client,req.Opts, ))

}
//request struct for the GetCdnV1Services
type GetCdnV1ServicesRequest struct{
    IdOrURL string
}

func NewGetCdnV1ServicesRequest()*GetCdnV1ServicesRequest{
    return &GetCdnV1ServicesRequest{}
}

//response struct for the GetCdnV1Services
type GetCdnV1ServicesResponse struct{
    GetResult services.GetResult
}

func NewGetCdnV1ServicesResponse(getResult services.GetResult,)*GetCdnV1ServicesResponse {
    return &GetCdnV1ServicesResponse{
            GetResult:getResult,
    }
}

// action function
func (oc *OpenstackClient) GetCdnV1Services(req *GetCdnV1ServicesRequest)(*GetCdnV1ServicesResponse){
    return NewGetCdnV1ServicesResponse(services.Get(oc.Client,req.IdOrURL, ))

}
//request struct for the UpdateCdnV1Services
type UpdateCdnV1ServicesRequest struct{
    IdOrURL string
    Opts services.UpdateOpts
}

func NewUpdateCdnV1ServicesRequest()*UpdateCdnV1ServicesRequest{
    return &UpdateCdnV1ServicesRequest{}
}

//response struct for the UpdateCdnV1Services
type UpdateCdnV1ServicesResponse struct{
    UpdateResult services.UpdateResult
}

func NewUpdateCdnV1ServicesResponse(updateResult services.UpdateResult,)*UpdateCdnV1ServicesResponse {
    return &UpdateCdnV1ServicesResponse{
            UpdateResult:updateResult,
    }
}

// action function
func (oc *OpenstackClient) UpdateCdnV1Services(req *UpdateCdnV1ServicesRequest)(*UpdateCdnV1ServicesResponse){
    return NewUpdateCdnV1ServicesResponse(services.Update(oc.Client,req.IdOrURL,req.Opts, ))

}
//request struct for the DeleteCdnV1Services
type DeleteCdnV1ServicesRequest struct{
    IdOrURL string
}

func NewDeleteCdnV1ServicesRequest()*DeleteCdnV1ServicesRequest{
    return &DeleteCdnV1ServicesRequest{}
}

//response struct for the DeleteCdnV1Services
type DeleteCdnV1ServicesResponse struct{
    DeleteResult services.DeleteResult
}

func NewDeleteCdnV1ServicesResponse(deleteResult services.DeleteResult,)*DeleteCdnV1ServicesResponse {
    return &DeleteCdnV1ServicesResponse{
            DeleteResult:deleteResult,
    }
}

// action function
func (oc *OpenstackClient) DeleteCdnV1Services(req *DeleteCdnV1ServicesRequest)(*DeleteCdnV1ServicesResponse){
    return NewDeleteCdnV1ServicesResponse(services.Delete(oc.Client,req.IdOrURL, ))

}