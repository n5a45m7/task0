package http

import (
	"app"
	"app/api"
	"bytes"
	"encoding/json"
	nhttp "net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAccountCreateHandler(t *testing.T) {
	createAccount := func() func(*testing.T) {
		return func(t *testing.T) {
			api := &_accApiMock{}
			handler := NewAccountCreateHandler(api).Handler

			data := accCreateRequest{
				CustomerID:    1,
				InitialCredit: 100.0,
			}
			b, err := json.Marshal(&data)
			if err != nil {
				t.Fatal(err)
			}
			req, err := nhttp.NewRequest("POST", "url is not important", bytes.NewReader(b))
			if err != nil {
				t.Fatal(err)
			}
			writer := httptest.NewRecorder()

			handler(writer, req)
			assert.Equal(t, nhttp.StatusOK, writer.Code)
		}
	}

	t.Run("200 ok", createAccount())
}

type _accApiMock struct {
}

func (m *_accApiMock) Create(request api.AccountCreateRequest) (app.Account, error) {
	return app.Account{}, nil
}
