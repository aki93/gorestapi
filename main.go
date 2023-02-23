package main

import (
	"fmt"
	"goapirest/controller"
	router "goapirest/http"
	"goapirest/repository"
	"goapirest/service"
	"net/http"
	"os"
)

// declaracion interfaces que se utilizaran
var (
	postRepository repository.PostRepository = repository.NewFirestorePostRepository()
	postService    service.PostService       = service.NewPostService(postRepository)
	postController controller.PostController = controller.NewPostController(postService)
	httpRouter     router.Router             = router.NewChiRouter()
)

func main() {

	const port string = ":8080"
	httpRouter.GET("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Hello world")
	})
	os.Setenv("GOOGLE_APPLICATION_CREDENTIALS", "C:/Users/Learsoft/Desktop/go/goapirest/firebaseKey/gorest-api-3cfa7-firebase-adminsdk-yeed4-6b467a1247.json")
	httpRouter.GET("/posts", postController.GetPosts)
	httpRouter.POST("/addpost", postController.AddPosts)

	httpRouter.SERVE(port)
}
