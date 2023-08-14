# Automate the installation of packages and software needed for this project.
echo "sudo apt update -y"
sudo apt update -y
echo "sudo apt upgrade -y"
sudo apt upgrade -y

# Install Golang
echo "sudo apt install golang -y"
sudo apt install golang -y
echo "go version"
go version

# Setup devevlopment environment
echo '''\
echo "export GOPATH=$HOME/go" >> ~/.bashrc
echo "export PATH=$PATH:$GOPATH/bin" >> ~/.bashrc
source ~/.bashrc
'''
echo "export GOPATH=$HOME/go" >> ~/.bashrc
echo "export PATH=$PATH:$GOPATH/bin" >> ~/.bashrc
source ~/.bashrc

# Initialize the Go project folder
echo'''\
cd server
go mod init main
'''
cd server
go mod init main
