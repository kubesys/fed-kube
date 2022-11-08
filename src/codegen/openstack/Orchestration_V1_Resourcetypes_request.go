package openstack

// Code generated by cloud manager.

import (
    "github.com/gophercloud/gophercloud/openstack/orchestration/v1/resourcetypes"
)
//request struct for the ListOrchestrationV1Resourcetypes
type ListOrchestrationV1ResourcetypesRequest struct{
    Opts resourcetypes.ListOpts
}

func NewListOrchestrationV1ResourcetypesRequest()*ListOrchestrationV1ResourcetypesRequest{
    return &ListOrchestrationV1ResourcetypesRequest{}
}

//response struct for the ListOrchestrationV1Resourcetypes
type ListOrchestrationV1ResourcetypesResponse struct{
    ListResult resourcetypes.ListResult
}

func NewListOrchestrationV1ResourcetypesResponse(listResult resourcetypes.ListResult,)*ListOrchestrationV1ResourcetypesResponse {
    return &ListOrchestrationV1ResourcetypesResponse{
            ListResult:listResult,
    }
}

// action function
func (oc *OpenstackClient) ListOrchestrationV1Resourcetypes(req *ListOrchestrationV1ResourcetypesRequest)(*ListOrchestrationV1ResourcetypesResponse){
    return NewListOrchestrationV1ResourcetypesResponse(resourcetypes.List(oc.Client,req.Opts, ))

}
//request struct for the GetSchemaOrchestrationV1Resourcetypes
type GetSchemaOrchestrationV1ResourcetypesRequest struct{
    ResourceType string
}

func NewGetSchemaOrchestrationV1ResourcetypesRequest()*GetSchemaOrchestrationV1ResourcetypesRequest{
    return &GetSchemaOrchestrationV1ResourcetypesRequest{}
}

//response struct for the GetSchemaOrchestrationV1Resourcetypes
type GetSchemaOrchestrationV1ResourcetypesResponse struct{
    GetSchemaResult resourcetypes.GetSchemaResult
}

func NewGetSchemaOrchestrationV1ResourcetypesResponse(getSchemaResult resourcetypes.GetSchemaResult,)*GetSchemaOrchestrationV1ResourcetypesResponse {
    return &GetSchemaOrchestrationV1ResourcetypesResponse{
            GetSchemaResult:getSchemaResult,
    }
}

// action function
func (oc *OpenstackClient) GetSchemaOrchestrationV1Resourcetypes(req *GetSchemaOrchestrationV1ResourcetypesRequest)(*GetSchemaOrchestrationV1ResourcetypesResponse){
    return NewGetSchemaOrchestrationV1ResourcetypesResponse(resourcetypes.GetSchema(oc.Client,req.ResourceType, ))

}
//request struct for the GenerateTemplateOrchestrationV1Resourcetypes
type GenerateTemplateOrchestrationV1ResourcetypesRequest struct{
    ResourceType string
    Opts resourcetypes.GenerateTemplateOpts
}

func NewGenerateTemplateOrchestrationV1ResourcetypesRequest()*GenerateTemplateOrchestrationV1ResourcetypesRequest{
    return &GenerateTemplateOrchestrationV1ResourcetypesRequest{}
}

//response struct for the GenerateTemplateOrchestrationV1Resourcetypes
type GenerateTemplateOrchestrationV1ResourcetypesResponse struct{
    TemplateResult resourcetypes.TemplateResult
}

func NewGenerateTemplateOrchestrationV1ResourcetypesResponse(templateResult resourcetypes.TemplateResult,)*GenerateTemplateOrchestrationV1ResourcetypesResponse {
    return &GenerateTemplateOrchestrationV1ResourcetypesResponse{
            TemplateResult:templateResult,
    }
}

// action function
func (oc *OpenstackClient) GenerateTemplateOrchestrationV1Resourcetypes(req *GenerateTemplateOrchestrationV1ResourcetypesRequest)(*GenerateTemplateOrchestrationV1ResourcetypesResponse){
    return NewGenerateTemplateOrchestrationV1ResourcetypesResponse(resourcetypes.GenerateTemplate(oc.Client,req.ResourceType,req.Opts, ))

}