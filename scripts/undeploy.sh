source $(dirname "$0")/services.sh
for service_name in "${services[@]}"; do
  kubectl delete -f k8s-deployments/$service_name.yml
done

kubectl delete -f k8s-deployments/ingress.yml