package route

import (
	authentification "FirstProject/Authentification"
	"FirstProject/rest"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func InializeRouter() *mux.Router {
	router := mux.NewRouter()

	router.Methods("GET", "PUT", "POST", "DELETE").Path("/api/v1/ping").HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		w.Write([]byte("server a l'ecoute"))
	})

	router.Methods("POST").Path("/api/v1/bibliotheques").Handler(Middleware4{Handler: http.HandlerFunc(rest.CreateBibliotheque)})
	router.Methods("POST").Path("/api/v1/clients").Handler(Middleware5{Handler: http.HandlerFunc(rest.CreateClient)})
	router.Methods("POST").Path("/api/v1/livres").Handler(Middleware6{Handler: http.HandlerFunc(rest.CreateLivre)})
	router.Methods("POST").Path("/api/v1/emprunts").Handler(Middleware7{Handler: http.HandlerFunc(rest.CreateEmprunt)})
	router.Methods("POST").Path("/api/v1/tenants").Handler(Middleware8{Handler: http.HandlerFunc(rest.CreateTenant)})

	router.Methods("GET").Path("/api/v1/bibliotheques/{id}").Handler(Middleware10{Handler: http.HandlerFunc(rest.GetBibliothequeById)})
	router.Methods("GET").Path("/api/v1/clients/{id}").Handler(Middleware11{Handler: http.HandlerFunc(rest.GetClientById)})
	router.Methods("GET").Path("/api/v1/livres/{id}").Handler(Middleware12{Handler: http.HandlerFunc(rest.GetLivreById)})
	router.Methods("GET").Path("/api/v1/emprunts/{id}").Handler(Middleware13{Handler: http.HandlerFunc(rest.GetEmpruntById)})
	router.Methods("GET").Path("/api/v1/tenants/{id}").Handler(Middleware14{Handler: http.HandlerFunc(rest.GetTenantById)})

	router.Methods("DELETE").Path("/api/v1/bibliotheques/sup/{id}").Handler(Middleware15{Handler: http.HandlerFunc(rest.DeleteBibliotheque)})
	router.Methods("DELETE").Path("/api/v1/clients/sup/{id}").Handler(Middleware16{Handler: http.HandlerFunc(rest.DeleteClientById)})
	router.Methods("DELETE").Path("/api/v1/emprunts/sup/{id}").Handler(Middleware17{Handler: http.HandlerFunc(rest.DeleteEmpruntById)})
	router.Methods("DELETE").Path("/api/v1/livres/sup/{id}").Handler(Middleware18{Handler: http.HandlerFunc(rest.DeleteLivreById)})
	router.Methods("DELETE").Path("/api/v1/tenants/sup/{id}").Handler(Middleware19{Handler: http.HandlerFunc(rest.DeleteTenantById)})

	router.Methods("PUT").Path("/api/v1/bibliotheques/up/{id}").Handler(Middleware20{Handler: http.HandlerFunc(rest.UpdateBibliotheque)})
	router.Methods("PUT").Path("/api/v1/clients/up/{id}").Handler(Middleware21{Handler: http.HandlerFunc(rest.UpdateClient)})
	router.Methods("PUT").Path("/api/v1/livres/up/{id}").Handler(Middleware22{Handler: http.HandlerFunc(rest.UpdateLivreById)})
	router.Methods("PUT").Path("/api/v1/emprunts/up/{id}").Handler(Middleware23{Handler: http.HandlerFunc(rest.UpdateEmpruntById)})
	router.Methods("PUT").Path("/api/v1/tenants/up/{id}").Handler(Middleware28{Handler: http.HandlerFunc(rest.UpdateTenantById)})

	router.Methods("POST").Path("/api/v1/users").Handler(Middleware24{Handler: http.HandlerFunc(rest.CreateUser)})
	router.Methods("GET").Path("/api/v1/users/{id}").Handler(Middleware1{handler: Middleware2{handler: http.HandlerFunc(rest.GetUserById)}})
	router.Methods("DELETE").Path("/api/v1/users/sup/{id}").Handler(Middleware26{Handler: http.HandlerFunc(rest.DeleteUser)})
	router.Methods("PUT").Path("/api/v1/users/up/{id}").Handler(Middleware3{Handler: http.HandlerFunc(rest.UpdateUser)})

	router.Methods("POST").Path("/api/v1/roles").HandlerFunc(rest.CreateRole)
	router.Methods("GET").Path("/api/v1/roles/{id}").HandlerFunc(rest.GetRoleById)
	router.Methods("DELETE").Path("/api/v1/roles/sup/{id}").HandlerFunc(rest.DeleteRoleById)
	router.Methods("PUT").Path("/api/v1/roles/up/{id}").HandlerFunc(rest.UpdateRoleById)
	//router.Methods("POST").Path("/api/v1/roles").HandlerFunc(rest.ListerRole)

	router.Methods("POST").Path("/api/v1/users/con").HandlerFunc(rest.ConnectUser)
	//router.Methods("POST").Path("/api/v1/users/con").HandlerFunc(rest.DeconnectUser)

	router.Methods("GET").Path("/api/v1/bibliotheques/lis").Handler(Middleware9{Handler: http.HandlerFunc(rest.ListerBibliotheque)})
	router.Methods("GET").Path("/api/v1/clients/lis").Handler(Middleware9{Handler: http.HandlerFunc(rest.ListerClient)})
	router.Methods("GET").Path("/api/v1/livres/lis").Handler(Middleware9{Handler: http.HandlerFunc(rest.ListerLivre)})
	router.Methods("GET").Path("/api/v1/emprunts/lis").Handler(Middleware9{Handler: http.HandlerFunc(rest.ListerEmprunt)})
	router.Methods("GET").Path("/api/v1/tenants/lis").Handler(Middleware9{Handler: http.HandlerFunc(rest.ListerTenant)})

	return router
}

