package logic_test

import (
	"fmt"
	"testing"

	"github.com/matt-hoiland/mocking/matching/logic"
	"github.com/matt-hoiland/mocking/matching/service"
	"github.com/matt-hoiland/mocking/matching/service/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// func TestAPI_Create_Testify(t *testing.T) {
// 	t.FailNow()
// }

func TestAPI_Retrieve_Testify(t *testing.T) {
	testCases := []GoMockTestCase{
		{
			Name:                "Happy",
			RequestMethod:       "GET",
			UserAgent:           "matt",
			SessionKey:          "12345",
			RequestID:           &idZero,
			ResponseCode:        200,
			ResponseCodeMessage: "OK",
			ResponseData:        &smallBody,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.Name, func(t *testing.T) {
			serviceMock := &mocks.Service{}
			serviceMock.On("MakeRequest", mock.MatchedBy(MatchRequest(tc))).Return(BuildResponse(tc)).Once()

			api := &logic.API{
				Service:   serviceMock,
				UserAgent: tc.UserAgent,
				Session:   tc.SessionKey,
			}
			data, err := api.Retrieve(*tc.RequestID)
			assert.Equal(t, *tc.ResponseData, data)
			if tc.ErrorString == "" {
				assert.Nil(t, err)
			} else {
				assert.NotNil(t, err)
				assert.EqualError(t, err, tc.ErrorString)
			}
			serviceMock.AssertExpectations(t)
		})
	}
}

// func TestAPI_Update_Testify(t *testing.T) {
// 	t.FailNow()
// }

// func TestAPI_Delete_Testify(t *testing.T) {
// 	t.FailNow()
// }

// Test Utilities

func MatchRequest(tc GoMockTestCase) func(*service.Request) bool {
	return func(request *service.Request) bool {
		if request.Method != tc.RequestMethod {
			return false
		}
		if len(request.Headers["UserAgent"]) != 1 || request.Headers["UserAgent"][0] != tc.UserAgent {
			return false
		}
		if len(request.Headers["SessionKey"]) != 1 || request.Headers["SessionKey"][0] != fmt.Sprintf("s=%s", tc.SessionKey) {
			return false
		}

		body := request.Body
		if body == nil && (tc.RequestData != nil || tc.RequestID != nil) {
			return false
		}
		if tc.RequestID != nil && body.ID == nil {
			return false
		}
		if tc.RequestID == nil && body.ID != nil {
			return false
		}
		if tc.RequestID != nil && body.ID != nil && *tc.RequestID != *body.ID {
			return false
		}
		if tc.RequestData != nil && body.Data == nil {
			return false
		}
		if tc.RequestData == nil && body.Data != nil {
			return false
		}
		if tc.RequestData != nil && body.Data != nil && *tc.RequestData != *body.Data {
			return false
		}

		return true
	}
}
