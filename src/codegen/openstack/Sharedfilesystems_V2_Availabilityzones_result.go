package openstack

// Code generated by cloud manager.

import (
    "github.com/gophercloud/gophercloud/openstack/sharedfilesystems/v2/availabilityzones"
)


//extract response info from pager for ListSharedfilesystemsV2Availabilityzones
func ExtractListSharedfilesystemsV2AvailabilityzonesResponse(response *ListSharedfilesystemsV2AvailabilityzonesResponse)([]availabilityzones.AvailabilityZone,error){
	page, err := response.Pager.AllPages()
	if err != nil {
		return nil, err
	}
	return availabilityzones.ExtractAvailabilityZones(page)
}