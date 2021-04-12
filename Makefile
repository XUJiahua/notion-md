fmt:
	go fmt ./...
# export TOKEN=<find your notion token from browser>
run:fmt
	go run . -i 8ae7005e8b154431940ab03c0a2ef08a -t ${TOKEN} -o /Users/jiahua/hugo-blogger/content/posts -p ../../images/
