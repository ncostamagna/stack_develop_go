package oauth

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"net/http"
	"fmt"
	"os"
	"github.com/mercadolibre/golang-restclient/rest"
)

// Testing Main
func TestMain(m *testing.M) {
	fmt.Println("about to start oauth tests")

	// mock request
	// resultados del mock
	rest.StartMockupServer()

	os.Exit(m.Run())
}

// Testeamos las constantess
func TestOauthConstants(t *testing.T) {
	assert.EqualValues(t, "X-Public", headerXPublic)
	assert.EqualValues(t, "X-Client-Id", headerXClientId)
	assert.EqualValues(t, "X-Caller-Id", headerXCallerId)
	assert.EqualValues(t, "access_token", paramAccessToken)
}

func TestIsPublicNilRequest(t *testing.T) {
	assert.True(t, IsPublic(nil))
}

func TestIsPublicNoError(t *testing.T) {
	request := http.Request{
		Header: make(http.Header),
	}
	assert.False(t, IsPublic(&request))

	request.Header.Add("X-Public", "true")
	assert.True(t, IsPublic(&request))
}

func TestGetCallerIdNilRequest(t *testing.T) {
	//TODO: Complete!
}

func TestGetCallerInvalidCallerFormat(t *testing.T) {
	//TODO: Complete!
}

func TestGetCallerNoError(t *testing.T) {
	//TODO: Complete!
}

func TestGetAccessTokenInvalidRestclientResponse(t *testing.T) {

	// limpiamos todos los muchs que tenemos
	// lo primero que debemos hacer
	rest.FlushMockups()

	// donde vamos a disparar el request
	// esto es para el rest de autentificacion de mercado libre
	rest.AddMockups(&rest.Mock{
		HTTPMethod:   http.MethodGet,
		URL:          "http://localhost:8080/oauth/access_token/AbC123",
		ReqBody:      ``,
		RespHTTPCode: -1,
		RespBody:     `{}`,
	})

	accessToken, err := getAccessToken("AbC123")
	assert.Nil(t, accessToken)
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusInternalServerError, err.Status)
	assert.EqualValues(t, "invalid restclient response when trying to get access token", err.Message)
}

//TODO: Add complete coverage for the getAccessToken function.
