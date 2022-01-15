package models

import (
	"strings"

	"github.com/gobuffalo/pop/slices"
)

func (ms *ModelSuite) Test_Member_Employee() {
	m := &Member{
		Name: "Member Name",
		Type: "employee",
		Role: "Software Engineer",
		Tags: slices.String{"golang", "kubernetes"},
	}

	verrs, err := DB.ValidateAndCreate(m)
	ms.NoError(err)
	ms.Equal(false, verrs.HasAny())
}

func (ms *ModelSuite) Test_Member_EmployeeWithoutRole() {
	m := &Member{
		Name: "Member Name",
		Type: "employee",
		Tags: slices.String{"golang", "kubernetes"},
	}

	verrs, err := DB.ValidateAndCreate(m)
	ms.NoError(err)
	ms.Equal(true, verrs.HasAny())
}

func (ms *ModelSuite) Test_Member_Contractor() {
	m := &Member{
		Name:             "Member Name",
		Type:             "contractor",
		ContractDuration: 150,
		Tags:             slices.String{"golang", "kubernetes"},
	}

	verrs, err := DB.ValidateAndCreate(m)
	ms.NoError(err)
	ms.Equal(false, verrs.HasAny())
}

func (ms *ModelSuite) Test_Member_ContractorWithoutContractDuration() {
	m := &Member{
		Name:             "Member Name",
		Type:             "contractor",
		ContractDuration: 150,
		Tags:             slices.String{"golang", "kubernetes"},
	}

	verrs, err := DB.ValidateAndCreate(m)
	ms.NoError(err)
	ms.Equal(false, verrs.HasAny())
}

func (ms *ModelSuite) Test_Member_IsTagsLowerCase() {
	m := &Member{
		Name:             "Member Name",
		Type:             "contractor",
		ContractDuration: 150,
		Tags:             slices.String{"GOLANG", "DocKEr", "kubernetes"},
	}

	verrs, err := DB.ValidateAndCreate(m)
	ms.NoError(err)
	ms.Equal(false, verrs.HasAny())

	DB.Reload(m)
	for _, v := range m.Tags {
		ms.Equal(strings.ToLower(v), v)
	}
}

func (ms *ModelSuite) Test_Member_WithoutName() {
	m := &Member{
		Type: "employee",
		Tags: slices.String{"golang", "kubernetes"},
	}

	verrs, err := DB.ValidateAndCreate(m)
	ms.NoError(err)
	ms.Equal(true, verrs.HasAny())
	ms.Contains(verrs.Error(), "Name can not be blank.")
}

func (ms *ModelSuite) Test_Member_WithoutType() {
	m := &Member{
		Name: "Member Name",
		Tags: slices.String{"golang", "kubernetes"},
	}

	verrs, err := DB.ValidateAndCreate(m)
	ms.NoError(err)
	ms.Equal(true, verrs.HasAny())
	ms.Contains(verrs.Error(), memberTypeInvalid)
}

func (ms *ModelSuite) Test_Member_WithoutTags() {
	m := &Member{
		Name: "Member Name",
		Type: "employee",
		Role: "Software Engineer",
	}

	verrs, err := DB.ValidateAndCreate(m)
	ms.NoError(err)
	ms.Equal(false, verrs.HasAny())
}
