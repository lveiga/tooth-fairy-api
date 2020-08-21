#!/usr/bin/env bash
log() {
  echo "STARTSCRIPT: $1"
}
buildServer() {
  log "Building server binary"
  go build -gcflags "all=-N -l" -o /server main.go
}
runServer() {
  log "Run server"
  log "Killing old server"
  killall dlv
  killall server
  log "Run in debug mode"
  dlv --listen=:2345 --headless=true --api-version=2 --accept-multiclient exec /server &
  inotifywait -e MODIFY /tmp/dlv_log/output.log &>/dev/null
  echo Delve PID: $(pidof dlv), Server PID: $(pidof server)
  pidof dlv > /tmp/dlv.pid
  pidof server > /tmp/server.pid
}
rerunServer() {
  log "Rerun server"
  buildServer
  runServer
}
lockBuild() {
  # check lock file existence
  if [ -f /tmp/server.lock ]
  then
    # waiting for the file to delete
    inotifywait -e DELETE /tmp/server.lock
  fi
  touch /tmp/server.lock
}
unlockBuild() {
  # remove lock file
  rm -f /tmp/server.lock
}
liveReloading() {
  log "Run liveReloading"
  inotifywait -e "MODIFY,DELETE,MOVED_TO,MOVED_FROM" -m -r --include '.go$' . | (
    # read changes from inotify, batch results to a second (read -t 1)
    while true; do
      read path action file
      ext=${file: -3}
      if [[ "$ext" == ".go" ]]; then
        echo "$file"
      fi
    done
  ) | (
    WAITING=""
    while true; do
      file=""
      read -t 1 file
      if test -z "$file"; then
        if test ! -z "$WAITING"; then
          echo "CHANGED"
          WAITING=""
        fi
      else
        log "File ${file} changed" >>/tmp/filechanges.log
        WAITING=1
      fi
    done
  ) | (
    # read statement release when some file has been changed
    while true; do
      read TMP
      log "File Changed. Reloading..."
      rerunServer
    done
  )
}
initializeFileChangeLogger() {
  echo "" > /tmp/filechanges.log
  tail -f /tmp/filechanges.log &
}
main() {
  initializeFileChangeLogger
  buildServer
  runServer
  liveReloading
}
main