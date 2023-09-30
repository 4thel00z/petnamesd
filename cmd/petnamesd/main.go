package main

import (
	"math/rand"
	"os"
	"time"

	petname "github.com/dustinkirkland/golang-petname"
	"github.com/kataras/iris/v12"
)

func init() {
	rand.Seed(time.Now().UTC().UnixNano())
}

// getEnv reads an environment variable and returns it if it exists; otherwise, it returns the provided default value.
func getEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}

func main() {
	// Initialize a new Iris application instance.
	app := iris.New()

	// Register a GET route to generate and return a pet name.
	app.Get("/petnames", func(ctx iris.Context) {
		// Read words and separator from query parameters, if they are not provided, use default values.
		words := ctx.URLParamIntDefault("words", 2)
		separator := ctx.URLParamDefault("separator", "-")

		// Generate a pet name.
		petName := petname.Generate(words, separator)

		// Create a map to hold the response.
		response := map[string]string{
			"pet_name": petName,
		}

		// Marshal the map to JSON and return it as the response.
		ctx.JSON(response)
	})

	// Read the host and port from the environment variables or use the defaults.
	host := getEnv("HOST", "")
	port := getEnv("PORT", "8080")
	address := host + ":" + port

	// Start the Iris web server on the specified address.
	app.Run(iris.Addr(address))
}
