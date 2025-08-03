 import ("backend/pkg/middleware"
 query "backend/pkg/db/queries")

func GetPostsHandler(w http.ResponseWriter, r *http.Request) {
	user, err := middleware.GetAuthenticatedUser(r)
	if err != nil {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	posts, err := query.GetPostsQuery(user.ID)
	if err != nil {
		http.Error(w, "Error retrieving posts", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(posts); err != nil {
		http.Error(w, "Error encoding posts to JSON", http.StatusInternalServerError)
	}
}
