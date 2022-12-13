package service

import(
	"testing"
	"app/domain"
	"context"
	//"fmt"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

//_________________________GOMOCK_____________________________//

func TestUserUseCase_CreateUser_GoMockAndTestify(t *testing.T){
	controller := gomock.NewController(t)
	defer controller.Finish()

	mockRepository := NewMockUserRepository(controller)

	mockRepository.
		EXPECT().
		Create(gomock.Any(), gomock.Any()).
		Return(nil).
		Times(1)

	UserUseCaseMock := NewUserUseCase(mockRepository)

	createdUser, err := UserUseCaseMock.Create(context.Background(), domain.User{
		Name: "Joana",
		Email: "joanavidon@gmail.com",
		Password: "jo123",
	})

	assert.Nil(t, err)
	assert.Exactly(t, "Joana", createdUser.Name)
	assert.Exactly(t, "joanavidon@gmail.com", createdUser.Email)
	assert.Exactly(t, "jo123", createdUser.Password)
	assert.NotEmpty(t, createdUser.ID)
}

//esta passsando, mas acho que não está certo
func TestUserUseCase_GetidUser_GoMockAndTestify(t *testing.T){
	controller := gomock.NewController(t)
	defer controller.Finish()

	mockRepository := NewMockUserRepository(controller)

	mockRepository.
		EXPECT().
		GetID(gomock.Any(), gomock.Any(), gomock.Any()).
		Return(domain.User{"2", "Joana", "joanavidon@gmail.com", "jo123"}, nil).
		Times(1)

	UserUseCaseMock := NewUserUseCase(mockRepository)

	gotUser,  err := UserUseCaseMock.GetID(context.Background(), "2", domain.User{
	ID: "2",
	Name: "Joana",
	Email: "joanavidon@gmail.com",
	Password: "jo123",
})

	assert.Nil(t, err)
	assert.Exactly(t, "2", gotUser.ID)
	assert.Exactly(t, "Joana", gotUser.Name)
	assert.Exactly(t, "joanavidon@gmail.com", gotUser.Email)
	assert.Exactly(t, "jo123", gotUser.Password)
}