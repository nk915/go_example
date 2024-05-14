FROM busybox 

WORKDIR /app

COPY msa /app/msa
COPY agent /app/agent
COPY run.sh /app/run.sh

# 파일 실행 권한 부여
RUN chmod 755 /app/run.sh

ENTRYPOINT [ "/app/run.sh", "msa"]