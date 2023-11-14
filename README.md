# be-tesis





Instalar Docker

Bajarse la imagen de postgres de Docker

docker run --name TesisDB -e POSTGRES_USER=postgres -e POSTGRES_PASSWORD=postgres -e POSTGRES_DB=TesisDB -d -p 5432:5432 postgres



docker stop TesisDB 

docker rm TesisDB