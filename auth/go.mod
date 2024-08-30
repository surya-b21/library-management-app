module github.com/surya-b21/library-management-app/auth

go 1.22.6

replace github.com/surya-b21/library-management-app/book => ../book

replace github.com/surya-b21/library-management-app/author => ../author

require (
	github.com/golang-jwt/jwt/v5 v5.2.1
	github.com/surya-b21/library-management-app/author v0.0.0-00010101000000-000000000000
	github.com/surya-b21/library-management-app/book v0.0.0-00010101000000-000000000000
	google.golang.org/protobuf v1.34.1
	gorm.io/driver/postgres v1.5.9
)

require (
	github.com/asaskevich/govalidator v0.0.0-20230301143203-a9d515a09cc2 // indirect
	github.com/go-openapi/errors v0.22.0 // indirect
	github.com/mitchellh/mapstructure v1.5.0 // indirect
	github.com/oklog/ulid v1.3.1 // indirect
	go.mongodb.org/mongo-driver v1.14.0 // indirect
	golang.org/x/net v0.26.0 // indirect
	golang.org/x/sys v0.21.0 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20240604185151-ef581f913117 // indirect
)

require (
	github.com/go-openapi/strfmt v0.23.0
	github.com/google/uuid v1.6.0
	github.com/jackc/pgpassfile v1.0.0 // indirect
	github.com/jackc/pgservicefile v0.0.0-20221227161230-091c0ba34f0a // indirect
	github.com/jackc/pgx/v5 v5.6.0 // indirect
	github.com/jackc/puddle/v2 v2.2.1 // indirect
	github.com/jinzhu/inflection v1.0.0 // indirect
	github.com/jinzhu/now v1.1.5 // indirect
	golang.org/x/crypto v0.24.0 // indirect
	golang.org/x/sync v0.8.0 // indirect
	golang.org/x/text v0.17.0 // indirect
	google.golang.org/grpc v1.66.0
	gorm.io/gorm v1.25.11
)
