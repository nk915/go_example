# 모든 테스트 실행
```
$ go test
```

# *_test.go 파일 테스트 실행
* 옵션: -v (상세한 결과 확인)
```
$ go test *_test.go -v
```

# 특정 테스트 함수(TestFunctionName) 실행
```
$ go test -run TestFunctionName -v
```