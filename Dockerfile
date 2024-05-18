# syntax=docker/dockerfile:1

##
## STEP 1 - BUILD
##

# specify the base image to  be used for the application, alpine or ubuntu
FROM golang:1.22.3-alpine3.19  AS build

# this instruction sets the USERNAME and PASSWORD arguments with the ARG keyword and then creates a user with the RUN adduser instruction.
#RUN useradd -u 1001 -m iamuser

ENV GOPATH /

# create a working directory inside the image
WORKDIR /app

# copy Go modules and dependencies to image
COPY go.mod ./

# download Go modules and dependencies
RUN go mod download
RUN go mod tidy

# copy directory files i.e all files ending with .go
COPY *.go ./
COPY . .

# if using a .env file but not recommended 
# COPY .env .

# compile application
# RUN go build -o /build

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o go-api /app/main.go

##
## STEP 2 - DEPLOY
##
FROM scratch

WORKDIR /app

COPY --from=build /app/go-api .

# if using a .env file but not recommended
# COPY .env /app

# COPY --from=build /etc/passwd /etc/passwd

# USER 1001

# tells Docker that the container listens on specified network ports at runtime
EXPOSE 8080

# command to be used to execute when the image is used to start a container
ENTRYPOINT [ "/app/go-api" ]