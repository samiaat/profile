package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"path/filepath"
)

func GetUsersHandler(w http.ResponseWriter, r *http.Request) {
	user, err := middleware.GetAuthenticatedUser(r)
	if err != nil {
		fmt.Println(err)
	}
	users, err := query.GetAllUsersExcluding(user.ID) // Implement this function
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	// Encode users to JSON and send the response
	if err := json.NewEncoder(w).Encode(users); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

// GetImageHandler serves the image file
func GetImageHandler(w http.ResponseWriter, r *http.Request) {
	imageName := r.URL.Query().Get("imageName") // Get the image name from the query parameter

	imagePath := filepath.Join("../../pkg/db/uploads", imageName)

	// Check if the file exists
	if _, err := os.Stat(imagePath); os.IsNotExist(err) {
		http.Error(w, "Image not found", http.StatusNotFound)
		return
	}

	// Serve the image file
	http.ServeFile(w, r, imagePath)
}
