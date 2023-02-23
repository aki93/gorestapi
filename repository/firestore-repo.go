package repository

import (
	"context"
	"goapirest/entity"
	"log"
	"os"

	"google.golang.org/api/iterator"

	"cloud.google.com/go/firestore"
)

// implementacion de la interface
type repo struct{}

// Constructor , para crear una nueva instancia, de tipo PostRepository (postRepository.go)
// esta devuelve el struct que implementa la interfaz PostRepository
func NewFirestorePostRepository() PostRepository {
	return &repo{} //&referencia a repo != *pointer a repo
}

// constante del projectId del firestore, indicado en el archivo que esta en firebaseKey
const (
	projectId      string = "gorest-api-3cfa7"
	collectionName string = "posts"
)

//implementacion de los metodos de postRepository.go

func (*repo) Save(post *entity.Post) (*entity.Post, error) {
	//os.Getenv("GOOGLE_APPLICATION_CREDENTIALS")
	//se crea contexto para firebase
	//Incoming requests to a server should create a Context, and outgoing calls to servers should accept a Context
	ctx := context.Background()
	client, err := firestore.NewClient(ctx, projectId)
	if err != nil {
		log.Fatalf("Failed to create a Firestore Client: %v", err)
		return nil, err
	}

	//ese cliente se cierra una vez que termina de procesar
	defer client.Close()
	//map[string : es la key de la collection(ver firebase)]
	//interface es el value
	_, _, err = client.Collection(collectionName).Add(ctx, map[string]interface{}{
		"ID":    post.ID,
		"Title": post.Title,
		"Text":  post.Text,
	})

	if err != nil {
		log.Fatalf("Failed adding new post: %v", err)
		return nil, err
	}

	return post, nil
}

func (*repo) FindAll() ([]entity.Post, error) {
	os.Getenv("GOOGLE_APPLICATION_CREDENTIALS")
	//se crea contexto para firebase
	//Incoming requests to a server should create a Context, and outgoing calls to servers should accept a Context
	ctx := context.Background()
	client, err := firestore.NewClient(ctx, projectId)
	if err != nil {
		log.Fatalf("Failed to create a Firestore Client: %v", err)
		return nil, err
	}

	//ese cliente se cierra una vez que termina de procesar
	defer client.Close()

	var posts []entity.Post

	dociterator := client.Collection(collectionName).Documents(ctx)
	for {
		doc, err := dociterator.Next()

		if err == iterator.Done {
			break
		}

		if err != nil {
			log.Fatalf("Failed to iterate list of Posts: %v v", err)
			return nil, err
		}

		post := entity.Post{
			ID:    doc.Data()["ID"].(int64),
			Title: doc.Data()["Title"].(string),
			Text:  doc.Data()["Text"].(string),
		}

		posts = append(posts, post)

	}

	return posts, nil

}
