build:
    @rm *.zip
    @rm lambda
    @env GOOS=linux GOARCH=amd64 go build -v src/lambda.go
    @zip lambda.zip lambda

#deploy:
#    @rm *.zip
#    @rm lambda
#    @env GOOS=linux GOARCH=amd64 go build -v src/lambda.go
#    @zip lambda.zip lambda
#    @aws lambda update-function-code --function-name basicLambda --zip-file 'file://lambda.zip'
#    @aws apigateway put-rest-api --rest-api-id vmyc8jslwb --body 'file://src/swagger.yaml'
