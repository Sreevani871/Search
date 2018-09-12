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
cd #HOME/search/src/cmd
go run main.go

# Sample Input
P Ford Car Review 
P Review Car
P Review Ford
P Toyota Car
P Honda Car
P Car
Q Ford
Q Car
Q Review
Q Ford Review
Q Ford Car
Q cooking French

# Output
Q1:P1 P3
Q2:P6 P1 P2 P4 P5
Q3:P2 P3 P1
Q4:P3 P1 P2
Q5:P1 P3 P6 P2 P4
Q6

