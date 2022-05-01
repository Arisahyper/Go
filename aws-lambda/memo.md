
go get -u "github.com/aws/aws-lambda-go/lambda"                                                                 
GOOS=linux go build -o main main.go
zip function.zip main
