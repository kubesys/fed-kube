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