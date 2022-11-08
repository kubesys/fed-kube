package openstack

// Code generated by cloud manager.

import (
	"github.com/gophercloud/gophercloud/openstack/workflow/v2/executions"
	"github.com/gophercloud/gophercloud/pagination"
)

// request struct for the CreateWorkflowV2Executions
type CreateWorkflowV2ExecutionsRequest struct {
	Opts executions.CreateOpts
}

func NewCreateWorkflowV2ExecutionsRequest() *CreateWorkflowV2ExecutionsRequest {
	return &CreateWorkflowV2ExecutionsRequest{}
}

// response struct for the CreateWorkflowV2Executions
type CreateWorkflowV2ExecutionsResponse struct {
	CreateResult executions.CreateResult
}

func NewCreateWorkflowV2ExecutionsResponse(createResult executions.CreateResult) *CreateWorkflowV2ExecutionsResponse {
	return &CreateWorkflowV2ExecutionsResponse{
		CreateResult: createResult,
	}
}

// action function
func (oc *OpenstackClient) CreateWorkflowV2Executions(req *CreateWorkflowV2ExecutionsRequest) *CreateWorkflowV2ExecutionsResponse {
	return NewCreateWorkflowV2ExecutionsResponse(executions.Create(oc.Client, req.Opts))

}

// request struct for the GetWorkflowV2Executions
type GetWorkflowV2ExecutionsRequest struct {
	Id string
}

func NewGetWorkflowV2ExecutionsRequest() *GetWorkflowV2ExecutionsRequest {
	return &GetWorkflowV2ExecutionsRequest{}
}

// response struct for the GetWorkflowV2Executions
type GetWorkflowV2ExecutionsResponse struct {
	GetResult executions.GetResult
}

func NewGetWorkflowV2ExecutionsResponse(getResult executions.GetResult) *GetWorkflowV2ExecutionsResponse {
	return &GetWorkflowV2ExecutionsResponse{
		GetResult: getResult,
	}
}

// action function
func (oc *OpenstackClient) GetWorkflowV2Executions(req *GetWorkflowV2ExecutionsRequest) *GetWorkflowV2ExecutionsResponse {
	return NewGetWorkflowV2ExecutionsResponse(executions.Get(oc.Client, req.Id))

}

// request struct for the DeleteWorkflowV2Executions
type DeleteWorkflowV2ExecutionsRequest struct {
	Id string
}

func NewDeleteWorkflowV2ExecutionsRequest() *DeleteWorkflowV2ExecutionsRequest {
	return &DeleteWorkflowV2ExecutionsRequest{}
}

// response struct for the DeleteWorkflowV2Executions
type DeleteWorkflowV2ExecutionsResponse struct {
	DeleteResult executions.DeleteResult
}

func NewDeleteWorkflowV2ExecutionsResponse(deleteResult executions.DeleteResult) *DeleteWorkflowV2ExecutionsResponse {
	return &DeleteWorkflowV2ExecutionsResponse{
		DeleteResult: deleteResult,
	}
}

// action function
func (oc *OpenstackClient) DeleteWorkflowV2Executions(req *DeleteWorkflowV2ExecutionsRequest) *DeleteWorkflowV2ExecutionsResponse {
	return NewDeleteWorkflowV2ExecutionsResponse(executions.Delete(oc.Client, req.Id))

}

// request struct for the ListWorkflowV2Executions
type ListWorkflowV2ExecutionsRequest struct {
	Opts executions.ListOpts
}

func NewListWorkflowV2ExecutionsRequest() *ListWorkflowV2ExecutionsRequest {
	return &ListWorkflowV2ExecutionsRequest{}
}

// response struct for the ListWorkflowV2Executions
type ListWorkflowV2ExecutionsResponse struct {
	Pager pagination.Pager
}

func NewListWorkflowV2ExecutionsResponse(pager pagination.Pager) *ListWorkflowV2ExecutionsResponse {
	return &ListWorkflowV2ExecutionsResponse{
		Pager: pager,
	}
}

// action function
func (oc *OpenstackClient) ListWorkflowV2Executions(req *ListWorkflowV2ExecutionsRequest) *ListWorkflowV2ExecutionsResponse {
	return NewListWorkflowV2ExecutionsResponse(executions.List(oc.Client, req.Opts))

}
