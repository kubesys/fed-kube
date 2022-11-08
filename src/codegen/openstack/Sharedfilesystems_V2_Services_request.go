package openstack

// Code generated by cloud manager.

import (
	"github.com/gophercloud/gophercloud/openstack/sharedfilesystems/v2/services"
	"github.com/gophercloud/gophercloud/pagination"
)

// request struct for the ListSharedfilesystemsV2Services
type ListSharedfilesystemsV2ServicesRequest struct {
	Opts services.ListOpts
}

func NewListSharedfilesystemsV2ServicesRequest() *ListSharedfilesystemsV2ServicesRequest {
	return &ListSharedfilesystemsV2ServicesRequest{}
}

// response struct for the ListSharedfilesystemsV2Services
type ListSharedfilesystemsV2ServicesResponse struct {
	Pager pagination.Pager
}

func NewListSharedfilesystemsV2ServicesResponse(pager pagination.Pager) *ListSharedfilesystemsV2ServicesResponse {
	return &ListSharedfilesystemsV2ServicesResponse{
		Pager: pager,
	}
}

// action function
func (oc *OpenstackClient) ListSharedfilesystemsV2Services(req *ListSharedfilesystemsV2ServicesRequest) *ListSharedfilesystemsV2ServicesResponse {
	return NewListSharedfilesystemsV2ServicesResponse(services.List(oc.Client, req.Opts))

}
