package http

import (
	"app/api"
	nhttp "net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUserGetInfoHandler(t *testing.T) {
	retrieveUserInfo := func() func(*testing.T) {
		userID := "1"
		return func(t *testing.T) {
			api := &_userApiMock{}
			handler := NewUserGetInfoHandler(api).Handler

			req, err := nhttp.NewRequest("GET", "url is not important", nil)
			if err != nil {
				t.Fatal(err)
			}
			q := req.URL.Query()
			q.Add("userID", userID)
			req.URL.RawQuery = q.Encode()
			writer := httptest.NewRecorder()

			handler(writer, req)
			assert.Equal(t, nhttp.StatusOK, writer.Code)
		}
	}

	retrieveUserInfoNotFound := func() func(*testing.T) {
		userID := ""
		return func(t *testing.T) {
			api := &_userApiMock{}
			handler := NewUserGetInfoHandler(api).Handler

			req, err := nhttp.NewRequest("GET", "url is not important", nil)
			if err != nil {
				t.Fatal(err)
			}
			q := req.URL.Query()
			q.Add("userID", userID)
			req.URL.RawQuery = q.Encode()
			writer := httptest.NewRecorder()

			handler(writer, req)
			assert.Equal(t, nhttp.StatusNotFound, writer.Code)
		}
	}

	t.Run("200 ok", retrieveUserInfo())
	t.Run("404 not found", retrieveUserInfoNotFound())
}

type _userApiMock struct {
}

func (m *_userApiMock) GetInfo(userID int) (api.UserInfo, error) {
	return api.UserInfo{}, nil
}
