hello:
	echo "Hello"

build:
	go build -o main main.go

run:
	go run main.go

dockerLabBuild:
	docker build --tag lab4 .

dockerLabRun:
	docker run -it -p 8080:8080 lab4
