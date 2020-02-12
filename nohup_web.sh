cp -R ./templates ./bin
mkdir ./bin/videos
cd bin

nohup ./web >nohup.out 2>&1 &

echo "deploy finished"