package openstack

// Code generated by cloud manager.

import (
    "github.com/gophercloud/gophercloud/openstack/baremetal/v1/nodes"
    "github.com/gophercloud/gophercloud/pagination"
)
//request struct for the ListBaremetalV1Nodes
type ListBaremetalV1NodesRequest struct{
    Opts nodes.ListOpts
}

func NewListBaremetalV1NodesRequest()*ListBaremetalV1NodesRequest{
    return &ListBaremetalV1NodesRequest{}
}

//response struct for the ListBaremetalV1Nodes
type ListBaremetalV1NodesResponse struct{
    Pager pagination.Pager
}

func NewListBaremetalV1NodesResponse(pager pagination.Pager,)*ListBaremetalV1NodesResponse {
    return &ListBaremetalV1NodesResponse{
            Pager:pager,
    }
}

// action function
func (oc *OpenstackClient) ListBaremetalV1Nodes(req *ListBaremetalV1NodesRequest)(*ListBaremetalV1NodesResponse){
    return NewListBaremetalV1NodesResponse(nodes.List(oc.Client,req.Opts, ))

}
//request struct for the ListDetailBaremetalV1Nodes
type ListDetailBaremetalV1NodesRequest struct{
    Opts nodes.ListOpts
}

func NewListDetailBaremetalV1NodesRequest()*ListDetailBaremetalV1NodesRequest{
    return &ListDetailBaremetalV1NodesRequest{}
}

//response struct for the ListDetailBaremetalV1Nodes
type ListDetailBaremetalV1NodesResponse struct{
    Pager pagination.Pager
}

func NewListDetailBaremetalV1NodesResponse(pager pagination.Pager,)*ListDetailBaremetalV1NodesResponse {
    return &ListDetailBaremetalV1NodesResponse{
            Pager:pager,
    }
}

// action function
func (oc *OpenstackClient) ListDetailBaremetalV1Nodes(req *ListDetailBaremetalV1NodesRequest)(*ListDetailBaremetalV1NodesResponse){
    return NewListDetailBaremetalV1NodesResponse(nodes.ListDetail(oc.Client,req.Opts, ))

}
//request struct for the GetBaremetalV1Nodes
type GetBaremetalV1NodesRequest struct{
    Id string
}

func NewGetBaremetalV1NodesRequest()*GetBaremetalV1NodesRequest{
    return &GetBaremetalV1NodesRequest{}
}

//response struct for the GetBaremetalV1Nodes
type GetBaremetalV1NodesResponse struct{
    GetResult nodes.GetResult
}

func NewGetBaremetalV1NodesResponse(getResult nodes.GetResult,)*GetBaremetalV1NodesResponse {
    return &GetBaremetalV1NodesResponse{
            GetResult:getResult,
    }
}

// action function
func (oc *OpenstackClient) GetBaremetalV1Nodes(req *GetBaremetalV1NodesRequest)(*GetBaremetalV1NodesResponse){
    return NewGetBaremetalV1NodesResponse(nodes.Get(oc.Client,req.Id, ))

}
//request struct for the CreateBaremetalV1Nodes
type CreateBaremetalV1NodesRequest struct{
    Opts nodes.CreateOpts
}

func NewCreateBaremetalV1NodesRequest()*CreateBaremetalV1NodesRequest{
    return &CreateBaremetalV1NodesRequest{}
}

//response struct for the CreateBaremetalV1Nodes
type CreateBaremetalV1NodesResponse struct{
    CreateResult nodes.CreateResult
}

func NewCreateBaremetalV1NodesResponse(createResult nodes.CreateResult,)*CreateBaremetalV1NodesResponse {
    return &CreateBaremetalV1NodesResponse{
            CreateResult:createResult,
    }
}

// action function
func (oc *OpenstackClient) CreateBaremetalV1Nodes(req *CreateBaremetalV1NodesRequest)(*CreateBaremetalV1NodesResponse){
    return NewCreateBaremetalV1NodesResponse(nodes.Create(oc.Client,req.Opts, ))

}
//request struct for the UpdateBaremetalV1Nodes
type UpdateBaremetalV1NodesRequest struct{
    Id string
    Opts nodes.UpdateOpts
}

