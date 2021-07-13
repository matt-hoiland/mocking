package logic_test

import (
	"bytes"
	"errors"
	"fmt"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/matt-hoiland/mocking/matching/logic"
	"github.com/matt-hoiland/mocking/matching/service"
	"github.com/matt-hoiland/mocking/matching/service/mocks"
	"github.com/stretchr/testify/assert"
)

var (
	smallBody = "{age:29}"
	idZero    = 0
)

func TestAPI_Create_GoMock(t *testing.T) {
	testCases := []GoMockTestCase{
		{
			Name:                "Happy",
			RequestMethod:       "POST",
			UserAgent:           "matt",
			SessionKey:          "12345",
			RequestData:         &smallBody,
			ResponseCode:        200,
			ResponseCodeMessage: "OK",
			ResponseID:          &idZero,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.Name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			requestMatcher := RequestMatcher{tc}

			serviceMock := mocks.NewMockService(ctrl)
			serviceMock.EXPECT().MakeRequest(&requestMatcher).Return(BuildResponse(tc)).Times(1)

			api := logic.API{
				Service:   serviceMock,
				UserAgent: tc.UserAgent,
				Session:   tc.SessionKey,
			}
			id, err := api.Create(*tc.RequestData)
			assert.Equal(t, *tc.ResponseID, id)
			if tc.ErrorString != "" {
				assert.NotNil(t, err)
				assert.EqualError(t, err, tc.ErrorString)
			} else {
				assert.Nil(t, err)
			}
		})

	}
}

// func TestAPI_Retrieve_GoMock(t *testing.T) {
// 	t.FailNow()
// }

// func TestAPI_Update_GoMock(t *testing.T) {
// 	t.FailNow()
// }

// func TestAPI_Delete_GoMock(t *testing.T) {
// 	t.FailNow()
// }

// Test Utilities

type GoMockTestCase struct {
	// TC Name
	Name string
	// Request Params
	RequestMethod string
	UserAgent     string
	SessionKey    string
	RequestID     *int
	RequestData   *string
	//Response Params
	ResponseCode        int
	ResponseCodeMessage string
	ErrorString         string
	ResponseID          *int
	ResponseData        *string
}

type RequestMatcher struct {
	TC GoMockTestCase
}

func (req *RequestMatcher) Matches(x interface{}) bool {
	request, ok := x.(*service.Request)
	if !ok {
		return false
	}
	return MatchRequest(req.TC)(request)
}

func (req *RequestMatcher) String() string {
	buf := bytes.Buffer{}
	fmt.Fprintf(&buf, "%s /api/resource", req.TC.RequestMethod)
	if req.TC.RequestID != nil {
		fmt.Fprintf(&buf, "/%d", *req.TC.RequestID)
	}
	fmt.Fprintf(&buf, "\nUserAgent: %s\n", req.TC.UserAgent)
	fmt.Fprintf(&buf, "SessionKey: s=%s\n", req.TC.SessionKey)
	if req.TC.RequestData != nil {
		fmt.Fprintf(&buf, "\n%s", *req.TC.RequestData)
	}
	return buf.String()
}

func BuildResponse(tc GoMockTestCase) *service.Response {
	res := &service.Response{
		Code:        tc.ResponseCode,
		CodeMessage: tc.ResponseCodeMessage,
	}
	if tc.ErrorString != "" {
		res.Error = errors.New(tc.ErrorString)
	}
	if tc.ResponseID != nil || tc.ResponseData != nil {
		res.Body = &service.DataBody{
			ID:   tc.ResponseID,
			Data: tc.ResponseData,
		}
	}
	return res
}
