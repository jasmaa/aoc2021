FROM ruby:2.7-alpine

WORKDIR /app
COPY . .

RUN bundle install
ENV APP_ENV "production"

ENTRYPOINT ["ruby", "server.rb"]