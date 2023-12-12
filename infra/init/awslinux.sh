#!/bin/bash

# Update system
sudo yum update -y

# Install Docker
sudo yum install docker -y

# Add 'ec2-user' to the 'docker' group
sudo usermod -a -G docker ec2-user

# Enable and start Docker service
sudo systemctl enable docker.service
sudo systemctl start docker.service

# Install Docker Compose
sudo yum install python3-pip -y
sudo pip3 install docker-compose

# Enable docker-compose command without specifying full path
echo 'export PATH=$PATH:/usr/local/bin' >> ~/.bash_profile
source ~/.bash_profile

# Reminder to log out and log back in
echo "Please log out and log back in for the group changes to take effect."
