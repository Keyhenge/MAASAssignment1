build:
	@rm -f *.zip
	@rm -f main
	@env GOOS=linux go build -v -tags netgo -installsuffix netgo -o main src/main.go
	@zip lambda.zip main

deploy:
	@rm -f *.zip
	@rm -f main
	@env GOOS=linux go build -v -tags netgo -installsuffix netgo -o main src/main.go
	@zip lambda.zip main
	@aws lambda update-function-code --function-name basicLambda --zip-file 'fileb://lambda.zip'
	#@aws apigateway put-rest-api --rest-api-id 6h22lym61k --body 'file://./src/gateway.yaml'
