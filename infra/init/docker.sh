#!/bin/bash

# Update system
sudo apt-get update -y
sudo apt-get upgrade -y

# Install Docker
sudo apt-get install docker.io -y

# Add 'ubuntu' to the 'docker' group
sudo usermod -a -G docker ubuntu

# Enable and start Docker service
sudo systemctl enable docker
sudo systemctl start docker

# Install Docker Compose
sudo apt-get install python3-pip -y
sudo pip3 install docker-compose

# Enable docker-compose command without specifying full path
echo 'export PATH=$PATH:/usr/local/bin' >> ~/.bashrc
source ~/.bashrc

# Reminder to log out and log back in
echo "Please log out and log back in for the group changes to take effect."
