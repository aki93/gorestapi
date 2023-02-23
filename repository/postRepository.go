package repository

import (
	"goapirest/entity"
)

//interfaz de repository, utilizando esta interfaz podemos implementar los metodos
//indistinamente de la DB que utilicemos, ej (firestore-repo.go)

type PostRepository interface {
	Save(post *entity.Post) (*entity.Post, error)
	FindAll() ([]entity.Post, error)
}
