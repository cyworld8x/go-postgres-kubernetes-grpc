Guideline to set up static app with nginx on Amazon Linux
sudo yum update -y
sudo yum install firewalld nginx git golang -y
sudo systemctl unmask firewalld
sudo systemctl enable firewalld
sudo systemctl start firewalld
sudo systemctl enable nginx
sudo systemctl start nginx

Open port

sudo firewall-cmd --permanent --add-port=5002/tcp
sudo firewall-cmd --reload
sudo netstat -nap | grep -i nginx