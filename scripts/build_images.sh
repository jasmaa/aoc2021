source $(dirname "$0")/services.sh
for service_name in "${services[@]}"; do
  docker build ./$service_name -t aoc2021/$service_name
done