docker build -t chat_stranger .

docker run -it \
    --rm \
    --publish 8080:8080 \
    --env-file .env \
    --name chat_stranger \
    chat_stranger
