docker build -t gin-project:1.0 -f ./Dockerfile .

docker run -d -p 9001:8888 --name gin_project_01 -t gin-project:1.0