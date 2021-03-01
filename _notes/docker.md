


1. init mod
```
  go mod init <project full name>
```


2. Create Dockerfile


3. Build docker image
```
  docker build -t go-rest-api-demo .
```

4. Run docker container
```
  docker run -p 8964:8964 go-rest-api-demo
```