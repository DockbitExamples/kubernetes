vi source/1.0/app.go
GOOS=linux GOARCH=amd64 go build -tags netgo -o app
export app_version=1.0
docker build --build-arg version=$app_version -t harbor.jpe1-apv1-prod.r-local.net/travel-poc/canary:1.1 .
