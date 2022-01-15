package actions

import (
	"encoding/json"
	"net/http"
	"strings"
	"team_manager/models"

	"github.com/gobuffalo/pop/slices"
)

func (as *ActionSuite) Test_MembersResource_List() {
	as.LoadFixture("employees")
	as.LoadFixture("contractors")

	res := as.JSON("/members").Get()
	as.Equal(http.StatusOK, res.Code)

	members := models.Members{}
	err := json.Unmarshal(res.Body.Bytes(), &members)
	as.NoError(err)
	as.Equal(6, len(members))
}

func (as *ActionSuite) Test_MembersResource_Show() {
	as.LoadFixture("employees")
	as.LoadFixture("contractors")

	target := &models.Member{}
	err := as.DB.First(target)
	as.NoError(err)

	res := as.JSON("/members/" + target.ID.String()).Get()
	as.Equal(http.StatusOK, res.Code)

	member := models.Member{}
	err = json.Unmarshal(res.Body.Bytes(), &member)
	as.NoError(err)
	as.Equal(target.Name, member.Name)
	as.Equal(target.Type, member.Type)
	as.Equal(len(target.Tags), len(member.Tags))
}

func (as *ActionSuite) Test_MembersResource_Create_Employee() {
	m := &models.Member{
		Name: "Member Name",
		Type: "employee",
		Role: "DevOps",
		Tags: slices.String{"GOLANG", "Kubernetes"},
	}
	res := as.JSON("/members").Post(m)
	as.Equal(http.StatusCreated, res.Code)

	employee := models.Member{}
	err := json.Unmarshal(res.Body.Bytes(), &employee)
	as.NoError(err)
	as.Equal(m.Name, employee.Name)
	as.Equal(m.Type, employee.Type)
	as.Equal(len(m.Tags), len(employee.Tags))
	for _, v := range employee.Tags {
		as.Equal(strings.ToLower(v), v)
	}
}

func (as *ActionSuite) Test_MembersResource_Create_EmployeeWithoutRole() {
	m := &models.Member{
		Name: "Member Name",
		Type: "employee",
		Tags: slices.String{"GOLANG", "Kubernetes"},
	}
	res := as.JSON("/members").Post(m)
	as.Equal(http.StatusUnprocessableEntity, res.Code)
}

func (as *ActionSuite) Test_MembersResource_Create_Contractor() {
	m := &models.Member{
		Name:             "Member Name",
		Type:             "contractor",
		ContractDuration: 500,
	}
	res := as.JSON("/members").Post(m)
	as.Equal(http.StatusCreated, res.Code)

	contractor := models.Member{}
	err := json.Unmarshal(res.Body.Bytes(), &contractor)
	as.NoError(err)
	as.Equal(m.Name, contractor.Name)
	as.Equal(m.Type, contractor.Type)
	as.Equal(len(m.Tags), len(contractor.Tags))
}

func (as *ActionSuite) Test_MembersResource_Create_Contractor_WithoutContractDuration() {
	m := &models.Member{
		Name: "Member Name",
		Type: "contractor",
	}
	res := as.JSON("/members").Post(m)
	as.Equal(http.StatusUnprocessableEntity, res.Code)
}

func (as *ActionSuite) Test_MembersResource_Create_InvalidType() {
	m := &models.Member{
		Name: "Member Name",
		Type: "invalid",
	}
	res := as.JSON("/members").Post(m)
	as.Equal(http.StatusUnprocessableEntity, res.Code)
}

func (as *ActionSuite) Test_MembersResource_Create_WithoutName() {
	m := &models.Member{
		Type: "contractor",
		Tags: slices.String{"php", "laravel"},
	}
	res := as.JSON("/members").Post(m)
	as.Equal(http.StatusUnprocessableEntity, res.Code)
}

func (as *ActionSuite) Test_MembersResource_Update() {
	as.LoadFixture("employees")
	as.LoadFixture("contractors")

	target := &models.Member{}
	err := as.DB.First(target)
	as.NoError(err)

	target.Name = "New Name"
	target.Tags = slices.String{"JS", "NoDeJs"}

	res := as.JSON("/members/" + target.ID.String()).Put(target)
	as.Equal(http.StatusOK, res.Code)

	member := models.Member{}
	err = json.Unmarshal(res.Body.Bytes(), &member)
	as.NoError(err)
	as.Equal(target.Name, member.Name)
	as.Equal(target.Type, member.Type)
	as.Equal(len(target.Tags), len(member.Tags))
	for _, v := range member.Tags {
		as.Equal(strings.ToLower(v), v)
	}
}

func (as *ActionSuite) Test_MembersResource_Destroy() {
	as.LoadFixture("employees")
	as.LoadFixture("contractors")

	target := &models.Member{}
	err := as.DB.First(target)
	as.NoError(err)

	res := as.JSON("/members/" + target.ID.String()).Delete()
	as.Equal(http.StatusNoContent, res.Code)
	as.Error(as.DB.Find(&models.Member{}, target.ID))
}
