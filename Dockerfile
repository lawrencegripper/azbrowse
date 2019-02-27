FROM microsoft/azure-cli
COPY azbrowse /
ENTRYPOINT ["/azbrowse"]