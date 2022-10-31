package openstack

// Code generated by cloud manager.





import (
    "github.com/gophercloud/gophercloud/openstack/messaging/v2/messages"
    "github.com/gophercloud/gophercloud/pagination"
)
//request struct for the ListMessagingV2Messages
type ListMessagingV2MessagesRequest struct{
    QueueName string
    Opts messages.ListOpts
}

func NewListMessagingV2MessagesRequest()*ListMessagingV2MessagesRequest{
    return &ListMessagingV2MessagesRequest{}
}

//response struct for the ListMessagingV2Messages
type ListMessagingV2MessagesResponse struct{
    Pager pagination.Pager
}

func NewListMessagingV2MessagesResponse(pager pagination.Pager,)*ListMessagingV2MessagesResponse {
    return &ListMessagingV2MessagesResponse{
            Pager:pager,
    }
}

// action function
func (oc *OpenstackClient) ListMessagingV2Messages(req *ListMessagingV2MessagesRequest)(*ListMessagingV2MessagesResponse){
    return NewListMessagingV2MessagesResponse(messages.List(oc.Client,req.QueueName,req.Opts, ))

}
//request struct for the CreateMessagingV2Messages
type CreateMessagingV2MessagesRequest struct{
    QueueName string
    Opts messages.BatchCreateOpts
}

func NewCreateMessagingV2MessagesRequest()*CreateMessagingV2MessagesRequest{
    return &CreateMessagingV2MessagesRequest{}
}

//response struct for the CreateMessagingV2Messages
type CreateMessagingV2MessagesResponse struct{
    CreateResult messages.CreateResult
}

func NewCreateMessagingV2MessagesResponse(createResult messages.CreateResult,)*CreateMessagingV2MessagesResponse {
    return &CreateMessagingV2MessagesResponse{
            CreateResult:createResult,
    }
}

// action function
func (oc *OpenstackClient) CreateMessagingV2Messages(req *CreateMessagingV2MessagesRequest)(*CreateMessagingV2MessagesResponse){
    return NewCreateMessagingV2MessagesResponse(messages.Create(oc.Client,req.QueueName,req.Opts, ))

}
//request struct for the DeleteMessagesMessagingV2Messages
type DeleteMessagesMessagingV2MessagesRequest struct{
    QueueName string
    Opts messages.DeleteMessagesOpts
}

func NewDeleteMessagesMessagingV2MessagesRequest()*DeleteMessagesMessagingV2MessagesRequest{
    return &DeleteMessagesMessagingV2MessagesRequest{}
}

//response struct for the DeleteMessagesMessagingV2Messages
type DeleteMessagesMessagingV2MessagesResponse struct{
    DeleteResult messages.DeleteResult
}

func NewDeleteMessagesMessagingV2MessagesResponse(deleteResult messages.DeleteResult,)*DeleteMessagesMessagingV2MessagesResponse {
    return &DeleteMessagesMessagingV2MessagesResponse{
            DeleteResult:deleteResult,
    }
}

// action function
func (oc *OpenstackClient) DeleteMessagesMessagingV2Messages(req *DeleteMessagesMessagingV2MessagesRequest)(*DeleteMessagesMessagingV2MessagesResponse){
    return NewDeleteMessagesMessagingV2MessagesResponse(messages.DeleteMessages(oc.Client,req.QueueName,req.Opts, ))

}
//request struct for the PopMessagesMessagingV2Messages
type PopMessagesMessagingV2MessagesRequest struct{
    QueueName string
    Opts messages.PopMessagesOpts
}

func NewPopMessagesMessagingV2MessagesRequest()*PopMessagesMessagingV2MessagesRequest{
    return &PopMessagesMessagingV2MessagesRequest{}
}

//response struct for the PopMessagesMessagingV2Messages
type PopMessagesMessagingV2MessagesResponse struct{
    PopResult messages.PopResult
}

func NewPopMessagesMessagingV2MessagesResponse(popResult messages.PopResult,)*PopMessagesMessagingV2MessagesResponse {
    return &PopMessagesMessagingV2MessagesResponse{
            PopResult:popResult,
    }
}

// action function
func (oc *OpenstackClient) PopMessagesMessagingV2Messages(req *PopMessagesMessagingV2MessagesRequest)(*PopMessagesMessagingV2MessagesResponse){
    return NewPopMessagesMessagingV2MessagesResponse(messages.PopMessages(oc.Client,req.QueueName,req.Opts, ))

}
//request struct for the GetMessagesMessagingV2Messages
type GetMessagesMessagingV2MessagesRequest struct{
    QueueName string
    Opts messages.GetMessagesOpts
}

func NewGetMessagesMessagingV2MessagesRequest()*GetMessagesMessagingV2MessagesRequest{
    return &GetMessagesMessagingV2MessagesRequest{}
}

//response struct for the GetMessagesMessagingV2Messages
type GetMessagesMessagingV2MessagesResponse struct{
    GetMessagesResult messages.GetMessagesResult
}

func NewGetMessagesMessagingV2MessagesResponse(getMessagesResult messages.GetMessagesResult,)*GetMessagesMessagingV2MessagesResponse {
    return &GetMessagesMessagingV2MessagesResponse{
            GetMessagesResult:getMessagesResult,
    }
}

// action function
func (oc *OpenstackClient) GetMessagesMessagingV2Messages(req *GetMessagesMessagingV2MessagesRequest)(*GetMessagesMessagingV2MessagesResponse){
    return NewGetMessagesMessagingV2MessagesResponse(messages.GetMessages(oc.Client,req.QueueName,req.Opts, ))

}
//request struct for the GetMessagingV2Messages
type GetMessagingV2MessagesRequest struct{
    QueueName string
    MessageID string
}

func NewGetMessagingV2MessagesRequest()*GetMessagingV2MessagesRequest{
    return &GetMessagingV2MessagesRequest{}
}

//response struct for the GetMessagingV2Messages
type GetMessagingV2MessagesResponse struct{
    GetResult messages.GetResult
}

func NewGetMessagingV2MessagesResponse(getResult messages.GetResult,)*GetMessagingV2MessagesResponse {
    return &GetMessagingV2MessagesResponse{
            GetResult:getResult,
    }
}

// action function
func (oc *OpenstackClient) GetMessagingV2Messages(req *GetMessagingV2MessagesRequest)(*GetMessagingV2MessagesResponse){
    return NewGetMessagingV2MessagesResponse(messages.Get(oc.Client,req.QueueName,req.MessageID, ))

}
//request struct for the DeleteMessagingV2Messages
type DeleteMessagingV2MessagesRequest struct{
    QueueName string
    MessageID string
    Opts messages.DeleteOpts
}

func NewDeleteMessagingV2MessagesRequest()*DeleteMessagingV2MessagesRequest{
    return &DeleteMessagingV2MessagesRequest{}
}

//response struct for the DeleteMessagingV2Messages
type DeleteMessagingV2MessagesResponse struct{
    DeleteResult messages.DeleteResult
}

func NewDeleteMessagingV2MessagesResponse(deleteResult messages.DeleteResult,)*DeleteMessagingV2MessagesResponse {
    return &DeleteMessagingV2MessagesResponse{
            DeleteResult:deleteResult,
    }
}

// action function
func (oc *OpenstackClient) DeleteMessagingV2Messages(req *DeleteMessagingV2MessagesRequest)(*DeleteMessagingV2MessagesResponse){
    return NewDeleteMessagingV2MessagesResponse(messages.Delete(oc.Client,req.QueueName,req.MessageID,req.Opts, ))

}