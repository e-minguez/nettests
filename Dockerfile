FROM fedora/python:latest

USER root
RUN curl -Lo /usr/local/bin/speedtest-cli https://raw.githubusercontent.com/sivel/speedtest-cli/master/speedtest_cli.py && \
    chmod a+x /usr/local/bin/speedtest-cli && \
    mkdir -p /data && \
    useradd foo && \
    chown foo:foo /data

ADD nettests /nettests

USER foo

VOLUME /data

CMD ["/nettests"]
