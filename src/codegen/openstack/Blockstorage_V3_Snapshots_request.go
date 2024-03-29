package openstack

// Code generated by cloud manager.

import (
    "github.com/gophercloud/gophercloud/openstack/blockstorage/v3/snapshots"
    "github.com/gophercloud/gophercloud/pagination"
)
//request struct for the CreateBlockstorageV3Snapshots
type CreateBlockstorageV3SnapshotsRequest struct{
    Opts snapshots.CreateOpts
}

func NewCreateBlockstorageV3SnapshotsRequest()*CreateBlockstorageV3SnapshotsRequest{
    return &CreateBlockstorageV3SnapshotsRequest{}
}

//response struct for the CreateBlockstorageV3Snapshots
type CreateBlockstorageV3SnapshotsResponse struct{
    CreateResult snapshots.CreateResult
}

func NewCreateBlockstorageV3SnapshotsResponse(createResult snapshots.CreateResult,)*CreateBlockstorageV3SnapshotsResponse {
    return &CreateBlockstorageV3SnapshotsResponse{
            CreateResult:createResult,
    }
}

// action function
func (oc *OpenstackClient) CreateBlockstorageV3Snapshots(req *CreateBlockstorageV3SnapshotsRequest)(*CreateBlockstorageV3SnapshotsResponse){
    return NewCreateBlockstorageV3SnapshotsResponse(snapshots.Create(oc.Client,req.Opts, ))

}
//request struct for the DeleteBlockstorageV3Snapshots
type DeleteBlockstorageV3SnapshotsRequest struct{
    Id string
}

func NewDeleteBlockstorageV3SnapshotsRequest()*DeleteBlockstorageV3SnapshotsRequest{
    return &DeleteBlockstorageV3SnapshotsRequest{}
}

//response struct for the DeleteBlockstorageV3Snapshots
type DeleteBlockstorageV3SnapshotsResponse struct{
    DeleteResult snapshots.DeleteResult
}

func NewDeleteBlockstorageV3SnapshotsResponse(deleteResult snapshots.DeleteResult,)*DeleteBlockstorageV3SnapshotsResponse {
    return &DeleteBlockstorageV3SnapshotsResponse{
            DeleteResult:deleteResult,
    }
}

// action function
func (oc *OpenstackClient) DeleteBlockstorageV3Snapshots(req *DeleteBlockstorageV3SnapshotsRequest)(*DeleteBlockstorageV3SnapshotsResponse){
    return NewDeleteBlockstorageV3SnapshotsResponse(snapshots.Delete(oc.Client,req.Id, ))

}
//request struct for the GetBlockstorageV3Snapshots
type GetBlockstorageV3SnapshotsRequest struct{
    Id string
}

func NewGetBlockstorageV3SnapshotsRequest()*GetBlockstorageV3SnapshotsRequest{
    return &GetBlockstorageV3SnapshotsRequest{}
}

//response struct for the GetBlockstorageV3Snapshots
type GetBlockstorageV3SnapshotsResponse struct{
    GetResult snapshots.GetResult
}

func NewGetBlockstorageV3SnapshotsResponse(getResult snapshots.GetResult,)*GetBlockstorageV3SnapshotsResponse {
    return &GetBlockstorageV3SnapshotsResponse{
            GetResult:getResult,
    }
}

// action function
func (oc *OpenstackClient) GetBlockstorageV3Snapshots(req *GetBlockstorageV3SnapshotsRequest)(*GetBlockstorageV3SnapshotsResponse){
    return NewGetBlockstorageV3SnapshotsResponse(snapshots.Get(oc.Client,req.Id, ))

}
//request struct for the ListBlockstorageV3Snapshots
type ListBlockstorageV3SnapshotsRequest struct{
    Opts snapshots.ListOpts
}

func NewListBlockstorageV3SnapshotsRequest()*ListBlockstorageV3SnapshotsRequest{
    return &ListBlockstorageV3SnapshotsRequest{}
}

//response struct for the ListBlockstorageV3Snapshots
type ListBlockstorageV3SnapshotsResponse struct{
    Pager pagination.Pager
}

func NewListBlockstorageV3SnapshotsResponse(pager pagination.Pager,)*ListBlockstorageV3SnapshotsResponse {
    return &ListBlockstorageV3SnapshotsResponse{
            Pager:pager,
    }
}

// action function
func (oc *OpenstackClient) ListBlockstorageV3Snapshots(req *ListBlockstorageV3SnapshotsRequest)(*ListBlockstorageV3SnapshotsResponse){
    return NewListBlockstorageV3SnapshotsResponse(snapshots.List(oc.Client,req.Opts, ))

}
//request struct for the UpdateBlockstorageV3Snapshots
type UpdateBlockstorageV3SnapshotsRequest struct{
    Id string
    Opts snapshots.UpdateOpts
}

func NewUpdateBlockstorageV3SnapshotsRequest()*UpdateBlockstorageV3SnapshotsRequest{
    return &UpdateBlockstorageV3SnapshotsRequest{}
}

//response struct for the UpdateBlockstorageV3Snapshots
type UpdateBlockstorageV3SnapshotsResponse struct{
    UpdateResult snapshots.UpdateResult
}

func NewUpdateBlockstorageV3SnapshotsResponse(updateResult snapshots.UpdateResult,)*UpdateBlockstorageV3SnapshotsResponse {
    return &UpdateBlockstorageV3SnapshotsResponse{
            UpdateResult:updateResult,
    }
}

// action function
func (oc *OpenstackClient) UpdateBlockstorageV3Snapshots(req *UpdateBlockstorageV3SnapshotsRequest)(*UpdateBlockstorageV3SnapshotsResponse){
    return NewUpdateBlockstorageV3SnapshotsResponse(snapshots.Update(oc.Client,req.Id,req.Opts, ))

}
//request struct for the UpdateMetadataBlockstorageV3Snapshots
type UpdateMetadataBlockstorageV3SnapshotsRequest struct{
    Id string
    Opts snapshots.UpdateMetadataOpts
}

func NewUpdateMetadataBlockstorageV3SnapshotsRequest()*UpdateMetadataBlockstorageV3SnapshotsRequest{
    return &UpdateMetadataBlockstorageV3SnapshotsRequest{}
}

//response struct for the UpdateMetadataBlockstorageV3Snapshots
type UpdateMetadataBlockstorageV3SnapshotsResponse struct{
    UpdateMetadataResult snapshots.UpdateMetadataResult
}

func NewUpdateMetadataBlockstorageV3SnapshotsResponse(updateMetadataResult snapshots.UpdateMetadataResult,)*UpdateMetadataBlockstorageV3SnapshotsResponse {
    return &UpdateMetadataBlockstorageV3SnapshotsResponse{
            UpdateMetadataResult:updateMetadataResult,
    }
}

// action function
func (oc *OpenstackClient) UpdateMetadataBlockstorageV3Snapshots(req *UpdateMetadataBlockstorageV3SnapshotsRequest)(*UpdateMetadataBlockstorageV3SnapshotsResponse){
    return NewUpdateMetadataBlockstorageV3SnapshotsResponse(snapshots.UpdateMetadata(oc.Client,req.Id,req.Opts, ))

}