type Middleware1 struct {
	handler http.Handler
}

func (m Middleware1) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	token := r.Header.Get("Authorization")
	fmt.Println(token)
	idUser := authentification.GetUser(token)

	if err := rest.AuthAndAuto(idUser); err != nil {
		return
	}
	fmt.Println("inside the middleware")
	m.handler.ServeHTTP(w, r)

	fmt.Println("inside the middleware, but after handler execution")
}

type Middleware2 struct {
	handler http.Handler
}

func (m Middleware2) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	token := r.Header.Get("Authorization")
	fmt.Println(token)
	idUser := authentification.GetUser(token)

	if err := rest.AuthAndAuto(idUser); err != nil {
		return
	}
	fmt.Println("inside the middleware")
	m.handler.ServeHTTP(w, r)

	fmt.Println("inside the middleware, but after handler execution")
}

type Middleware3 struct {
	Handler http.Handler
}

func (m Middleware3) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Hello je suis a l'ecoute")
	token := r.Header.Get("Authorization")
	fmt.Println(token)
	idUser := authentification.GetUser(token)

	if err := rest.AuthAndAuto(idUser); err != nil {
		return
	}
	fmt.Println("inside the middleware")
	m.Handler.ServeHTTP(w, r)

	fmt.Println("inside the middleware, but after handler execution")

}

type Middleware4 struct {
	Handler http.Handler
}

func (m Middleware4) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Hello je suis a l'ecoute")
	token := r.Header.Get("Authorization")
	fmt.Println(token)
	idUser := authentification.GetUser(token)

	if err := rest.AuthAndAuto(idUser); err != nil {
		return
	}
	fmt.Println("inside the middleware")
	m.Handler.ServeHTTP(w, r)

	fmt.Println("inside the middleware, but after handler execution")

}

type Middleware5 struct {
	Handler http.Handler
}

func (m Middleware5) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Hello je suis a l'ecoute")
	token := r.Header.Get("Authorization")
	fmt.Println(token)
	idUser := authentification.GetUser(token)

	if err := rest.AuthAndAuto(idUser); err != nil {
		return
	}
	fmt.Println("inside the middleware")
	m.Handler.ServeHTTP(w, r)

	fmt.Println("inside the middleware, but after handler execution")

}

type Middleware6 struct {
	Handler http.Handler
}

func (m Middleware6) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Hello je suis a l'ecoute")
	token := r.Header.Get("Authorization")
	fmt.Println(token)
	idUser := authentification.GetUser(token)

	if err := rest.AuthAndAuto(idUser); err != nil {
		return
	}
	fmt.Println("inside the middleware")
	m.Handler.ServeHTTP(w, r)

	fmt.Println("inside the middleware, but after handler execution")

}

type Middleware7 struct {
	Handler http.Handler
}

func (m Middleware7) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Hello je suis a l'ecoute")
	token := r.Header.Get("Authorization")
	fmt.Println(token)
	idUser := authentification.GetUser(token)

	if err := rest.AuthAndAuto(idUser); err != nil {
		return
	}
	fmt.Println("inside the middleware")
	m.Handler.ServeHTTP(w, r)

	fmt.Println("inside the middleware, but after handler execution")

}

type Middleware8 struct {
	Handler http.Handler
}

func (m Middleware8) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Hello je suis a l'ecoute")
	token := r.Header.Get("Authorization")
	fmt.Println(token)
	idUser := authentification.GetUser(token)

	if err := rest.AuthAndAuto(idUser); err != nil {
		return
	}
	fmt.Println("inside the middleware")
	m.Handler.ServeHTTP(w, r)

	fmt.Println("inside the middleware, but after handler execution")

}

