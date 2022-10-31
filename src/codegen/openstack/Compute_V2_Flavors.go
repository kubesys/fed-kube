package openstack

// Code generated by cloud manager.





import (
    "github.com/gophercloud/gophercloud/openstack/compute/v2/flavors"
    "github.com/gophercloud/gophercloud/pagination"
)
//request struct for the ListDetailComputeV2Flavors
type ListDetailComputeV2FlavorsRequest struct{
    Opts flavors.ListOpts
}

func NewListDetailComputeV2FlavorsRequest()*ListDetailComputeV2FlavorsRequest{
    return &ListDetailComputeV2FlavorsRequest{}
}

//response struct for the ListDetailComputeV2Flavors
type ListDetailComputeV2FlavorsResponse struct{
    Pager pagination.Pager
}

func NewListDetailComputeV2FlavorsResponse(pager pagination.Pager,)*ListDetailComputeV2FlavorsResponse {
    return &ListDetailComputeV2FlavorsResponse{
            Pager:pager,
    }
}

// action function
func (oc *OpenstackClient) ListDetailComputeV2Flavors(req *ListDetailComputeV2FlavorsRequest)(*ListDetailComputeV2FlavorsResponse){
    return NewListDetailComputeV2FlavorsResponse(flavors.ListDetail(oc.Client,req.Opts, ))

}
//request struct for the CreateComputeV2Flavors
type CreateComputeV2FlavorsRequest struct{
    Opts flavors.CreateOpts
}

func NewCreateComputeV2FlavorsRequest()*CreateComputeV2FlavorsRequest{
    return &CreateComputeV2FlavorsRequest{}
}

//response struct for the CreateComputeV2Flavors
type CreateComputeV2FlavorsResponse struct{
    CreateResult flavors.CreateResult
}

func NewCreateComputeV2FlavorsResponse(createResult flavors.CreateResult,)*CreateComputeV2FlavorsResponse {
    return &CreateComputeV2FlavorsResponse{
            CreateResult:createResult,
    }
}

// action function
func (oc *OpenstackClient) CreateComputeV2Flavors(req *CreateComputeV2FlavorsRequest)(*CreateComputeV2FlavorsResponse){
    return NewCreateComputeV2FlavorsResponse(flavors.Create(oc.Client,req.Opts, ))

}
//request struct for the GetComputeV2Flavors
type GetComputeV2FlavorsRequest struct{
    Id string
}

func NewGetComputeV2FlavorsRequest()*GetComputeV2FlavorsRequest{
    return &GetComputeV2FlavorsRequest{}
}

//response struct for the GetComputeV2Flavors
type GetComputeV2FlavorsResponse struct{
    GetResult flavors.GetResult
}

func NewGetComputeV2FlavorsResponse(getResult flavors.GetResult,)*GetComputeV2FlavorsResponse {
    return &GetComputeV2FlavorsResponse{
            GetResult:getResult,
    }
}

// action function
func (oc *OpenstackClient) GetComputeV2Flavors(req *GetComputeV2FlavorsRequest)(*GetComputeV2FlavorsResponse){
    return NewGetComputeV2FlavorsResponse(flavors.Get(oc.Client,req.Id, ))

}
//request struct for the DeleteComputeV2Flavors
type DeleteComputeV2FlavorsRequest struct{
    Id string
}

func NewDeleteComputeV2FlavorsRequest()*DeleteComputeV2FlavorsRequest{
    return &DeleteComputeV2FlavorsRequest{}
}

//response struct for the DeleteComputeV2Flavors
type DeleteComputeV2FlavorsResponse struct{
    DeleteResult flavors.DeleteResult
}

func NewDeleteComputeV2FlavorsResponse(deleteResult flavors.DeleteResult,)*DeleteComputeV2FlavorsResponse {
    return &DeleteComputeV2FlavorsResponse{
            DeleteResult:deleteResult,
    }
}

// action function
func (oc *OpenstackClient) DeleteComputeV2Flavors(req *DeleteComputeV2FlavorsRequest)(*DeleteComputeV2FlavorsResponse){
    return NewDeleteComputeV2FlavorsResponse(flavors.Delete(oc.Client,req.Id, ))

}
//request struct for the ListAccessesComputeV2Flavors
type ListAccessesComputeV2FlavorsRequest struct{
    Id string
}

func NewListAccessesComputeV2FlavorsRequest()*ListAccessesComputeV2FlavorsRequest{
    return &ListAccessesComputeV2FlavorsRequest{}
}

//response struct for the ListAccessesComputeV2Flavors
type ListAccessesComputeV2FlavorsResponse struct{
    Pager pagination.Pager
}

