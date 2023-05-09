package main

func main() {
	server := NewServer(":8000")
	server.Handle("GET", "/", HandleRoot)
	server.Handle("POST", "/create", PostRequest)
	server.Handle("POST", "/user", UserPostRequest)
	server.Handle("POST", "/home", server.AddMidleware(HandleHome, CheckAuth(), Logging()))
	server.Listen()
}