type Middleware9 struct {
	Handler http.Handler
}

func (m Middleware9) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Hello je suis a l'ecoute")
	token := r.Header.Get("Authorization")
	fmt.Println(token)
	idUser := authentification.GetUser(token)

	if err := rest.AuthAndAuto(idUser); err != nil {
		return
	}
	fmt.Println("inside the middleware")
	m.Handler.ServeHTTP(w, r)

	fmt.Println("inside the middleware, but after handler execution")

}

type Middleware10 struct {
	Handler http.Handler
}

func (m Middleware10) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Hello je suis a l'ecoute")
	token := r.Header.Get("Authorization")
	fmt.Println(token)
	idUser := authentification.GetUser(token)

	if err := rest.AuthAndAuto(idUser); err != nil {
		return
	}
	fmt.Println("inside the middleware")
	m.Handler.ServeHTTP(w, r)

	fmt.Println("inside the middleware, but after handler execution")

}

type Middleware11 struct {
	Handler http.Handler
}

func (m Middleware11) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Hello je suis a l'ecoute")
	token := r.Header.Get("Authorization")
	fmt.Println(token)
	idUser := authentification.GetUser(token)

	if err := rest.AuthAndAuto(idUser); err != nil {
		return
	}
	fmt.Println("inside the middleware")
	m.Handler.ServeHTTP(w, r)

	fmt.Println("inside the middleware, but after handler execution")

}

type Middleware12 struct {
	Handler http.Handler
}

func (m Middleware12) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Hello je suis a l'ecoute")
	token := r.Header.Get("Authorization")
	fmt.Println(token)
	idUser := authentification.GetUser(token)

	if err := rest.AuthAndAuto(idUser); err != nil {
		return
	}
	fmt.Println("inside the middleware")
	m.Handler.ServeHTTP(w, r)

	fmt.Println("inside the middleware, but after handler execution")

}

type Middleware13 struct {
	Handler http.Handler
}

func (m Middleware13) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Hello je suis a l'ecoute")
	token := r.Header.Get("Authorization")
	fmt.Println(token)
	idUser := authentification.GetUser(token)

	if err := rest.AuthAndAuto(idUser); err != nil {
		return
	}
	fmt.Println("inside the middleware")
	m.Handler.ServeHTTP(w, r)

	fmt.Println("inside the middleware, but after handler execution")

}

type Middleware14 struct {
	Handler http.Handler
}

func (m Middleware14) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Hello je suis a l'ecoute")
	token := r.Header.Get("Authorization")
	fmt.Println(token)
	idUser := authentification.GetUser(token)

	if err := rest.AuthAndAuto(idUser); err != nil {
		return
	}
	fmt.Println("inside the middleware")
	m.Handler.ServeHTTP(w, r)

	fmt.Println("inside the middleware, but after handler execution")

}

type Middleware15 struct {
	Handler http.Handler
}

func (m Middleware15) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Hello je suis a l'ecoute")
	token := r.Header.Get("Authorization")
	fmt.Println(token)
	idUser := authentification.GetUser(token)

	if err := rest.AuthAndAuto(idUser); err != nil {
		return
	}
	fmt.Println("inside the middleware")
	m.Handler.ServeHTTP(w, r)

	fmt.Println("inside the middleware, but after handler execution")

}

type Middleware16 struct {
	Handler http.Handler
}

func (m Middleware16) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Hello je suis a l'ecoute")
	token := r.Header.Get("Authorization")
	fmt.Println(token)
	idUser := authentification.GetUser(token)

	if err := rest.AuthAndAuto(idUser); err != nil {
		return
	}
	fmt.Println("inside the middleware")
	m.Handler.ServeHTTP(w, r)

	fmt.Println("inside the middleware, but after handler execution")

}

type Middleware17 struct {
	Handler http.Handler
}

func (m Middleware17) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Hello je suis a l'ecoute")
	token := r.Header.Get("Authorization")
	fmt.Println(token)
	idUser := authentification.GetUser(token)

	if err := rest.AuthAndAuto(idUser); err != nil {
		return
	}
	fmt.Println("inside the middleware")
	m.Handler.ServeHTTP(w, r)

	fmt.Println("inside the middleware, but after handler execution")

}

type Middleware18 struct {
	Handler http.Handler
}

func (m Middleware18) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Hello je suis a l'ecoute")
	token := r.Header.Get("Authorization")
	fmt.Println(token)
	idUser := authentification.GetUser(token)

	if err := rest.AuthAndAuto(idUser); err != nil {
		return
	}
	fmt.Println("inside the middleware")
	m.Handler.ServeHTTP(w, r)

	fmt.Println("inside the middleware, but after handler execution")

}

