FROM ubuntu:18.04
# Install Azure CLI and root CAs 
RUN apt-get update && apt-get install apt-transport-https lsb-release software-properties-common dirmngr ca-certificates -y && \
    AZ_REPO=$(lsb_release -cs) && \
    echo "deb [arch=amd64] https://packages.microsoft.com/repos/azure-cli/ $AZ_REPO main" | \
    tee /etc/apt/sources.list.d/azure-cli.list && \
    apt-key --keyring /etc/apt/trusted.gpg.d/Microsoft.gpg adv \
     --keyserver packages.microsoft.com \
     --recv-keys BC528686B50D79E339D3721CEB3E94ADBE1229CF && \
    apt-get update && apt-get install -y azure-cli
COPY azbrowse /
ENTRYPOINT ["/azbrowse"]