# automate_firmware

This code is made for the raspberry pi platform that hosts automate (although it should be generic enough to work on any linux system -- but other non-linux platforms are secondary target -- at least for now). 

This code installs the features of the automate hardware system.  This means this is 
1. a collection of scripts to install automate 
and 
2. a collection of additional utilities to enhance automate.

The additional utilities are meant to enhance automates functionality is ways that allow the user to take advantage of the fact that the user is in fact on real hardware, and has a linux device.  However, while  this functionality may or may not be able to be to live in automate_core (sometimes), the idea is that these are system tools, that should not be directly a concern of automate_core.

Current extended features:

-secure automatic usb script execution 
  - This is targeted to allow  creative uses of the fact that the user has root to automate_system
  - Allows things like, for example, plugging in usb to unlock a system, back it up, etc (since we are on linux!) 
  - but doesnt compromise security for that coolness (uses crypto library generated hash)


-fs_mount_mqtt
  - allows you to call a script or write a file when an mqtt topic is publishes
  - this allows you to use scripts/executables that can use mqtt (without having to include mqtt/http in your library)
  - really nice for easy scripting 
  
-led_control
  - super simple script that uses fs_mount_mqtt to toggle a gpio pin on the raspberry pi (via /dev) based on mqtt topic
  - gpio pin should be hooked up to led (this is really gpio pin toggle software lol)

