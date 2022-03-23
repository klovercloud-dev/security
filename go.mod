module github.com/klovercloud-ci-cd/security

go 1.16

require (
	github.com/google/uuid v1.1.2
	github.com/joho/godotenv v1.3.0
	github.com/labstack/echo/v4 v4.6.1
	github.com/stretchr/testify v1.7.1
	go.mongodb.org/mongo-driver v1.8.1
)

require github.com/google/go-cmp v0.5.5 // indirect

require (
	github.com/golang-jwt/jwt v3.2.2+incompatible
	github.com/labstack/gommon v0.3.1 // indirect
	github.com/swaggo/echo-swagger v1.2.0
	github.com/swaggo/swag v1.7.8
	golang.org/x/crypto v0.0.0-20211108221036-ceb1ce70b4fa
)
