package openstack

// Code generated by cloud manager.





import (
    "github.com/gophercloud/gophercloud/openstack/orchestration/v1/stacktemplates"
)
//request struct for the GetOrchestrationV1Stacktemplates
type GetOrchestrationV1StacktemplatesRequest struct{
    StackName string
    StackID string
}

func NewGetOrchestrationV1StacktemplatesRequest()*GetOrchestrationV1StacktemplatesRequest{
    return &GetOrchestrationV1StacktemplatesRequest{}
}

//response struct for the GetOrchestrationV1Stacktemplates
type GetOrchestrationV1StacktemplatesResponse struct{
    GetResult stacktemplates.GetResult
}

func NewGetOrchestrationV1StacktemplatesResponse(getResult stacktemplates.GetResult,)*GetOrchestrationV1StacktemplatesResponse {
    return &GetOrchestrationV1StacktemplatesResponse{
            GetResult:getResult,
    }
}

// action function
func (oc *OpenstackClient) GetOrchestrationV1Stacktemplates(req *GetOrchestrationV1StacktemplatesRequest)(*GetOrchestrationV1StacktemplatesResponse){
    return NewGetOrchestrationV1StacktemplatesResponse(stacktemplates.Get(oc.Client,req.StackName,req.StackID, ))

}
//request struct for the ValidateOrchestrationV1Stacktemplates
type ValidateOrchestrationV1StacktemplatesRequest struct{
    Opts stacktemplates.ValidateOpts
}

func NewValidateOrchestrationV1StacktemplatesRequest()*ValidateOrchestrationV1StacktemplatesRequest{
    return &ValidateOrchestrationV1StacktemplatesRequest{}
}

//response struct for the ValidateOrchestrationV1Stacktemplates
type ValidateOrchestrationV1StacktemplatesResponse struct{
    ValidateResult stacktemplates.ValidateResult
}

func NewValidateOrchestrationV1StacktemplatesResponse(validateResult stacktemplates.ValidateResult,)*ValidateOrchestrationV1StacktemplatesResponse {
    return &ValidateOrchestrationV1StacktemplatesResponse{
            ValidateResult:validateResult,
    }
}

// action function
func (oc *OpenstackClient) ValidateOrchestrationV1Stacktemplates(req *ValidateOrchestrationV1StacktemplatesRequest)(*ValidateOrchestrationV1StacktemplatesResponse){
    return NewValidateOrchestrationV1StacktemplatesResponse(stacktemplates.Validate(oc.Client,req.Opts, ))

}