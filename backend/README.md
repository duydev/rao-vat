## Auth docker repo

gcloud auth configure-docker asia-southeast1-docker.pkg.dev

## Build image

docker build -t rao-vat .

## Push to GCP

docker tag rao-vat asia-southeast1-docker.pkg.dev/rao-vat-428611/docker-repo/rao-vat
docker push asia-southeast1-docker.pkg.dev/rao-vat-428611/docker-repo/rao-vat

## Deploy cloud run

gcloud run deploy --image asia-southeast1-docker.pkg.dev/rao-vat-428611/docker-repo/rao-vat --platform managed --region asia-southeast1
