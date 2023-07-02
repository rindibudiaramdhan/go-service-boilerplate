package resources

import (
	"context"
	"io"
	"net/http"

	"github.com/go-chi/chi/v5"
)

type PostsResource struct{}

func (PostsResource) Routes() chi.Router {
	r := chi.NewRouter()

	r.Get("/", listPosts)   // GET /posts - Read a list of posts.
	r.Post("/", createPost) // POST /posts - Create a new post.

	r.Route("/{id}", func(r chi.Router) {
		r.Use(postCtx)
		r.Get("/", getPost)       // GET /posts/{id} - Read a single post by :id.
		r.Put("/", updatePost)    // PUT /posts/{id} - Update a single post by :id.
		r.Delete("/", deletePost) // DELETE /posts/{id} - Delete a single post by :id.
	})

	return r
}

// `docgen` cannot link to source code of handlers that are attached to a struct.
// Keep struct methods separate for testing, etc.
func (PostsResource) List(w http.ResponseWriter, r *http.Request) {
	listPosts(w, r)
}

func (PostsResource) Create(w http.ResponseWriter, r *http.Request) {
	createPost(w, r)
}

func (PostsResource) Get(w http.ResponseWriter, r *http.Request) {
	getPost(w, r)
}

func (PostsResource) Update(w http.ResponseWriter, r *http.Request) {
	updatePost(w, r)
}

func (PostsResource) Delete(w http.ResponseWriter, r *http.Request) {
	deletePost(w, r)
}

// Request Handler - GET /posts - Read a list of posts.
func listPosts(w http.ResponseWriter, r *http.Request) {
	resp, err := http.Get("https://jsonplaceholder.typicode.com/posts")

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	defer resp.Body.Close()

	w.Header().Set("Content-Type", "application/json")

	if _, err := io.Copy(w, resp.Body); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

// Request Handler - POST /posts - Create a new post.
func createPost(w http.ResponseWriter, r *http.Request) {
	resp, err := http.Post("https://jsonplaceholder.typicode.com/posts", "application/json", r.Body)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	defer resp.Body.Close()

	w.Header().Set("Content-Type", "application/json")

	if _, err := io.Copy(w, resp.Body); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

type key string

const (
	keyPrincipalID key = "id"
	// ...
)

func postCtx(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), keyPrincipalID, chi.URLParam(r, "id"))
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

// Request Handler - GET /posts/{id} - Read a single post by :id.
func getPost(w http.ResponseWriter, r *http.Request) {
	id := r.Context().Value("id").(string)

	resp, err := http.Get("https://jsonplaceholder.typicode.com/posts/" + id)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	defer resp.Body.Close()

	w.Header().Set("Content-Type", "application/json")

	if _, err := io.Copy(w, resp.Body); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

// Request Handler - PUT /posts/{id} - Update a single post by :id.
func updatePost(w http.ResponseWriter, r *http.Request) {
	id := r.Context().Value("id").(string)
	client := &http.Client{}

	req, err := http.NewRequest("PUT", "https://jsonplaceholder.typicode.com/posts/"+id, r.Body)
	req.Header.Add("Content-Type", "application/json")

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	resp, err := client.Do(req)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	defer resp.Body.Close()

	w.Header().Set("Content-Type", "application/json")

	if _, err := io.Copy(w, resp.Body); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

// Request Handler - DELETE /posts/{id} - Delete a single post by :id.
func deletePost(w http.ResponseWriter, r *http.Request) {
	id := r.Context().Value("id").(string)
	client := &http.Client{}

	req, err := http.NewRequest("DELETE", "https://jsonplaceholder.typicode.com/posts/"+id, nil)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	resp, err := client.Do(req)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	defer resp.Body.Close()

	w.Header().Set("Content-Type", "application/json")

	if _, err := io.Copy(w, resp.Body); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
