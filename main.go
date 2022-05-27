package main

func main() {
	server := NewServer(":3000")
	server.Handle("/", server.AddMiddleware(HandleRoot, Logging()))
	server.Handle("/API", server.AddMiddleware(HandleHome, CheckAuth(), Logging()))
	server.Listen()
}
