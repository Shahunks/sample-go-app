FROM ubuntu:jammy-20220428

MAINTAINER applicant@lightspeedhq.com

ENV NPM_TOKEN=30961fa8-cd15-49e1-96b9-9c8b40caa0e2

EXPOSE 10000

RUN apt update
RUN apt install -y golang ca-certificates

RUN mkdir -p /dx-test
COPY ./* /dx-test
WORKDIR /dx-test
RUN echo $(ls -la)

RUN CGO_ENABLED=0 GOARCH=amd64 GOOS="linux" go build -o ${BINARY_NAME:-$(basename $(pwd))}.amd64
RUN CGO_ENABLED=0 GOARCH=arm64 GOOS="linux" go build -o ${BINARY_NAME:-$(basename $(pwd))}.arm64
RUN cp dx-test.* /bin/

RUN apt -y dist-upgrade

ENTRYPOINT ["/bin/dx-test.arm64"]
