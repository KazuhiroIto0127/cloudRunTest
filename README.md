# cloudRunTest

## imageをbuildして実行

docker build -t cloudrunexp .
docker run -e "PORT=3000" -p 3000:3000 -t cloudrunexp

## google cloud platformにデプロイする。
gcloudをインストールした後に下記実行
- gcloud init
- gcloud run deploy
