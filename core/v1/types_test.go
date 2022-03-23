package v1

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/stretchr/testify/assert"
	"reflect"
	"testing"
)

// ROLE TEST
func InitRole(str string) Role {
	res := Role{}
	err := json.Unmarshal([]byte(str), &res)
	if err != nil {
		return Role{}
	}
	return res
}

func TestRole_Validate(t *testing.T) {
	type TestCase struct {
		data     Role
		expected error
		actual   error
	}
	var testdata []TestCase

	jsonData := []string{
		`{
			"name": "ADMIN"	
		}`,
		`{
			"name": ""	
		}`,
	}
	expectedData := []error{nil, errors.New("[ERROR]: Blank role name")}

	for i := range jsonData {
		testcase := TestCase{
			data:     InitRole(jsonData[i]),
			expected: expectedData[i],
		}
		testdata = append(testdata, testcase)
	}

	for i := range jsonData {
		testdata[i].actual = testdata[i].data.Validate()
		if !reflect.DeepEqual(testdata[i].expected, testdata[i].actual) {
			fmt.Println(testdata[i].actual)
			assert.ElementsMatch(t, testdata[i].expected, testdata[i].actual)
		}
	}
}

// PERMISSION TEST
func InitPermission(str string) Permission {
	res := Permission{}
	err := json.Unmarshal([]byte(str), &res)
	if err != nil {
		return Permission{}
	}
	return res
}

func TestPermission_Validate(t *testing.T) {

	type TestCase struct {
		data     Permission
		expected error
		actual   error
	}
	var testdata []TestCase
	jsonData := []string{
		`{
			"name": "CREATE"	
		}`,
		`{
			"name": ""	
		}`,
		`{
			"name": "abc"	
		}`,
	}

	expectedData := []error{nil, errors.New("[ERROR]: Blank permission name"), errors.New("[ERROR]: Invalid permission name")}

	for i := range jsonData {
		testcase := TestCase{
			data:     InitPermission(jsonData[i]),
			expected: expectedData[i],
		}
		testdata = append(testdata, testcase)
	}
	for i := range jsonData {
		testdata[i].actual = testdata[i].data.Validate()
		if !reflect.DeepEqual(testdata[i].expected, testdata[i].actual) {
			fmt.Println(testdata[i].actual)
			assert.ElementsMatch(t, testdata[i].expected, testdata[i].actual)
		}
	}
}

// UserRegistrationDto test
func InitUserRegistrationDto(str string) UserRegistrationDto {
	res := UserRegistrationDto{}
	err := json.Unmarshal([]byte(str), &res)
	if err != nil {
		return UserRegistrationDto{}
	}
	return res
}

