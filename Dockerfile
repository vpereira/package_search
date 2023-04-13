FROM opensuse/tumbleweed:latest

RUN zypper --non-interactive install go1.19 git gcc  make
# Set the working directory to /go/src/app
WORKDIR /go/src/app

# Copy the current directory contents into the container at /go/src/app
COPY . .

RUN go mod tidy

CMD ["tail", "-f"]
