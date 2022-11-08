package openstack

// Code generated by cloud manager.

import (
	"github.com/gophercloud/gophercloud/openstack/identity/v2/users"
)

// extract response info from pager for ListIdentityV2Users
func ExtractListIdentityV2UsersResponse(response *ListIdentityV2UsersResponse) ([]users.User, error) {
	page, err := response.Pager.AllPages()
	if err != nil {
		return nil, err
	}
	return users.ExtractUsers(page)
}

// call result's extract function
func ExtractCreateIdentityV2UsersResponse(response *CreateIdentityV2UsersResponse) (interface{}, error) {
	return response.CreateResult.Body, response.CreateResult.Err
}

// call result's extract function
func ExtractGetIdentityV2UsersResponse(response *GetIdentityV2UsersResponse) (interface{}, error) {
	return response.GetResult.Body, response.GetResult.Err
}

// call result's extract function
func ExtractUpdateIdentityV2UsersResponse(response *UpdateIdentityV2UsersResponse) (interface{}, error) {
	return response.UpdateResult.Body, response.UpdateResult.Err
}

// call result's extract function
func ExtractDeleteIdentityV2UsersResponse(response *DeleteIdentityV2UsersResponse) (interface{}, error) {
	return response.DeleteResult.Body, response.DeleteResult.Err
}

// extract response info from pager for ListRolesIdentityV2Users
func ExtractListRolesIdentityV2UsersResponse(response *ListRolesIdentityV2UsersResponse) ([]users.Role, error) {
	page, err := response.Pager.AllPages()
	if err != nil {
		return nil, err
	}
	return users.ExtractRoles(page)
}
