package errormsg

//struct que se utiliza para printear mensajes de error
//ej controller.go
//linea 31 : json.NewEncoder(response).Encode(errormsg.ServiceError{Message: "Error getting posts"})

type ServiceError struct {
	Message string `json : "message"`
}
