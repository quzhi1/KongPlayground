FROM --platform=linux/amd64 kong:2.8.1-alpine

ADD bin/go-hello /usr/local/bin/go-hello

# Prove that go-hello works
RUN /usr/local/bin/go-hello -dump

USER kong