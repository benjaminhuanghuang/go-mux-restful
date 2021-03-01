
# Golang Crash Course (2020)
https://www.youtube.com/watch?v=kd-8mb6HfGA&list=PL3eAkoh7fypqUQUQPn-bXtfiYT_ZSVKmB&ab_channel=PragmaticReviews

## Code
https://gitlab.com/pragmaticreviews/golang-mux-api/tree/clean-arhitecture
https://gitlab.com/pragmaticreviews/golang-mux-api/tree/rest-api-testing



- Independent of framework
- Testable
- Independent of UI
- Independent of Database
- Independent of any external agency



## evolution
1. 有所的代码写在 route.go中
2. 引入 firebase
3. 支持多种database
4. 提取 service, controller layer

router -> controller -> service -> repository


## mock for service testing

## http test for controller testing

- Create new HTTP request
- Assing HTTP Request handler Function (controller function)
- Record the HTTP Response
- Dispatch the HTTP Request
- Assert HTTP status
- Decode HTTP response
- Assert HTTP response
- Cleanup database


