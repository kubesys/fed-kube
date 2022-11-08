package openstack

// Code generated by cloud manager.

import (
	"github.com/gophercloud/gophercloud/openstack/compute/v2/extensions/resetstate"
)

// request struct for the ResetStateComputeV2ExtensionsResetstate
type ResetStateComputeV2ExtensionsResetstateRequest struct {
	Id    string
	State resetstate.ServerState
}

func NewResetStateComputeV2ExtensionsResetstateRequest() *ResetStateComputeV2ExtensionsResetstateRequest {
	return &ResetStateComputeV2ExtensionsResetstateRequest{}
}

// response struct for the ResetStateComputeV2ExtensionsResetstate
type ResetStateComputeV2ExtensionsResetstateResponse struct {
	ResetResult resetstate.ResetResult
}

func NewResetStateComputeV2ExtensionsResetstateResponse(resetResult resetstate.ResetResult) *ResetStateComputeV2ExtensionsResetstateResponse {
	return &ResetStateComputeV2ExtensionsResetstateResponse{
		ResetResult: resetResult,
	}
}

// action function
func (oc *OpenstackClient) ResetStateComputeV2ExtensionsResetstate(req *ResetStateComputeV2ExtensionsResetstateRequest) *ResetStateComputeV2ExtensionsResetstateResponse {
	return NewResetStateComputeV2ExtensionsResetstateResponse(resetstate.ResetState(oc.Client, req.Id, req.State))

}
