package service

import(
	"testing"
	"app/domain"
	"context"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestUser_Create(t *testing.T){
	controller := gomock.NewController(t)
	defer controller.Finish()

	mockRepository := NewMockUserRepository(controller)

	mockRepository.
		EXPECT().
		Create(gomock.Any(), gomock.Any()).
		Return( domain.User{
			Name: "Joana",
			Email: "joanavidon@gmail.com",
			Password: "jo123"},
			 nil).
		Times(1)

	UserServiceMock := NewUserService(mockRepository)

	createdUser, err := UserServiceMock.Create(context.Background(), domain.User{
		Name: "Joana",
		Email: "joanavidon@gmail.com",
		Password: "jo123",
	})

	createdUser.ID = "1"

	assert.Exactly(t, "Joana", createdUser.Name)
	assert.Exactly(t, "joanavidon@gmail.com", createdUser.Email)
	assert.Exactly(t, "jo123", createdUser.Password)
	assert.Exactly(t, "1", createdUser.ID)
	assert.NotEmpty(t, createdUser.ID)
	assert.Nil(t, err)
}

func TestUser_GetID(t *testing.T){
	controller := gomock.NewController(t)
	defer controller.Finish()

	mockRepository :=  NewMockUserRepository(controller)

	mockRepository.
		EXPECT().
		Create(gomock.Any(), gomock.Any()).
		Return(domain.User{ Name: "JoaAna", Email: "joanaAvidon@gmail.com", Password: "joA123"}, nil)

	mockRepository.
		EXPECT().
		GetID(gomock.Any(), gomock.Any(), gomock.Any()).
		Return(domain.User{ID: "1", Name: "Joana", Email: "joanavidon@gmail.com", Password: "jo123"}, nil)
	
	UserServiceMock := NewUserService(mockRepository)

	_, err := UserServiceMock.Create(context.Background(), domain.User{
		Name: "JoanaA",
		Email: "joanaAvidon@gmail.com",
		Password: "joA123",
	})

	gotUser,  err := UserServiceMock.GetID(context.Background(), "1", domain.User{})

	assert.Nil(t, err)
	assert.Exactly(t, "1", gotUser.ID)
	assert.Exactly(t, "Joana", gotUser.Name)
	assert.Exactly(t, "joanavidon@gmail.com", gotUser.Email)
}

func TestUser_Login(t *testing.T){
	controller := gomock.NewController(t)
	defer controller.Finish()

	mockRepository :=  NewMockUserRepository(controller)

	mockRepository.
		EXPECT().
		Login(gomock.Any(), gomock.Any(), gomock.Any()).
		Return(false).
		Times(1)

	UserServiceMock := NewUserService(mockRepository)

	result := UserServiceMock.Login(context.Background(), 
	domain.User{
		Email: "joanavidon@gmail.com",
		Password: "jo123",
		},
	domain.Login{
		Email: "joanavidon@gmail.com",
		Password: "jo1234",
		})

	assert.Exactly(t, false, result)
}