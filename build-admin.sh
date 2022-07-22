go build -o admin-build main.go
rsync admin-build ubuntu@13.228.244.196:/home/ubuntu
rm admin-build