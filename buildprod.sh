cd ~/go_work/src/video_server/api
env GOOS=linux GOARCH=amd64 go build -o ../bin/api

cd ~/go_work/src/video_server/scheduler
env GOOS=linux GOARCH=amd64 go build -o ../bin/scheduler

cd ~/go_work/src/video_server/streamserver
env GOOS=linux GOARCH=amd64 go build -o ../bin/streamserver

cd ~/go_work/src/video_server/web
env GOOS=linux GOARCH=amd64 go build -o ../bin/web

