package openstack

// Code generated by cloud manager.





import (
    "github.com/gophercloud/gophercloud/openstack/sharedfilesystems/v2/snapshots"
)


//extract response info from pager for ListDetailSharedfilesystemsV2Snapshots
func ExtractListDetailSharedfilesystemsV2SnapshotsResponse(response *ListDetailSharedfilesystemsV2SnapshotsResponse)([]snapshots.Snapshot,error){
	page, err := response.Pager.AllPages()
	if err != nil {
		return nil, err
	}
	return snapshots.ExtractSnapshots(page)
}