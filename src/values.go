package main

import (
  "time"
  "fmt"
  "net/http"
)

const URLS_REGEX = `https://www.mangaread.org/.*/data/[^"]*`
const BASE_URL = "https://www.mangaread.org/manga/"
const CH_PREFIX = "chapter-"
const TIMEOUT = time.Minute * 1
const HTML_DIR = ""
const CH_HTML_PATTERN = "mgr-gorg_ch_html_*"
const BASE_HTML_PATTERN = "mgr-gorg_base_html_*"
var client = http.Client { Timeout: TIMEOUT }

func BASE_REGEX(manga string) string { return fmt.Sprintf(`https://www.mangaread.org/manga/%v/chapter-[^"]*`, manga) }