func NewUpdateBaremetalV1NodesRequest()*UpdateBaremetalV1NodesRequest{
    return &UpdateBaremetalV1NodesRequest{}
}

//response struct for the UpdateBaremetalV1Nodes
type UpdateBaremetalV1NodesResponse struct{
    UpdateResult nodes.UpdateResult
}

func NewUpdateBaremetalV1NodesResponse(updateResult nodes.UpdateResult,)*UpdateBaremetalV1NodesResponse {
    return &UpdateBaremetalV1NodesResponse{
            UpdateResult:updateResult,
    }
}

// action function
func (oc *OpenstackClient) UpdateBaremetalV1Nodes(req *UpdateBaremetalV1NodesRequest)(*UpdateBaremetalV1NodesResponse){
    return NewUpdateBaremetalV1NodesResponse(nodes.Update(oc.Client,req.Id,req.Opts, ))

}
//request struct for the DeleteBaremetalV1Nodes
type DeleteBaremetalV1NodesRequest struct{
    Id string
}

func NewDeleteBaremetalV1NodesRequest()*DeleteBaremetalV1NodesRequest{
    return &DeleteBaremetalV1NodesRequest{}
}

//response struct for the DeleteBaremetalV1Nodes
type DeleteBaremetalV1NodesResponse struct{
    DeleteResult nodes.DeleteResult
}

func NewDeleteBaremetalV1NodesResponse(deleteResult nodes.DeleteResult,)*DeleteBaremetalV1NodesResponse {
    return &DeleteBaremetalV1NodesResponse{
            DeleteResult:deleteResult,
    }
}

// action function
func (oc *OpenstackClient) DeleteBaremetalV1Nodes(req *DeleteBaremetalV1NodesRequest)(*DeleteBaremetalV1NodesResponse){
    return NewDeleteBaremetalV1NodesResponse(nodes.Delete(oc.Client,req.Id, ))

}
//request struct for the ValidateBaremetalV1Nodes
type ValidateBaremetalV1NodesRequest struct{
    Id string
}

func NewValidateBaremetalV1NodesRequest()*ValidateBaremetalV1NodesRequest{
    return &ValidateBaremetalV1NodesRequest{}
}

//response struct for the ValidateBaremetalV1Nodes
type ValidateBaremetalV1NodesResponse struct{
    ValidateResult nodes.ValidateResult
}

func NewValidateBaremetalV1NodesResponse(validateResult nodes.ValidateResult,)*ValidateBaremetalV1NodesResponse {
    return &ValidateBaremetalV1NodesResponse{
            ValidateResult:validateResult,
    }
}

// action function
func (oc *OpenstackClient) ValidateBaremetalV1Nodes(req *ValidateBaremetalV1NodesRequest)(*ValidateBaremetalV1NodesResponse){
    return NewValidateBaremetalV1NodesResponse(nodes.Validate(oc.Client,req.Id, ))

}
//request struct for the InjectNMIBaremetalV1Nodes
type InjectNMIBaremetalV1NodesRequest struct{
    Id string
}

func NewInjectNMIBaremetalV1NodesRequest()*InjectNMIBaremetalV1NodesRequest{
    return &InjectNMIBaremetalV1NodesRequest{}
}

//response struct for the InjectNMIBaremetalV1Nodes
type InjectNMIBaremetalV1NodesResponse struct{
    InjectNMIResult nodes.InjectNMIResult
}

func NewInjectNMIBaremetalV1NodesResponse(injectNMIResult nodes.InjectNMIResult,)*InjectNMIBaremetalV1NodesResponse {
    return &InjectNMIBaremetalV1NodesResponse{
            InjectNMIResult:injectNMIResult,
    }
}

// action function
func (oc *OpenstackClient) InjectNMIBaremetalV1Nodes(req *InjectNMIBaremetalV1NodesRequest)(*InjectNMIBaremetalV1NodesResponse){
    return NewInjectNMIBaremetalV1NodesResponse(nodes.InjectNMI(oc.Client,req.Id, ))

}
//request struct for the SetBootDeviceBaremetalV1Nodes
type SetBootDeviceBaremetalV1NodesRequest struct{
    Id string
    BootDevice nodes.BootDeviceOpts
}

