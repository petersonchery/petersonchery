package rest

	//gestionnaire de route
	func GestRoute(w http.ResponseWriter, r *http.Request){
		w.Write([]byte("Welcome Peet"))
	}
	
	func GestRoute1(w http.ResponseWriter, r *http.Request){
		w.Write ([]byte("I am here"))
	}
	