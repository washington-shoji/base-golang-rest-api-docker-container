# Developer's Guide and Notes

## Adding Go to WSL (Windows Subsystem for Linux)

### Download and Install Golang from official repo with wget
Check the offial repo for latest/version 
https://go.dev/dl/

#### Run the following commands

wget https://go.dev/dl/go1.22.3.linux-amd64.tar.gz -O go.tar.gz

#### Extract the Go package

sudo tar -xzvf go.tar.gz -C /usr/local

#### Set the directory and paths (using bash)

mkdir go

nano ~/.bashrc

#### Add the following commands at the end of the bash script then save (ctrl + O) then enter and exit (ctrl + X)

export GOROOT=/usr/local/go
export GOPATH=$HOME/go
export PATH=$PATH:/usr/local/go/bin

#### Check if Go is working, if it has the right version

go version

## Go libs for this project

### Intallation commands

go get github.com/labstack/echo/v4
go get github.com/labstack/echo/v4/middleware
go get github.com/lib/pq
go get github.com/joho/godotenv
go get github.com/google/uuid

### Adding to GitHub repository


echo "# project-name" >> README.md
git init
git add README.md
git commit -m "first commit"
git branch -M main
git remote add origin project-url
git push -u origin main

## Docker image creation reference

https://docs.docker.com/language/golang/build-images/
https://snyk.io/blog/containerizing-go-applications-with-docker/

### Build Docker image locally

sudo docker build -t <Docker image name>:<TAG> -f Dockerfile .

sudo docker build -t base-golang-rest-api-docker-container:multistage -f Dockerfile .

### Check the list of Docker images 

sudo docker image ls

### Run docker image with TAG locally

sudo docker run --read-only <Docker Image>:<TAG>
