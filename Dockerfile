FROM debian:10

ENV APP_HOME /srv
WORKDIR $APP_HOME

COPY engra $APP_HOME
COPY dict $APP_HOME/dict

CMD ["./engra", "serve"]