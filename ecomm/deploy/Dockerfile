FROM golang:1.16-alpine

ARG APP_NAME
ENV APP_NAME ${APP_NAME}
ENV PATH="bin;%PATH%"
WORKDIR /app

COPY bin/ ./
COPY resources/config*.yml ./resources

CMD "${APP_NAME}"
