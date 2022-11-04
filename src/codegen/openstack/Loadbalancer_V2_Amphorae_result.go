package openstack

// Code generated by cloud manager.





import (
    "github.com/gophercloud/gophercloud/openstack/loadbalancer/v2/amphorae"
)


//extract response info from pager for ListLoadbalancerV2Amphorae
func ExtractListLoadbalancerV2AmphoraeResponse(response *ListLoadbalancerV2AmphoraeResponse)([]amphorae.Amphora,error){
	page, err := response.Pager.AllPages()
	if err != nil {
		return nil, err
	}
	return amphorae.ExtractAmphorae(page)
}