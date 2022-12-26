package service

import (
	"app/domain"
	"context"
	"testing"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

var pizzeria = domain.Shop{
	Name:     "Pizzaria",
	Flavors:  [6]string{"Marguerita", "Calabresa"},
	Score:    9.7,
	Price:    55.60,
	Link:     "www.ifood.com.br",
	Favorite: true,
}

var pizzeria1 = domain.Shop{
	Name:     "PizzaAria",
	Flavors:  [6]string{"MargueritaA", "CalabresaA"},
	Score:    8.5,
	Price:    85.60,
	Link:     "www.ifoodA.com.br",
	Favorite: true,
}

var pizzeria2 = domain.Shop{
	Name:     "Pizzeria2",
	Flavors:  [6]string{"Champignon"},
	Score:    4.5,
	Price:    75.0,
	Link:     "www.ifood.com.br/pizzeria2",
	Favorite: false,
}

func TestShop_CreateShop(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()
	mockRepository := NewMockShopRepository(controller)

	mockRepository.EXPECT().
		CreateShop(gomock.Any(), gomock.Any()).
		Return(pizzeria, nil)

	ShopServiceMock := NewShopService(mockRepository)
	createdShop, err := ShopServiceMock.CreateShop(context.Background(), pizzeria)

	createdShop.ID = "1"

	assert.Exactly(t, "Pizzaria", createdShop.Name)
	assert.Exactly(t, [6]string{"Marguerita", "Calabresa"}, createdShop.Flavors)
	assert.Exactly(t, 9.7, createdShop.Score)
	assert.Exactly(t, 55.60, createdShop.Price)
	assert.Exactly(t, "www.ifood.com.br", createdShop.Link)
	assert.Exactly(t, true, createdShop.Favorite)
	assert.NotEmpty(t, createdShop.ID)
	assert.Nil(t, err)
}

func TestShop_GetAll(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()
	mockRepository := NewMockShopRepository(controller)

	mockRepository.EXPECT().
		CreateShop(gomock.Any(), gomock.Any()).
		Return(pizzeria1, nil)

	mockRepository.EXPECT().
		GetAll(gomock.Any(), gomock.Any()).
		Return([]domain.Shop{pizzeria, pizzeria1}, nil)

	ShopServiceMock := NewShopService(mockRepository)
	_, err := ShopServiceMock.CreateShop(context.Background(), pizzeria1)
	shops, err := ShopServiceMock.GetAll(context.Background(), domain.Shop{})

	if shops[0] != pizzeria {
		t.Fatal()
	}
	if shops[1] != pizzeria1 {
		t.Fatal()
	}
	assert.Nil(t, err)
}

func TestShop_GetByID(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()
	mockRepository := NewMockShopRepository(controller)

	mockRepository.EXPECT().
		GetByID(gomock.Any(), gomock.Any(), gomock.Any()).
		Return(pizzeria, nil)

	ShopServiceMock := NewShopService(mockRepository)
	shop, err := ShopServiceMock.GetByID(context.Background(), "1", domain.Shop{})

	assert.Exactly(t, "Pizzaria", shop.Name)
	assert.Exactly(t, [6]string{"Marguerita", "Calabresa"}, shop.Flavors)
	assert.Exactly(t, 9.7, shop.Score)
	assert.Exactly(t, 55.60, shop.Price)
	assert.Exactly(t, "www.ifood.com.br", shop.Link)
	assert.Exactly(t, true, shop.Favorite)
	assert.Nil(t, err)
}

func TestShop_GetByName(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()
	mockRepository := NewMockShopRepository(controller)

	mockRepository.EXPECT().
		GetByName(gomock.Any(), gomock.Any(), gomock.Any()).
		Return(pizzeria, nil)

	ShopServiceMock := NewShopService(mockRepository)
	createdShop, err := ShopServiceMock.GetByName(context.Background(), "pizzeria", domain.Shop{})

	assert.Exactly(t, "Pizzaria", createdShop.Name)
	assert.Exactly(t, [6]string{"Marguerita", "Calabresa"}, createdShop.Flavors)
	assert.Exactly(t, 9.7, createdShop.Score)
	assert.Exactly(t, 55.60, createdShop.Price)
	assert.Exactly(t, "www.ifood.com.br", createdShop.Link)
	assert.Exactly(t, true, createdShop.Favorite)
	assert.Nil(t, err)
}

func TestShop_GetByScore(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()
	mockRepository := NewMockShopRepository(controller)

	mockRepository.EXPECT().
		GetByScore(gomock.Any(), gomock.Any(), gomock.Any()).
		Return([]domain.Shop{pizzeria}, nil)

	ShopServiceMock := NewShopService(mockRepository)
	shops, err := ShopServiceMock.GetByScore(context.Background(), 9.7, domain.Shop{})

	if shops[0] != pizzeria {
		t.Fatal()
	}
	assert.Nil(t, err)
}

func TestShop_GetByPrice(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()
	mockRepository := NewMockShopRepository(controller)

	mockRepository.EXPECT().
		GetByPrice(gomock.Any(), gomock.Any(), gomock.Any()).
		Return([]domain.Shop{pizzeria, pizzeria1}, nil)

	ShopServiceMock := NewShopService(mockRepository)
	_, err := ShopServiceMock.GetByPrice(context.Background(), 85.60, domain.Shop{})

	assert.Nil(t, err)
}

func TestShop_Update(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()
	mockRepository := NewMockShopRepository(controller)


	mockRepository.EXPECT().
		CreateShop(gomock.Any(), gomock.Any()).
		Return(pizzeria2, nil)

	mockRepository.EXPECT().
		Update(gomock.Any(), gomock.Any(), gomock.Any()).
		Return(nil)

	mockRepository.EXPECT().
		GetByID(gomock.Any(), gomock.Any(), gomock.Any()).
		Return(pizzeria2, nil)

	ShopServiceMock := NewShopService(mockRepository)

	createdShop, err := ShopServiceMock.CreateShop(context.Background(), pizzeria2)
	createdShop.ID = "New"
	
			
	err = ShopServiceMock.Update(context.Background(), "New", domain.Shop{
		Flavors:  [6]string{"Champignon", "Toscana"}})

	shop, err := ShopServiceMock.GetByID(context.Background(), "New", domain.Shop{})

	assert.Exactly(t, "Pizzeria2", shop.Name)
	assert.Exactly(t, [6]string{"Champignon", "Toscana"}, shop.Flavors)
	assert.Exactly(t, 4.5, shop.Score)
	assert.Exactly(t, 75.0, shop.Price)
	assert.Exactly(t, "www.ifood.com.br/pizzeria2", shop.Link)
	assert.Exactly(t, false, shop.Favorite)
	assert.Nil(t, err)
}

/* func TestShop_UpdateScore(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()
	mockRepository := NewMockShopRepository(controller)

	mockRepository.EXPECT().
		UpdateScore(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).
		Return(nil)

	ShopServiceMock := NewShopService(mockRepository)

	err := ShopServiceMock.UpdateScore(context.Background(), "1",
		9.0, domain.Shop{})

	assert.Exactly(t, "Pizzaria", pizzeria.Name)
	assert.Exactly(t, [6]string{"Marguerita", "Calabresa"}, pizzeria.Flavors)
	assert.Exactly(t, 9.0, pizzeria.Score)
	assert.Exactly(t, 55.60, pizzeria.Price)
	assert.Exactly(t, "www.ifood.com.br", pizzeria.Link)
	assert.Exactly(t, true, pizzeria.Favorite)
	//assert.NotEmpty(t, pizzeria.ID)
	assert.Nil(t, err)
} */

/* func TestShop_UpdatePrice(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()
	mockRepository := NewMockShopRepository(controller)

	mockRepository.EXPECT().
		UpdatePrice(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).
		Return(nil)

	mockRepository.EXPECT().
		GetByID(gomock.Any(), gomock.Any(), gomock.Any()).
		Return(pizzeria2, nil)
	
	ShopServiceMock := NewShopService(mockRepository)

	err := ShopServiceMock.UpdatePrice(context.Background(), "New",
		80.0, domain.Shop{})

	shop, err := ShopServiceMock.GetByID(context.Background(), "New", domain.Shop{})

	assert.Exactly(t, "Pizzeria2", shop.Name)
	assert.Exactly(t, [6]string{"Champignon"}, shop.Flavors)
	assert.Exactly(t, 4.5, shop.Score)
	assert.Exactly(t, 80.0, shop.Price)
	assert.Exactly(t, "www.ifood.com.br/pizzeria2", shop.Link)
	assert.Exactly(t, false, shop.Favorite)
	assert.Nil(t, err)
} */

func Test_Delete(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()
	mockRepository := NewMockShopRepository(controller)

	mockRepository.EXPECT().
		Delete(gomock.Any(), gomock.Any()).
		Return(nil)

	mockRepository.EXPECT().
		GetAll(gomock.Any(), gomock.Any()).
		Return([]domain.Shop{pizzeria1}, nil)

	ShopServiceMock := NewShopService(mockRepository)
	err := ShopServiceMock.Delete(context.Background(), "1")
	_, err = ShopServiceMock.GetAll(context.Background(), domain.Shop{})

	assert.Nil(t, err)
}

func Test_ListScores(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()
	mockRepository := NewMockShopRepository(controller)

	mockRepository.EXPECT().
		ListScores(gomock.Any(), gomock.Any()).
		Return(nil, []string{"pizzeria1", "pizzeria2"})

	ShopServiceMock := NewShopService(mockRepository)
	_, _ = ShopServiceMock.ListScores(context.Background(), domain.Shop{})
}
