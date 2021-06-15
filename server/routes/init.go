package routes

import (
  "context"
  // "html/template"
  "net/http"
  "time"
  "encoding/json"

	"github.com/gorilla/mux"
  "github.com/go-project/utils"
  "go.mongodb.org/mongo-driver/bson"
  "github.com/go-project/models"
  // "go.mongodb.org/mongo-driver/mongo"
  // "go.mongodb.org/mongo-driver/mongo/options"
)


func Routes() {
  r := mux.NewRouter()
  // r.GET("/", models.getHandler)
  r.HandleFunc("/", getHandler).Methods("GET", "OPTIONS")
  r.HandleFunc("/addNote", postHandler).Methods("POST")
  r.HandleFunc("/deleteNote", postDelete).Methods("POST")
  http.ListenAndServe(":8080", r)
}

func getHandler(w http.ResponseWriter, r *http.Request) {
  w.Header().Set("Access-Control-Allow-Origin", "*")

  w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	var notes []utils.Note
	notesCollection := models.DB.Database("birdpedia").Collection("birds")
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()
  cur, err := notesCollection.Find(ctx, bson.M{})
  if err = cur.All(ctx, &notes); err != nil {
      panic(err)
  }
	defer cur.Close(ctx)

  err = json.NewEncoder(w).Encode(notes)
  if err != nil {
      http.Error(w, err.Error(), 500)
  }
   
}

func postHandler(w http.ResponseWriter, r *http.Request) {
  // var notes utils.Note
  notesCollection := models.DB.Database("birdpedia").Collection("birds")
  ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

  note := utils.Note{}

  r.ParseForm()

  note.Title = r.Form.Get("title")
  note.Description = r.Form.Get("description")

  notesCollection.InsertOne(ctx, bson.D{
      {Key: "title", Value: note.Title},
      {Key: "description", Value: note.Description},
  })

  http.Redirect(w, r, "/", http.StatusFound)
}

func postDelete(w http.ResponseWriter, r *http.Request) {
  // var notes utils.Note
  notesCollection := models.DB.Database("birdpedia").Collection("birds")
  ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

  note := utils.Note{}

  r.ParseForm()

  note.Title = r.Form.Get("title")

  notesCollection.DeleteOne(ctx, bson.M{"title": note.Title})

  http.Redirect(w, r, "/", http.StatusFound)
}