func NewSetBootDeviceBaremetalV1NodesRequest()*SetBootDeviceBaremetalV1NodesRequest{
    return &SetBootDeviceBaremetalV1NodesRequest{}
}

//response struct for the SetBootDeviceBaremetalV1Nodes
type SetBootDeviceBaremetalV1NodesResponse struct{
    SetBootDeviceResult nodes.SetBootDeviceResult
}

func NewSetBootDeviceBaremetalV1NodesResponse(setBootDeviceResult nodes.SetBootDeviceResult,)*SetBootDeviceBaremetalV1NodesResponse {
    return &SetBootDeviceBaremetalV1NodesResponse{
            SetBootDeviceResult:setBootDeviceResult,
    }
}

// action function
func (oc *OpenstackClient) SetBootDeviceBaremetalV1Nodes(req *SetBootDeviceBaremetalV1NodesRequest)(*SetBootDeviceBaremetalV1NodesResponse){
    return NewSetBootDeviceBaremetalV1NodesResponse(nodes.SetBootDevice(oc.Client,req.Id,req.BootDevice, ))

}
//request struct for the GetBootDeviceBaremetalV1Nodes
type GetBootDeviceBaremetalV1NodesRequest struct{
    Id string
}

func NewGetBootDeviceBaremetalV1NodesRequest()*GetBootDeviceBaremetalV1NodesRequest{
    return &GetBootDeviceBaremetalV1NodesRequest{}
}

//response struct for the GetBootDeviceBaremetalV1Nodes
type GetBootDeviceBaremetalV1NodesResponse struct{
    BootDeviceResult nodes.BootDeviceResult
}

func NewGetBootDeviceBaremetalV1NodesResponse(bootDeviceResult nodes.BootDeviceResult,)*GetBootDeviceBaremetalV1NodesResponse {
    return &GetBootDeviceBaremetalV1NodesResponse{
            BootDeviceResult:bootDeviceResult,
    }
}

// action function
func (oc *OpenstackClient) GetBootDeviceBaremetalV1Nodes(req *GetBootDeviceBaremetalV1NodesRequest)(*GetBootDeviceBaremetalV1NodesResponse){
    return NewGetBootDeviceBaremetalV1NodesResponse(nodes.GetBootDevice(oc.Client,req.Id, ))

}
//request struct for the GetSupportedBootDevicesBaremetalV1Nodes
type GetSupportedBootDevicesBaremetalV1NodesRequest struct{
    Id string
}

func NewGetSupportedBootDevicesBaremetalV1NodesRequest()*GetSupportedBootDevicesBaremetalV1NodesRequest{
    return &GetSupportedBootDevicesBaremetalV1NodesRequest{}
}

//response struct for the GetSupportedBootDevicesBaremetalV1Nodes
type GetSupportedBootDevicesBaremetalV1NodesResponse struct{
    SupportedBootDeviceResult nodes.SupportedBootDeviceResult
}

func NewGetSupportedBootDevicesBaremetalV1NodesResponse(supportedBootDeviceResult nodes.SupportedBootDeviceResult,)*GetSupportedBootDevicesBaremetalV1NodesResponse {
    return &GetSupportedBootDevicesBaremetalV1NodesResponse{
            SupportedBootDeviceResult:supportedBootDeviceResult,
    }
}

// action function
func (oc *OpenstackClient) GetSupportedBootDevicesBaremetalV1Nodes(req *GetSupportedBootDevicesBaremetalV1NodesRequest)(*GetSupportedBootDevicesBaremetalV1NodesResponse){
    return NewGetSupportedBootDevicesBaremetalV1NodesResponse(nodes.GetSupportedBootDevices(oc.Client,req.Id, ))

}
//request struct for the ChangeProvisionStateBaremetalV1Nodes
type ChangeProvisionStateBaremetalV1NodesRequest struct{
    Id string
    Opts nodes.ProvisionStateOpts
}

