package openstack

// Code generated by cloud manager.

import (
    "github.com/gophercloud/gophercloud/openstack/db/v1/users"
)


// call result's extract function
func ExtractCreateDbV1UsersResponse(response *CreateDbV1UsersResponse)(interface{}, error){
    return response.CreateResult.Body, response.CreateResult.Err
}



//extract response info from pager for ListDbV1Users
func ExtractListDbV1UsersResponse(response *ListDbV1UsersResponse)([]users.User,error){
	page, err := response.Pager.AllPages()
	if err != nil {
		return nil, err
	}
	return users.ExtractUsers(page)
}


// call result's extract function
func ExtractDeleteDbV1UsersResponse(response *DeleteDbV1UsersResponse)(interface{}, error){
    return response.DeleteResult.Body, response.DeleteResult.Err
}
