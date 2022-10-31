package openstack

// Code generated by cloud manager.





import (
    "github.com/gophercloud/gophercloud/openstack/compute/v2/extensions/bootfromvolume"
    "github.com/gophercloud/gophercloud/openstack/compute/v2/servers"
)
//request struct for the CreateComputeV2ExtensionsBootfromvolume
type CreateComputeV2ExtensionsBootfromvolumeRequest struct{
    Opts servers.CreateOpts
}

func NewCreateComputeV2ExtensionsBootfromvolumeRequest()*CreateComputeV2ExtensionsBootfromvolumeRequest{
    return &CreateComputeV2ExtensionsBootfromvolumeRequest{}
}

//response struct for the CreateComputeV2ExtensionsBootfromvolume
type CreateComputeV2ExtensionsBootfromvolumeResponse struct{
    CreateResult servers.CreateResult
}

func NewCreateComputeV2ExtensionsBootfromvolumeResponse(createResult servers.CreateResult,)*CreateComputeV2ExtensionsBootfromvolumeResponse {
    return &CreateComputeV2ExtensionsBootfromvolumeResponse{
            CreateResult:createResult,
    }
}

// action function
func (oc *OpenstackClient) CreateComputeV2ExtensionsBootfromvolume(req *CreateComputeV2ExtensionsBootfromvolumeRequest)(*CreateComputeV2ExtensionsBootfromvolumeResponse){
    return NewCreateComputeV2ExtensionsBootfromvolumeResponse(bootfromvolume.Create(oc.Client,req.Opts, ))

}