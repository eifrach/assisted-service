package requestid

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-openapi/swag"
	"github.com/google/uuid"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type mockTransport struct {
	mock.Mock
}

func (m *mockTransport) RoundTrip(r *http.Request) (*http.Response, error) {
	m.Called(r)
	return nil, nil
}

// This is an old package test. While all of our testing infrastructure was switched to use ginkgo -
// this test remains until it would get converted.
// This ginkgo wrapper was added to allow running this packge with ginkgo flags.
func TestRequestID(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Request id tests")
}

func TestTransport(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name    string
		prepare func(t *testing.T, tr *mockTransport) *http.Request
	}{
		{
			name: "happy flow",
			prepare: func(t *testing.T, tr *mockTransport) *http.Request {
				const requestID = "1234"

				match := mock.MatchedBy(func(req *http.Request) bool {
					return req.Header.Get(headerKey) == requestID
				})

				tr.On("RoundTrip", match).Return(nil, nil).Once()

				ctx := context.WithValue(context.Background(), ctxKey, requestID)
				req := httptest.NewRequest(http.MethodGet, "http://example.org", nil)
				req = req.WithContext(ctx)
				return req
			},
		},
		{
			name: "no request id in context",
			prepare: func(t *testing.T, tr *mockTransport) *http.Request {
				match := mock.MatchedBy(func(req *http.Request) bool {
					return req.Header.Get(headerKey) == ""
				})

				tr.On("RoundTrip", match).Return(nil, nil).Once()

				req := httptest.NewRequest(http.MethodGet, "http://example.org", nil)
				return req
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			var tr mockTransport
			defer tr.AssertExpectations(t)
			req := tt.prepare(t, &tr)
			_, _ = Transport(&tr).RoundTrip(req)
		})
	}
}

func TestMiddleware(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name      string
		requestID *string
	}{
		{
			name:      "RequestID exist",
			requestID: swag.String("1234"),
		},
		{
			name:      "no request  in context",
			requestID: nil,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			// create a handler to use as "next" which will verify the request
			nextHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				val := r.Context().Value(ctxKey)
				assert.NotNil(t, val)
				valStr, ok := val.(string)
				if !ok {
					t.Error("not string")
				}

				if tt.requestID != nil {
					// if request-id passed in header, that should be in the context
					assert.Equal(t, valStr, *tt.requestID)
				} else {
					// if no request-id passed in header, valid uuid should be generated by middleware
					assert.True(t, IsValidUUID(valStr))
				}

			})

			// create the handler to test, using our custom "next" handler
			h := Middleware(nextHandler)

			// create a mock request to use
			req := httptest.NewRequest("GET", "http://testing", nil)
			if tt.requestID != nil {
				req.Header.Set(headerKey, *tt.requestID)
			}

			// call the handler using a mock response recorder (we'll not use that anyway)
			h.ServeHTTP(httptest.NewRecorder(), req)
		})
	}
}

func IsValidUUID(u string) bool {
	_, err := uuid.Parse(u)
	return err == nil
}