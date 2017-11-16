
extern crate mqtt3;

use std::env;
use std::net::TcpStream;
use std::io::{Read, Write, BufReader, BufWriter};
use std::process::exit;
use mqtt3::{MqttRead, MqttWrite, Packet, Connect, Publish, Subscribe, Protocol, QoS, PacketIdentifier};
use std::sync::Arc;

mod mqtt_client;

/*enum MqttMountInfoKind {
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
}*/




fn main() {

    //let thing = MqttMountInfo { topic: String::from("/room1/temp"), mount_type: MqttMountInfoKind::WriteFile };

    //let manager = MqttMountManager::new();
    //manager.add_mqtt_mount(String::from("some path"), MqttMountInfo{ topic: String::from("wow"), mount_type: MqttMountInfoKind::CallScript });

    //let mut opts = ClientOptions::new();
    //opts.set_reconnect(ReconnectMethod::ReconnectAfter(Duration::from_secs(1)));
    //let mut client = opts.connect("127.0.0.1:1883", netopt).expect("Can't connect to server");
    println!("hello world");


    let client = mqtt_client::MqttClient::new(String::from("127.0.0.1:1883"));


    let packet = reader.read_packet().unwrap();
    println!("{:?}", packet);

    subscribe_to_topic(&mut writer, "humidity".to_owned());

    let packet = reader.read_packet().unwrap();
    println!("{:?}", packet);

    loop {
        let packet = reader.read_packet().unwrap();
        println!("{:?}", packet);
    }

}
