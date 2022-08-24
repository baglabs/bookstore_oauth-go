package oauth

import (
	"fmt"
	"net/http"
	"os"
	"testing"

	"github.com/baglabs/bookstore_oauth-go/oauth/github.com/mercadolibre/golang-restclient/rest"
	"github.com/stretchr/testify/assert"
)

func TestMain(m *testing.M) {
	fmt.Println("about to start oauth test")

	rest.StartMockupServer()

	os.Exit(m.Run())
}

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

	request.Header.Add(headerXPublic, "true")
	assert.True(t, IsPublic(&request))
}

func TestCallerIdNilRequest(t *testing.T) {
	//TODO: Complete!
}

func TestCallerIdInvalidCallerFormat(t *testing.T) {
	//TODO: Complete!
}

func TestCallerIdNoError(t *testing.T) {
	//TODO: Complete!
}

func TestGetAccessTokenInvalidRestClientResponse(t *testing.T) {
	rest.FlushMockups()
	rest.AddMockups(&rest.Mock{
		HTTPMethod:   http.MethodGet,
		URL:          "http://localhost:8080/oauth/access_token/abc123",
		ReqBody:      "",
		RespHTTPCode: -1,
		RespBody:     "{}",
	})

	accessToken, err := getAccessToken("abc123")
	assert.Nil(t, accessToken)
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusInternalServerError, err.Status())
	assert.EqualValues(t, "invalid restclient response when trying to get access token", err.Message())
}

//TODO: Add complete coverage for the getAccessToken function
