package domain

// Movie
type Movie struct {
	ID    int64  `json:"id" bson:"_id"`
	Title string `json:"title" bson:"title"`
	Year  string `json:"year" bson:"year"`
}
