package openstack

// Code generated by cloud manager.





import (
    "github.com/gophercloud/gophercloud/openstack/db/v1/databases"
)


//extract response info from pager for ListDbV1Databases
func ExtractListDbV1DatabasesResponse(response *ListDbV1DatabasesResponse)([]databases.Database,error){
	page, err := response.Pager.AllPages()
	if err != nil {
		return nil, err
	}
	return databases.ExtractDBs(page)
}