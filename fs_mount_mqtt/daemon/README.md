Dameon program that acts as a tcp server to listen for commands from the fs_mount_mqtt client.

Acts as a client to an mqtt broker.  

Based upon client commands, saved topics subscriptions, and when a topic is  emitted from the broker for one of its subscriptions,
it either

-executes a bash script with $1=  topic, and $2= topic_message

-writes a file with the value of topic_message
