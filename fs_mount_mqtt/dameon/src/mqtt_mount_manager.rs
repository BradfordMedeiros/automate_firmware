
use std::collections::HashMap;

pub enum MqttMountInfoKind {
    WriteFile,
    CallScript,
}
pub struct MqttMountInfo {
    pub topic : String,
    pub mount_type : MqttMountInfoKind,
}

pub struct MqttMountManager {
    mqtt_mounts: HashMap<String, MqttMountInfo>,
}

impl MqttMountManager {
    pub fn new () -> MqttMountManager {
        MqttMountManager { mqtt_mounts : HashMap::new()  }
    }
    pub fn add_mqtt_mount(mut self, topic: String, mount_info: MqttMountInfo) {
        self.mqtt_mounts.insert(topic, mount_info);
    }
    pub fn remove_mqtt_mount(mut self, topic: String, mount_info: MqttMountInfo){
        self.mqtt_mounts.remove(&topic);
    }
}
