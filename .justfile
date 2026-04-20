[linux, macos]
run *args:
  #! /bin/env sh

  go build -ldflags="-s -w" -o ./bin/mgr-gorg ./src/
  ./bin/mgr-gorg {{ args }}
