package main

import (
	"io"
	"log"
	"os"
	"strings"
)

func main() {
	name := "agni"

	str := `<!DOCTYPE html>
<html lang="en" dir="ltr">
  <head>
    <meta charset="utf-8">
    <title>Hello World</title>
  </head>
  <body>
    <h1>` + name + `<h1>
  </body>
</html>`

	f, err := os.Create("index.html")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	io.Copy(f, strings.NewReader(str))
}
