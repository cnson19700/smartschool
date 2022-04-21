go build -o api-build api/main.go
rsync api-build ubuntu@13.228.244.196:/home/ubuntu
rsync -r resources/ ubuntu@13.228.244.196:/home/ubuntu/resources
rm api-build