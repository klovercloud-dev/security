package v1

import (
	"encoding/json"
	"fmt"
	_ "github.com/klovercloud-ci/enums"
	"reflect"
	"testing"
	"github.com/stretchr/testify/assert"
)

// Initialize UserRegistrationDto from JSON data
func InitUserRegistrationDto(str string) UserRegistrationDto {
	res := UserRegistrationDto{}
	err := json.Unmarshal([]byte(str), &res)
	if err != nil {
		return UserRegistrationDto{}
	}
	return res
}

// Test function for testing UserRegistrationDto to User and UserResourcePermission serialization.
func TestGetUserAndResourcePermissionBody(t *testing.T) {
	type resultData struct {
		user User
		userResourcePermission UserResourcePermission
	}

	type TestCase struct {
		data     UserRegistrationDto
		expected resultData
		actual   resultData
	}

	var testdata []TestCase

	jsonData := []string{
		`{
			"_id": "123",
			"first_name": "Shabrul",
			"last_name": "TheBOSS",
			"email": "shabrul_theboss@klovercloud.com",
			"password":	"IAmTheBoss",
			"status": "Married",
			"created_date": "2022-01-10",
			"updated_date": "2022-01-10",
			"token": "1231654sdaf61v1a1rg",
			"refresh_token": "4dsaf4e9a4f89484",
			"auth_type": "Secured",
			"resource_permission":	{
				"user_id": "123",
				"resources": [
					{
						"name": "resources1",
						"roles": [
							{
								"name": "admin",
								"permissions": [
									{
										"name": "create"
									},
									{
										"name": "read"
									}
								]
							}
						]
					}
				]
			}
		}`,
	}

	expec := []resultData{
		{
			User{
				ID:           "123",
				FirstName:    "Shabrul",
				LastName:     "TheBOSS",
				Email:        "shabrul_theboss@klovercloud.com",
				Password:     "IAmTheBoss",
				Status:       "Married",
				CreatedDate:  "2022-01-10",
				UpdatedDate:  "2022-01-10",
				Token:        "1231654sdaf61v1a1rg",
				RefreshToken: "4dsaf4e9a4f89484",
				AuthType:     "Secured",
			},
			UserResourcePermission{
				UserId:    "123",
				Resources: []struct {
					Name  string `json:"name" bson:"name"`
					Roles []Role `json:"roles" bson:"roles"`
				}{
					{
						Name: "resources1",
						Roles: []Role{
							{
								Name: "admin",
								Permissions: []Permission{
									{
										Name: "create",
									},
									{
										Name: "read",
									},
								},
							},
						},
					},
				},
			},
		},
	}

	for i := 0; i < len(jsonData); i++ {
		testcase := TestCase{
			data:     InitUserRegistrationDto(jsonData[i]),
			expected: expec[i],
		}
		testdata = append(testdata, testcase)
	}

	for i := 0; i < len(jsonData); i++ {
		user, userResourcePermission := GetUserAndResourcePermissionBody(testdata[i].data)
		testdata[i].actual = struct {
			user                   User
			userResourcePermission UserResourcePermission
		}{
			user: user,
			userResourcePermission: userResourcePermission,
		}
		if !reflect.DeepEqual(testdata[i].expected, testdata[i].actual) {
			fmt.Println(testdata[i].actual)
			assert.ElementsMatch(t, testdata[i].expected, testdata[i].actual)
		}
	}
}