func NewChangeProvisionStateBaremetalV1NodesRequest()*ChangeProvisionStateBaremetalV1NodesRequest{
    return &ChangeProvisionStateBaremetalV1NodesRequest{}
}

//response struct for the ChangeProvisionStateBaremetalV1Nodes
type ChangeProvisionStateBaremetalV1NodesResponse struct{
    ChangeStateResult nodes.ChangeStateResult
}

func NewChangeProvisionStateBaremetalV1NodesResponse(changeStateResult nodes.ChangeStateResult,)*ChangeProvisionStateBaremetalV1NodesResponse {
    return &ChangeProvisionStateBaremetalV1NodesResponse{
            ChangeStateResult:changeStateResult,
    }
}

// action function
func (oc *OpenstackClient) ChangeProvisionStateBaremetalV1Nodes(req *ChangeProvisionStateBaremetalV1NodesRequest)(*ChangeProvisionStateBaremetalV1NodesResponse){
    return NewChangeProvisionStateBaremetalV1NodesResponse(nodes.ChangeProvisionState(oc.Client,req.Id,req.Opts, ))

}
//request struct for the ChangePowerStateBaremetalV1Nodes
type ChangePowerStateBaremetalV1NodesRequest struct{
    Id string
    Opts nodes.PowerStateOpts
}

func NewChangePowerStateBaremetalV1NodesRequest()*ChangePowerStateBaremetalV1NodesRequest{
    return &ChangePowerStateBaremetalV1NodesRequest{}
}

//response struct for the ChangePowerStateBaremetalV1Nodes
type ChangePowerStateBaremetalV1NodesResponse struct{
    ChangePowerStateResult nodes.ChangePowerStateResult
}

func NewChangePowerStateBaremetalV1NodesResponse(changePowerStateResult nodes.ChangePowerStateResult,)*ChangePowerStateBaremetalV1NodesResponse {
    return &ChangePowerStateBaremetalV1NodesResponse{
            ChangePowerStateResult:changePowerStateResult,
    }
}

// action function
func (oc *OpenstackClient) ChangePowerStateBaremetalV1Nodes(req *ChangePowerStateBaremetalV1NodesRequest)(*ChangePowerStateBaremetalV1NodesResponse){
    return NewChangePowerStateBaremetalV1NodesResponse(nodes.ChangePowerState(oc.Client,req.Id,req.Opts, ))

}
//request struct for the SetRAIDConfigBaremetalV1Nodes
type SetRAIDConfigBaremetalV1NodesRequest struct{
    Id string
    RaidConfigOptsBuilder nodes.RAIDConfigOpts
}

func NewSetRAIDConfigBaremetalV1NodesRequest()*SetRAIDConfigBaremetalV1NodesRequest{
    return &SetRAIDConfigBaremetalV1NodesRequest{}
}

//response struct for the SetRAIDConfigBaremetalV1Nodes
type SetRAIDConfigBaremetalV1NodesResponse struct{
    ChangeStateResult nodes.ChangeStateResult
}

func NewSetRAIDConfigBaremetalV1NodesResponse(changeStateResult nodes.ChangeStateResult,)*SetRAIDConfigBaremetalV1NodesResponse {
    return &SetRAIDConfigBaremetalV1NodesResponse{
            ChangeStateResult:changeStateResult,
    }
}

// action function
func (oc *OpenstackClient) SetRAIDConfigBaremetalV1Nodes(req *SetRAIDConfigBaremetalV1NodesRequest)(*SetRAIDConfigBaremetalV1NodesResponse){
    return NewSetRAIDConfigBaremetalV1NodesResponse(nodes.SetRAIDConfig(oc.Client,req.Id,req.RaidConfigOptsBuilder, ))

}
//request struct for the ListBIOSSettingsBaremetalV1Nodes
type ListBIOSSettingsBaremetalV1NodesRequest struct{
    Id string
    Opts nodes.ListBIOSSettingsOpts
}

func NewListBIOSSettingsBaremetalV1NodesRequest()*ListBIOSSettingsBaremetalV1NodesRequest{
    return &ListBIOSSettingsBaremetalV1NodesRequest{}
}

