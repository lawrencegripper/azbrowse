# syntax = docker/dockerfile:1.2
FROM ghcr.io/lawrencegripper/azbrowse/snapbase:latest as builder

# Multi-stage build, only need the snaps from the builder. Copy them one at a
# time so they can be cached.
FROM golang:1.20-bullseye
LABEL org.opencontainers.image.source https://github.com/lawrencegripper/azbrowse

COPY --from=builder /snap/core /snap/core
COPY --from=builder /snap/core18 /snap/core18
COPY --from=builder /snap/core20 /snap/core20
COPY --from=builder /snap/snapcraft /snap/snapcraft
COPY --from=builder /snap/bin/snapcraft /snap/bin/snapcraft

# Avoid warnings by switching to noninteractive
ENV DEBIAN_FRONTEND=noninteractive

# Envs
ENV GO111MODULE=on
ENV DEVCONTAINER="TRUE"
## Snap envs: Set the proper environment.
ENV LANG="en_US.UTF-8"
ENV LANGUAGE="en_US:en"
ENV LC_ALL="en_US.UTF-8"
ENV PATH="/snap/bin:$PATH"
ENV SNAP="/snap/snapcraft/current"
ENV SNAP_NAME="snapcraft"
ENV SNAP_ARCH="amd64"
# Versions
ENV DOCKER_VERSION=23.0.1
ARG GO_PLS_VERSION=0.11.0
ARG DLV_VERSION=1.20.1
ARG GO_RELEASER_VERSION=1.15.2
ARG GOLANGCI_LINT_VERSION=1.51.1

RUN \
    apt-get update \
    && apt-get -y install --no-install-recommends apt-utils dialog fuse nano locales ruby-full gnupg2 snapd sudo locales && locale-gen en_US.UTF-8 \
    #
    # Verify git, process tools, lsb-release (common in install instructions for CLIs) installed
    && apt-get -y install git iproute2 procps lsb-release unzip \
    # 
    # Install Release Tools
    #
    # --> RPM used by goreleaser
    && apt install -y rpm \
    # Install python3
    && apt-get -y install python3 python3-pip \
    && apt-get -y install bash-completion \
    # Clean up
    && apt-get autoremove -y \
    && apt-get clean -y \
    && rm -rf /var/lib/apt/lists/*

# Setup locale as required by snapd: https://stackoverflow.com/questions/28405902/how-to-set-the-locale-inside-a-debian-ubuntu-docker-container
RUN sed -i -e 's/# en_US.UTF-8 UTF-8/en_US.UTF-8 UTF-8/' /etc/locale.gen && \
    dpkg-reconfigure --frontend=noninteractive locales && \
    update-locale LANG=en_US.UTF-8

# Install docker used by go releaser
RUN curl -fsSLO https://download.docker.com/linux/static/stable/x86_64/docker-${DOCKER_VERSION}.tgz \
  && tar xzvf docker-${DOCKER_VERSION}.tgz --strip 1 \
                 -C /usr/local/bin docker/docker \
  && rm docker-${DOCKER_VERSION}.tgz

# This Dockerfile adds a non-root 'vscode' user with sudo access. However, for Linux,
# this user's GID/UID must match your local user UID/GID to avoid permission issues
# with bind mounts. Update USER_UID / USER_GID if yours is not 1000. See
# https://aka.ms/vscode-remote/containers/non-root-user for details.
ARG USERNAME=vscode
RUN echo '%sudo ALL=(ALL) NOPASSWD:ALL' >> /etc/sudoers
ARG USER_UID=1000
ARG USER_GID=$USER_UID
ARG DOCKER_GID=997
RUN useradd --create-home --uid $USER_UID -m $USERNAME && usermod -aG sudo $USERNAME && \
    groupadd --gid $DOCKER_GID docker && \
    usermod -aG docker $USERNAME
USER $USERNAME

# Install python deps for deving
RUN pip3 install rope black

# Install ruby bunder gem to support inline gems in ruby scripts
RUN mkdir /home/$USERNAME/.ruby
ENV GEM_HOME=/home/$USERNAME/.ruby/
RUN gem install bundler
ENV PATH="/home/$USERNAME/.ruby/bin:$PATH"


# azure-cli-no-mount
COPY scripts/azure-cli.sh /tmp/
RUN /tmp/azure-cli.sh

# Install Go tools (with cache https://github.com/moby/buildkit/blob/master/frontend/dockerfile/docs/syntax.md#example-cache-go-packages)
RUN \
    # --> Delve for debugging
    go install github.com/go-delve/delve/cmd/dlv@v${DLV_VERSION}\
    # --> Go language server
    && go install golang.org/x/tools/gopls@v${GO_PLS_VERSION} \
    # --> Go symbols and outline for go to symbol support and test support 
    && go install github.com/acroca/go-symbols@v0.1.1 && go install github.com/ramya-rao-a/go-outline@7182a932836a71948db4a81991a494751eccfe77 \
    # --> GolangCI-lint
    && curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sudo sh -s -- -b $(go env GOPATH)/bin v${GOLANGCI_LINT_VERSION} \
    # --> Go releaser 
    && go install github.com/goreleaser/goreleaser@v${GO_RELEASER_VERSION} \
    && sudo rm -rf /go/src/ && sudo rm -rf /go/pkg/mod/**
