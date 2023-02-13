package clientcontroller

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"

	"github.com/gorilla/mux"
)

type mockService struct{}

func (s *mockService) FindByID(ctx context.Context, id int) (*Client, error) {
	if id == 2 {
		return &Client{ID: 2, Name: "John Doe"}, nil
	}
	if id == 3 {
		return nil, errors.New("sql: no rows in result set")
	}
	return nil, errors.New("error finding client")
}

func TestHandleClientGet(t *testing.T) {
	tests := []struct {
		name         string
		rawCid       string
		expectedCode int
		expectedErr  string
	}{
		{
			name:         "Client ID less than 2",
			rawCid:       "1",
			expectedCode: http.StatusBadRequest,
			expectedErr:  "Bad client id",
		},
		{
			name:         "Invalid client ID",
			rawCid:       "invalid",
			expectedCode: http.StatusBadRequest,
			expectedErr:  "strconv.Atoi: parsing \"invalid\": invalid syntax",
		},
		{
			name:         "Client not found",
			rawCid:       "3",
			expectedCode: http.StatusNotFound,
			expectedErr:  "sql: no rows in result set",
		},
		{
			name:         "Error finding client",
			rawCid:       "4",
			expectedCode: http.StatusInternalServerError,
			expectedErr:  "error finding client",
		},
		{
			name:         "Successful request",
			rawCid:       "2",
			expectedCode: http.StatusOK,
			expectedErr:  "",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			req, err := http.NewRequest("GET", "/clients/"+test.rawCid, nil)
			if err != nil {
				t.Fatal(err)
			}

			rr := httptest.NewRecorder()
			router := mux.NewRouter()
			router.HandleFunc("/clients/{id}", (*ClientController).handleClientGet)

			router.ServeHTTP(rr, req)

			if status := rr.Code; status != test.expectedCode {
				t.Errorf("handler returned wrong status code: got %v want %v", status, test.expectedCode)
			}

			if rr.Code == http.
