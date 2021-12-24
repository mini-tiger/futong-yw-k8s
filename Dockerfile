FROM 172.16.71.10:15005/dev/ubuntu:focal
WORKDIR /ftcloud/yw/futong-yw-k8s
RUN mkdir -p /ftcloud/yw/futong-yw-k8s/base
COPY out/main_ftk8s .
COPY base/i18n ./base/i18n
COPY base/rsakey ./base/rsakey
ENTRYPOINT ["./main_ftk8s"]