//response struct for the ListBIOSSettingsBaremetalV1Nodes
type ListBIOSSettingsBaremetalV1NodesResponse struct{
    ListBIOSSettingsResult nodes.ListBIOSSettingsResult
}

func NewListBIOSSettingsBaremetalV1NodesResponse(listBIOSSettingsResult nodes.ListBIOSSettingsResult,)*ListBIOSSettingsBaremetalV1NodesResponse {
    return &ListBIOSSettingsBaremetalV1NodesResponse{
            ListBIOSSettingsResult:listBIOSSettingsResult,
    }
}

// action function
func (oc *OpenstackClient) ListBIOSSettingsBaremetalV1Nodes(req *ListBIOSSettingsBaremetalV1NodesRequest)(*ListBIOSSettingsBaremetalV1NodesResponse){
    return NewListBIOSSettingsBaremetalV1NodesResponse(nodes.ListBIOSSettings(oc.Client,req.Id,req.Opts, ))

}
//request struct for the GetBIOSSettingBaremetalV1Nodes
type GetBIOSSettingBaremetalV1NodesRequest struct{
    Id string
    Setting string
}

func NewGetBIOSSettingBaremetalV1NodesRequest()*GetBIOSSettingBaremetalV1NodesRequest{
    return &GetBIOSSettingBaremetalV1NodesRequest{}
}

//response struct for the GetBIOSSettingBaremetalV1Nodes
type GetBIOSSettingBaremetalV1NodesResponse struct{
    GetBIOSSettingResult nodes.GetBIOSSettingResult
}

func NewGetBIOSSettingBaremetalV1NodesResponse(getBIOSSettingResult nodes.GetBIOSSettingResult,)*GetBIOSSettingBaremetalV1NodesResponse {
    return &GetBIOSSettingBaremetalV1NodesResponse{
            GetBIOSSettingResult:getBIOSSettingResult,
    }
}

// action function
func (oc *OpenstackClient) GetBIOSSettingBaremetalV1Nodes(req *GetBIOSSettingBaremetalV1NodesRequest)(*GetBIOSSettingBaremetalV1NodesResponse){
    return NewGetBIOSSettingBaremetalV1NodesResponse(nodes.GetBIOSSetting(oc.Client,req.Id,req.Setting, ))

}
//request struct for the GetVendorPassthruMethodsBaremetalV1Nodes
type GetVendorPassthruMethodsBaremetalV1NodesRequest struct{
    Id string
}

func NewGetVendorPassthruMethodsBaremetalV1NodesRequest()*GetVendorPassthruMethodsBaremetalV1NodesRequest{
    return &GetVendorPassthruMethodsBaremetalV1NodesRequest{}
}

//response struct for the GetVendorPassthruMethodsBaremetalV1Nodes
type GetVendorPassthruMethodsBaremetalV1NodesResponse struct{
    VendorPassthruMethodsResult nodes.VendorPassthruMethodsResult
}

func NewGetVendorPassthruMethodsBaremetalV1NodesResponse(vendorPassthruMethodsResult nodes.VendorPassthruMethodsResult,)*GetVendorPassthruMethodsBaremetalV1NodesResponse {
    return &GetVendorPassthruMethodsBaremetalV1NodesResponse{
            VendorPassthruMethodsResult:vendorPassthruMethodsResult,
    }
}

// action function
func (oc *OpenstackClient) GetVendorPassthruMethodsBaremetalV1Nodes(req *GetVendorPassthruMethodsBaremetalV1NodesRequest)(*GetVendorPassthruMethodsBaremetalV1NodesResponse){
    return NewGetVendorPassthruMethodsBaremetalV1NodesResponse(nodes.GetVendorPassthruMethods(oc.Client,req.Id, ))

}
//request struct for the GetAllSubscriptionsBaremetalV1Nodes
type GetAllSubscriptionsBaremetalV1NodesRequest struct{
    Id string
    Method nodes.CallVendorPassthruOpts
}

func NewGetAllSubscriptionsBaremetalV1NodesRequest()*GetAllSubscriptionsBaremetalV1NodesRequest{
    return &GetAllSubscriptionsBaremetalV1NodesRequest{}
}