func NewListAccessesComputeV2FlavorsResponse(pager pagination.Pager,)*ListAccessesComputeV2FlavorsResponse {
    return &ListAccessesComputeV2FlavorsResponse{
            Pager:pager,
    }
}

// action function
func (oc *OpenstackClient) ListAccessesComputeV2Flavors(req *ListAccessesComputeV2FlavorsRequest)(*ListAccessesComputeV2FlavorsResponse){
    return NewListAccessesComputeV2FlavorsResponse(flavors.ListAccesses(oc.Client,req.Id, ))

}
//request struct for the AddAccessComputeV2Flavors
type AddAccessComputeV2FlavorsRequest struct{
    Id string
    Opts flavors.AddAccessOpts
}

func NewAddAccessComputeV2FlavorsRequest()*AddAccessComputeV2FlavorsRequest{
    return &AddAccessComputeV2FlavorsRequest{}
}

//response struct for the AddAccessComputeV2Flavors
type AddAccessComputeV2FlavorsResponse struct{
    AddAccessResult flavors.AddAccessResult
}

func NewAddAccessComputeV2FlavorsResponse(addAccessResult flavors.AddAccessResult,)*AddAccessComputeV2FlavorsResponse {
    return &AddAccessComputeV2FlavorsResponse{
            AddAccessResult:addAccessResult,
    }
}

// action function
func (oc *OpenstackClient) AddAccessComputeV2Flavors(req *AddAccessComputeV2FlavorsRequest)(*AddAccessComputeV2FlavorsResponse){
    return NewAddAccessComputeV2FlavorsResponse(flavors.AddAccess(oc.Client,req.Id,req.Opts, ))

}
//request struct for the RemoveAccessComputeV2Flavors
type RemoveAccessComputeV2FlavorsRequest struct{
    Id string
    Opts flavors.RemoveAccessOpts
}

func NewRemoveAccessComputeV2FlavorsRequest()*RemoveAccessComputeV2FlavorsRequest{
    return &RemoveAccessComputeV2FlavorsRequest{}
}

//response struct for the RemoveAccessComputeV2Flavors
type RemoveAccessComputeV2FlavorsResponse struct{
    RemoveAccessResult flavors.RemoveAccessResult
}

func NewRemoveAccessComputeV2FlavorsResponse(removeAccessResult flavors.RemoveAccessResult,)*RemoveAccessComputeV2FlavorsResponse {
    return &RemoveAccessComputeV2FlavorsResponse{
            RemoveAccessResult:removeAccessResult,
    }
}

// action function
func (oc *OpenstackClient) RemoveAccessComputeV2Flavors(req *RemoveAccessComputeV2FlavorsRequest)(*RemoveAccessComputeV2FlavorsResponse){
    return NewRemoveAccessComputeV2FlavorsResponse(flavors.RemoveAccess(oc.Client,req.Id,req.Opts, ))

}
//request struct for the ListExtraSpecsComputeV2Flavors
type ListExtraSpecsComputeV2FlavorsRequest struct{
    FlavorID string
}

func NewListExtraSpecsComputeV2FlavorsRequest()*ListExtraSpecsComputeV2FlavorsRequest{
    return &ListExtraSpecsComputeV2FlavorsRequest{}
}

//response struct for the ListExtraSpecsComputeV2Flavors
type ListExtraSpecsComputeV2FlavorsResponse struct{
    ListExtraSpecsResult flavors.ListExtraSpecsResult
}

func NewListExtraSpecsComputeV2FlavorsResponse(listExtraSpecsResult flavors.ListExtraSpecsResult,)*ListExtraSpecsComputeV2FlavorsResponse {
    return &ListExtraSpecsComputeV2FlavorsResponse{
            ListExtraSpecsResult:listExtraSpecsResult,
    }
}

// action function
func (oc *OpenstackClient) ListExtraSpecsComputeV2Flavors(req *ListExtraSpecsComputeV2FlavorsRequest)(*ListExtraSpecsComputeV2FlavorsResponse){
    return NewListExtraSpecsComputeV2FlavorsResponse(flavors.ListExtraSpecs(oc.Client,req.FlavorID, ))

}
//request struct for the GetExtraSpecComputeV2Flavors
type GetExtraSpecComputeV2FlavorsRequest struct{
    FlavorID string
    Key string
}

func NewGetExtraSpecComputeV2FlavorsRequest()*GetExtraSpecComputeV2FlavorsRequest{
    return &GetExtraSpecComputeV2FlavorsRequest{}
}

//response struct for the GetExtraSpecComputeV2Flavors
type GetExtraSpecComputeV2FlavorsResponse struct{
    GetExtraSpecResult flavors.GetExtraSpecResult
}

