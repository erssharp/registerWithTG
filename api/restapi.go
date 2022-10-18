package api

import(
	"github.com/gorilla/mux"
	"net/http"
	"registerWithTG/api/controllers"
)

func main(){
	r := mux.NewRouter()
	empsr := r.PathPrefix("/employee").Subrouter()
	empc = controllers.
}

