
# Golang Crash Course (2020)
https://www.youtube.com/watch?v=kd-8mb6HfGA&list=PL3eAkoh7fypqUQUQPn-bXtfiYT_ZSVKmB&ab_channel=PragmaticReviews


https://gitlab.com/pragmaticreviews/golang-mux-api/tree/clean-arhitecture



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
