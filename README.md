# Search

# Install
cd /tmp

wget https://storage.googleapis.com/golang/go1.7.3.linux-amd64.tar.gz

sudo tar -C /usr/local -xzf /tmp/go1.7.3.linux-amd64.tar.gz

# Clone
cd $HOME

git clone git@github.com:Sreevani871/Search.git

export GOPATH=$HOME/search

# Run 
cd $HOME/search/src/cmd

go run main.go

# Run Unit Tests
cd $HOME/search/src/search

go test -v

