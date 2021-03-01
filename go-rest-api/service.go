package main

import (
	"fmt"
	"net/http"

	"./controller"
	router "./http"
)

var (
	httpRouter     router.Router             = router.NewMuxRouter()
	postController controller.PostController = controller.NewPostController()
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
