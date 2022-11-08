package openstack

// Code generated by cloud manager.

import (
	"github.com/gophercloud/gophercloud/openstack/blockstorage/extensions/availabilityzones"
	"github.com/gophercloud/gophercloud/pagination"
)

// request struct for the ListBlockstorageExtensionsAvailabilityzones
type ListBlockstorageExtensionsAvailabilityzonesRequest struct {
}

func NewListBlockstorageExtensionsAvailabilityzonesRequest() *ListBlockstorageExtensionsAvailabilityzonesRequest {
	return &ListBlockstorageExtensionsAvailabilityzonesRequest{}
}

// response struct for the ListBlockstorageExtensionsAvailabilityzones
type ListBlockstorageExtensionsAvailabilityzonesResponse struct {
	Pager pagination.Pager
}

func NewListBlockstorageExtensionsAvailabilityzonesResponse(pager pagination.Pager) *ListBlockstorageExtensionsAvailabilityzonesResponse {
	return &ListBlockstorageExtensionsAvailabilityzonesResponse{
		Pager: pager,
	}
}

// action function
func (oc *OpenstackClient) ListBlockstorageExtensionsAvailabilityzones(req *ListBlockstorageExtensionsAvailabilityzonesRequest) *ListBlockstorageExtensionsAvailabilityzonesResponse {
	return NewListBlockstorageExtensionsAvailabilityzonesResponse(availabilityzones.List(oc.Client))

}
