FROM golang
WORKDIR /app
COPY go.mod ./
COPY go.sum ./
RUN go mod download
# COPY /*/*.go ./
COPY . .
RUN go build -o ./otus ./cmd
EXPOSE 8080
CMD [ "./otus" ]