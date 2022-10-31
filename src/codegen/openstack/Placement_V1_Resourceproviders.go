package openstack

// Code generated by cloud manager.





import (
    "github.com/gophercloud/gophercloud/openstack/placement/v1/resourceproviders"
    "github.com/gophercloud/gophercloud/pagination"
)
//request struct for the ListPlacementV1Resourceproviders
type ListPlacementV1ResourceprovidersRequest struct{
    Opts resourceproviders.ListOpts
}

func NewListPlacementV1ResourceprovidersRequest()*ListPlacementV1ResourceprovidersRequest{
    return &ListPlacementV1ResourceprovidersRequest{}
}

//response struct for the ListPlacementV1Resourceproviders
type ListPlacementV1ResourceprovidersResponse struct{
    Pager pagination.Pager
}

func NewListPlacementV1ResourceprovidersResponse(pager pagination.Pager,)*ListPlacementV1ResourceprovidersResponse {
    return &ListPlacementV1ResourceprovidersResponse{
            Pager:pager,
    }
}

// action function
func (oc *OpenstackClient) ListPlacementV1Resourceproviders(req *ListPlacementV1ResourceprovidersRequest)(*ListPlacementV1ResourceprovidersResponse){
    return NewListPlacementV1ResourceprovidersResponse(resourceproviders.List(oc.Client,req.Opts, ))

}
//request struct for the CreatePlacementV1Resourceproviders
type CreatePlacementV1ResourceprovidersRequest struct{
    Opts resourceproviders.CreateOpts
}

func NewCreatePlacementV1ResourceprovidersRequest()*CreatePlacementV1ResourceprovidersRequest{
    return &CreatePlacementV1ResourceprovidersRequest{}
}

//response struct for the CreatePlacementV1Resourceproviders
type CreatePlacementV1ResourceprovidersResponse struct{
    CreateResult resourceproviders.CreateResult
}

func NewCreatePlacementV1ResourceprovidersResponse(createResult resourceproviders.CreateResult,)*CreatePlacementV1ResourceprovidersResponse {
    return &CreatePlacementV1ResourceprovidersResponse{
            CreateResult:createResult,
    }
}

// action function
func (oc *OpenstackClient) CreatePlacementV1Resourceproviders(req *CreatePlacementV1ResourceprovidersRequest)(*CreatePlacementV1ResourceprovidersResponse){
    return NewCreatePlacementV1ResourceprovidersResponse(resourceproviders.Create(oc.Client,req.Opts, ))

}
//request struct for the DeletePlacementV1Resourceproviders
type DeletePlacementV1ResourceprovidersRequest struct{
    ResourceProviderID string
}

func NewDeletePlacementV1ResourceprovidersRequest()*DeletePlacementV1ResourceprovidersRequest{
    return &DeletePlacementV1ResourceprovidersRequest{}
}

//response struct for the DeletePlacementV1Resourceproviders
type DeletePlacementV1ResourceprovidersResponse struct{
    DeleteResult resourceproviders.DeleteResult
}

func NewDeletePlacementV1ResourceprovidersResponse(deleteResult resourceproviders.DeleteResult,)*DeletePlacementV1ResourceprovidersResponse {
    return &DeletePlacementV1ResourceprovidersResponse{
            DeleteResult:deleteResult,
    }
}

// action function
func (oc *OpenstackClient) DeletePlacementV1Resourceproviders(req *DeletePlacementV1ResourceprovidersRequest)(*DeletePlacementV1ResourceprovidersResponse){
    return NewDeletePlacementV1ResourceprovidersResponse(resourceproviders.Delete(oc.Client,req.ResourceProviderID, ))

}
//request struct for the GetPlacementV1Resourceproviders
type GetPlacementV1ResourceprovidersRequest struct{
    ResourceProviderID string
}

func NewGetPlacementV1ResourceprovidersRequest()*GetPlacementV1ResourceprovidersRequest{
    return &GetPlacementV1ResourceprovidersRequest{}
}

//response struct for the GetPlacementV1Resourceproviders
type GetPlacementV1ResourceprovidersResponse struct{
    GetResult resourceproviders.GetResult
}

func NewGetPlacementV1ResourceprovidersResponse(getResult resourceproviders.GetResult,)*GetPlacementV1ResourceprovidersResponse {
    return &GetPlacementV1ResourceprovidersResponse{
            GetResult:getResult,
    }
}

// action function
func (oc *OpenstackClient) GetPlacementV1Resourceproviders(req *GetPlacementV1ResourceprovidersRequest)(*GetPlacementV1ResourceprovidersResponse){
    return NewGetPlacementV1ResourceprovidersResponse(resourceproviders.Get(oc.Client,req.ResourceProviderID, ))

}
//request struct for the UpdatePlacementV1Resourceproviders
type UpdatePlacementV1ResourceprovidersRequest struct{
    ResourceProviderID string
    Opts resourceproviders.UpdateOpts
}

func NewUpdatePlacementV1ResourceprovidersRequest()*UpdatePlacementV1ResourceprovidersRequest{
    return &UpdatePlacementV1ResourceprovidersRequest{}
}

//response struct for the UpdatePlacementV1Resourceproviders
type UpdatePlacementV1ResourceprovidersResponse struct{
    UpdateResult resourceproviders.UpdateResult
}

func NewUpdatePlacementV1ResourceprovidersResponse(updateResult resourceproviders.UpdateResult,)*UpdatePlacementV1ResourceprovidersResponse {
    return &UpdatePlacementV1ResourceprovidersResponse{
            UpdateResult:updateResult,
    }
}

