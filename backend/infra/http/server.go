package http

func StartServer() {
	r := SetupRouter()

	r.Run(":8080")
}
