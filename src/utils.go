package main

import (
	ex "github.com/coregx/coregex"

	"context"
	"sync"
	"fmt"
	"io"
	"os"
	"path"
)

func compile_regex(URLS_REGEX string) *ex.Regex {
  regex, err := ex.Compile(URLS_REGEX)
  if err != nil {
    fmt.Printf("Error while compiling regex: %v\n", err.Error())
    os.Exit(1)
  }

  return regex
}

func get_base_html(ctx context.Context, manga string, ch string) (*os.File, error) {
  ch_url := BASE_URL + manga + CH_PREFIX + ch + "/"

  response, err := client.Get(ch_url)
  if err != nil { return nil, err }
  defer response.Body.Close()

  body, err := io.ReadAll(response.Body)
  if err != nil { return nil, err }

  file, err := os.CreateTemp(BASE_HTML_DIR, BASE_HTML_PATTERN)
  if err != nil { return nil, err }

  _, err = file.Write(body)
  if err != nil {
    file.Close()
    return nil, err
  }

  _, err = file.Seek(0, 0)
  if err != nil {
    file.Close()
    return nil, err
  }

  return file, nil
}

func get_urls(file *os.File, regex *ex.Regex) ([]string, error) {
  defer file.Close()

  content, err := io.ReadAll(file)
  if err != nil { return nil, err }

  urls := regex.FindAllString(string(content), -1)

  return urls, nil
}

func get_page(wg *sync.WaitGroup, url string, save_path string) error {
  defer wg.Done()

  response, err := client.Get(url)
  if err != nil { return err }
  defer response.Body.Close()

  data, err := io.ReadAll(response.Body)
  if err != nil { return err }

  return os.WriteFile(save_path, data, 0644)
}

func get_save_path(manga string, ch string, page int) string {
  return path.Join(manga, ch, fmt.Sprintf("%03d.jpg", page))
}
