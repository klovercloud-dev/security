package v1

//
//import (
//	"encoding/json"
//	"fmt"
//	"github.com/klovercloud-ci/enums"
//	_ "github.com/klovercloud-ci/enums"
//	"github.com/stretchr/testify/assert"
//	"reflect"
//	"testing"
//	"time"
//)
//
//// Initialize UserRegistrationDto from JSON data
//func InitUserRegistrationDto(str string) UserRegistrationDto {
//	res := UserRegistrationDto{}
//	err := json.Unmarshal([]byte(str), &res)
//	if err != nil {
//		return UserRegistrationDto{}
//	}
//	fmt.Println(res)
//	return res
//}
//
//// Test function for testing UserRegistrationDto to User and UserResourcePermissionDto serialization.
//func TestGetUserAndResourcePermissionBody(t *testing.T) {
//	type resultData struct {
//		user User
//		userResourcePermission UserResourcePermissionDto
//	}
//
//	type TestCase struct {
//		data     UserRegistrationDto
//		expected resultData
//		actual   resultData
//	}
//
//	var testdata []TestCase
//
//	loc, _ := time.LoadLocation("EST")
//	timeData := time.Date(2022, 1, 11, 14,9,0,0, loc).UTC()
//	fmt.Println(timeData)
//
//	// UserResourcePermissionDto test data
//	ids := []string{"123"}
//	firstNames := []string{"Shabrul"}
//	lastNames := []string{"TheBOSS"}
//	emails := []string{"shabrul_theboss@klovercloud.com"}
//	phones := []string{"01707007007"}
//	passwords := []string{"IAmTheBoss"}
//	status := []enums.STATUS{"ACTIVE"}
//	createdDates := []time.Time{timeData}
//	updatedDates := []time.Time{timeData}
//	authTypes := []enums.AUTH_TYPE{"PASSWORD"}
//	// UserResourcePermissionDto -> Resource_Permission data
//	userIds := []string{"123"}
//		// UserResourcePermissionDto -> Resource_Permission ->  Resources data
//		resourcesNames := [][]string{{"resources1"}}
//			// UserResourcePermissionDto -> Resource_Permission ->  Resources -> RoleDto data
//			rolesNames := [][]string{{"admin"}}
//				// UserResourcePermissionDto -> Resource_Permission ->  Resources -> RoleDto -> Permission data
//				permissionNames := [][]string{{"create", "read"}}
//
//	jsonData := []string{
//		`{
//			"_id": "123",
//			"first_name": "Shabrul",
//			"last_name": "TheBOSS",
//			"email": "shabrul_theboss@klovercloud.com",
//			"phone": "01707007007",
//			"password":	"IAmTheBoss",
//			"status": "ACTIVE",
//			"created_date": "2022-01-11 19:09:00 +0000 UTC",
//			"updated_date": "2022-01-11 19:09:00 +0000 UTC",
//			"auth_type": "PASSWORD",
//			"resource_permission":	{
//				"user_id": "123",
//				"resources": [
//					{
//						"name": "resources1",
//						"roles": [
//							{
//								"name": "admin",
//								"permissions": [
//									{
//										"name": "create"
//									},
//									{
//										"name": "read"
//									}
//								]
//							}
//						]
//					}
//				]
//			}
//		}`,
//	}
//
//	expec := []resultData{
//		{
//			User{
//				ID:           "123",
//				FirstName:    "Shabrul",
//				LastName:     "TheBOSS",
//				Email:        "shabrul_theboss@klovercloud.com",
//				Phone: 		  "01707007007",
//				Password:     "IAmTheBoss",
//				Status:       "ACTIVE",
//				CreatedDate:  timeData,
//				UpdatedDate:  timeData,
//				AuthType:     "PASSWORD",
//			},
//			UserResourcePermissionDto{
//				UserId:    "123",
//				Resources: []struct {
//					Name  string `json:"name" bson:"name"`
//					RoleDto []RoleDto `json:"roles" bson:"roles"`
//				}{
//					{
//						Name: "resources1",
//						RoleDto: []RoleDto{
//							{
//								Name: "admin",
//								Permissions: []Permission{
//									{
//										Name: "create",
//									},
//									{
//										Name: "read",
//									},
//								},
//							},
//						},
//					},
//				},
//			},
//		},
//	}
//
//	for i := 0; i < len(jsonData); i++ {
//
//		testcase := TestCase{
//			data:     UserRegistrationDto{
//				ID:                 ids[i],
//				FirstName:          firstNames[i],
//				LastName:           lastNames[i],
//				Email:              emails[i],
//				Phone:              phones[i],
//				Password:           passwords[i],
//				Status:             status[i],
//				CreatedDate:        createdDates[i],
//				UpdatedDate:        updatedDates[i],
//				AuthType:           authTypes[i],
//				ResourcePermission: UserResourcePermissionDto{
//					UserId:    userIds[i],
//					Resources: []struct {
//						Name  string `json:"name" bson:"name"`
//						RoleDto []RoleDto `json:"roles" bson:"roles"`
//					}{
//						{
//							Name: resourcesNames[i][0],
//							RoleDto: []RoleDto{
//								{
//									Name:        rolesNames[i][0],
//									Permissions: []Permission{
//										{
//											Name: permissionNames[i][0],
//										},
//										{
//											Name: permissionNames[i][1],
//										},
//									},
//								},
//							},
//						},
//					},
//				},
//			},
//			expected: expec[i],
//		}
//		testdata = append(testdata, testcase)
//	}
//
//	for i := 0; i < len(jsonData); i++ {
//		user, userResourcePermission := GetUserAndResourcePermissionBody(testdata[i].data)
//		testdata[i].actual = struct {
//			user                   User
//			userResourcePermission UserResourcePermissionDto
//		}{
//			user: user,
//			userResourcePermission: userResourcePermission,
//		}
//		if !reflect.DeepEqual(testdata[i].expected, testdata[i].actual) {
//			fmt.Println(testdata[i].actual)
//			assert.ElementsMatch(t, testdata[i].expected, testdata[i].actual)
//		}
//	}
//}
