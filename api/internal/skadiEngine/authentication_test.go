package skadiEngine

import (
	"github.com/gin-gonic/gin"
	"github.com/ribeirosaimon/skadi/api/internal/entity"
	"github.com/ribeirosaimon/skadi/domain/config"
	"github.com/ribeirosaimon/skadi/domain/noSql"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

const testAuthPath = "/test-auth"

func NewAuthorizationController() *config.SkadiRouter {
	testDto := struct {
		Content string
	}{Content: "test"}

	return config.NewSkadiRouter(testAuthPath).
		AddController(config.NewSkadiController(http.MethodGet, "/admin", func(ctx *gin.Context) {
			entity.HandleSuccess(ctx, 200, testDto)
		}, noSql.ADMINISTRATOR))
}

func TestAuthSkadiApi(t *testing.T) {
	go StartSkadiApi("test")
	RegisterRouter([]*config.SkadiRouter{NewAuthorizationController()})

	server := httptest.NewServer(skadi)
	defer server.Close()
	resp, err := http.Get(server.URL + testAuthPath + "/admin")
	defer resp.Body.Close()

	assert.Nil(t, err)
	assert.Equal(t, http.StatusUnauthorized, resp.StatusCode)
}