type Middleware19 struct {
	Handler http.Handler
}

func (m Middleware19) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Hello je suis a l'ecoute")
	token := r.Header.Get("Authorization")
	fmt.Println(token)
	idUser := authentification.GetUser(token)

	if err := rest.AuthAndAuto(idUser); err != nil {
		return
	}
	fmt.Println("inside the middleware")
	m.Handler.ServeHTTP(w, r)

	fmt.Println("inside the middleware, but after handler execution")

}

type Middleware20 struct {
	Handler http.Handler
}

func (m Middleware20) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Hello je suis a l'ecoute")
	token := r.Header.Get("Authorization")
	fmt.Println(token)
	idUser := authentification.GetUser(token)

	if err := rest.AuthAndAuto(idUser); err != nil {
		return
	}
	fmt.Println("inside the middleware")
	m.Handler.ServeHTTP(w, r)

	fmt.Println("inside the middleware, but after handler execution")

}

type Middleware21 struct {
	Handler http.Handler
}

func (m Middleware21) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Hello je suis a l'ecoute")
	token := r.Header.Get("Authorization")
	fmt.Println(token)
	idUser := authentification.GetUser(token)

	if err := rest.AuthAndAuto(idUser); err != nil {
		return
	}
	fmt.Println("inside the middleware")
	m.Handler.ServeHTTP(w, r)

	fmt.Println("inside the middleware, but after handler execution")

}

type Middleware22 struct {
	Handler http.Handler
}

func (m Middleware22) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Hello je suis a l'ecoute")
	token := r.Header.Get("Authorization")
	fmt.Println(token)
	idUser := authentification.GetUser(token)

	if err := rest.AuthAndAuto(idUser); err != nil {
		return
	}
	fmt.Println("inside the middleware")
	m.Handler.ServeHTTP(w, r)

	fmt.Println("inside the middleware, but after handler execution")

}

type Middleware23 struct {
	Handler http.Handler
}

func (m Middleware23) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Hello je suis a l'ecoute")
	token := r.Header.Get("Authorization")
	fmt.Println(token)
	idUser := authentification.GetUser(token)

	if err := rest.AuthAndAuto(idUser); err != nil {
		return
	}
	fmt.Println("inside the middleware")
	m.Handler.ServeHTTP(w, r)

	fmt.Println("inside the middleware, but after handler execution")

}

type Middleware24 struct {
	Handler http.Handler
}

func (m Middleware24) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Hello je suis a l'ecoute")
	token := r.Header.Get("Authorization")
	fmt.Println(token)
	idUser := authentification.GetUser(token)

	if err := rest.AuthAndAuto(idUser); err != nil {
		return
	}
	fmt.Println("inside the middleware")
	m.Handler.ServeHTTP(w, r)

	fmt.Println("inside the middleware, but after handler execution")

}

type Middleware25 struct {
	Handler http.Handler
}

func (m Middleware25) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Hello je suis a l'ecoute")
	token := r.Header.Get("Authorization")
	fmt.Println(token)
	idUser := authentification.GetUser(token)

	if err := rest.AuthAndAuto(idUser); err != nil {
		return
	}
	fmt.Println("inside the middleware")
	m.Handler.ServeHTTP(w, r)

	fmt.Println("inside the middleware, but after handler execution")

}

type Middleware26 struct {
	Handler http.Handler
}

func (m Middleware26) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Hello je suis a l'ecoute")
	token := r.Header.Get("Authorization")
	fmt.Println(token)
	idUser := authentification.GetUser(token)

	if err := rest.AuthAndAuto(idUser); err != nil {
		return
	}
	fmt.Println("inside the middleware")
	m.Handler.ServeHTTP(w, r)

	fmt.Println("inside the middleware, but after handler execution")

}

type Middleware27 struct {
	Handler http.Handler
}

func (m Middleware27) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Hello je suis a l'ecoute")
	token := r.Header.Get("Authorization")
	fmt.Println(token)
	idUser := authentification.GetUser(token)

	if err := rest.AuthAndAuto(idUser); err != nil {
		return
	}
	fmt.Println("inside the middleware")
	m.Handler.ServeHTTP(w, r)

	fmt.Println("inside the middleware, but after handler execution")

}

type Middleware28 struct {
	Handler http.Handler
}

func (m Middleware28) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Hello je suis a l'ecoute")
	token := r.Header.Get("Authorization")
	fmt.Println(token)
	idUser := authentification.GetUser(token)

	if err := rest.AuthAndAuto(idUser); err != nil {
		return
	}
	fmt.Println("inside the middleware")
	m.Handler.ServeHTTP(w, r)

	fmt.Println("inside the middleware, but after handler execution")

}
