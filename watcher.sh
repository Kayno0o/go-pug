#!/bin/bash

# server commands
SERVER_COMMAND="make go-prod"

server() {
  pkill -f "$SERVER_COMMAND"
  $SERVER_COMMAND &
}

make jade-build
make unocss
server

# Monitor for changes in the directories
inotifywait -m -r -e close_write,create,delete "./" |
while read -r directory event file
do
  path="$directory$file"
  if [[ $path != *"/app/jade/"* && $path == *".go" ]]; then
    echo ""
    echo "Restarting Go server..."
    server
  fi

  if [[ $path == *".jade" ]]; then
    make jade-build
    make unocss
    server
  fi

  if [[ $path == "unocss.config.ts" || $path == *"/client/css/"*".css" ]]; then
    make unocss
  fi
done
