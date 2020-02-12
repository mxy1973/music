cp -R ./templates ./bin/

mkdir ./bin/videos

cd bin

nohup ./api >nohup.out 2>&1 &
nohup ./scheduler  >nohup.out 2>&1 &
nohup ./streamserver >nohup.out 2&1 &
nohup ./web >nohup.out 2>&1 &

echo "deploy finished"