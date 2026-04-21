package main

import (
	ex "github.com/coregx/coregex"

	"context"
	"fmt"
	"os"
	"path"
	"sync"
  "strings"
  "slices"
)

func main() {
  var chs []string

  args := os.Args
  if len(args) < 2 {
    fmt.Printf("USAGE: %v MANAGA [CHAPTERS]...\n", args[0])
    os.Exit(0)
  } else if len(args) > 2 {
    chs = args[2:]
  }

  ctx := context.Background()
  manga := args[1]

  urls_regex, err := ex.CompilePOSIX(URLS_REGEX)
  if err != nil {
    fmt.Printf("Invalid urls regex: %v\n", urls_regex)
    os.Exit(1)
  }

  if len(args) == 2 {
    base_regex, err := ex.CompilePOSIX(BASE_REGEX(manga))
    if err != nil {
      fmt.Printf("Invalid urls regex: %v\n", urls_regex)
      os.Exit(1)
    }

    html, err := get_html(ctx, manga, "")
    if err != nil {
      fmt.Printf("Error while getting base html: %v\n", err.Error())
      os.Exit(1)
    }
    urls, err := get_urls(html, base_regex)
    if err != nil {
      fmt.Printf("Error while getting urls: %v\n", err.Error())
      os.Exit(1)
    } else {
      fmt.Printf("Found %v chapters\n", len(urls))
    }

    for _, url := range urls {
      chs = append(chs, strings.TrimPrefix(path.Base(url), CH_PREFIX))
    }
    slices.Reverse(chs)
  }

  var wg sync.WaitGroup

  os.MkdirAll(path.Join(manga, "archive"), 0755)

  for _, ch := range chs {
    fmt.Printf("Chapter: %v\n", ch)
    os.MkdirAll(path.Join(manga, ch), 0755)

    html, err := get_html(ctx, manga, ch)
    if err != nil {
      fmt.Printf("Error while getting chapter html: %v\n", err.Error())
      continue
    }

    fmt.Print("Started.. ")

    urls, err := get_urls(html, urls_regex)
    if err != nil {
      fmt.Printf("Error while getting urls: %v\n", err.Error())
      continue
    } else {
      fmt.Printf("Found %v pages.. ", len(urls))
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
