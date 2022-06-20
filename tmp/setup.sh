go mod vendor
./hack/update-codegen.sh
minikube start --image-mirror-country='cn'