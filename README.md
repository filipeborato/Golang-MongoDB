# Golang-MongoDB

### Instalation and Testing

Runing project:
```
go mod vendor
```
```
go run main.go
```
Create database in mongo with collection:
```
db = sport
collection = news_information
```

Testing Route:

```
http://localhost:8080/v1/apipopulate-test
```

Get Many News:

```
http://localhost:8080/v1/api/news
```

Get One News:

```
http://localhost:8080/v1/api/news/:id
```

### Create the .env in the project root with the key:

```
MONGODB_URI="mongodb://{host}:{port}"
```

#### Run Populate ScheduleJob:


```
 go run main.go populate-news
```

