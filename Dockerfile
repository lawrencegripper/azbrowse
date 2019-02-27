FROM ubuntu
RUN curl -L https://aka.ms/InstallAzureCli | bash
COPY azbrowse /
ENTRYPOINT ["/azbrowse"]