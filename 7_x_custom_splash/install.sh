sudo apt-get install fbi
sudo cp ./splash.png /opt/splash.png
sudo cp splashscreen.service /etc/systemd/system
sudo systemctl enable splashscreen
