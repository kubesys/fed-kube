package openstack

// Code generated by cloud manager.





import (
    "github.com/gophercloud/gophercloud/openstack/compute/v2/extensions/migrate"
)
//request struct for the MigrateComputeV2ExtensionsMigrate
type MigrateComputeV2ExtensionsMigrateRequest struct{
    Id string
}

func NewMigrateComputeV2ExtensionsMigrateRequest()*MigrateComputeV2ExtensionsMigrateRequest{
    return &MigrateComputeV2ExtensionsMigrateRequest{}
}

//response struct for the MigrateComputeV2ExtensionsMigrate
type MigrateComputeV2ExtensionsMigrateResponse struct{
    MigrateResult migrate.MigrateResult
}

func NewMigrateComputeV2ExtensionsMigrateResponse(migrateResult migrate.MigrateResult,)*MigrateComputeV2ExtensionsMigrateResponse {
    return &MigrateComputeV2ExtensionsMigrateResponse{
            MigrateResult:migrateResult,
    }
}

// action function
func (oc *OpenstackClient) MigrateComputeV2ExtensionsMigrate(req *MigrateComputeV2ExtensionsMigrateRequest)(*MigrateComputeV2ExtensionsMigrateResponse){
    return NewMigrateComputeV2ExtensionsMigrateResponse(migrate.Migrate(oc.Client,req.Id, ))

}
//request struct for the LiveMigrateComputeV2ExtensionsMigrate
type LiveMigrateComputeV2ExtensionsMigrateRequest struct{
    Id string
    Opts migrate.LiveMigrateOpts
}

func NewLiveMigrateComputeV2ExtensionsMigrateRequest()*LiveMigrateComputeV2ExtensionsMigrateRequest{
    return &LiveMigrateComputeV2ExtensionsMigrateRequest{}
}

//response struct for the LiveMigrateComputeV2ExtensionsMigrate
type LiveMigrateComputeV2ExtensionsMigrateResponse struct{
    MigrateResult migrate.MigrateResult
}

func NewLiveMigrateComputeV2ExtensionsMigrateResponse(migrateResult migrate.MigrateResult,)*LiveMigrateComputeV2ExtensionsMigrateResponse {
    return &LiveMigrateComputeV2ExtensionsMigrateResponse{
            MigrateResult:migrateResult,
    }
}

// action function
func (oc *OpenstackClient) LiveMigrateComputeV2ExtensionsMigrate(req *LiveMigrateComputeV2ExtensionsMigrateRequest)(*LiveMigrateComputeV2ExtensionsMigrateResponse){
    return NewLiveMigrateComputeV2ExtensionsMigrateResponse(migrate.LiveMigrate(oc.Client,req.Id,req.Opts, ))

}