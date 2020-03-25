# explore-golang


-- Starter App
1. Setup a Go application with a web server. Use the standard net/http package https://golang.org/pkg/net/http/#example_ListenAndServe. Have a single main.go file.
2. Write a ping handler function that returns 200. Register this handler for “/ping” route with the server in the main function. Start the server in main with a hardcoded port. Test your server by calling this api on localhost using Postman.
-- Setup (Editor, Database)
1. Run the standard formatting and linting tools gofmt https://golang.org/cmd/gofmt/ and vet https://golang.org/cmd/vet/. Observe the errors and recommendations suggested and update your code accordingly.
2. Setup your editor to run gofmt automatically on save.
3. Use https://github.com/sirupsen/logrus to add info logs for every request and response.
4. Install a Postgres database using Docker. Use the official Postgres Docker image with tag 11. https://hub.docker.com/_/postgres. Once you’ve pulled and run this image, run the psql command on the container and explore Postgres functionality to create a database, create a table, insert into and select from the table.
-- Dependency management
https://blog.golang.org/using-go-modules
1. Initialise a Go module in your project and save the go.mod file created in your project.
2. Use go get to fetch any dependency packages you are using, for example logrus
3. Update your go.mod file to list the dependencies along with their versions, using `go mod tidy`.
-- Sample App
1. Write a function that accepts a UUID and returns the movie information for the given movie ID. Fetch movie information from Moviebuff using the Moviebuff SDK https://github.com/RealImage/moviebuff-sdk-go
2. Update the function to save the movie information in a Postgres database. Create a movie table in your database manually, with the movie ID as primary key. In the application, use the Postgres driver https://github.com/jackc/pgx to create a connection to the database. Whenever the movie function is called, create a new connection and execute a query to insert or update the movie in the database.
3. Update the function to run DB migrations programmatically instead of manually setting up the database. Use the https://github.com/golang-migrate/migrate library for Postgres https://github.com/golang-migrate/migrate/tree/master/database/postgres
4. Write a handler for the route /movies/{id} that accepts a movie ID in the request parameter and returns movie information as JSON. If the movie is available in your database, fetch the information from there. If the movie is not available in your database, fetch it from Moviebuff, save it locally and then serve.