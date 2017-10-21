package main

import (
	"gitlab.com/vwxyzjn/fucommencement-backend/backend"
)

func main() {
	backend.Migrate()
	backend.Setup()
	// r.Run() // listen and serve on 0.0.0.0:8080
}
