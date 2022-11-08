package openstack

// Code generated by cloud manager.

import (
	"github.com/gophercloud/gophercloud/openstack/db/v1/configurations"
	"github.com/gophercloud/gophercloud/pagination"
)

// request struct for the ListDbV1Configurations
type ListDbV1ConfigurationsRequest struct {
}

func NewListDbV1ConfigurationsRequest() *ListDbV1ConfigurationsRequest {
	return &ListDbV1ConfigurationsRequest{}
}

// response struct for the ListDbV1Configurations
type ListDbV1ConfigurationsResponse struct {
	Pager pagination.Pager
}

func NewListDbV1ConfigurationsResponse(pager pagination.Pager) *ListDbV1ConfigurationsResponse {
	return &ListDbV1ConfigurationsResponse{
		Pager: pager,
	}
}

// action function
func (oc *OpenstackClient) ListDbV1Configurations(req *ListDbV1ConfigurationsRequest) *ListDbV1ConfigurationsResponse {
	return NewListDbV1ConfigurationsResponse(configurations.List(oc.Client))

}

// request struct for the CreateDbV1Configurations
type CreateDbV1ConfigurationsRequest struct {
	Opts configurations.CreateOpts
}

func NewCreateDbV1ConfigurationsRequest() *CreateDbV1ConfigurationsRequest {
	return &CreateDbV1ConfigurationsRequest{}
}

// response struct for the CreateDbV1Configurations
type CreateDbV1ConfigurationsResponse struct {
	CreateResult configurations.CreateResult
}

func NewCreateDbV1ConfigurationsResponse(createResult configurations.CreateResult) *CreateDbV1ConfigurationsResponse {
	return &CreateDbV1ConfigurationsResponse{
		CreateResult: createResult,
	}
}

// action function
func (oc *OpenstackClient) CreateDbV1Configurations(req *CreateDbV1ConfigurationsRequest) *CreateDbV1ConfigurationsResponse {
	return NewCreateDbV1ConfigurationsResponse(configurations.Create(oc.Client, req.Opts))

}

// request struct for the GetDbV1Configurations
type GetDbV1ConfigurationsRequest struct {
	ConfigID string
}

func NewGetDbV1ConfigurationsRequest() *GetDbV1ConfigurationsRequest {
	return &GetDbV1ConfigurationsRequest{}
}

// response struct for the GetDbV1Configurations
type GetDbV1ConfigurationsResponse struct {
	GetResult configurations.GetResult
}

func NewGetDbV1ConfigurationsResponse(getResult configurations.GetResult) *GetDbV1ConfigurationsResponse {
	return &GetDbV1ConfigurationsResponse{
		GetResult: getResult,
	}
}

// action function
func (oc *OpenstackClient) GetDbV1Configurations(req *GetDbV1ConfigurationsRequest) *GetDbV1ConfigurationsResponse {
	return NewGetDbV1ConfigurationsResponse(configurations.Get(oc.Client, req.ConfigID))

}

// request struct for the UpdateDbV1Configurations
type UpdateDbV1ConfigurationsRequest struct {
	ConfigID string
	Opts     configurations.UpdateOpts
}

func NewUpdateDbV1ConfigurationsRequest() *UpdateDbV1ConfigurationsRequest {
	return &UpdateDbV1ConfigurationsRequest{}
}

// response struct for the UpdateDbV1Configurations
type UpdateDbV1ConfigurationsResponse struct {
	UpdateResult configurations.UpdateResult
}

func NewUpdateDbV1ConfigurationsResponse(updateResult configurations.UpdateResult) *UpdateDbV1ConfigurationsResponse {
	return &UpdateDbV1ConfigurationsResponse{
		UpdateResult: updateResult,
	}
}

// action function
func (oc *OpenstackClient) UpdateDbV1Configurations(req *UpdateDbV1ConfigurationsRequest) *UpdateDbV1ConfigurationsResponse {
	return NewUpdateDbV1ConfigurationsResponse(configurations.Update(oc.Client, req.ConfigID, req.Opts))

}

// request struct for the ReplaceDbV1Configurations
type ReplaceDbV1ConfigurationsRequest struct {
	ConfigID string
	Opts     configurations.UpdateOpts
}

func NewReplaceDbV1ConfigurationsRequest() *ReplaceDbV1ConfigurationsRequest {
	return &ReplaceDbV1ConfigurationsRequest{}
}

// response struct for the ReplaceDbV1Configurations
type ReplaceDbV1ConfigurationsResponse struct {
	ReplaceResult configurations.ReplaceResult
}

func NewReplaceDbV1ConfigurationsResponse(replaceResult configurations.ReplaceResult) *ReplaceDbV1ConfigurationsResponse {
	return &ReplaceDbV1ConfigurationsResponse{
		ReplaceResult: replaceResult,
	}
}

// action function
func (oc *OpenstackClient) ReplaceDbV1Configurations(req *ReplaceDbV1ConfigurationsRequest) *ReplaceDbV1ConfigurationsResponse {
	return NewReplaceDbV1ConfigurationsResponse(configurations.Replace(oc.Client, req.ConfigID, req.Opts))

}

// request struct for the DeleteDbV1Configurations
type DeleteDbV1ConfigurationsRequest struct {
	ConfigID string
}

func NewDeleteDbV1ConfigurationsRequest() *DeleteDbV1ConfigurationsRequest {
	return &DeleteDbV1ConfigurationsRequest{}
}

