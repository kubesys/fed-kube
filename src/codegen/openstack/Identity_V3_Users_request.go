package openstack

// Code generated by cloud manager.

import (
    "github.com/gophercloud/gophercloud/openstack/identity/v3/users"
    "github.com/gophercloud/gophercloud/pagination"
)
//request struct for the ListIdentityV3Users
type ListIdentityV3UsersRequest struct{
    Opts users.ListOpts
}

func NewListIdentityV3UsersRequest()*ListIdentityV3UsersRequest{
    return &ListIdentityV3UsersRequest{}
}

//response struct for the ListIdentityV3Users
type ListIdentityV3UsersResponse struct{
    Pager pagination.Pager
}

func NewListIdentityV3UsersResponse(pager pagination.Pager,)*ListIdentityV3UsersResponse {
    return &ListIdentityV3UsersResponse{
            Pager:pager,
    }
}

// action function
func (oc *OpenstackClient) ListIdentityV3Users(req *ListIdentityV3UsersRequest)(*ListIdentityV3UsersResponse){
    return NewListIdentityV3UsersResponse(users.List(oc.Client,req.Opts, ))

}
//request struct for the GetIdentityV3Users
type GetIdentityV3UsersRequest struct{
    Id string
}

func NewGetIdentityV3UsersRequest()*GetIdentityV3UsersRequest{
    return &GetIdentityV3UsersRequest{}
}

//response struct for the GetIdentityV3Users
type GetIdentityV3UsersResponse struct{
    GetResult users.GetResult
}

func NewGetIdentityV3UsersResponse(getResult users.GetResult,)*GetIdentityV3UsersResponse {
    return &GetIdentityV3UsersResponse{
            GetResult:getResult,
    }
}

// action function
func (oc *OpenstackClient) GetIdentityV3Users(req *GetIdentityV3UsersRequest)(*GetIdentityV3UsersResponse){
    return NewGetIdentityV3UsersResponse(users.Get(oc.Client,req.Id, ))

}
//request struct for the CreateIdentityV3Users
type CreateIdentityV3UsersRequest struct{
    Opts users.CreateOpts
}

func NewCreateIdentityV3UsersRequest()*CreateIdentityV3UsersRequest{
    return &CreateIdentityV3UsersRequest{}
}

//response struct for the CreateIdentityV3Users
type CreateIdentityV3UsersResponse struct{
    CreateResult users.CreateResult
}

func NewCreateIdentityV3UsersResponse(createResult users.CreateResult,)*CreateIdentityV3UsersResponse {
    return &CreateIdentityV3UsersResponse{
            CreateResult:createResult,
    }
}

// action function
func (oc *OpenstackClient) CreateIdentityV3Users(req *CreateIdentityV3UsersRequest)(*CreateIdentityV3UsersResponse){
    return NewCreateIdentityV3UsersResponse(users.Create(oc.Client,req.Opts, ))

}
//request struct for the UpdateIdentityV3Users
type UpdateIdentityV3UsersRequest struct{
    UserID string
    Opts users.UpdateOpts
}

func NewUpdateIdentityV3UsersRequest()*UpdateIdentityV3UsersRequest{
    return &UpdateIdentityV3UsersRequest{}
}

//response struct for the UpdateIdentityV3Users
type UpdateIdentityV3UsersResponse struct{
    UpdateResult users.UpdateResult
}

func NewUpdateIdentityV3UsersResponse(updateResult users.UpdateResult,)*UpdateIdentityV3UsersResponse {
    return &UpdateIdentityV3UsersResponse{
            UpdateResult:updateResult,
    }
}

// action function
func (oc *OpenstackClient) UpdateIdentityV3Users(req *UpdateIdentityV3UsersRequest)(*UpdateIdentityV3UsersResponse){
    return NewUpdateIdentityV3UsersResponse(users.Update(oc.Client,req.UserID,req.Opts, ))

}
//request struct for the ChangePasswordIdentityV3Users
type ChangePasswordIdentityV3UsersRequest struct{
    UserID string
    Opts users.ChangePasswordOpts
}

func NewChangePasswordIdentityV3UsersRequest()*ChangePasswordIdentityV3UsersRequest{
    return &ChangePasswordIdentityV3UsersRequest{}
}

