package main

import (
  "time"
  "net/http"
)

const URLS_REGEX = `https://www.mangaread.org/.*/data/[^"]*`
const BASE_URL = "https://www.mangaread.org/manga/"
const CH_PREFIX = "/chapter-"
const TIMEOUT = time.Minute * 1
const BASE_HTML_DIR = ""
const BASE_HTML_PATTERN = "mgr-gorg_base_html_*"
var client = http.Client { Timeout: TIMEOUT }
