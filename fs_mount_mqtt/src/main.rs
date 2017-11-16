use std::collections::HashMap;

enum MqttMountInfoKind {
    WriteFile,
    CallScript,
}

struct MqttMountInfo {
    topic : String,
    mount_type : MqttMountInfoKind,
}


struct MqttMountManager {
    mqtt_mounts: HashMap<String, MqttMountInfo>,
}
impl MqttMountManager {
    fn new () -> MqttMountManager {
        MqttMountManager { mqtt_mounts : HashMap::new()  }
    }
    fn add_mqtt_mount(mut self, topic: String, mount_info: MqttMountInfo) {
        self.mqtt_mounts.insert(topic, mount_info);
    }
    fn remove_mqtt_mount(mut self, topic: String, mount_info: MqttMountInfo){
        self.mqtt_mounts.remove(&topic);
    }
}


fn main() {

    let thing = MqttMountInfo { topic: String::from("/room1/temp"), mount_type: MqttMountInfoKind::WriteFile };

    let manager = MqttMountManager::new();
    manager.add_mqtt_mount(String::from("some path"), MqttkMountInfo{ topic: String::from("wow"), mount_type: MqttMountInfoKind::CallScript });
    println!("hello world");
}
