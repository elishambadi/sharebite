wget https://go.dev/dl/go1.23.4.linux-amd64.tar.gz
rm -rf /usr/local/go && tar -C /usr/local -xzf go1.23.4.linux-amd64.tar.gz
export PATH=$PATH:/usr/local/go/bin
go version

go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest

# Install node and nvm
curl -o- https://raw.githubusercontent.com/nvm-sh/nvm/v0.40.1/install.sh | bash
export NVM_DIR="$([ -z "${XDG_CONFIG_HOME-}" ] && printf %s "${HOME}/.nvm" || printf %s "${XDG_CONFIG_HOME}/nvm")"
[ -s "$NVM_DIR/nvm.sh" ] && \. "$NVM_DIR/nvm.sh" # This loads nvm

nvm install --lts

npm install -D tailwindcss
npx tailwindcss init

npm run build

# install templ - templating for go htmls
go install github.com/a-h/templ/cmd/templ@latest

# install postgres
sudo apt update
sudo apt install postgresql postgresql-contrib

go install github.com/sqlc-dev/sqlc/cmd/sqlc@latest
go get github.com/jackc/pgx/v4

sudo apt-get update
sudo apt-get install docker-compose-plugin