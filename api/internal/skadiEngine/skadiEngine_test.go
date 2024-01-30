package skadiEngine

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/ribeirosaimon/skadi/api/internal/entity"
	"github.com/ribeirosaimon/skadi/domain/config"
	"github.com/stretchr/testify/assert"
)

const testPath = "/test"

func NewTestController() *config.SkadiRouter {

	testDto := struct {
		Content string
	}{Content: "test"}

	return config.NewSkadiRouter(testPath).
		AddController(config.NewSkadiController(http.MethodGet, "/controller-test", func(ctx *gin.Context) {
			entity.HandleSuccess(ctx, 200, testDto)
		}))
}

func TestStartSkadiApi(t *testing.T) {
	go StartSkadiApi("test")
	RegisterRouter([]*config.SkadiRouter{NewTestController()})
	t.Run("Up Engine", func(t *testing.T) {
		server := httptest.NewServer(skadi)
		defer server.Close()
		resp, err := http.Get(server.URL + "/test/controller-test")
		defer resp.Body.Close()

		assert.Nil(t, err)
		assert.Equal(t, http.StatusOK, resp.StatusCode)
	})
}
