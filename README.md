Setup
-
#### install the gitlab-runner
wget https://gitlab-runner-downloads.s3.amazonaws.com/master/binaries/gitlab-runner-linux-amd64 \
sudo mv gitlab-runner-linux-amd64 /usr/local/bin \
sudo chmod +x /usr/local/bin/gitlab-runner-linux-amd64 \

#### add the gitlab-runner user to the docker group
sudo usermod -aG docker gitlab-runner
    
#### execute the gitlab-ci script by running
gitlab-runner exec docker app