//response struct for the ChangePasswordIdentityV3Users
type ChangePasswordIdentityV3UsersResponse struct{
    ChangePasswordResult users.ChangePasswordResult
}

func NewChangePasswordIdentityV3UsersResponse(changePasswordResult users.ChangePasswordResult,)*ChangePasswordIdentityV3UsersResponse {
    return &ChangePasswordIdentityV3UsersResponse{
            ChangePasswordResult:changePasswordResult,
    }
}

// action function
func (oc *OpenstackClient) ChangePasswordIdentityV3Users(req *ChangePasswordIdentityV3UsersRequest)(*ChangePasswordIdentityV3UsersResponse){
    return NewChangePasswordIdentityV3UsersResponse(users.ChangePassword(oc.Client,req.UserID,req.Opts, ))

}
//request struct for the DeleteIdentityV3Users
type DeleteIdentityV3UsersRequest struct{
    UserID string
}

func NewDeleteIdentityV3UsersRequest()*DeleteIdentityV3UsersRequest{
    return &DeleteIdentityV3UsersRequest{}
}

//response struct for the DeleteIdentityV3Users
type DeleteIdentityV3UsersResponse struct{
    DeleteResult users.DeleteResult
}

func NewDeleteIdentityV3UsersResponse(deleteResult users.DeleteResult,)*DeleteIdentityV3UsersResponse {
    return &DeleteIdentityV3UsersResponse{
            DeleteResult:deleteResult,
    }
}

// action function
func (oc *OpenstackClient) DeleteIdentityV3Users(req *DeleteIdentityV3UsersRequest)(*DeleteIdentityV3UsersResponse){
    return NewDeleteIdentityV3UsersResponse(users.Delete(oc.Client,req.UserID, ))

}
//request struct for the ListGroupsIdentityV3Users
type ListGroupsIdentityV3UsersRequest struct{
    UserID string
}

func NewListGroupsIdentityV3UsersRequest()*ListGroupsIdentityV3UsersRequest{
    return &ListGroupsIdentityV3UsersRequest{}
}

//response struct for the ListGroupsIdentityV3Users
type ListGroupsIdentityV3UsersResponse struct{
    Pager pagination.Pager
}

func NewListGroupsIdentityV3UsersResponse(pager pagination.Pager,)*ListGroupsIdentityV3UsersResponse {
    return &ListGroupsIdentityV3UsersResponse{
            Pager:pager,
    }
}

// action function
func (oc *OpenstackClient) ListGroupsIdentityV3Users(req *ListGroupsIdentityV3UsersRequest)(*ListGroupsIdentityV3UsersResponse){
    return NewListGroupsIdentityV3UsersResponse(users.ListGroups(oc.Client,req.UserID, ))

}
//request struct for the AddToGroupIdentityV3Users
type AddToGroupIdentityV3UsersRequest struct{
    GroupID string
    UserID string
}

func NewAddToGroupIdentityV3UsersRequest()*AddToGroupIdentityV3UsersRequest{
    return &AddToGroupIdentityV3UsersRequest{}
}

//response struct for the AddToGroupIdentityV3Users
type AddToGroupIdentityV3UsersResponse struct{
    AddToGroupResult users.AddToGroupResult
}

func NewAddToGroupIdentityV3UsersResponse(addToGroupResult users.AddToGroupResult,)*AddToGroupIdentityV3UsersResponse {
    return &AddToGroupIdentityV3UsersResponse{
            AddToGroupResult:addToGroupResult,
    }
}

// action function
func (oc *OpenstackClient) AddToGroupIdentityV3Users(req *AddToGroupIdentityV3UsersRequest)(*AddToGroupIdentityV3UsersResponse){
    return NewAddToGroupIdentityV3UsersResponse(users.AddToGroup(oc.Client,req.GroupID,req.UserID, ))

}
//request struct for the IsMemberOfGroupIdentityV3Users
type IsMemberOfGroupIdentityV3UsersRequest struct{
    GroupID string
    UserID string
}

func NewIsMemberOfGroupIdentityV3UsersRequest()*IsMemberOfGroupIdentityV3UsersRequest{
    return &IsMemberOfGroupIdentityV3UsersRequest{}
}

