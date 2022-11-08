package openstack

// Code generated by cloud manager.

import (
	"github.com/gophercloud/gophercloud/openstack/objectstorage/v1/objects"
)

// extract response info from pager for ListObjectstorageV1Objects
func ExtractListObjectstorageV1ObjectsResponse(response *ListObjectstorageV1ObjectsResponse) ([]objects.Object, error) {
	page, err := response.Pager.AllPages()
	if err != nil {
		return nil, err
	}
	return objects.ExtractInfo(page)
}

// call result's extract function
func ExtractDownloadObjectstorageV1ObjectsResponse(response *DownloadObjectstorageV1ObjectsResponse) (interface{}, error) {
	return response.DownloadResult.Body, response.DownloadResult.Err
}

// call result's extract function
func ExtractCreateObjectstorageV1ObjectsResponse(response *CreateObjectstorageV1ObjectsResponse) (interface{}, error) {
	return response.CreateResult.Body, response.CreateResult.Err
}

// call result's extract function
func ExtractCopyObjectstorageV1ObjectsResponse(response *CopyObjectstorageV1ObjectsResponse) (interface{}, error) {
	return response.CopyResult.Body, response.CopyResult.Err
}

// call result's extract function
func ExtractDeleteObjectstorageV1ObjectsResponse(response *DeleteObjectstorageV1ObjectsResponse) (interface{}, error) {
	return response.DeleteResult.Body, response.DeleteResult.Err
}

// call result's extract function
func ExtractGetObjectstorageV1ObjectsResponse(response *GetObjectstorageV1ObjectsResponse) (interface{}, error) {
	return response.GetResult.Body, response.GetResult.Err
}

// call result's extract function
func ExtractUpdateObjectstorageV1ObjectsResponse(response *UpdateObjectstorageV1ObjectsResponse) (interface{}, error) {
	return response.UpdateResult.Body, response.UpdateResult.Err
}

// call result's extract function
func ExtractBulkDeleteObjectstorageV1ObjectsResponse(response *BulkDeleteObjectstorageV1ObjectsResponse) (interface{}, error) {
	return response.BulkDeleteResult.Body, response.BulkDeleteResult.Err
}
