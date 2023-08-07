package presenters

import (
	"context"

	"github.com/99designs/gqlgen/graphql"
	"github.com/labstack/echo/v4"
	core_models "github.com/pepusz/go_redirect/models"
	"github.com/vektah/gqlparser/v2/gqlerror"
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

func (c *Context) SetCurrentUser(user *core_models.User) {
	c.Set("currentUser", user)
}

func (c *Context) GetCurrentUser() *core_models.User {
	currentUser := c.Get("currentUser")
	if currentUser != nil {
		return currentUser.(*core_models.User)
	}
	return nil
}

func IsAuthenticated(ctx context.Context) bool {
	c := ctx.Value("echoContext").(*Context)
	currentUser := c.GetCurrentUser()

	if currentUser == nil {
		graphql.AddError(ctx, &gqlerror.Error{
			Message: "Unauthenticated",
			Extensions: map[string]interface{}{
				"code":             "1.3.1", // SessionUnauthenticatedCode
				"errorCodeMessage": "Unauthenticated",
			},
		})
		return false
	}
	return true
}
