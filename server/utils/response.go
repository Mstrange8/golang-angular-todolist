package utils


type Note struct {
    Title       string `bson:"title,omitempty"`
    Description string `bson:"description,omitempty"`
}