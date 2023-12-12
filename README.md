# Starting a new project

Update the nginx.conf for both staging and production, this should match the domain names.

```
/infra/production/nginx/nginx.conf
/infra/staging/nginx/nginx.conf
```

# Project Setup Guide for Ubuntu
This guide outlines the steps to set up a project on Ubuntu, including the installation of Git, cloning a repository, running an initialisation script, and starting the project with Docker Compose.

## Prerequisites
An Ubuntu instance.
Sudo privileges on your system.

### Step 1: Install Git
Git is essential for version control and managing source code changes. Install Git with the following commands:

Update
```
sudo apt update
```
Install git
```
sudo apt install git
```
### Step 2: Clone the Repository
Clone your project's repository using Git:

```
git clone [URL_of_the_Repository]
cd [Repository_Name]
```
Replace [URL_of_the_Repository] with the URL of your Git repository and [Repository_Name] with the name of the cloned directory.

### Step 3: Run the Docker Initialisation Script
This script will setup docker and docker-compose on the server.

Change permissions
```
chmod +x ./infra/init/docker.sh
```
Run docker
```
./infra/init/docker.sh
```

### Step 4: Restart the Session
To ensure all changes are applied, log out and log back into your session. This can be achieved by exiting the SSH session and reconnecting, or simply by restarting the terminal session.

### Step 5: Initiate the project

cd into the directory and add a .env

```
cd [Repository_Name]
touch .env
```
cd into either production or staging, depending on the server and compose the docker container

```
cd infra/production or infra/staging
docker-compose up
```
This will start the containers as defined in your docker-compose.yml file.