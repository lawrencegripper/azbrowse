FROM mcr.microsoft.com/azure-cli
COPY azbrowse /
ENTRYPOINT ["/azbrowse"]
LABEL org.opencontainers.image.source https://github.com/lawrencegripper/azbrowse