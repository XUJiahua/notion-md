fmt:
	go fmt ./...
# export TOKEN=<find your notion token from browser>
run:fmt
	go run . -i cab2ea6d530341769e5dc9a269a1097e -t ${TOKEN} -o /Users/jiahua/hugo-blogger/content/posts -p ../../images/ -v
