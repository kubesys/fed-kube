package openstack

// Code generated by cloud manager.

import (
    "github.com/gophercloud/gophercloud/openstack/compute/v2/extensions/lockunlock"
)
//request struct for the LockComputeV2ExtensionsLockunlock
type LockComputeV2ExtensionsLockunlockRequest struct{
    Id string
}

func NewLockComputeV2ExtensionsLockunlockRequest()*LockComputeV2ExtensionsLockunlockRequest{
    return &LockComputeV2ExtensionsLockunlockRequest{}
}

//response struct for the LockComputeV2ExtensionsLockunlock
type LockComputeV2ExtensionsLockunlockResponse struct{
    LockResult lockunlock.LockResult
}

func NewLockComputeV2ExtensionsLockunlockResponse(lockResult lockunlock.LockResult,)*LockComputeV2ExtensionsLockunlockResponse {
    return &LockComputeV2ExtensionsLockunlockResponse{
            LockResult:lockResult,
    }
}

// action function
func (oc *OpenstackClient) LockComputeV2ExtensionsLockunlock(req *LockComputeV2ExtensionsLockunlockRequest)(*LockComputeV2ExtensionsLockunlockResponse){
    return NewLockComputeV2ExtensionsLockunlockResponse(lockunlock.Lock(oc.Client,req.Id, ))

}
//request struct for the UnlockComputeV2ExtensionsLockunlock
type UnlockComputeV2ExtensionsLockunlockRequest struct{
    Id string
}

func NewUnlockComputeV2ExtensionsLockunlockRequest()*UnlockComputeV2ExtensionsLockunlockRequest{
    return &UnlockComputeV2ExtensionsLockunlockRequest{}
}

//response struct for the UnlockComputeV2ExtensionsLockunlock
type UnlockComputeV2ExtensionsLockunlockResponse struct{
    UnlockResult lockunlock.UnlockResult
}

func NewUnlockComputeV2ExtensionsLockunlockResponse(unlockResult lockunlock.UnlockResult,)*UnlockComputeV2ExtensionsLockunlockResponse {
    return &UnlockComputeV2ExtensionsLockunlockResponse{
            UnlockResult:unlockResult,
    }
}

// action function
func (oc *OpenstackClient) UnlockComputeV2ExtensionsLockunlock(req *UnlockComputeV2ExtensionsLockunlockRequest)(*UnlockComputeV2ExtensionsLockunlockResponse){
    return NewUnlockComputeV2ExtensionsLockunlockResponse(lockunlock.Unlock(oc.Client,req.Id, ))

}