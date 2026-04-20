package main

import (
  "sync"
  "fmt"
)

type Chapter struct {
  Manga string
  Chapter string
  Total_Pages int
  Urls []string
}

func (ch *Chapter) Download(wg *sync.WaitGroup) {
  defer wg.Done()

  var page_wg sync.WaitGroup

  for i, url := range ch.Urls {
    page_wg.Add(1)
    go func() {
      err := get_page(&page_wg, url, get_save_path(ch.Manga, ch.Chapter, i + 1))
      if err != nil {
        fmt.Printf("Error while getting page: %v\n", err.Error())
      }
    }()
  }

  page_wg.Wait()
  fmt.Print("Downloaded.. ")

  err := zip_folder(ch)
  if err != nil {
    fmt.Printf("Error while zipping chapter: %v\n", err.Error())
  }
  fmt.Print("Zipped..")
  fmt.Print("\n\n")
}