func NewGetExtraSpecComputeV2FlavorsResponse(getExtraSpecResult flavors.GetExtraSpecResult,)*GetExtraSpecComputeV2FlavorsResponse {
    return &GetExtraSpecComputeV2FlavorsResponse{
            GetExtraSpecResult:getExtraSpecResult,
    }
}

// action function
func (oc *OpenstackClient) GetExtraSpecComputeV2Flavors(req *GetExtraSpecComputeV2FlavorsRequest)(*GetExtraSpecComputeV2FlavorsResponse){
    return NewGetExtraSpecComputeV2FlavorsResponse(flavors.GetExtraSpec(oc.Client,req.FlavorID,req.Key, ))

}
//request struct for the CreateExtraSpecsComputeV2Flavors
type CreateExtraSpecsComputeV2FlavorsRequest struct{
    FlavorID string
    Opts flavors.ExtraSpecsOpts
}

func NewCreateExtraSpecsComputeV2FlavorsRequest()*CreateExtraSpecsComputeV2FlavorsRequest{
    return &CreateExtraSpecsComputeV2FlavorsRequest{}
}

//response struct for the CreateExtraSpecsComputeV2Flavors
type CreateExtraSpecsComputeV2FlavorsResponse struct{
    CreateExtraSpecsResult flavors.CreateExtraSpecsResult
}

func NewCreateExtraSpecsComputeV2FlavorsResponse(createExtraSpecsResult flavors.CreateExtraSpecsResult,)*CreateExtraSpecsComputeV2FlavorsResponse {
    return &CreateExtraSpecsComputeV2FlavorsResponse{
            CreateExtraSpecsResult:createExtraSpecsResult,
    }
}

// action function
func (oc *OpenstackClient) CreateExtraSpecsComputeV2Flavors(req *CreateExtraSpecsComputeV2FlavorsRequest)(*CreateExtraSpecsComputeV2FlavorsResponse){
    return NewCreateExtraSpecsComputeV2FlavorsResponse(flavors.CreateExtraSpecs(oc.Client,req.FlavorID,req.Opts, ))

}
//request struct for the UpdateExtraSpecComputeV2Flavors
type UpdateExtraSpecComputeV2FlavorsRequest struct{
    FlavorID string
    Opts flavors.ExtraSpecsOpts
}

func NewUpdateExtraSpecComputeV2FlavorsRequest()*UpdateExtraSpecComputeV2FlavorsRequest{
    return &UpdateExtraSpecComputeV2FlavorsRequest{}
}

//response struct for the UpdateExtraSpecComputeV2Flavors
type UpdateExtraSpecComputeV2FlavorsResponse struct{
    UpdateExtraSpecResult flavors.UpdateExtraSpecResult
}

func NewUpdateExtraSpecComputeV2FlavorsResponse(updateExtraSpecResult flavors.UpdateExtraSpecResult,)*UpdateExtraSpecComputeV2FlavorsResponse {
    return &UpdateExtraSpecComputeV2FlavorsResponse{
            UpdateExtraSpecResult:updateExtraSpecResult,
    }
}

// action function
func (oc *OpenstackClient) UpdateExtraSpecComputeV2Flavors(req *UpdateExtraSpecComputeV2FlavorsRequest)(*UpdateExtraSpecComputeV2FlavorsResponse){
    return NewUpdateExtraSpecComputeV2FlavorsResponse(flavors.UpdateExtraSpec(oc.Client,req.FlavorID,req.Opts, ))

}
//request struct for the DeleteExtraSpecComputeV2Flavors
type DeleteExtraSpecComputeV2FlavorsRequest struct{
    FlavorID string
    Key string
}

func NewDeleteExtraSpecComputeV2FlavorsRequest()*DeleteExtraSpecComputeV2FlavorsRequest{
    return &DeleteExtraSpecComputeV2FlavorsRequest{}
}

//response struct for the DeleteExtraSpecComputeV2Flavors
type DeleteExtraSpecComputeV2FlavorsResponse struct{
    DeleteExtraSpecResult flavors.DeleteExtraSpecResult
}

func NewDeleteExtraSpecComputeV2FlavorsResponse(deleteExtraSpecResult flavors.DeleteExtraSpecResult,)*DeleteExtraSpecComputeV2FlavorsResponse {
    return &DeleteExtraSpecComputeV2FlavorsResponse{
            DeleteExtraSpecResult:deleteExtraSpecResult,
    }
}

// action function
func (oc *OpenstackClient) DeleteExtraSpecComputeV2Flavors(req *DeleteExtraSpecComputeV2FlavorsRequest)(*DeleteExtraSpecComputeV2FlavorsResponse){
    return NewDeleteExtraSpecComputeV2FlavorsResponse(flavors.DeleteExtraSpec(oc.Client,req.FlavorID,req.Key, ))

}