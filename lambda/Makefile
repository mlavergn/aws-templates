###############################################
#
# Makefile
#
###############################################

SHELL := /bin/bash
PATH := $(PATH):~/Library/Python/3.5/bin
GOOS := linux 

AWS_ACCOUNT := 0
AWS_ROLE := lamba-execute-role

setup:
	go get github.com/aws/aws-lambda-go/lambda
	pip3 install awscli --upgrade --user

build:
	go build -o main main.go
	zip deployment.zip main

awsconsole:
	echo "AWS -> IAM -> Role -> AWSLambdaRole"
	open https://console.aws.amazon.com/

awsver:
	aws --version

awssetup:
	aws configure

aws:
	aws lambda create-function \
		--region us-east-1 \
		--function-name DiscoverMovies \
		--zip-file fileb://./deployment.zip \
		--runtime go1.x \
		--role arn:aws:iam::$(AWS_ACCOUNT):role/$(AWS_ROLE) \
		--handler main
