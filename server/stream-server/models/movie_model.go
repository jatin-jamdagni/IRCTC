// package models

// import "go.mongodb.org/mongo-driver/bson/primitive"

// type Genre struct {
// 	GenreID   int    `bson:"genre_id" json:"genre_id" validate:"required"`
// 	GenreName string ` bson:"genre_name" json:"genre_name" validate:"required,min=2,max=100"`
// }

// type Ranking struct {
// 	RankingValue int    `bson:"ranking_value" json:"ranking_value" validate:"required"`
// 	RankingName  string `bson:"ranking_name" json:"ranking_name" validate:"required"`
// }

// type Movie struct {
// 	ID          primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
// 	ImdbID      string             `bson:"imdb_id" json:"imdb_id" validate:"required"`
// 	Title       string             `bson:"title" json:"title" validate:"required,min=2,max=500"`
// 	PosterPath  string             `bson:"poster_path,omitempty" json:"poster_path" validate:"required,url"`
// 	YoutubeID   string             `bson:"youtube_id,omitempty" json:"youtube_id" validate:"required"`
// 	Genre       []Genre            `bson:"genre,omitempty" json:"genre,omitempty" validate:"required,dive"`
// 	AdminReview string             `bson:"admin_review,omitempty" json:"admin_review"`
// 	Ranking     Ranking            `bson:"ranking,omitempty" json:"ranking" validate:"required"`
// }

package models

import "go.mongodb.org/mongo-driver/bson/primitive"

// --- Genre subdocument ---
type Genre struct {
	GenreID   int    `bson:"genre_id" json:"genre_id" validate:"required"`
	GenreName string `bson:"genre_name" json:"genre_name" validate:"required,min=2,max=100"`
	// ✅ Fix: Removed extra space in json tag (`" bson:` → `"bson:`) and ensured field types match MongoDB storage
}

// --- Ranking subdocument ---
type Ranking struct {
	RankingValue int    `bson:"ranking_value" json:"ranking_value" validate:"required"`
	RankingName  string `bson:"ranking_name" json:"ranking_name" validate:"required"`
	// ✅ Fix: Ensure this matches MongoDB subdocument structure. Was fine, but added comment for clarity.
}

// --- Main Movie model ---
type Movie struct {
	ID          primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	ImdbID      string             `bson:"imdb_id" json:"imdb_id" validate:"required"`
	Title       string             `bson:"title" json:"title" validate:"required,min=2,max=500"`
	PosterPath  string             `bson:"poster_path,omitempty" json:"poster_path" validate:"required,url"`
	YoutubeID   string             `bson:"youtube_id,omitempty" json:"youtube_id" validate:"required"`
	Genre       []string           `bson:"genre,omitempty" json:"genre,omitempty" validate:"required,dive,required"`
	AdminReview string             `bson:"admin_review,omitempty" json:"admin_review"`
	Ranking     Ranking            `bson:"ranking,omitempty" json:"ranking" validate:"required"`
}
