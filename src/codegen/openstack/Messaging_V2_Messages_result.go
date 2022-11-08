package openstack

// Code generated by cloud manager.

import (
    "github.com/gophercloud/gophercloud/openstack/messaging/v2/messages"
)


//extract response info from pager for ListMessagingV2Messages
func ExtractListMessagingV2MessagesResponse(response *ListMessagingV2MessagesResponse)([]messages.Message,error){
	page, err := response.Pager.AllPages()
	if err != nil {
		return nil, err
	}
	return messages.ExtractMessages(page)
}


// call result's extract function
func ExtractCreateMessagingV2MessagesResponse(response *CreateMessagingV2MessagesResponse)(interface{}, error){
    return response.CreateResult.Body, response.CreateResult.Err
}



// call result's extract function
func ExtractDeleteMessagesMessagingV2MessagesResponse(response *DeleteMessagesMessagingV2MessagesResponse)(interface{}, error){
    return response.DeleteResult.Body, response.DeleteResult.Err
}



// call result's extract function
func ExtractPopMessagesMessagingV2MessagesResponse(response *PopMessagesMessagingV2MessagesResponse)(interface{}, error){
    return response.PopResult.Body, response.PopResult.Err
}



// call result's extract function
func ExtractGetMessagesMessagingV2MessagesResponse(response *GetMessagesMessagingV2MessagesResponse)(interface{}, error){
    return response.GetMessagesResult.Body, response.GetMessagesResult.Err
}



// call result's extract function
func ExtractGetMessagingV2MessagesResponse(response *GetMessagingV2MessagesResponse)(interface{}, error){
    return response.GetResult.Body, response.GetResult.Err
}



// call result's extract function
func ExtractDeleteMessagingV2MessagesResponse(response *DeleteMessagingV2MessagesResponse)(interface{}, error){
    return response.DeleteResult.Body, response.DeleteResult.Err
}
