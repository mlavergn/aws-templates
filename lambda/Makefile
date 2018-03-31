###############################################
#
# Makefile
#
###############################################

SHELL := /bin/bash
PATH := $(PATH):~/Library/Python/3.5/bin

AWS_ACCOUNT := 0
AWS_ROLE := lamba-execute-role

all:
	GOOS=linux go build -o main main.go
	zip deployment.zip main

setup:
	go get github.com/aws/aws-lambda-go/lambda
	pip3 install awscli --upgrade --user

console:
	echo "AWS -> IAM -> Role -> AWSLambdaRole"
	open https://console.aws.amazon.com/

test:
	echo "AWS -> Lambda -> Demo -> Test"
	open https://console.aws.amazon.com/

awsver:
	aws --version

awsconfig:
	echo "AWS -> IAM -> Users -> <username or Add user> -> Securitry Credentials -> Access Keys"
	aws configure

awscreate:
	aws lambda create-function \
		--region us-east-1 \
		--function-name Demo \
		--zip-file fileb://./deployment.zip \
		--runtime go1.x \
		--role arn:aws:iam::$(AWS_ACCOUNT):role/$(AWS_ROLE) \
		--handler main

awslist:
	aws lambda list-functions

awsupdate:
	aws lambda update-function-code \
		--function-name Demo \
		--zip-file fileb://./deployment.zip