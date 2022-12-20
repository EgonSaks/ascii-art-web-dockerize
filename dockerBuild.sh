clear
docker image build -t go-web-server .
docker container run -p 8080:8080 -it go-web-server