func TestUserRegistrationDto_Validate(t *testing.T) {
	type TestCase struct {
		data     UserRegistrationDto
		expected error
		actual   error
	}
	var testdata []TestCase
	jsonData := []string{
		`{
			"id": "1",
			"first_name": "steve",
			"last_name": "jobs",
			"email": "stevejobs@gmail.com",
			"auth_type": "password",
			"resources": [
				{
				  "name": "pipeline",
					  "roles": [
						{
						  "name": "ADMIN"
						}
					  ]
				}
			  ]
		}`,
		`{
			"id": "",
			"first_name": "steve",
			"last_name": "jobs",
			"email": "stevejobs@gmail.com",
			"auth_type": "password",
			"resources": [
				{
				  "name": "pipeline",
					  "roles": [
						{
						  "name": "ADMIN"
						}
					  ]
				}
			  ]
		}`,
		`{
			"id": "1",
			"first_name": "",
			"last_name": "jobs",
			"email": "stevejobs@gmail.com",
			"auth_type": "password",
			"resources": [
				{
				  "name": "pipeline",
					  "roles": [
						{
						  "name": "ADMIN"
						}
					  ]
				}
			  ]
		}`,
		`{
			"id": "1",
			"first_name": "steve",
			"last_name": "",
			"email": "stevejobs@gmail.com",
			"auth_type": "password",
			"resources": [
				{
				  "name": "pipeline",
					  "roles": [
						{
						  "name": "ADMIN"
						}
					  ]
				}
			  ]
		}`,
		`{
			"id": "1",
			"first_name": "steve",
			"last_name": "jobs",
			"email": "",
			"auth_type": "password",
			"resources": [
				{
				  "name": "pipeline",
					  "roles": [
						{
						  "name": "ADMIN"
						}
					  ]
				}
			  ]
		}`,
		`{
			"id": "1",
			"first_name": "steve",
			"last_name": "jobs",
			"email": "stevejobs@gmail.com",
			"auth_type": "abc",
			"resources": [
				{
				  "name": "pipeline",
					  "roles": [
						{
						  "name": "ADMIN"
						}
					  ]
				}
			  ]
		}`,
		`{
			"id": "1",
			"first_name": "steve",
			"last_name": "jobs",
			"email": "stevejobs@gmail.com",
			"auth_type": "password",
			"resource_permission": {
				"resources": [
					{
					  "name": "",
						  "roles": [
							{
							  "name": "ADMIN"
							}
						  ]
					}
			]
			}
		}`,
		`{
			"id": "1",
			"first_name": "steve",
			"last_name": "jobs",
			"email": "stevejobs@gmail.com",
			"auth_type": "password",
			"resource_permission": {
				"resources": [
					{
					  "name": "pipeline",
						  "roles": [
							{
							  "name": ""
							}
						  ]
					}
			]
			}
		}`,
	}

	expectedData := []error{nil, errors.New("user id is required"), errors.New("first name is required"), errors.New("last name is required"), errors.New("email is required"), errors.New("invalid user AuthType"), errors.New("[ERROR]: Blank resource name"), errors.New("[ERROR]: Blank role name")}

	for i := range jsonData {
		testcase := TestCase{
			data:     InitUserRegistrationDto(jsonData[i]),
			expected: expectedData[i],
		}
		testdata = append(testdata, testcase)

	}

	for i := range jsonData {
		testdata[i].actual = testdata[i].data.Validate()
		if !reflect.DeepEqual(testdata[i].expected, testdata[i].actual) {
			fmt.Println(testdata[i].actual)
			assert.ElementsMatch(t, testdata[i].expected, testdata[i].actual)
		}
	}
}

//UserResourcePermission Test
func InitUserResourcePermission(str string) UserResourcePermission {
	res := UserResourcePermission{}
	err := json.Unmarshal([]byte(str), &res)
	if err != nil {
		return UserResourcePermission{}
	}
	return res
}
func TestUserResourcePermission_Validate(t *testing.T) {
	type TestCase struct {
		data     UserResourcePermission
		expected error
		actual   error
	}
	var testdata []TestCase

	jsonData := []string{
		`{
			  "resources": [
				{
				  "name": "pipeline",
					  "roles": [
						{
						  "name": "ADMIN"
						}
					  ]
				}
			  ]
		}`,
		`{
			  "resources": [
				{
				  "name": "",
					  "roles": [
						{
						  "name": "ADMIN"
						}
					  ]
				}
			  ]
		}`,
		`{
			  "resources": [
				{
				  "name": "pipeline",
					  "roles": [
						{
						  "name": ""
						}
					  ]
				}
			  ]
		}`,
	}
	expectedData := []error{nil, errors.New("[ERROR]: Blank resource name"), errors.New("[ERROR]: Blank role name")}

	for i := range jsonData {
		testcase := TestCase{
			data:     InitUserResourcePermission(jsonData[i]),
			expected: expectedData[i],
		}
		testdata = append(testdata, testcase)
	}

	for i := range jsonData {
		testdata[i].actual = testdata[i].data.Validate()
		if !reflect.DeepEqual(testdata[i].expected, testdata[i].actual) {
			fmt.Println(testdata[i].actual)
			assert.ElementsMatch(t, testdata[i].expected, testdata[i].actual)
		}
	}
}
