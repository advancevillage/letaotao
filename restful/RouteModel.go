package restful

import hr "github.com/julienschmidt/httprouter"

type Route struct {
	Method 		string 		`json:"method"`
	Path 		string 		`json:"path"`
	HandlerFunc hr.Handle	`json:"handlerfunc"`
}










