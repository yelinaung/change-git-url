change-git-url
=============

SSH to HTTP url, vice versa. Tested with GitHub and BitBucket urls.

How to Use
----------

`change-git-url` can be easily installed as an executable. Download the latest [compiled binaries](https://github.com/yelinaung/change-git-url/releases/tag/v0.0.1) and put it anywhere in your executable path.

Or, if you've done Go development before and your $GOPATH/bin directory is already in your PATH:

```go
$ go get github.com/yelinaung/change-git-url
$ cd $GOPATH/src/github.com/yelinaung/change-git-url
$ go run change-git-url.go git@github.com:yelinaung/change-git-url.git
https://github.com/yelinaung/change-git-url.git
$ go run change-git-url.go https://github.com/yelinaung/change-git-url.git
git@github.com:yelinaung/change-git-url.git
```

License
-------
MIT
