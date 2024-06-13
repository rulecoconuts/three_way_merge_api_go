docker build -t 3_way_merge .
docker run -p 3000:3000 -p 8000:8000 -t -i 3_way_merge