FROM golang:1.19.3-buster AS development

ENV NODE_VERSION v16.17.1
ENV NVM_DIR /home/go/.nvm
ENV NPM_FETCH_RETRIES 2
ENV NPM_FETCH_RETRY_FACTOR 10
ENV NPM_FETCH_RETRY_MINTIMEOUT 10000
ENV NPM_FETCH_RETRY_MAXTIMEOUT 60000

RUN apt-key adv --keyserver pgp.mit.edu --recv-keys 3A79BD29

RUN go install golang.org/x/tools/gopls@v0.11.0

RUN go install github.com/go-delve/delve/cmd/dlv@v1.20.1
RUN go install golang.org/x/tools/cmd/godoc@v0.5.0

RUN echo "deb http://repo.mysql.com/apt/ubuntu/ bionic mysql-8.0" | tee /etc/apt/sources.list.d/mysql.list > /dev/null

RUN apt update && apt upgrade -y

RUN apt install \
              sudo \
              mysql-client \
              git \
              curl \
              make \
              ncurses-dev \
              build-essential \
              tree \
              nano \
              tmux \
              tmuxinator \
              xclip \
              apt-transport-https \
              ca-certificates \
              gnupg-agent \
              software-properties-common \
              build-essential \
              libssl-dev -y

RUN git clone --depth 1 --branch v9.0.1224 https://github.com/vim/vim.git /tmp/vim-installation && \
                  cd /tmp/vim-installation/src/ && \
                  ./configure && \
                  make && \
                  make install && \
                  rm -rf /tmp/vim-installation

RUN useradd -ms /bin/bash go && echo "go:pass" | chpasswd && adduser go sudo

USER go

# Install Node.js NPM and Yarn through NVM
RUN mkdir -p $NVM_DIR && \
              curl -o- https://raw.githubusercontent.com/nvm-sh/nvm/v0.38.0/install.sh | bash \
              && . $NVM_DIR/nvm.sh \
              && nvm install ${NODE_VERSION} \
              && nvm use ${NODE_VERSION} \
              && nvm alias ${NODE_VERSION} \
              && npm config set fetch-retries ${NPM_FETCH_RETRIES} \
              && npm config set fetch-retry-factor ${NPM_FETCH_RETRY_FACTOR} \
              && npm config set fetch-retry-mintimeout ${NPM_FETCH_RETRY_MINTIMEOUT} \
              && npm config set fetch-retry-maxtimeout ${NPM_FETCH_RETRY_MAXTIMEOUT} \
              && ln -s `npm bin --global` /home/go/.node-bin \
              && npm install -g yarn \
              && npm install -g npm

# Install FZF
RUN git clone --depth 1 https://github.com/junegunn/fzf.git $HOME/.fzf && $HOME/.fzf/install

# Customizations
RUN git clone --bare -b godevenv https://github.com/i3onilha/.dotfiles.git $HOME/.dotfiles && \
              git clone -b heavenly2 https://github.com/i3onilha/.vim.git $HOME/.vim && \
              git clone https://github.com/i3onilha/.tmux.git $HOME/.tmux && \
              ln -sf .vim/.vimrc $HOME && \
              ln -sf .tmux/.tmux.conf $HOME && \
              cp $HOME/.tmux/.tmux.conf.local $HOME && \
              cd ~/.vim && \
              git submodule init && \
              git submodule update && \
              curl -o- https://raw.githubusercontent.com/crusoexia/vim-monokai/master/colors/monokai.vim > ~/.vim/colors/monokai.vim && \
              cd ~ && \
              git --git-dir=$HOME/.dotfiles/ --work-tree=$HOME config --local status.showUntrackedFiles no && \
              git --git-dir=$HOME/.dotfiles/ --work-tree=$HOME reset HEAD . && \
              git --git-dir=$HOME/.dotfiles/ --work-tree=$HOME checkout -- .

WORKDIR /home/go/sourcecode

COPY . .

FROM golang:1.19.3-bullseye AS builder

WORKDIR /home/go/sourcecode

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

RUN go build -o main .

FROM scratch

COPY --from=builder /home/go/sourcecode/main /app/main

EXPOSE 8080

CMD ["/app/main"]
