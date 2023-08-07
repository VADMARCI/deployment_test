package presenters

import (
	"github.com/labstack/echo/v4"
	"github.com/pepusz/go_redirect/models"
)

type Context struct {
	echo.Context
}

func (c *Context) GetMainFactory() *MainFactory {
	return c.Get("mainFactory").(*MainFactory)
}

func (c *Context) SetMainFactory(factory *MainFactory) {
	c.Set("mainFactory", factory)
}

func (c *Context) SetCurrentUser(user *models.User) {
	c.Set("currentUser", user)
}
func (c *Context) SetClientId(clientID string) {
	c.Set("clientID", clientID)
}

func (c *Context) GetCurrentUser() *models.User {
	currentUser := c.Get("currentUser")
	if currentUser != nil {
		return currentUser.(*models.User)
	}
	return nil
}
func (c *Context) GetClientID() string {
	clientID := c.Get("clientID")
	if clientID != nil {
		return clientID.(string)
	}
	return ""
}
func (c *Context) Authorize() bool {
	currentUser := c.GetCurrentUser()

	if currentUser == nil {

		return false
	}
	// if !currentUser.Approved {
	// 	mainFactory.Gateways.ErrorGateway.CreateError("Unapproved", "0-2")
	// 	return false
	// }
	return true
}
