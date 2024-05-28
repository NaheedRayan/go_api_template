package main

// https://hugo-johnsson.medium.com/rest-api-with-golang-and-mux-e934f581b8b5
import (
	"encoding/json"
	"flag"
	"fmt"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)
type Post struct {
  	ID string `json:"id"`
  	Title string `json:"title"`
  	Body string `json:"body"`
}

var posts []Post

func Hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w , "Hello from the GO server")
}

func getPosts(w http.ResponseWriter, r *http.Request) {
  	w.Header().Set("Content-Type", "application/json")
  	json.NewEncoder(w).Encode(posts)	
}
func createPost(w http.ResponseWriter, r *http.Request) {
  	w.Header().Set("Content-Type", "application/json")
  	var post Post
  	_ = json.NewDecoder(r.Body).Decode(&post)
  	post.ID = strconv.Itoa(rand.Intn(1000000))
  	posts = append(posts, post)
  	json.NewEncoder(w).Encode(&post)
}
func getPost(w http.ResponseWriter, r *http.Request) {
  	w.Header().Set("Content-Type", "application/json")
  	params := mux.Vars(r)
  	for _, item := range posts {
  	  if item.ID == params["id"] {
  	    json.NewEncoder(w).Encode(item)
  	    return
  	  }
  	}
  	json.NewEncoder(w).Encode(&Post{})
}
func updatePost(w http.ResponseWriter, r *http.Request) {
  	w.Header().Set("Content-Type", "application/json")
  	params := mux.Vars(r)
  	for index, item := range posts {
  	  if item.ID == params["id"] {
  	    posts = append(posts[:index], posts[index+1:]...)
  	    var post Post
  	    _ = json.NewDecoder(r.Body).Decode(&post)
  	    post.ID = params["id"]
  	    posts = append(posts, post)
  	    json.NewEncoder(w).Encode(&post)
  	    return
  	  }
  	}
  	json.NewEncoder(w).Encode(posts)
}
func deletePost(w http.ResponseWriter, r *http.Request) {
  	w.Header().Set("Content-Type", "application/json")
  	params := mux.Vars(r)
  	for index, item := range posts {
  	  if item.ID == params["id"] {
  	    posts = append(posts[:index], posts[index+1:]...)
  	    break
  	  }
  	}
  	json.NewEncoder(w).Encode(posts)
}
func main() {

	// for argument passing
	// go run main.go --listenAddr :2000
	listenAddr := flag.String("listenAddr" , ":3000" , "the server address")
	flag.Parse()// Parse the command-line flags

	
	router := mux.NewRouter()
	posts = append(posts, Post{ID: "1", Title: "My first post",  Body: "This is the content of my first post" })
	posts = append(posts, Post{ID: "2", Title: "My second post", Body: "This is the content of my second post"})

	// router.HandleFunc("/", Hello).Methods("GET")
	// router.HandleFunc("/posts", getPosts).Methods("GET")
	// router.HandleFunc("/posts", createPost).Methods("POST")
	// router.HandleFunc("/posts/{id}", getPost).Methods("GET")
	// router.HandleFunc("/posts/{id}", updatePost).Methods("PUT")
	// router.HandleFunc("/posts/{id}", deletePost).Methods("DELETE")
	router.HandleFunc("GET /", Hello)
	router.HandleFunc("GET /posts", getPosts)
	router.HandleFunc("POST /posts", createPost)
	router.HandleFunc("GET /posts/{id}", getPost)
	router.HandleFunc("PUT /posts/{id}", updatePost)
	router.HandleFunc("DELETE /posts/{id}", deletePost)

	fmt.Println("Server running at port", *listenAddr)
	// http.ListenAndServe(":8000", router)
	http.ListenAndServe(*listenAddr, router)
}