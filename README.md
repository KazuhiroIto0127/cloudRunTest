# cloudRunTest

## はじめに

## frontendのイメージを作成してcreate-react-app

 1. docker-compose build
 2. docker-compose run --rm frontend sh -c "npm install -g create-react-app && create-react-app . --template typescript"


## backendのimageをbuildして実行

docker build -t cloudrunexp .
docker run -p 8080:8080 -t cloudrunexp

curl http://localhost:8080
hello world!

## deploy

### backend

google cloud platformにデプロイする。
gcloudをインストールした後に下記実行
- gcloud init
- gcloud builds submit --tag gcr.io/PROJECT_ID/helloworld
- gcloud run deploy --image gcr.io/PROJECT_ID/helloworld
- gcloud run deploy

### frontend
firebase hosting

- firebase deploy
