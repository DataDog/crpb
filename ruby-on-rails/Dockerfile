FROM ruby:3.1.2
RUN apt-get update && apt-get install -y nodejs
WORKDIR /app
COPY Gemfile .
COPY Gemfile.lock .
RUN bundle install
COPY . .
EXPOSE 3000

COPY --from=datadog/serverless-init:beta8 /datadog-init /app/datadog-init

ENV DD_SERVICE=datadog-demo-run-ruby-on-rails
ENV DD_ENV=datadog-demo
ENV DD_VERSION=1
ENV PORT=3000

ENTRYPOINT ["/app/datadog-init"]
CMD ["rails", "server", "-b", "0.0.0.0"]