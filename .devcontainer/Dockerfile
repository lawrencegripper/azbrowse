#-------------------------------------------------------------------------------------------------------------
# Copyright (c) Microsoft Corporation. All rights reserved.
# Licensed under the MIT License. See https://go.microsoft.com/fwlink/?linkid=2090316 for license information.
#-------------------------------------------------------------------------------------------------------------

FROM golang:1.13.5-stretch

# Avoid warnings by switching to noninteractive
ENV DEBIAN_FRONTEND=noninteractive

# This Dockerfile adds a non-root 'vscode' user with sudo access. However, for Linux,
# this user's GID/UID must match your local user UID/GID to avoid permission issues
# with bind mounts. Update USER_UID / USER_GID if yours is not 1000. See
# https://aka.ms/vscode-remote/containers/non-root-user for details.
ARG USERNAME=vscode
ARG USER_UID=1000
ARG USER_GID=$USER_UID

ENV GO111MODULE=on

# Configure apt, install packages and tools
RUN apt-get update \
    && apt-get -y install --no-install-recommends apt-utils dialog \
    #
    # Verify git, process tools, lsb-release (common in install instructions for CLIs) installed
    && apt-get -y install git iproute2 procps lsb-release \
    #
    # Install Azure CLI
    && curl -sL https://aka.ms/InstallAzureCLIDeb | bash \
    #
    # Install Integration testing Tools 
    #
    # --> xvfb for integration testing (gocui requires a valid tty which isn't available in ci)
    && apt-get install --no-install-recommends -y xvfb libgl1-mesa-dri \
    # --> xterm for integration testing
    && apt-get install -y xterm \
    # 
    # Install Release Tools
    #
    # --> RPM used by goreleaser
    && apt install -y rpm \
    #
    # Clean up
    && apt-get autoremove -y \
    && apt-get clean -y \
    && rm -rf /var/lib/apt/lists/*

ENV GIT_PROMPT_START='\033[1;36mazb-devcon>\033[0m\033[0;33m\w\a\033[0m'

# Save command line history 
RUN echo "export HISTFILE=/root/commandhistory/.bash_history" >> "/root/.bashrc" \
    && echo "export PROMPT_COMMAND='history -a'" >> "/root/.bashrc" \
    && mkdir -p /root/commandhistory \
    && touch /root/commandhistory/.bash_history
    
# Git command prompt
RUN git clone https://github.com/magicmonty/bash-git-prompt.git ~/.bash-git-prompt --depth=1 \
    && echo "if [ -f \"$HOME/.bash-git-prompt/gitprompt.sh\" ]; then GIT_PROMPT_ONLY_IN_REPO=1 && source $HOME/.bash-git-prompt/gitprompt.sh; fi" >> "/root/.bashrc"


# Alow config for xvfb
ARG RESOLUTION="1920x1080x24"
ENV XVFB_RES="${RESOLUTION}"
ARG XARGS=""
ENV XVFB_ARGS="${XARGS}"

ENV DEVCONTAINER="TRUE"

# Install docker used by go releaser
RUN bash -c "cd /tmp && curl -fsSLO https://download.docker.com/linux/static/stable/x86_64/docker-19.03.5.tgz && tar --strip-components=1 -xvzf docker-19.03.5.tgz -C /usr/local/bin"

# Install Go tools
RUN \
    # --> Delve for debugging
    go get github.com/go-delve/delve/cmd/dlv@v1.3.2 \
    # --> Go language server
    && go get golang.org/x/tools/gopls@v0.2.1 \
    # --> Go symbols and outline for go to symbol support and test support 
    && go get github.com/acroca/go-symbols@v0.1.1 && go get github.com/ramya-rao-a/go-outline@7182a932836a71948db4a81991a494751eccfe77 \
    # --> GolangCI-lint
    && curl -sfL https://install.goreleaser.com/github.com/golangci/golangci-lint.sh | sed 's/tar -/tar --no-same-owner -/g' | sh -s -- -b $(go env GOPATH)/bin \
    # --> Go releaser 
    && curl -sfL https://install.goreleaser.com/github.com/goreleaser/goreleaser.sh | sh -s -- "v0.123.3"\
    && rm -rf /go/src/ && rm -rf /go/pkg