//response struct for the IsMemberOfGroupIdentityV3Users
type IsMemberOfGroupIdentityV3UsersResponse struct{
    IsMemberOfGroupResult users.IsMemberOfGroupResult
}

func NewIsMemberOfGroupIdentityV3UsersResponse(isMemberOfGroupResult users.IsMemberOfGroupResult,)*IsMemberOfGroupIdentityV3UsersResponse {
    return &IsMemberOfGroupIdentityV3UsersResponse{
            IsMemberOfGroupResult:isMemberOfGroupResult,
    }
}

// action function
func (oc *OpenstackClient) IsMemberOfGroupIdentityV3Users(req *IsMemberOfGroupIdentityV3UsersRequest)(*IsMemberOfGroupIdentityV3UsersResponse){
    return NewIsMemberOfGroupIdentityV3UsersResponse(users.IsMemberOfGroup(oc.Client,req.GroupID,req.UserID, ))

}
//request struct for the RemoveFromGroupIdentityV3Users
type RemoveFromGroupIdentityV3UsersRequest struct{
    GroupID string
    UserID string
}

func NewRemoveFromGroupIdentityV3UsersRequest()*RemoveFromGroupIdentityV3UsersRequest{
    return &RemoveFromGroupIdentityV3UsersRequest{}
}

//response struct for the RemoveFromGroupIdentityV3Users
type RemoveFromGroupIdentityV3UsersResponse struct{
    RemoveFromGroupResult users.RemoveFromGroupResult
}

func NewRemoveFromGroupIdentityV3UsersResponse(removeFromGroupResult users.RemoveFromGroupResult,)*RemoveFromGroupIdentityV3UsersResponse {
    return &RemoveFromGroupIdentityV3UsersResponse{
            RemoveFromGroupResult:removeFromGroupResult,
    }
}

// action function
func (oc *OpenstackClient) RemoveFromGroupIdentityV3Users(req *RemoveFromGroupIdentityV3UsersRequest)(*RemoveFromGroupIdentityV3UsersResponse){
    return NewRemoveFromGroupIdentityV3UsersResponse(users.RemoveFromGroup(oc.Client,req.GroupID,req.UserID, ))

}
//request struct for the ListProjectsIdentityV3Users
type ListProjectsIdentityV3UsersRequest struct{
    UserID string
}

func NewListProjectsIdentityV3UsersRequest()*ListProjectsIdentityV3UsersRequest{
    return &ListProjectsIdentityV3UsersRequest{}
}

//response struct for the ListProjectsIdentityV3Users
type ListProjectsIdentityV3UsersResponse struct{
    Pager pagination.Pager
}

func NewListProjectsIdentityV3UsersResponse(pager pagination.Pager,)*ListProjectsIdentityV3UsersResponse {
    return &ListProjectsIdentityV3UsersResponse{
            Pager:pager,
    }
}

// action function
func (oc *OpenstackClient) ListProjectsIdentityV3Users(req *ListProjectsIdentityV3UsersRequest)(*ListProjectsIdentityV3UsersResponse){
    return NewListProjectsIdentityV3UsersResponse(users.ListProjects(oc.Client,req.UserID, ))

}
//request struct for the ListInGroupIdentityV3Users
type ListInGroupIdentityV3UsersRequest struct{
    GroupID string
    Opts users.ListOpts
}

func NewListInGroupIdentityV3UsersRequest()*ListInGroupIdentityV3UsersRequest{
    return &ListInGroupIdentityV3UsersRequest{}
}

//response struct for the ListInGroupIdentityV3Users
type ListInGroupIdentityV3UsersResponse struct{
    Pager pagination.Pager
}

func NewListInGroupIdentityV3UsersResponse(pager pagination.Pager,)*ListInGroupIdentityV3UsersResponse {
    return &ListInGroupIdentityV3UsersResponse{
            Pager:pager,
    }
}

// action function
func (oc *OpenstackClient) ListInGroupIdentityV3Users(req *ListInGroupIdentityV3UsersRequest)(*ListInGroupIdentityV3UsersResponse){
    return NewListInGroupIdentityV3UsersResponse(users.ListInGroup(oc.Client,req.GroupID,req.Opts, ))

}