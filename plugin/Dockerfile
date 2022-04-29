FROM kong:2.8.1-alpine

ADD bin/go-hello /usr/local/bin/go-hello

USER kong

# Prove that go-hello works
RUN ["/usr/local/bin/go-hello", "-dump"]
