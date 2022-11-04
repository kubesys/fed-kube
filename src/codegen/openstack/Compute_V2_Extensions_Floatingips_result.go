package openstack

// Code generated by cloud manager.





import (
    "github.com/gophercloud/gophercloud/openstack/compute/v2/extensions/floatingips"
)


//extract response info from pager for ListComputeV2ExtensionsFloatingips
func ExtractListComputeV2ExtensionsFloatingipsResponse(response *ListComputeV2ExtensionsFloatingipsResponse)([]floatingips.FloatingIP,error){
	page, err := response.Pager.AllPages()
	if err != nil {
		return nil, err
	}
	return floatingips.ExtractFloatingIPs(page)
}