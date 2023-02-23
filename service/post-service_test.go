package service

import (
	"goapirest/entity"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockRepository struct {
	mock.Mock
}

func (mock *MockRepository) Save(post *entity.Post) (*entity.Post, error) {
	arg := mock.Called()
	result := arg.Get(0)
	return result.(*entity.Post), arg.Error(1)
}

func (mock *MockRepository) FindAll() ([]entity.Post, error) {
	arg := mock.Called()
	result := arg.Get(0)
	return result.([]entity.Post), arg.Error(1)
}

func TestValidateEmptyPost(t *testing.T) {

	//le pasamos nil en este metodo, porque no utiliza un repo
	testService := NewPostService(nil)

	err := testService.Validate(nil)

	assert.NotNil(t, err)

	assert.Equal(t, err.Error(), "The post is empty")
}

func TestValidateEmptyTitlePost(t *testing.T) {
	post := entity.Post{ID: 1, Title: "", Text: "B "}

	//le pasamos nil en este metodo, porque no utiliza un repo
	testService := NewPostService(nil)

	//si al param post no agregamos su referencia &
	//tira error , ver argumento del metodo , utiliza *entity.Post
	err := testService.Validate(&post)

	assert.NotNil(t, err)

	assert.Equal(t, "The title is empty", err.Error())
}

func TestFindAll(t *testing.T) {
	post := entity.Post{ID: 1, Title: "A", Text: "A"}

	var identifier int64 = 1

	mockRepo := new(MockRepository)

	mockRepo.On("FindAll").Return([]entity.Post{post}, nil)

	testService := NewPostService(mockRepo)

	result, err := testService.FindAll()

	mockRepo.AssertExpectations(t)

	assert.Equal(t, identifier, result[0].ID)
	assert.Equal(t, "A", result[0].Title)
	assert.Equal(t, "A", result[0].Text)

	assert.Equal(t, nil, err)

}

func TestCreate(t *testing.T) {
	post := entity.Post{ID: 1, Title: "A", Text: "A"}

	//var identifier int64 = 1

	mockRepo := new(MockRepository)

	mockRepo.On("Save").Return(&post, nil)

	testService := NewPostService(mockRepo)

	result, err := testService.Create(&post)

	mockRepo.AssertExpectations(t)

	assert.NotNil(t, result.ID)
	assert.Equal(t, result.Title, "A")
	assert.Equal(t, "A", result.Text)
	assert.Nil(t, err)

}
