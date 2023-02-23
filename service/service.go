package service

import (
	"errors"
	"goapirest/entity"
	"goapirest/repository"
	"math/rand"
)

// interface PostService == interface service en spring, se declaran los metodos
type PostService interface {
	Validate(post *entity.Post) error
	Create(post *entity.Post) (*entity.Post, error)
	FindAll() ([]entity.Post, error)
}

// struct que implementa la interface PostService
type service struct{}

// instancia de NewFirePostRepository , simil  @Autowired NewFirestorePostRepository en spring
// la funcionalidad de esta variable depende de la interfaz que le pasemos al constructor
var (
	repo repository.PostRepository
)

// Constructor , para crear una nueva instancia, de tipo PostService
// esta devuelve el struct que implementa la interfaz PostService
// por param le pasamos la interfaz PostRepository y se la asignamos a la variable
// repo usara la implementacion de la interfaz que le pasemos
// ver declaracion de interfaces en main.go
func NewPostService(repository repository.PostRepository) PostService {
	repo = repository
	return &service{}
}

//implementacion de los metodos de la interfaz PostService

func (*service) Validate(post *entity.Post) error {
	if post == nil {
		err := errors.New("The post is empty")
		return err
	}

	if post.Title == "" {
		err := errors.New("The title is empty")
		return err
	}

	return nil
}
func (*service) Create(post *entity.Post) (*entity.Post, error) {
	post.ID = rand.Int63()
	return repo.Save(post)

}
func (*service) FindAll() ([]entity.Post, error) {
	return repo.FindAll()
}