// response struct for the DeleteDbV1Configurations
type DeleteDbV1ConfigurationsResponse struct {
	DeleteResult configurations.DeleteResult
}

func NewDeleteDbV1ConfigurationsResponse(deleteResult configurations.DeleteResult) *DeleteDbV1ConfigurationsResponse {
	return &DeleteDbV1ConfigurationsResponse{
		DeleteResult: deleteResult,
	}
}

// action function
func (oc *OpenstackClient) DeleteDbV1Configurations(req *DeleteDbV1ConfigurationsRequest) *DeleteDbV1ConfigurationsResponse {
	return NewDeleteDbV1ConfigurationsResponse(configurations.Delete(oc.Client, req.ConfigID))

}

// request struct for the ListDatastoreParamsDbV1Configurations
type ListDatastoreParamsDbV1ConfigurationsRequest struct {
	DatastoreID string
	VersionID   string
}

func NewListDatastoreParamsDbV1ConfigurationsRequest() *ListDatastoreParamsDbV1ConfigurationsRequest {
	return &ListDatastoreParamsDbV1ConfigurationsRequest{}
}

// response struct for the ListDatastoreParamsDbV1Configurations
type ListDatastoreParamsDbV1ConfigurationsResponse struct {
	Pager pagination.Pager
}

func NewListDatastoreParamsDbV1ConfigurationsResponse(pager pagination.Pager) *ListDatastoreParamsDbV1ConfigurationsResponse {
	return &ListDatastoreParamsDbV1ConfigurationsResponse{
		Pager: pager,
	}
}

// action function
func (oc *OpenstackClient) ListDatastoreParamsDbV1Configurations(req *ListDatastoreParamsDbV1ConfigurationsRequest) *ListDatastoreParamsDbV1ConfigurationsResponse {
	return NewListDatastoreParamsDbV1ConfigurationsResponse(configurations.ListDatastoreParams(oc.Client, req.DatastoreID, req.VersionID))

}

// request struct for the GetDatastoreParamDbV1Configurations
type GetDatastoreParamDbV1ConfigurationsRequest struct {
	DatastoreID string
	VersionID   string
	ParamID     string
}

func NewGetDatastoreParamDbV1ConfigurationsRequest() *GetDatastoreParamDbV1ConfigurationsRequest {
	return &GetDatastoreParamDbV1ConfigurationsRequest{}
}

// response struct for the GetDatastoreParamDbV1Configurations
type GetDatastoreParamDbV1ConfigurationsResponse struct {
	ParamResult configurations.ParamResult
}

func NewGetDatastoreParamDbV1ConfigurationsResponse(paramResult configurations.ParamResult) *GetDatastoreParamDbV1ConfigurationsResponse {
	return &GetDatastoreParamDbV1ConfigurationsResponse{
		ParamResult: paramResult,
	}
}

// action function
func (oc *OpenstackClient) GetDatastoreParamDbV1Configurations(req *GetDatastoreParamDbV1ConfigurationsRequest) *GetDatastoreParamDbV1ConfigurationsResponse {
	return NewGetDatastoreParamDbV1ConfigurationsResponse(configurations.GetDatastoreParam(oc.Client, req.DatastoreID, req.VersionID, req.ParamID))

}

// request struct for the ListGlobalParamsDbV1Configurations
type ListGlobalParamsDbV1ConfigurationsRequest struct {
	VersionID string
}

func NewListGlobalParamsDbV1ConfigurationsRequest() *ListGlobalParamsDbV1ConfigurationsRequest {
	return &ListGlobalParamsDbV1ConfigurationsRequest{}
}

// response struct for the ListGlobalParamsDbV1Configurations
type ListGlobalParamsDbV1ConfigurationsResponse struct {
	Pager pagination.Pager
}

func NewListGlobalParamsDbV1ConfigurationsResponse(pager pagination.Pager) *ListGlobalParamsDbV1ConfigurationsResponse {
	return &ListGlobalParamsDbV1ConfigurationsResponse{
		Pager: pager,
	}
}

// action function
func (oc *OpenstackClient) ListGlobalParamsDbV1Configurations(req *ListGlobalParamsDbV1ConfigurationsRequest) *ListGlobalParamsDbV1ConfigurationsResponse {
	return NewListGlobalParamsDbV1ConfigurationsResponse(configurations.ListGlobalParams(oc.Client, req.VersionID))

}

// request struct for the GetGlobalParamDbV1Configurations
type GetGlobalParamDbV1ConfigurationsRequest struct {
	VersionID string
	ParamID   string
}

func NewGetGlobalParamDbV1ConfigurationsRequest() *GetGlobalParamDbV1ConfigurationsRequest {
	return &GetGlobalParamDbV1ConfigurationsRequest{}
}

// response struct for the GetGlobalParamDbV1Configurations
type GetGlobalParamDbV1ConfigurationsResponse struct {
	ParamResult configurations.ParamResult
}

func NewGetGlobalParamDbV1ConfigurationsResponse(paramResult configurations.ParamResult) *GetGlobalParamDbV1ConfigurationsResponse {
	return &GetGlobalParamDbV1ConfigurationsResponse{
		ParamResult: paramResult,
	}
}

// action function
func (oc *OpenstackClient) GetGlobalParamDbV1Configurations(req *GetGlobalParamDbV1ConfigurationsRequest) *GetGlobalParamDbV1ConfigurationsResponse {
	return NewGetGlobalParamDbV1ConfigurationsResponse(configurations.GetGlobalParam(oc.Client, req.VersionID, req.ParamID))

}
