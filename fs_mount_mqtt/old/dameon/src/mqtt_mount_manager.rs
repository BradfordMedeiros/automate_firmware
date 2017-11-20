use mqtt::*;
use uuid::Uuid;
use std::collections::HashMap;
use std::sync::mpsc::{Sender, Receiver};
use std::sync::mpsc;

pub enum MqttMountInfoKind {
    WriteFile,
    CallScript,
}
pub struct MqttMountInfo {
    pub file_path: String,
    pub mount_type : MqttMountInfoKind,
    id: String,
}
impl MqttMountInfo {
    pub fn new (file_path: String, mount_type: MqttMountInfoKind) -> MqttMountInfo {MqttMountInfo { file_path, mount_type, id: Uuid::new_v4().to_string() }}
}

pub struct MqttMountManager {
    mqtt_mounts: HashMap<String, Vec<MqttMountInfo>>,
   // mqtt_client: MqttClient,
}


impl MqttMountManager {
    pub fn new () -> MqttMountManager {

        let mut mqtt_mounts = HashMap::new();


        let result = MqttConnection::connect("localhost:1883", "fs_mount_mqtt", None, None, None, false, 0);
        if let Ok(conn @ MqttConnection(_)) = result {

        } else {
            println!("Can't connect to broker!")
        }

        let manager =  MqttMountManager {
            mqtt_mounts : mqtt_mounts,
            //mqtt_client,
            /*mqtt_client: MqttClient::new(String::from("127.0.0.1:9000"), &mut || {
                println!("--got message");

            })*/
        };

        //manager.on_topic(String::from("sometopic"), String::from("some_message"));
        manager
    }
    pub fn add_mqtt_mount(&mut self, topic: String, mount_info: MqttMountInfo) {
        self.mqtt_mounts.entry(topic).or_insert(Vec::new()).push(mount_info);
    }
    pub fn remove_mqtt_mount(uuid: String){
        //&self.mqtt_mounts.remove(&topic);
    }
    pub fn reset(&mut self){
        //println!("reset placeholder");
        &self.mqtt_mounts.clear();
    }
    pub fn on_topic(&self, topic: String, message: String) {
        // call  topics  and stuff
    }
    pub fn to_string(&self) -> String {
        let mut val = String::from("");
        for (key, value) in &self.mqtt_mounts {
            //println!("key is: {}, value is { }", key, value.file_path);
            let topic = key;

            for item in value {
                let mountType = match item.mount_type {
                    MqttMountInfoKind::CallScript => "script",
                    MqttMountInfoKind::WriteFile => "file",
                };
                val.push_str("topic: ");
                val.push_str(topic);

                val.push_str("  type: ");
                val.push_str(mountType);

                val.push_str("  path: ");
                val.push_str(item.file_path.as_str());

                val.push_str("  id:  ");
                val.push_str(item.id.as_str());
                val.push_str("\n");
            }
        }
        val
    }
}
