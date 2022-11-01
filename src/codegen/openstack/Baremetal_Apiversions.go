package openstack

// Code generated by cloud manager.





import (
    "github.com/gophercloud/gophercloud/openstack/baremetal/apiversions"
)
//request struct for the ListBaremetalApiversions
type ListBaremetalApiversionsRequest struct{
}

func NewListBaremetalApiversionsRequest()*ListBaremetalApiversionsRequest{
    return &ListBaremetalApiversionsRequest{}
}

//response struct for the ListBaremetalApiversions
type ListBaremetalApiversionsResponse struct{
    ListResult apiversions.ListResult
}

func NewListBaremetalApiversionsResponse(listResult apiversions.ListResult,)*ListBaremetalApiversionsResponse {
    return &ListBaremetalApiversionsResponse{
            ListResult:listResult,
    }
}

// action function
func (oc *OpenstackClient) ListBaremetalApiversions(req *ListBaremetalApiversionsRequest)(*ListBaremetalApiversionsResponse){
    return NewListBaremetalApiversionsResponse(apiversions.List(oc.Client, ))

}
//request struct for the GetBaremetalApiversions
type GetBaremetalApiversionsRequest struct{
    V string
}

func NewGetBaremetalApiversionsRequest()*GetBaremetalApiversionsRequest{
    return &GetBaremetalApiversionsRequest{}
}

//response struct for the GetBaremetalApiversions
type GetBaremetalApiversionsResponse struct{
    GetResult apiversions.GetResult
}

func NewGetBaremetalApiversionsResponse(getResult apiversions.GetResult,)*GetBaremetalApiversionsResponse {
    return &GetBaremetalApiversionsResponse{
            GetResult:getResult,
    }
}

// action function
func (oc *OpenstackClient) GetBaremetalApiversions(req *GetBaremetalApiversionsRequest)(*GetBaremetalApiversionsResponse){
    return NewGetBaremetalApiversionsResponse(apiversions.Get(oc.Client,req.V, ))

}