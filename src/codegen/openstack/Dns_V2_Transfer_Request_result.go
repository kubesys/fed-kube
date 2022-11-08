package openstack

// Code generated by cloud manager.

import (
	"github.com/gophercloud/gophercloud/openstack/dns/v2/transfer/request"
)

// extract response info from pager for ListDnsV2TransferRequest
func ExtractListDnsV2TransferRequestResponse(response *ListDnsV2TransferRequestResponse) ([]request.TransferRequest, error) {
	page, err := response.Pager.AllPages()
	if err != nil {
		return nil, err
	}
	return request.ExtractTransferRequests(page)
}

// call result's extract function
func ExtractGetDnsV2TransferRequestResponse(response *GetDnsV2TransferRequestResponse) (interface{}, error) {
	return response.GetResult.Body, response.GetResult.Err
}

// call result's extract function
func ExtractCreateDnsV2TransferRequestResponse(response *CreateDnsV2TransferRequestResponse) (interface{}, error) {
	return response.CreateResult.Body, response.CreateResult.Err
}

// call result's extract function
func ExtractUpdateDnsV2TransferRequestResponse(response *UpdateDnsV2TransferRequestResponse) (interface{}, error) {
	return response.UpdateResult.Body, response.UpdateResult.Err
}

// call result's extract function
func ExtractDeleteDnsV2TransferRequestResponse(response *DeleteDnsV2TransferRequestResponse) (interface{}, error) {
	return response.DeleteResult.Body, response.DeleteResult.Err
}
