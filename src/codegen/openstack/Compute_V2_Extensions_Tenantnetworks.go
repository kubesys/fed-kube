package openstack

// Code generated by cloud manager.





import (
    "github.com/gophercloud/gophercloud/openstack/compute/v2/extensions/tenantnetworks"
    "github.com/gophercloud/gophercloud/pagination"
)
//request struct for the ListComputeV2ExtensionsTenantnetworks
type ListComputeV2ExtensionsTenantnetworksRequest struct{
}

func NewListComputeV2ExtensionsTenantnetworksRequest()*ListComputeV2ExtensionsTenantnetworksRequest{
    return &ListComputeV2ExtensionsTenantnetworksRequest{}
}

//response struct for the ListComputeV2ExtensionsTenantnetworks
type ListComputeV2ExtensionsTenantnetworksResponse struct{
    Pager pagination.Pager
}

func NewListComputeV2ExtensionsTenantnetworksResponse(pager pagination.Pager,)*ListComputeV2ExtensionsTenantnetworksResponse {
    return &ListComputeV2ExtensionsTenantnetworksResponse{
            Pager:pager,
    }
}

// action function
func (oc *OpenstackClient) ListComputeV2ExtensionsTenantnetworks(req *ListComputeV2ExtensionsTenantnetworksRequest)(*ListComputeV2ExtensionsTenantnetworksResponse){
    return NewListComputeV2ExtensionsTenantnetworksResponse(tenantnetworks.List(oc.Client, ))

}
//request struct for the GetComputeV2ExtensionsTenantnetworks
type GetComputeV2ExtensionsTenantnetworksRequest struct{
    Id string
}

func NewGetComputeV2ExtensionsTenantnetworksRequest()*GetComputeV2ExtensionsTenantnetworksRequest{
    return &GetComputeV2ExtensionsTenantnetworksRequest{}
}

//response struct for the GetComputeV2ExtensionsTenantnetworks
type GetComputeV2ExtensionsTenantnetworksResponse struct{
    GetResult tenantnetworks.GetResult
}

func NewGetComputeV2ExtensionsTenantnetworksResponse(getResult tenantnetworks.GetResult,)*GetComputeV2ExtensionsTenantnetworksResponse {
    return &GetComputeV2ExtensionsTenantnetworksResponse{
            GetResult:getResult,
    }
}

// action function
func (oc *OpenstackClient) GetComputeV2ExtensionsTenantnetworks(req *GetComputeV2ExtensionsTenantnetworksRequest)(*GetComputeV2ExtensionsTenantnetworksResponse){
    return NewGetComputeV2ExtensionsTenantnetworksResponse(tenantnetworks.Get(oc.Client,req.Id, ))

}