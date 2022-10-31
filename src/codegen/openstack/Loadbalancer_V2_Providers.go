package openstack

// Code generated by cloud manager.





import (
    "github.com/gophercloud/gophercloud/openstack/loadbalancer/v2/providers"
    "github.com/gophercloud/gophercloud/pagination"
)
//request struct for the ListLoadbalancerV2Providers
type ListLoadbalancerV2ProvidersRequest struct{
    Opts providers.ListOpts
}

func NewListLoadbalancerV2ProvidersRequest()*ListLoadbalancerV2ProvidersRequest{
    return &ListLoadbalancerV2ProvidersRequest{}
}

//response struct for the ListLoadbalancerV2Providers
type ListLoadbalancerV2ProvidersResponse struct{
    Pager pagination.Pager
}

func NewListLoadbalancerV2ProvidersResponse(pager pagination.Pager,)*ListLoadbalancerV2ProvidersResponse {
    return &ListLoadbalancerV2ProvidersResponse{
            Pager:pager,
    }
}

// action function
func (oc *OpenstackClient) ListLoadbalancerV2Providers(req *ListLoadbalancerV2ProvidersRequest)(*ListLoadbalancerV2ProvidersResponse){
    return NewListLoadbalancerV2ProvidersResponse(providers.List(oc.Client,req.Opts, ))

}