package openstack

// Code generated by cloud manager.

import (
)


// call result's extract function
func ExtractTriggerClusteringV1WebhooksResponse(response *TriggerClusteringV1WebhooksResponse)(interface{}, error){
    return response.TriggerResult.Body, response.TriggerResult.Err
}
