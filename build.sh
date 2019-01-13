#! /bin/bash

# Build web UI

cd ~/gopath/src/rushflow/web

go install

cp ~/gopath/bin/web ~/gopath/bin/video_server_web_ui/web

cp -R ~/gopath/src/rushflow/templates ~/gopath/bin/video_server_web_ui/

