FROM golang:1.19.3-buster AS development

ENV NODE_VERSION v16.17.1
ENV NVM_DIR /home/go/.nvm
ENV NPM_FETCH_RETRIES 2
ENV NPM_FETCH_RETRY_FACTOR 10
ENV NPM_FETCH_RETRY_MINTIMEOUT 10000
ENV NPM_FETCH_RETRY_MAXTIMEOUT 60000

RUN go install golang.org/x/tools/gopls@latest

RUN apt update && apt upgrade -y

RUN apt install curl

RUN apt install \
              git \
              vim \
              tmux \
              tmuxinator \
              xclip \
              apt-transport-https \
              ca-certificates \
              gnupg-agent \
              software-properties-common \
              build-essential \
              libssl-dev -y

RUN useradd -ms /bin/bash go

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
RUN git clone --bare -b last-stable https://github.com/jean-bonilha/.dotfiles.git $HOME/.dotfiles && \
              git clone -b heavenly2 https://github.com/jean-bonilha/.vim.git $HOME/.vim && \
              git clone https://github.com/jean-bonilha/.tmux.git $HOME/.tmux && \
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

RUN CGO_ENABLED=0 go test -v

RUN go build .

FROM scratch

COPY --from=builder /home/go/sourcecode/main /app/main

EXPOSE 8080

CMD ["/app/main"]
