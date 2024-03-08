#!/bin/bash

# server commands
SERVER_COMMAND="make go-build"

server() {
  pkill -f "$SERVER_COMMAND"
  $SERVER_COMMAND &
}

make jade-build
server

# Monitor for changes in the directories
inotifywait -m -r -e close_write,create,delete "./" "../client" |
while read -r directory event file
do
  path="$directory$file"
  if [[ $path != *"/jade/"* && $path == *".go" ]]; then
    echo ""
    echo "Restarting Go server..."
    server
  fi

  if [[ $path == *".pug" ]]; then
    make jade-build
    server
  fi
done
