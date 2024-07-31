gcc -shared -o libexample.so example.c -fPIC


go clean -testcache; LD_LIBRARY_PATH=/home/nk915/github/kng/go_example/design/cgo/example_01/example/lib go test