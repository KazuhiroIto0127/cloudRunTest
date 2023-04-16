# cloudRunTest

## imageをbuildして実行

docker build -t cloudrunexp .
docker run -p 8080:8080 -t cloudrunexp

curl http://localhost:8080
hello world!

## backend
google cloud platformにデプロイする。
gcloudをインストールした後に下記実行
- gcloud init
- gcloud builds submit --tag gcr.io/PROJECT_ID/helloworld
- gcloud run deploy --image gcr.io/PROJECT_ID/helloworld
- gcloud run deploy

## frontend
firebase hosting


