package app

import (
	"fmt"
	"log"
	"net/http"
	"postapi/app/models"
)

func (a *App) IndexHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Welcome to Post API")
	}
}

func (a *App) CreatePostHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		req := models.PostRequest{}
		err := parse(w, r, &req)
		if err != nil {
			log.Printf("Cannot parse post body. err=%v \n", err)
			sendResponse(w, r, nil, http.StatusBadRequest)
			return
		}

		// Create the post
		p := &models.Post{
			ID:      0,
			Title:   req.Title,
			Content: req.Content,
			Author:  req.Author,
		}

		// Save in DB
		err = a.DB.CreatePost(p)
		if err != nil {
			log.Printf("Cannot save post in DB. err=%v \n", err)
			sendResponse(w, r, nil, http.StatusInternalServerError)
			return
		}

		resp := mapPostToJSON(p)
		sendResponse(w, r, resp, http.StatusOK)

	}
}

func (a *App) GetPostsHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		posts, err := a.DB.GetPosts()
		if err != nil {
			log.Printf("Cannot get posts, err=%v \n", err)
			sendResponse(w, r, nil, http.StatusInternalServerError)
			return
		}

		var resp = make([]models.JsonPost, len(posts))
		for idx, post := range posts {
			resp[idx] = mapPostToJSON(post)
		}

		sendResponse(w, r, resp, http.StatusOK)
	}
}
