FROM php:8.1-fpm
COPY --from=datadog/serverless-init /datadog-init /app/datadog-init
RUN docker-php-ext-install sockets

WORKDIR /app
COPY app.php /app/

ENV DD_SERVICE=datadog-demo-run-php
ENV DD_ENV=datadog-demo
ENV DD_VERSION=1
ENV DD_TRACE_CLI_ENABLED=1

ENTRYPOINT ["/app/datadog-init"]
CMD ["php", "app.php"]