//response struct for the GetAllSubscriptionsBaremetalV1Nodes
type GetAllSubscriptionsBaremetalV1NodesResponse struct{
    GetAllSubscriptionsVendorPassthruResult nodes.GetAllSubscriptionsVendorPassthruResult
}

func NewGetAllSubscriptionsBaremetalV1NodesResponse(getAllSubscriptionsVendorPassthruResult nodes.GetAllSubscriptionsVendorPassthruResult,)*GetAllSubscriptionsBaremetalV1NodesResponse {
    return &GetAllSubscriptionsBaremetalV1NodesResponse{
            GetAllSubscriptionsVendorPassthruResult:getAllSubscriptionsVendorPassthruResult,
    }
}

// action function
func (oc *OpenstackClient) GetAllSubscriptionsBaremetalV1Nodes(req *GetAllSubscriptionsBaremetalV1NodesRequest)(*GetAllSubscriptionsBaremetalV1NodesResponse){
    return NewGetAllSubscriptionsBaremetalV1NodesResponse(nodes.GetAllSubscriptions(oc.Client,req.Id,req.Method, ))

}
//request struct for the GetSubscriptionBaremetalV1Nodes
type GetSubscriptionBaremetalV1NodesRequest struct{
    Id string
    Method nodes.CallVendorPassthruOpts
    SubscriptionOpts nodes.GetSubscriptionOpts
}

func NewGetSubscriptionBaremetalV1NodesRequest()*GetSubscriptionBaremetalV1NodesRequest{
    return &GetSubscriptionBaremetalV1NodesRequest{}
}

//response struct for the GetSubscriptionBaremetalV1Nodes
type GetSubscriptionBaremetalV1NodesResponse struct{
    SubscriptionVendorPassthruResult nodes.SubscriptionVendorPassthruResult
}

func NewGetSubscriptionBaremetalV1NodesResponse(subscriptionVendorPassthruResult nodes.SubscriptionVendorPassthruResult,)*GetSubscriptionBaremetalV1NodesResponse {
    return &GetSubscriptionBaremetalV1NodesResponse{
            SubscriptionVendorPassthruResult:subscriptionVendorPassthruResult,
    }
}

// action function
func (oc *OpenstackClient) GetSubscriptionBaremetalV1Nodes(req *GetSubscriptionBaremetalV1NodesRequest)(*GetSubscriptionBaremetalV1NodesResponse){
    return NewGetSubscriptionBaremetalV1NodesResponse(nodes.GetSubscription(oc.Client,req.Id,req.Method,req.SubscriptionOpts, ))

}
//request struct for the DeleteSubscriptionBaremetalV1Nodes
type DeleteSubscriptionBaremetalV1NodesRequest struct{
    Id string
    Method nodes.CallVendorPassthruOpts
    SubscriptionOpts nodes.DeleteSubscriptionOpts
}

func NewDeleteSubscriptionBaremetalV1NodesRequest()*DeleteSubscriptionBaremetalV1NodesRequest{
    return &DeleteSubscriptionBaremetalV1NodesRequest{}
}

//response struct for the DeleteSubscriptionBaremetalV1Nodes
type DeleteSubscriptionBaremetalV1NodesResponse struct{
    DeleteSubscriptionVendorPassthruResult nodes.DeleteSubscriptionVendorPassthruResult
}

func NewDeleteSubscriptionBaremetalV1NodesResponse(deleteSubscriptionVendorPassthruResult nodes.DeleteSubscriptionVendorPassthruResult,)*DeleteSubscriptionBaremetalV1NodesResponse {
    return &DeleteSubscriptionBaremetalV1NodesResponse{
            DeleteSubscriptionVendorPassthruResult:deleteSubscriptionVendorPassthruResult,
    }
}

// action function
func (oc *OpenstackClient) DeleteSubscriptionBaremetalV1Nodes(req *DeleteSubscriptionBaremetalV1NodesRequest)(*DeleteSubscriptionBaremetalV1NodesResponse){
    return NewDeleteSubscriptionBaremetalV1NodesResponse(nodes.DeleteSubscription(oc.Client,req.Id,req.Method,req.SubscriptionOpts, ))

}
//request struct for the CreateSubscriptionBaremetalV1Nodes
type CreateSubscriptionBaremetalV1NodesRequest struct{
    Id string
    Method nodes.CallVendorPassthruOpts
    SubscriptionOpts nodes.CreateSubscriptionOpts
}

