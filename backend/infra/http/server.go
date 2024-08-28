package http

func StartServer() {
	r := SetupRouter()

	err := r.Run(":8080")

	if err != nil {
		return
	}
}
