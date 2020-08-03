package server

func Init() {
	r := setupRouter()
	r.Run(":8080")
}
