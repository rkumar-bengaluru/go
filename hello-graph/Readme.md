
# project initialization
go mod init github.com/rkumar-bengaluru/hello-graph

# intialize graphql
go run github.com/99designs/gqlgen init

# everytime schema changes.
go run github.com/99designs/gqlgen generate

# run the server
go run server.go