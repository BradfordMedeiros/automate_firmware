# turn on and start services to host pi as an access point and be able to run dns

sudo systemctl enable hostapd
sudo systemctl enable dnsmasq

sudo service hostapd start
sudo service dnsmasq start
