FROM golang:1.12.0

ARG app_env
ENV APP_ENV $app_env

COPY ./ /go/src/github.com/unicef/dignity-platform/backend
WORKDIR /go/src/github.com/unicef/dignity-platform/backend

RUN go get ./
RUN go build

CMD if [ ${APP_ENV} = production ]; \
	then \
	app; \
	else \
	go get github.com/pilu/fresh && \
	fresh; \
	fi

EXPOSE 3000
