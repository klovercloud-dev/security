package v1

import (
	"errors"
	v1 "github.com/klovercloud-ci/core/v1"
	"github.com/stretchr/testify/assert"
	"reflect"
	"testing"
)

func TestCheckAuthority(t *testing.T) {
	type TestCase struct {
		actual error
		expected error
	}

	var testCases [] TestCase

	testCases= append(testCases,TestCase{
		actual:   checkAuthority(v1.UserResourcePermission{Resources: []v1.ResourceWiseRoles{{Name: "user", Roles: []v1.Role{{Name: "ADMIN", Permissions: []v1.Permission{{Name: "CREATE"}}}}}}},"user","ADMIN",""),
		expected: nil,
	})
	testCases= append(testCases,TestCase{
		actual:   checkAuthority(v1.UserResourcePermission{Resources: []v1.ResourceWiseRoles{{Name: "pipeline", Roles: []v1.Role{{Name: "ADMIN", Permissions: []v1.Permission{{Name: "CREATE"}}}}}}},"user","ADMIN",""),
		expected: errors.New("[ERROR]: Insufficient permission"),
	})
	testCases= append(testCases,TestCase{
		actual:   checkAuthority(v1.UserResourcePermission{Resources: []v1.ResourceWiseRoles{{Name: "user", Roles: []v1.Role{{Name: "ADMIN", Permissions: []v1.Permission{{Name: "CREATE"}}}}}}},"user","","CREATE"),
		expected:nil,
	})
	testCases= append(testCases,TestCase{
		actual:   checkAuthority(v1.UserResourcePermission{Resources: []v1.ResourceWiseRoles{{Name: "pipeline", Roles: []v1.Role{{Name: "ADMIN", Permissions: []v1.Permission{{Name: "CREATE"}}}}}}},"user","",""),
		expected: errors.New("[ERROR]: Insufficient permission"),
	})

	for _, each := range testCases {
		if !reflect.DeepEqual(each.expected, each.actual) {
			//log.Println(each.expected, " ", each.actual)
			assert.ElementsMatch(t, each.expected, each.actual)
		}
	}
}