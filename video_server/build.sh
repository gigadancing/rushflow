#! /bin/bash

# Build web UI

cd ~/gopath/src/rushflow/video_server/web

go install

cp ~/gopath/bin/web ~/gopath/bin/video_server_web_ui/web

cp -R ~/gopath/src/video_server/templates ~/gopath/bin/video_server_web_ui/

