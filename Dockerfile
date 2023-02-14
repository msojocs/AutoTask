FROM openjdk:8-jre
COPY ./tmp/auto_task /app/auto_task
RUN chmod +x /app/auto_task
WORKDIR /app
CMD ["/app/auto_task"]