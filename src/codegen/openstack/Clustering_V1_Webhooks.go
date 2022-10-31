package openstack

// Code generated by cloud manager.





import (
    "github.com/gophercloud/gophercloud/openstack/clustering/v1/webhooks"
)
//request struct for the TriggerClusteringV1Webhooks
type TriggerClusteringV1WebhooksRequest struct{
    Id string
    Opts webhooks.TriggerOpts
}

func NewTriggerClusteringV1WebhooksRequest()*TriggerClusteringV1WebhooksRequest{
    return &TriggerClusteringV1WebhooksRequest{}
}

//response struct for the TriggerClusteringV1Webhooks
type TriggerClusteringV1WebhooksResponse struct{
    TriggerResult webhooks.TriggerResult
}

func NewTriggerClusteringV1WebhooksResponse(triggerResult webhooks.TriggerResult,)*TriggerClusteringV1WebhooksResponse {
    return &TriggerClusteringV1WebhooksResponse{
            TriggerResult:triggerResult,
    }
}

// action function
func (oc *OpenstackClient) TriggerClusteringV1Webhooks(req *TriggerClusteringV1WebhooksRequest)(*TriggerClusteringV1WebhooksResponse){
    return NewTriggerClusteringV1WebhooksResponse(webhooks.Trigger(oc.Client,req.Id,req.Opts, ))

}