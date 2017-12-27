# stop services allowing the pi to be an access point
sudo service hostapd stop
sudo service dnsmasq stop

# prevent services from restarting when rebooted
sudo systemctl disable hostapd
sudo systemctl disable dnsmasq

# reconnect to available wifi as if booting
sudo systemctl restart dhcpcd