// action function
func (oc *OpenstackClient) UpdatePlacementV1Resourceproviders(req *UpdatePlacementV1ResourceprovidersRequest)(*UpdatePlacementV1ResourceprovidersResponse){
    return NewUpdatePlacementV1ResourceprovidersResponse(resourceproviders.Update(oc.Client,req.ResourceProviderID,req.Opts, ))

}
//request struct for the GetUsagesPlacementV1Resourceproviders
type GetUsagesPlacementV1ResourceprovidersRequest struct{
    ResourceProviderID string
}

func NewGetUsagesPlacementV1ResourceprovidersRequest()*GetUsagesPlacementV1ResourceprovidersRequest{
    return &GetUsagesPlacementV1ResourceprovidersRequest{}
}

//response struct for the GetUsagesPlacementV1Resourceproviders
type GetUsagesPlacementV1ResourceprovidersResponse struct{
    GetUsagesResult resourceproviders.GetUsagesResult
}

func NewGetUsagesPlacementV1ResourceprovidersResponse(getUsagesResult resourceproviders.GetUsagesResult,)*GetUsagesPlacementV1ResourceprovidersResponse {
    return &GetUsagesPlacementV1ResourceprovidersResponse{
            GetUsagesResult:getUsagesResult,
    }
}

// action function
func (oc *OpenstackClient) GetUsagesPlacementV1Resourceproviders(req *GetUsagesPlacementV1ResourceprovidersRequest)(*GetUsagesPlacementV1ResourceprovidersResponse){
    return NewGetUsagesPlacementV1ResourceprovidersResponse(resourceproviders.GetUsages(oc.Client,req.ResourceProviderID, ))

}
//request struct for the GetInventoriesPlacementV1Resourceproviders
type GetInventoriesPlacementV1ResourceprovidersRequest struct{
    ResourceProviderID string
}

func NewGetInventoriesPlacementV1ResourceprovidersRequest()*GetInventoriesPlacementV1ResourceprovidersRequest{
    return &GetInventoriesPlacementV1ResourceprovidersRequest{}
}

//response struct for the GetInventoriesPlacementV1Resourceproviders
type GetInventoriesPlacementV1ResourceprovidersResponse struct{
    GetInventoriesResult resourceproviders.GetInventoriesResult
}

func NewGetInventoriesPlacementV1ResourceprovidersResponse(getInventoriesResult resourceproviders.GetInventoriesResult,)*GetInventoriesPlacementV1ResourceprovidersResponse {
    return &GetInventoriesPlacementV1ResourceprovidersResponse{
            GetInventoriesResult:getInventoriesResult,
    }
}

// action function
func (oc *OpenstackClient) GetInventoriesPlacementV1Resourceproviders(req *GetInventoriesPlacementV1ResourceprovidersRequest)(*GetInventoriesPlacementV1ResourceprovidersResponse){
    return NewGetInventoriesPlacementV1ResourceprovidersResponse(resourceproviders.GetInventories(oc.Client,req.ResourceProviderID, ))

}
//request struct for the GetAllocationsPlacementV1Resourceproviders
type GetAllocationsPlacementV1ResourceprovidersRequest struct{
    ResourceProviderID string
}

func NewGetAllocationsPlacementV1ResourceprovidersRequest()*GetAllocationsPlacementV1ResourceprovidersRequest{
    return &GetAllocationsPlacementV1ResourceprovidersRequest{}
}

//response struct for the GetAllocationsPlacementV1Resourceproviders
type GetAllocationsPlacementV1ResourceprovidersResponse struct{
    GetAllocationsResult resourceproviders.GetAllocationsResult
}

func NewGetAllocationsPlacementV1ResourceprovidersResponse(getAllocationsResult resourceproviders.GetAllocationsResult,)*GetAllocationsPlacementV1ResourceprovidersResponse {
    return &GetAllocationsPlacementV1ResourceprovidersResponse{
            GetAllocationsResult:getAllocationsResult,
    }
}

// action function
func (oc *OpenstackClient) GetAllocationsPlacementV1Resourceproviders(req *GetAllocationsPlacementV1ResourceprovidersRequest)(*GetAllocationsPlacementV1ResourceprovidersResponse){
    return NewGetAllocationsPlacementV1ResourceprovidersResponse(resourceproviders.GetAllocations(oc.Client,req.ResourceProviderID, ))

}
//request struct for the GetTraitsPlacementV1Resourceproviders
type GetTraitsPlacementV1ResourceprovidersRequest struct{
    ResourceProviderID string
}

func NewGetTraitsPlacementV1ResourceprovidersRequest()*GetTraitsPlacementV1ResourceprovidersRequest{
    return &GetTraitsPlacementV1ResourceprovidersRequest{}
}

//response struct for the GetTraitsPlacementV1Resourceproviders
type GetTraitsPlacementV1ResourceprovidersResponse struct{
    GetTraitsResult resourceproviders.GetTraitsResult
}

func NewGetTraitsPlacementV1ResourceprovidersResponse(getTraitsResult resourceproviders.GetTraitsResult,)*GetTraitsPlacementV1ResourceprovidersResponse {
    return &GetTraitsPlacementV1ResourceprovidersResponse{
            GetTraitsResult:getTraitsResult,
    }
}

// action function
func (oc *OpenstackClient) GetTraitsPlacementV1Resourceproviders(req *GetTraitsPlacementV1ResourceprovidersRequest)(*GetTraitsPlacementV1ResourceprovidersResponse){
    return NewGetTraitsPlacementV1ResourceprovidersResponse(resourceproviders.GetTraits(oc.Client,req.ResourceProviderID, ))

}