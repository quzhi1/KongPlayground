# FROM --platform=linux/amd64 kong:2.8.1-alpine
FROM kong:2.3.3-alpine

# USER root

ADD bin/go-hello /usr/local/bin/go-hello

# RUN ["chmod", "777", "/usr/local/bin/go-hello"]
# RUN ["chmod", "-R", "777", "/usr/local/kong"]

USER kong

# Prove that go-hello works
RUN ["/usr/local/bin/go-hello", "-dump"]
