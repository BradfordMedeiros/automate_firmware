
fs_mount_mqtt
--------------
Dameon utility intended to allow mounting of mqtt topics to the filesystem from the local automate broker. 

------
sudo service fs_mount_mqtt start/stop/restart/status/etc.  

Service must be running to use fs_mount_mqtt commands. If service is not running commands will fail.
A list of topics will be maintained, so do not have to monitor reboot/start/etc.  Only need to call once. 
Successive calls to the same topic with same path/script path will yield a no-op. 


usage: 
----
fs_mount_mqtt --start/--stop -t &lt;topict&gt; -p &lt;path&gt;
 This creates a file that will be updated with the latest value whenever a new value is published 


fs_mount_mqtt --start/--stop -t &lt;topic&gt; -s &lt;script_path&gt;
This will call the script at script_path with $2 as the topic, and $3 with the value of the mqtt topic whnever a value is published.

fs_mount-mqtt --reset
This will unsubscribe to all topics, and effectively reset the daemon.  Generally do not use this during production. 

-------------------
For publishing to a the broker, this is good. 
https://github.com/mqttjs/MQTT.js
