```
minikube -p minikube docker-env | Invoke-Expression
docker build . -t aoc2021/foo
kubectl apply -f deployment.yml
minikube service foo
```