package controller

import (
	"encoding/json"
	"goapirest/entity"
	"goapirest/errormsg"
	"goapirest/service"
	"net/http"
)

// instancia de NewPostService , simil  @Autowired NewPostService en spring
// la funcionalidad de esta variable depende de la interfaz que le pasemos al constructor
var (
	postService service.PostService
)

// struct que implementa la interface PostController
type controller struct{}

// Constructor , para crear una nueva instancia, de tipo PostController
// esta devuelve el struct que implementa la interfaz PostController
// por param le pasamos la interfaz PostService y se la asignamos a la variable
// postService usara la implementacion de la interfaz que le pasemos
// ver declaracion de interfaces en main.go
func NewPostController(service service.PostService) PostController {
	postService = service
	return &controller{}
}

// interface de PostController , declaramos los metodos del controlador
type PostController interface {
	GetPosts(response http.ResponseWriter, request *http.Request)
	AddPosts(response http.ResponseWriter, request *http.Request)
}

//implementacion de los metodos de PostController

func (*controller) GetPosts(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-type", "application/json")
	posts, err := postService.FindAll()

	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(response).Encode(errormsg.ServiceError{Message: "Error getting posts"})
	}

	response.WriteHeader(http.StatusOK)
	json.NewEncoder(response).Encode(posts)

}

// metodo para hacer posteos
// por param envio interfaz writer de Responses (paquete http) y una Struct de request (paquete http)
func (*controller) AddPosts(response http.ResponseWriter, request *http.Request) {
	//creo la variable de tipo(Struct)Post
	var post entity.Post
	//mediante el decoder , decodeo el body de la request y lo mapeo mediante (.Decode) a la variable post que cree arriba
	//si o si devuelve err, si esta ok el err va a ser nil
	err := json.NewDecoder(request.Body).Decode(&post)

	if err != nil {
		//le indicamos el status code a la response
		response.WriteHeader(http.StatusInternalServerError)
		//asi escribe la respuesta
		json.NewEncoder(response).Encode(errormsg.ServiceError{Message: "Error unmarshalling request"})
		return
	}

	valerr := postService.Validate(&post)

	if valerr != nil {
		//le indicamos el status code a la response
		response.WriteHeader(http.StatusInternalServerError)
		//asi escribe la respuesta
		json.NewEncoder(response).Encode(errormsg.ServiceError{Message: valerr.Error()})
		return
	}

	//agregamos el post al slice(array) posts
	result, createErr := postService.Create(&post)

	if createErr != nil {
		//le indicamos el status code a la response
		response.WriteHeader(http.StatusInternalServerError)
		//asi escribe la respuesta
		json.NewEncoder(response).Encode(errormsg.ServiceError{Message: "Error savig post"})
		return
	}
	response.WriteHeader(http.StatusOK)
	//pasamos post a json
	json.NewEncoder(response).Encode(result)
}
