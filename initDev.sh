# Automate the installation of packages and software needed for this project.
sudo apt update -y
sudo apt upgrade -y
sudo apt --fix-broken install
sudo apt autoremove
sudo apt clean
sudo apt autoclean

# Install make
sudo apt install make -y

# Install Golang
sudo apt install golang -y
go version

# Install docker compose
sudo apt install docker.io -y

# Install Docker
sudo apt install -y apt-transport-https ca-certificates curl software-properties-common
curl -fsSL https://download.docker.com/linux/ubuntu/gpg | sudo gpg --dearmor -o /usr/share/keyrings/docker-archive-keyring.gpg
echo "deb [signed-by=/usr/share/keyrings/docker-archive-keyring.gpg] https://download.docker.com/linux/ubuntu $(lsb_release -cs) stable" | sudo tee /etc/apt/sources.list.d/docker.list > /dev/null
sudo apt update -y
sudo apt install -y docker-ce docker-ce-cli containerd.io
sudo usermod -aG docker $USER
sudo systemctl start docker
sudo systemctl enable docker
sudo docker --version
sudo docker info

# Setup devevlopment environment
echo "export GOPATH=$HOME/go" >> ~/.bashrc
echo "export PATH=$PATH:$GOPATH/bin" >> ~/.bashrc
source ~/.bashrc

# Initialize the Go project folder
cd server
go mod init main
