
# notion-md

[中文说明](https://xujiahua.github.io/posts/notion-markdown-hugo-78139407-85c5-4a75-b288-b54d5c3df34b/)

convert notion pages into markdowns (and download images).

markdown is hugo compatible.

## install

```
go install github.com/xujiahua/notion-md
```

## usage

```
$ notion-md -h
convert notion pages into markdowns

Usage:
  notion-md [flags]

Flags:
      --config string   config file (default is $HOME/.notion-md.yaml)
  -h, --help            help for notion-md
  -i, --id string       id of root page which contains subpages
  -v, --listview        use listview hold blogs, contain category, tags
  -o, --output string   output directory of markdowns and images (default "./output")
  -p, --prefix string   hugo markdown image prefix (relative path to image folder)
  -t, --token string    notion token

```

### example

```
notion-md -i 8ae7005e8b154431940ab03c0a2ef08a -t ${TOKEN}
```

### example listview

```
notion-md -i cab2ea6d530341769e5dc9a269a1097e -t ${TOKEN} -o /Users/jiahua/hugo-blogger/content/posts -p ../../images/ -v
```

credit: 

https://github.com/kjk/notionapi
