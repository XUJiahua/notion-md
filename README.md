
# notion-md

convert notion pages into markdowns (and download images).

markdown is hugo compatible.

## usage

```
$ go run main.go -h
convert notion pages into markdowns

Usage:
  notion-md [flags]

Flags:
      --config string   config file (default is $HOME/.notion-md.yaml)
  -h, --help            help for notion-md
  -i, --id string       id of root page which contains subpages
  -o, --output string   output directory of markdowns and images (default "./output")
  -t, --token string    notion token

```

example

```
go run . -i 8ae7005e8b154431940ab03c0a2ef08a -t ${TOKEN}
```



credit: 

https://github.com/kjk/notionapi
