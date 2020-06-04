package design

import (
	. "github.com/goadesign/goa/design/apidsl"
)

var _ = API("RESTAPI", func() {
	Title("GOA DESIGN")
	Scheme("http", "https")
	Host("localhost:8080")
	BasePath("/api")
})
