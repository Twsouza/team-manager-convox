package actions

import (
	"fmt"
	"log"
	"net/http"
	"team_manager/models"

	"github.com/gobuffalo/buffalo"
	"github.com/gobuffalo/pop/v5"
	"github.com/gobuffalo/x/responder"
)

// MembersResource is the resource for the Member model (CRUD)
type MembersResource struct {
	buffalo.Resource
}

// List gets all Members.
// @Summary List members
// @ID list-members
// @Param page query integer false "Go to the page"
// @Param per_page query integer false "How many member per pages"
// @Produce json,xml
// @Success 200 {object} models.Members
// @Failure 500
// @Router /members [get]
func (v MembersResource) List(c buffalo.Context) error {
	// Get the DB connection from the context
	tx, ok := c.Value("tx").(*pop.Connection)
	if !ok {
		return fmt.Errorf("no transaction found")
	}

	members := models.Members{}

	// Paginate results. Params "page" and "per_page" control pagination.
	// Default values are "page=1" and "per_page=20".
	q := tx.PaginateFromParams(c.Params())

	// Retrieve all Members from the DB
	if err := q.All(&members); err != nil {
		return err
	}

	return responder.Wants("json", func(c buffalo.Context) error {
		return c.Render(200, r.JSON(members))
	}).Wants("xml", func(c buffalo.Context) error {
		return c.Render(200, r.XML(members))
	}).Respond(c)
}

// Show gets the data for one Member.
// @Summary Show a member
// @ID show-member
// @Produce json,xml
// @Param member_id path string true "Member ID"
// @Success 200 {object} models.Members
// @Failure 404,500
// @Router /members/{member_id} [get]
func (v MembersResource) Show(c buffalo.Context) error {
	// Get the DB connection from the context
	tx, ok := c.Value("tx").(*pop.Connection)
	if !ok {
		return fmt.Errorf("no transaction found")
	}

	// Allocate an empty Member
	member := &models.Member{}

	// To find the Member the parameter member_id is used.
	if err := tx.Find(member, c.Param("member_id")); err != nil {
		return c.Error(http.StatusNotFound, err)
	}

	return responder.Wants("json", func(c buffalo.Context) error {
		return c.Render(200, r.JSON(member))
	}).Wants("xml", func(c buffalo.Context) error {
		return c.Render(200, r.XML(member))
	}).Respond(c)
}

// Create adds a Member to the DB.
// @Summary Create a new member
// @Description Create a new member, employee only accepts role, contractor only accepts contract_duration
// @ID create-member
// @Accept json,xml
// @Produce json,xml
// @Param member body models.Member true "Member Payload"
// @Success 201 {object} models.Members
// @Failure 422 {object} validate.Errors
// @Failure 500
// @Router /members [post]
func (v MembersResource) Create(c buffalo.Context) error {
	// Allocate an empty Member
	member := &models.Member{}

	// Bind member to the request payload
	if err := c.Bind(member); err != nil {
		return err
	}

	// Get the DB connection from the context
	tx, ok := c.Value("tx").(*pop.Connection)
	if !ok {
		return fmt.Errorf("no transaction found")
	}

	log.Printf("%+v", member)

	// Validate the data from the request
	verrs, err := tx.ValidateAndCreate(member)
	if err != nil {
		return err
	}

	if verrs.HasAny() {
		return responder.Wants("json", func(c buffalo.Context) error {
			return c.Render(http.StatusUnprocessableEntity, r.JSON(verrs))
		}).Wants("xml", func(c buffalo.Context) error {
			return c.Render(http.StatusUnprocessableEntity, r.XML(verrs))
		}).Respond(c)
	}

	return responder.Wants("json", func(c buffalo.Context) error {
		return c.Render(http.StatusCreated, r.JSON(member))
	}).Wants("xml", func(c buffalo.Context) error {
		return c.Render(http.StatusCreated, r.XML(member))
	}).Respond(c)
}

// Update changes a Member in the DB.
// @Summary Update a member
// @ID update-member
// @Accept json,xml
// @Produce json,xml
// @Param member_id path string true "Member ID"
// @Success 200 {object} models.Members
// @Failure 422 {object} validate.Errors
// @Failure 500
// @Router /members/{member_id} [put]
func (v MembersResource) Update(c buffalo.Context) error {
	// Get the DB connection from the context
	tx, ok := c.Value("tx").(*pop.Connection)
	if !ok {
		return fmt.Errorf("no transaction found")
	}

	// Allocate an empty Member
	member := &models.Member{}

	if err := tx.Find(member, c.Param("member_id")); err != nil {
		return c.Error(http.StatusNotFound, err)
	}

	// Bind Member to the request payload
	if err := c.Bind(member); err != nil {
		return err
	}

	verrs, err := tx.ValidateAndUpdate(member)
	if err != nil {
		return err
	}

	if verrs.HasAny() {
		return responder.Wants("json", func(c buffalo.Context) error {
			return c.Render(http.StatusUnprocessableEntity, r.JSON(verrs))
		}).Wants("xml", func(c buffalo.Context) error {
			return c.Render(http.StatusUnprocessableEntity, r.XML(verrs))
		}).Respond(c)
	}

	return responder.Wants("json", func(c buffalo.Context) error {
		return c.Render(http.StatusOK, r.JSON(member))
	}).Wants("xml", func(c buffalo.Context) error {
		return c.Render(http.StatusOK, r.XML(member))
	}).Respond(c)
}

// Destroy deletes a Member from the DB.
// @Summary Delete a member
// @ID delete-member
// @Accept json,xml
// @Produce json,xml
// @Param member_id path string true "Member ID"
// @Success 200 {object} models.Members
// @Failure 404,500
// @Router /members/{member_id} [delete]
func (v MembersResource) Destroy(c buffalo.Context) error {
	// Get the DB connection from the context
	tx, ok := c.Value("tx").(*pop.Connection)
	if !ok {
		return fmt.Errorf("no transaction found")
	}

	// Allocate an empty Member
	member := &models.Member{}

	// To find the Member the parameter member_id is used.
	if err := tx.Find(member, c.Param("member_id")); err != nil {
		return c.Error(http.StatusNotFound, err)
	}

	if err := tx.Destroy(member); err != nil {
		return err
	}

	return responder.Wants("json", func(c buffalo.Context) error {
		return c.Render(http.StatusNoContent, nil)
	}).Wants("xml", func(c buffalo.Context) error {
		return c.Render(http.StatusNoContent, nil)
	}).Respond(c)
}
