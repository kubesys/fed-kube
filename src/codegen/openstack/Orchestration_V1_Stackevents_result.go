package openstack

// Code generated by cloud manager.





import (
    "github.com/gophercloud/gophercloud/openstack/orchestration/v1/stackevents"
)


//extract response info from pager for ListOrchestrationV1Stackevents
func ExtractListOrchestrationV1StackeventsResponse(response *ListOrchestrationV1StackeventsResponse)([]stackevents.Event,error){
	page, err := response.Pager.AllPages()
	if err != nil {
		return nil, err
	}
	return stackevents.ExtractEvents(page)
}


//extract response info from pager for ListResourceEventsOrchestrationV1Stackevents
func ExtractListResourceEventsOrchestrationV1StackeventsResponse(response *ListResourceEventsOrchestrationV1StackeventsResponse)([]stackevents.Event,error){
	page, err := response.Pager.AllPages()
	if err != nil {
		return nil, err
	}
	return stackevents.ExtractResourceEvents(page)
}