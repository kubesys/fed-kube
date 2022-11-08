package openstack

// Code generated by cloud manager.

import (
    "github.com/gophercloud/gophercloud/openstack/imageservice/v2/images"
)


//extract response info from pager for ListImageserviceV2Images
func ExtractListImageserviceV2ImagesResponse(response *ListImageserviceV2ImagesResponse)([]images.Image,error){
	page, err := response.Pager.AllPages()
	if err != nil {
		return nil, err
	}
	return images.ExtractImages(page)
}


// call result's extract function
func ExtractCreateImageserviceV2ImagesResponse(response *CreateImageserviceV2ImagesResponse)(interface{}, error){
    return response.CreateResult.Body, response.CreateResult.Err
}



// call result's extract function
func ExtractDeleteImageserviceV2ImagesResponse(response *DeleteImageserviceV2ImagesResponse)(interface{}, error){
    return response.DeleteResult.Body, response.DeleteResult.Err
}



// call result's extract function
func ExtractGetImageserviceV2ImagesResponse(response *GetImageserviceV2ImagesResponse)(interface{}, error){
    return response.GetResult.Body, response.GetResult.Err
}



// call result's extract function
func ExtractUpdateImageserviceV2ImagesResponse(response *UpdateImageserviceV2ImagesResponse)(interface{}, error){
    return response.UpdateResult.Body, response.UpdateResult.Err
}
