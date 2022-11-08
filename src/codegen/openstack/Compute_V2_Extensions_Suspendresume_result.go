package openstack

// Code generated by cloud manager.

import ()

// call result's extract function
func ExtractSuspendComputeV2ExtensionsSuspendresumeResponse(response *SuspendComputeV2ExtensionsSuspendresumeResponse) (interface{}, error) {
	return response.SuspendResult.Body, response.SuspendResult.Err
}

// call result's extract function
func ExtractResumeComputeV2ExtensionsSuspendresumeResponse(response *ResumeComputeV2ExtensionsSuspendresumeResponse) (interface{}, error) {
	return response.UnsuspendResult.Body, response.UnsuspendResult.Err
}
