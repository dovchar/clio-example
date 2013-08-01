package main

import (
    . "github.com/pallada/clio/core"
    "./app/routes"
)

func main() {
    routes.BooksRoutes()
    Run (4567)
}
