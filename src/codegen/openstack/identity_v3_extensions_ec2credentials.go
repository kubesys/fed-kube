package openstack

// Code generated by cloud manager.





import (
    "github.com/gophercloud/gophercloud/openstack/identity/v3/extensions/ec2credentials"
    "github.com/gophercloud/gophercloud/pagination"
)
//request struct for the ListIdentityV3ExtensionsEc2credentials
type ListIdentityV3ExtensionsEc2credentialsRequest struct{
    UserID string
}

func NewListIdentityV3ExtensionsEc2credentialsRequest()*ListIdentityV3ExtensionsEc2credentialsRequest{
    return &ListIdentityV3ExtensionsEc2credentialsRequest{}
}

//response struct for the ListIdentityV3ExtensionsEc2credentials
type ListIdentityV3ExtensionsEc2credentialsResponse struct{
    Pager pagination.Pager
}

func NewListIdentityV3ExtensionsEc2credentialsResponse(pager pagination.Pager,)*ListIdentityV3ExtensionsEc2credentialsResponse {
    return &ListIdentityV3ExtensionsEc2credentialsResponse{
            Pager:pager,
    }
}

// action function
func (oc *OpenstackClient) ListIdentityV3ExtensionsEc2credentials(req *ListIdentityV3ExtensionsEc2credentialsRequest)(*ListIdentityV3ExtensionsEc2credentialsResponse){
    return NewListIdentityV3ExtensionsEc2credentialsResponse(ec2credentials.List(oc.Client,req.UserID, ))

}
//request struct for the GetIdentityV3ExtensionsEc2credentials
type GetIdentityV3ExtensionsEc2credentialsRequest struct{
    UserID string
    Id string
}

func NewGetIdentityV3ExtensionsEc2credentialsRequest()*GetIdentityV3ExtensionsEc2credentialsRequest{
    return &GetIdentityV3ExtensionsEc2credentialsRequest{}
}

//response struct for the GetIdentityV3ExtensionsEc2credentials
type GetIdentityV3ExtensionsEc2credentialsResponse struct{
    GetResult ec2credentials.GetResult
}

func NewGetIdentityV3ExtensionsEc2credentialsResponse(getResult ec2credentials.GetResult,)*GetIdentityV3ExtensionsEc2credentialsResponse {
    return &GetIdentityV3ExtensionsEc2credentialsResponse{
            GetResult:getResult,
    }
}

// action function
func (oc *OpenstackClient) GetIdentityV3ExtensionsEc2credentials(req *GetIdentityV3ExtensionsEc2credentialsRequest)(*GetIdentityV3ExtensionsEc2credentialsResponse){
    return NewGetIdentityV3ExtensionsEc2credentialsResponse(ec2credentials.Get(oc.Client,req.UserID,req.Id, ))

}
//request struct for the CreateIdentityV3ExtensionsEc2credentials
type CreateIdentityV3ExtensionsEc2credentialsRequest struct{
    UserID string
    Opts ec2credentials.CreateOpts
}

func NewCreateIdentityV3ExtensionsEc2credentialsRequest()*CreateIdentityV3ExtensionsEc2credentialsRequest{
    return &CreateIdentityV3ExtensionsEc2credentialsRequest{}
}

//response struct for the CreateIdentityV3ExtensionsEc2credentials
type CreateIdentityV3ExtensionsEc2credentialsResponse struct{
    CreateResult ec2credentials.CreateResult
}

func NewCreateIdentityV3ExtensionsEc2credentialsResponse(createResult ec2credentials.CreateResult,)*CreateIdentityV3ExtensionsEc2credentialsResponse {
    return &CreateIdentityV3ExtensionsEc2credentialsResponse{
            CreateResult:createResult,
    }
}

// action function
func (oc *OpenstackClient) CreateIdentityV3ExtensionsEc2credentials(req *CreateIdentityV3ExtensionsEc2credentialsRequest)(*CreateIdentityV3ExtensionsEc2credentialsResponse){
    return NewCreateIdentityV3ExtensionsEc2credentialsResponse(ec2credentials.Create(oc.Client,req.UserID,req.Opts, ))

}
//request struct for the DeleteIdentityV3ExtensionsEc2credentials
type DeleteIdentityV3ExtensionsEc2credentialsRequest struct{
    UserID string
    Id string
}

func NewDeleteIdentityV3ExtensionsEc2credentialsRequest()*DeleteIdentityV3ExtensionsEc2credentialsRequest{
    return &DeleteIdentityV3ExtensionsEc2credentialsRequest{}
}

//response struct for the DeleteIdentityV3ExtensionsEc2credentials
type DeleteIdentityV3ExtensionsEc2credentialsResponse struct{
    DeleteResult ec2credentials.DeleteResult
}

func NewDeleteIdentityV3ExtensionsEc2credentialsResponse(deleteResult ec2credentials.DeleteResult,)*DeleteIdentityV3ExtensionsEc2credentialsResponse {
    return &DeleteIdentityV3ExtensionsEc2credentialsResponse{
            DeleteResult:deleteResult,
    }
}

// action function
func (oc *OpenstackClient) DeleteIdentityV3ExtensionsEc2credentials(req *DeleteIdentityV3ExtensionsEc2credentialsRequest)(*DeleteIdentityV3ExtensionsEc2credentialsResponse){
    return NewDeleteIdentityV3ExtensionsEc2credentialsResponse(ec2credentials.Delete(oc.Client,req.UserID,req.Id, ))

}