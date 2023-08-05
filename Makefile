c:	
	# build c
	@gcc -c foo.c 
	@ar rcs libfoo.a foo.o

cgo:
	@make c
	@rm -rf ffi
	# build cgo
	@go build -v -buildmode=c-shared -o ffi/lib.a *.go

ffigen:
	@make cgo
	# build ffigen
	@dart run ffigen

rundart:
	@make ffigen

	# run dart
	@dart run bin/ffitest.dart