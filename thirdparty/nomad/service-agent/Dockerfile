FROM busybox 

WORKDIR /app

COPY agent /app/agent
COPY service /app/service

ENTRYPOINT [ "/app/agent" ]