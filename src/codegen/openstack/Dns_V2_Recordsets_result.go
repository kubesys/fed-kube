package openstack

// Code generated by cloud manager.

import (
    "github.com/gophercloud/gophercloud/openstack/dns/v2/recordsets"
)


//extract response info from pager for ListByZoneDnsV2Recordsets
func ExtractListByZoneDnsV2RecordsetsResponse(response *ListByZoneDnsV2RecordsetsResponse)([]recordsets.RecordSet,error){
	page, err := response.Pager.AllPages()
	if err != nil {
		return nil, err
	}
	return recordsets.ExtractRecordSets(page)
}


// call result's extract function
func ExtractGetDnsV2RecordsetsResponse(response *GetDnsV2RecordsetsResponse)(interface{}, error){
    return response.GetResult.Body, response.GetResult.Err
}



// call result's extract function
func ExtractCreateDnsV2RecordsetsResponse(response *CreateDnsV2RecordsetsResponse)(interface{}, error){
    return response.CreateResult.Body, response.CreateResult.Err
}



// call result's extract function
func ExtractUpdateDnsV2RecordsetsResponse(response *UpdateDnsV2RecordsetsResponse)(interface{}, error){
    return response.UpdateResult.Body, response.UpdateResult.Err
}



// call result's extract function
func ExtractDeleteDnsV2RecordsetsResponse(response *DeleteDnsV2RecordsetsResponse)(interface{}, error){
    return response.DeleteResult.Body, response.DeleteResult.Err
}
