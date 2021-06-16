package routes

import (
  "context"
  // "html/template"
  "net/http"
  "time"
  "encoding/json"
  "fmt"

	"github.com/gorilla/mux"
  "github.com/go-project/utils"
  "go.mongodb.org/mongo-driver/bson"
  "github.com/go-project/models"
  // "go.mongodb.org/mongo-driver/mongo"
  // "go.mongodb.org/mongo-driver/mongo/options"
)

// func CORS(next http.Handler) http.Handler {
// 	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

// 		// Set headers
// 		w.Header().Set("Access-Control-Allow-Headers:", "content-type application/json")
// 		w.Header().Set("Access-Control-Allow-Origin", "*")
// 		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")

// 		if r.Method == "OPTIONS" {
// 			w.WriteHeader(http.StatusOK)
// 			return
// 		}

// 		fmt.Println("ok")

// 		// Next
// 		next.ServeHTTP(w, r)
// 		return
// 	})
// }


func Routes() {
  r := mux.NewRouter()
  // r.GET("/", models.getHandler)
  r.HandleFunc("/", getHandler).Methods("GET", "OPTIONS")
  r.HandleFunc("/addNote", postHandler).Methods("POST", "OPTIONS")
  r.HandleFunc("/deleteNote", postDelete).Methods("POST", "OPTIONS")
  http.ListenAndServe(":8080", r)
}

func getHandler(w http.ResponseWriter, r *http.Request) {
  w.Header().Set("Access-Control-Allow-Origin", "*")
  w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
  w.Header().Set("Access-Control-Allow-Headers", "application/json text/plain */*")

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
  // if r.Method == "OPTIONS" {
  //   w.WriteHeader(http.StatusOK)
  //   return
  // }
  // var notes utils.Note
  fmt.Println(r.Method)
  // w.Header().Set("Content-Type", "application/json")
  // w.Header().Set("Access-Control-Allow-Origin", "*")
  // w.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS")

  // w.Header().Set("Access-Control-Allow-Headers", "content-type application/json")
  // notesCollection := models.DB.Database("birdpedia").Collection("birds")
  // ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	// defer cancel()

  // note := utils.Note{}

  // var s utils.Note

  fmt.Println(r)
  fmt.Println(w)
  fmt.Println(r.Body)

  // err := json.NewDecoder(strings.NewReader(r))
  // if err != nil {
  //   http.Error(w, err.Error(), 500)
  // }

  // fmt.Println(s)

  // r.ParseForm()

  // note.Title = r.Form.Get("title")
  // note.Description = r.Form.Get("description")

  // notesCollection.InsertOne(ctx, bson.D{
  //     {Key: "title", Value: note.Title},
  //     {Key: "description", Value: note.Description},
  // })

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
