# TopTen

Service has one REST endpoint (/top/ten/words) that accepts a piece of text and returns the top 10 most frequenty occurring words along with their count

# How to run

```
go run .
```

This will start the server and listen on the default port which is 8080

```
go run . -port <port>
```

To start server and listen on specified port