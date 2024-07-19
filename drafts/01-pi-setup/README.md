# Pi Setup

👋 you can use a app locally that will connect remotely to the Ollama running on the Pi, or you can run everything on the Pi (with Docker Compose for example).

> Connect
```bash
ssh k33g@bob.local
```

## Install or update Ollama

```bash
# install
curl -fsSL https://ollama.com/install.sh | sh

ollama --version
ollama version is 0.2.5
```

## Install Docker

- https://docs.docker.com/engine/install/debian/
- https://docs.docker.com/engine/install/linux-postinstall/#manage-docker-as-a-non-root-user

### Prepare

```bash
# Add Docker's official GPG key:

sudo apt-get update
sudo apt-get install ca-certificates curl
sudo install -m 0755 -d /etc/apt/keyrings
sudo curl -fsSL https://download.docker.com/linux/debian/gpg -o /etc/apt/keyrings/docker.asc
sudo chmod a+r /etc/apt/keyrings/docker.asc

# Add the repository to Apt sources:

echo \
  "deb [arch=$(dpkg --print-architecture) signed-by=/etc/apt/keyrings/docker.asc] https://download.docker.com/linux/debian \
  $(. /etc/os-release && echo "$VERSION_CODENAME") stable" | \
  sudo tee /etc/apt/sources.list.d/docker.list > /dev/null
sudo apt-get update
```

### Install
```bash
sudo apt-get install docker-ce docker-ce-cli containerd.io docker-buildx-plugin docker-compose-plugin
```

### Test
```bash
sudo docker run hello-world
```

### Manage Docker as a non-root user
```bash
sudo groupadd docker
sudo usermod -aG docker $USER
newgrp docker

# Verify that you can run docker commands without sudo.
docker run hello-world
```

## Install Remote Development extension

https://www.raspberrypi.com/news/coding-on-raspberry-pi-remotely-with-visual-studio-code/

Next you can connect to your Raspberry Pi. Launch the VS Code command palette using Ctrl+Shift+P on Linux or Windows, or Cmd+Shift+P on macOS. Search for and select Remote SSH: Connect current window to host (there’s also a connect to host option that will create a new window).
