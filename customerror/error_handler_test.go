package customerror

import (
	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
	"net/http"
	"routines/commons/utils"
	"testing"
)

func TestCustomErrorHandler(t *testing.T) {
	req, err := http.NewRequest("GET", "/", nil)
	assert.NoError(t, err)
	assert.NotNil(t, req)

	appWithErrResp := func(err error) *fiber.App {
		app := fiber.New(fiber.Config{ErrorHandler: CustomErrorHandler})
		app.Get("/", func(ctx *fiber.Ctx) error {
			return err
		})
		return app
	}

	t.Run("should return status as 400", func(t *testing.T) {
		expectedData := "Bad request err message"
		app := appWithErrResp(BadRequest(expectedData))
		resp, err := app.Test(req, -1)
		assert.NoError(t, err)
		assert.Equal(t, 400, resp.StatusCode)
		actual := make(map[string]any)
		assert.NoError(t, utils.ReadInto(resp.Body, &actual))
		assert.Equal(t, expectedData, actual["errorMessage"])
	})

	t.Run("should return status as 405 from fiber", func(t *testing.T) {
		expectedData := "method not allowed"
		app := appWithErrResp(&fiber.Error{
			Code:    http.StatusMethodNotAllowed,
			Message: expectedData,
		})
		resp, err := app.Test(req, -1)
		assert.NoError(t, err)
		assert.Equal(t, 405, resp.StatusCode)
		actual := make(map[string]any)
		assert.NoError(t, utils.ReadInto(resp.Body, &actual))
		assert.Equal(t, expectedData, actual["errorMessage"])
	})

	t.Run("should return status as 500 from routine customerror", func(t *testing.T) {
		expectedData := "internal server customerror"
		app := appWithErrResp(InternalError(expectedData))
		resp, err := app.Test(req, -1)
		assert.NoError(t, err)
		assert.Equal(t, 500, resp.StatusCode)
		actual := make(map[string]any)
		assert.NoError(t, utils.ReadInto(resp.Body, &actual))
		assert.Equal(t, expectedData, actual["errorMessage"])
	})
}
