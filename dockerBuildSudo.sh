clear
sudo docker image build -t go-web-server .
sudo docker container run -p 8080:8080 -it go-web-server
