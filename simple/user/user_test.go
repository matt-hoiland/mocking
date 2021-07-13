package user_test

import (
	"testing"

	"github.com/matt-hoiland/mocking/simple/doer/mocks"
	"github.com/matt-hoiland/mocking/simple/user"

	"github.com/golang/mock/gomock"
)

func TestUser_AccessService_GoMock(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockDoer := mocks.NewMockDoer(ctrl)
	mockDoer.EXPECT().Do("five", 5).Return(nil).MaxTimes(1)

	user := &user.User{mockDoer}
	user.AccessService()
	user.AccessService()
}

func TestUser_AccessService_Testify(t *testing.T) {
	mockDoer := &mocks.Doer{}
	mockDoer.On("Do", "five", 5).Return(nil).Once()

	user := &user.User{mockDoer}
	user.AccessService()
	user.AccessService()

	mockDoer.AssertExpectations(t)
}