func NewCreateSubscriptionBaremetalV1NodesRequest()*CreateSubscriptionBaremetalV1NodesRequest{
    return &CreateSubscriptionBaremetalV1NodesRequest{}
}

//response struct for the CreateSubscriptionBaremetalV1Nodes
type CreateSubscriptionBaremetalV1NodesResponse struct{
    SubscriptionVendorPassthruResult nodes.SubscriptionVendorPassthruResult
}

func NewCreateSubscriptionBaremetalV1NodesResponse(subscriptionVendorPassthruResult nodes.SubscriptionVendorPassthruResult,)*CreateSubscriptionBaremetalV1NodesResponse {
    return &CreateSubscriptionBaremetalV1NodesResponse{
            SubscriptionVendorPassthruResult:subscriptionVendorPassthruResult,
    }
}

// action function
func (oc *OpenstackClient) CreateSubscriptionBaremetalV1Nodes(req *CreateSubscriptionBaremetalV1NodesRequest)(*CreateSubscriptionBaremetalV1NodesResponse){
    return NewCreateSubscriptionBaremetalV1NodesResponse(nodes.CreateSubscription(oc.Client,req.Id,req.Method,req.SubscriptionOpts, ))

}
//request struct for the SetMaintenanceBaremetalV1Nodes
type SetMaintenanceBaremetalV1NodesRequest struct{
    Id string
    Opts nodes.MaintenanceOpts
}

func NewSetMaintenanceBaremetalV1NodesRequest()*SetMaintenanceBaremetalV1NodesRequest{
    return &SetMaintenanceBaremetalV1NodesRequest{}
}

//response struct for the SetMaintenanceBaremetalV1Nodes
type SetMaintenanceBaremetalV1NodesResponse struct{
    SetMaintenanceResult nodes.SetMaintenanceResult
}

func NewSetMaintenanceBaremetalV1NodesResponse(setMaintenanceResult nodes.SetMaintenanceResult,)*SetMaintenanceBaremetalV1NodesResponse {
    return &SetMaintenanceBaremetalV1NodesResponse{
            SetMaintenanceResult:setMaintenanceResult,
    }
}

// action function
func (oc *OpenstackClient) SetMaintenanceBaremetalV1Nodes(req *SetMaintenanceBaremetalV1NodesRequest)(*SetMaintenanceBaremetalV1NodesResponse){
    return NewSetMaintenanceBaremetalV1NodesResponse(nodes.SetMaintenance(oc.Client,req.Id,req.Opts, ))

}
//request struct for the UnsetMaintenanceBaremetalV1Nodes
type UnsetMaintenanceBaremetalV1NodesRequest struct{
    Id string
}

func NewUnsetMaintenanceBaremetalV1NodesRequest()*UnsetMaintenanceBaremetalV1NodesRequest{
    return &UnsetMaintenanceBaremetalV1NodesRequest{}
}

//response struct for the UnsetMaintenanceBaremetalV1Nodes
type UnsetMaintenanceBaremetalV1NodesResponse struct{
    SetMaintenanceResult nodes.SetMaintenanceResult
}

func NewUnsetMaintenanceBaremetalV1NodesResponse(setMaintenanceResult nodes.SetMaintenanceResult,)*UnsetMaintenanceBaremetalV1NodesResponse {
    return &UnsetMaintenanceBaremetalV1NodesResponse{
            SetMaintenanceResult:setMaintenanceResult,
    }
}

// action function
func (oc *OpenstackClient) UnsetMaintenanceBaremetalV1Nodes(req *UnsetMaintenanceBaremetalV1NodesRequest)(*UnsetMaintenanceBaremetalV1NodesResponse){
    return NewUnsetMaintenanceBaremetalV1NodesResponse(nodes.UnsetMaintenance(oc.Client,req.Id, ))

}