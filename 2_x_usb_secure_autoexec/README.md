
#USB_SECURE_AUTO_EXEC
------------------
Scripts which automatically executes the scripts mounted when a usb is plugged in. 
Build script configures the USB PATH, EXECUTABLE_LOCATION (binary that executes the scripts), and the AUTOMATE_KEY LOCATION.

The file stored at the AUTOMATE_KEY_LOCATION is compared with the  contents of $USB_PATH/automate.key.  If they match, all files ending in .sh are executed as bash script.
