package main
import (
	"net/http"
	"log"
	"github.com/smeet07/cinemabookingGO/internal/booking"
	"github.com/smeet07/cinemabookingGO/internal/utils"
	"github.com/smeet07/cinemabookingGO/internal/adapters/redis"

)
func main(){
	mux:=http.NewServeMux()
	mux.HandleFunc("GET /movies",listMovies)
	mux.Handle("GET /",http.FileServer(http.Dir("./static")))
	store:=booking.NewRedisStore(redis.NewClient("localhost:6379"))
	svc:=booking.NewService(store)
	h:=booking.NewHandler(svc)
	mux.HandleFunc("GET /movies/{movieID}/seats",h.ListSeats)	
	if err:=http.ListenAndServe(":8080", mux); err!=nil{
		log.Fatal(err)
	}
}
var movies=[]MovieResponse{
	{ID:"inception",Title:"Inception",Rows:5,SeatsPerRow:8},
	{ID:"interstellar",Title:"Interstellar",Rows:4,SeatsPerRow:6},
}
func listMovies(w http.ResponseWriter,r *http.Request){
	utils.WriteJSON(w,http.StatusOK,movies)
}

type MovieResponse struct{
	ID string `json:"id"`
	Title string `json:"title"`
	Rows int `json:"rows"`
	SeatsPerRow int `json:"seats_per_row"`
}
