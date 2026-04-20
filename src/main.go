package main

import (
	"context"
	"fmt"
	"os"
	"path"
	"sync"
)

func main() {
  args := os.Args
  if len(args) < 3 {
    fmt.Printf("USAGE: %v [manga] [chapters...]\n", args[0])
    os.Exit(0)
  }

  ctx := context.Background()
  manga := args[1]
  chs := args[2:]
  regex := compile_regex(URLS_REGEX)

  var wg sync.WaitGroup

  os.MkdirAll(path.Join(manga, "archive"), 0755)

  for _, ch := range chs {
    fmt.Printf("Chapter: %v\n", ch)
    os.MkdirAll(path.Join(manga, ch), 0755)

    html, err := get_base_html(ctx, manga, ch)
    if err != nil {
      fmt.Printf("Error while getting base html: %v\n", err.Error())
      continue
    }

    fmt.Print("Started.. ")

    urls, err := get_urls(html, regex)
    if err != nil {
      fmt.Printf("Error while getting urls: %v\n", err.Error())
      continue
    }

    chapter := &Chapter {
      Chapter: ch,
      Manga: manga,
      Total_Pages: len(urls),
      Urls: urls,
    }

    wg.Add(1)
    chapter.Download(&wg)
  }

  wg.Wait()
}
