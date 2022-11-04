package openstack

// Code generated by cloud manager.





import (
    "github.com/gophercloud/gophercloud/openstack/loadbalancer/v2/l7policies"
)


//extract response info from pager for ListLoadbalancerV2L7policies
func ExtractListLoadbalancerV2L7policiesResponse(response *ListLoadbalancerV2L7policiesResponse)([]l7policies.L7Policy,error){
	page, err := response.Pager.AllPages()
	if err != nil {
		return nil, err
	}
	return l7policies.ExtractL7Policies(page)
}


//extract response info from pager for ListRulesLoadbalancerV2L7policies
func ExtractListRulesLoadbalancerV2L7policiesResponse(response *ListRulesLoadbalancerV2L7policiesResponse)([]l7policies.Rule,error){
	page, err := response.Pager.AllPages()
	if err != nil {
		return nil, err
	}
	return l7policies.ExtractRules(page)
}