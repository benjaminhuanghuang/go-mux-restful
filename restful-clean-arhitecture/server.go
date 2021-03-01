package main

import (
	"fmt"
	"net/http"

	"./controller"
	router "./http"
	"./repository"
	"./service"
)

var (
	postRepository repository.PostRepository = repository.NewFirestorePostRepository()
	postService    service.PostService       = service.NewPostService(postRepository)
	postController controller.PostController = controller.NewPostController(postService)
	httpRouter     router.Router             = router.NewMuxRouter()
)

func main() {
	const port string = ":8000"

	// Create router
	httpRouter.GET("/", func(resp http.ResponseWriter, req *http.Request) {
		fmt.Fprintln(resp, "dddd")
	})

	httpRouter.GET("/posts", postController.GtePosts)
	httpRouter.POST("/posts", postController.AddPost)
	// Start server
	httpRouter.SERVE(port)
}
