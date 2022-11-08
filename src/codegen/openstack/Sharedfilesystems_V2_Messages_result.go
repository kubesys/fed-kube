package openstack

// Code generated by cloud manager.

import (
	"github.com/gophercloud/gophercloud/openstack/sharedfilesystems/v2/messages"
)

// call result's extract function
func ExtractDeleteSharedfilesystemsV2MessagesResponse(response *DeleteSharedfilesystemsV2MessagesResponse) (interface{}, error) {
	return response.DeleteResult.Body, response.DeleteResult.Err
}

// extract response info from pager for ListSharedfilesystemsV2Messages
func ExtractListSharedfilesystemsV2MessagesResponse(response *ListSharedfilesystemsV2MessagesResponse) ([]messages.Message, error) {
	page, err := response.Pager.AllPages()
	if err != nil {
		return nil, err
	}
	return messages.ExtractMessages(page)
}

// call result's extract function
func ExtractGetSharedfilesystemsV2MessagesResponse(response *GetSharedfilesystemsV2MessagesResponse) (interface{}, error) {
	return response.GetResult.Body, response.GetResult.Err
}
