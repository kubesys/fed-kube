package openstack

// Code generated by cloud manager.

import (
    "github.com/gophercloud/gophercloud/openstack/compute/v2/images"
)


//extract response info from pager for ListDetailComputeV2Images
func ExtractListDetailComputeV2ImagesResponse(response *ListDetailComputeV2ImagesResponse)([]images.Image,error){
	page, err := response.Pager.AllPages()
	if err != nil {
		return nil, err
	}
	return images.ExtractImages(page)
}


// call result's extract function
func ExtractGetComputeV2ImagesResponse(response *GetComputeV2ImagesResponse)(interface{}, error){
    return response.GetResult.Body, response.GetResult.Err
}



// call result's extract function
func ExtractDeleteComputeV2ImagesResponse(response *DeleteComputeV2ImagesResponse)(interface{}, error){
    return response.DeleteResult.Body, response.DeleteResult.Err
}