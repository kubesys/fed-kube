package openstack

// Code generated by cloud manager.

import (
    "github.com/gophercloud/gophercloud/openstack/orchestration/v1/stackevents"
    "github.com/gophercloud/gophercloud/pagination"
)
//request struct for the FindOrchestrationV1Stackevents
type FindOrchestrationV1StackeventsRequest struct{
    StackName string
}

func NewFindOrchestrationV1StackeventsRequest()*FindOrchestrationV1StackeventsRequest{
    return &FindOrchestrationV1StackeventsRequest{}
}

//response struct for the FindOrchestrationV1Stackevents
type FindOrchestrationV1StackeventsResponse struct{
    FindResult stackevents.FindResult
}

func NewFindOrchestrationV1StackeventsResponse(findResult stackevents.FindResult,)*FindOrchestrationV1StackeventsResponse {
    return &FindOrchestrationV1StackeventsResponse{
            FindResult:findResult,
    }
}

// action function
func (oc *OpenstackClient) FindOrchestrationV1Stackevents(req *FindOrchestrationV1StackeventsRequest)(*FindOrchestrationV1StackeventsResponse){
    return NewFindOrchestrationV1StackeventsResponse(stackevents.Find(oc.Client,req.StackName, ))

}
//request struct for the ListOrchestrationV1Stackevents
type ListOrchestrationV1StackeventsRequest struct{
    StackName string
    StackID string
    Opts stackevents.ListOpts
}

func NewListOrchestrationV1StackeventsRequest()*ListOrchestrationV1StackeventsRequest{
    return &ListOrchestrationV1StackeventsRequest{}
}

//response struct for the ListOrchestrationV1Stackevents
type ListOrchestrationV1StackeventsResponse struct{
    Pager pagination.Pager
}

func NewListOrchestrationV1StackeventsResponse(pager pagination.Pager,)*ListOrchestrationV1StackeventsResponse {
    return &ListOrchestrationV1StackeventsResponse{
            Pager:pager,
    }
}

// action function
func (oc *OpenstackClient) ListOrchestrationV1Stackevents(req *ListOrchestrationV1StackeventsRequest)(*ListOrchestrationV1StackeventsResponse){
    return NewListOrchestrationV1StackeventsResponse(stackevents.List(oc.Client,req.StackName,req.StackID,req.Opts, ))

}
//request struct for the ListResourceEventsOrchestrationV1Stackevents
type ListResourceEventsOrchestrationV1StackeventsRequest struct{
    StackName string
    StackID string
    ResourceName string
    Opts stackevents.ListResourceEventsOpts
}

func NewListResourceEventsOrchestrationV1StackeventsRequest()*ListResourceEventsOrchestrationV1StackeventsRequest{
    return &ListResourceEventsOrchestrationV1StackeventsRequest{}
}

//response struct for the ListResourceEventsOrchestrationV1Stackevents
type ListResourceEventsOrchestrationV1StackeventsResponse struct{
    Pager pagination.Pager
}

func NewListResourceEventsOrchestrationV1StackeventsResponse(pager pagination.Pager,)*ListResourceEventsOrchestrationV1StackeventsResponse {
    return &ListResourceEventsOrchestrationV1StackeventsResponse{
            Pager:pager,
    }
}

// action function
func (oc *OpenstackClient) ListResourceEventsOrchestrationV1Stackevents(req *ListResourceEventsOrchestrationV1StackeventsRequest)(*ListResourceEventsOrchestrationV1StackeventsResponse){
    return NewListResourceEventsOrchestrationV1StackeventsResponse(stackevents.ListResourceEvents(oc.Client,req.StackName,req.StackID,req.ResourceName,req.Opts, ))

}
//request struct for the GetOrchestrationV1Stackevents
type GetOrchestrationV1StackeventsRequest struct{
    StackName string
    StackID string
    ResourceName string
    EventID string
}

func NewGetOrchestrationV1StackeventsRequest()*GetOrchestrationV1StackeventsRequest{
    return &GetOrchestrationV1StackeventsRequest{}
}

//response struct for the GetOrchestrationV1Stackevents
type GetOrchestrationV1StackeventsResponse struct{
    GetResult stackevents.GetResult
}

func NewGetOrchestrationV1StackeventsResponse(getResult stackevents.GetResult,)*GetOrchestrationV1StackeventsResponse {
    return &GetOrchestrationV1StackeventsResponse{
            GetResult:getResult,
    }
}

// action function
func (oc *OpenstackClient) GetOrchestrationV1Stackevents(req *GetOrchestrationV1StackeventsRequest)(*GetOrchestrationV1StackeventsResponse){
    return NewGetOrchestrationV1StackeventsResponse(stackevents.Get(oc.Client,req.StackName,req.StackID,req.ResourceName,req.EventID, ))

}