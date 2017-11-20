

// to do :
// -  listen on tcp socket
// -  on message print
// - then write client
// - on client write proper arg pasing
// - on server process messages to add correct topics to manager and subscribe
// - on server on publish, call a function
// - need functions to be able to execute file, print to file
// - provide functoinality to reset the daemon
// - create build scripts, one to install it as dameon

#[macro_use]
extern crate serde_derive;
extern crate uuid;
extern crate serde;
extern crate serde_json;
extern crate mqtt;


mod mqtt_mount_manager;

use std::net::{TcpListener, TcpStream};
use std::thread;
use std::io::Read;
use std::io::Write;
use std::str;
use mqtt_mount_manager::*;
use std::sync::mpsc::{Sender, Receiver};
use std::sync::mpsc;



#[derive(Serialize, Deserialize)]
struct SimpleCommand {
    action_type: String,
}

#[derive(Serialize, Deserialize)]
struct AddTopicScriptCommand {
    action_type: String,
    topic: String,
    script_path: String
}

#[derive(Serialize, Deserialize)]
struct AddTopicFileCommand {
    action_type: String,
    topic: String,
    file_path: String,
}

#[derive(Serialize, Deserialize)]
struct DeleteTopicCommand {
    action_type: String,
    id: u32,
}



fn handle_client(mut stream: TcpStream, manager: &mut MqttMountManager) {



    // read 20 bytes at a time from stream echoing back to stream
    loop {
        let mut read = [0; 1028];
        match stream.read(&mut read) {
            Ok(n) => {
                if n == 0 {
                    // connection was closed
                    break;
                }
                let data = &read[0..n];
                println!("data is:  {:?}", data);
                let mut string = String::from(str::from_utf8(&data).unwrap());

                let deserialized: SimpleCommand = serde_json::from_str(&string).unwrap();

                match deserialized.action_type.as_str() {
                    "list" => {
                        stream.write(&manager.to_string().as_bytes()).unwrap();

                    },
                    "reset" => {
                        println!("reset command");
                        &manager.reset();
                        stream.write(String::from("ok").as_bytes()).unwrap();

                    },
                    "add_topic_script" => {
                        let deserialized: AddTopicScriptCommand = serde_json::from_str(&string).unwrap();
                        let mount = MqttMountInfo::new(deserialized.script_path, MqttMountInfoKind::CallScript);
                        &manager.add_mqtt_mount(deserialized.topic, mount);
                        stream.write(String::from("ok").as_bytes()).unwrap();
                    },
                    "add_topic_path" => {
                        let deserialized: AddTopicFileCommand = serde_json::from_str(&string).unwrap();
                        let mount = MqttMountInfo::new(deserialized.file_path, MqttMountInfoKind::WriteFile);
                        &manager.add_mqtt_mount(deserialized.topic, mount);
                        stream.write(String::from("ok").as_bytes()).unwrap();
                    },
                    "delete" => {
                        println!("delete command");
                        stream.write(String::from("ok").as_bytes()).unwrap();

                    },
                    _ => {
                        stream.write(String::from("Command not found").as_bytes()).unwrap();
                    },
                }

            }
            Err(err) => {
                panic!(err);
            }
        }
    }
}




fn main() {

    //let thing = MqttMountInfo { topic: String::from("/room1/temp"), mount_type: MqttMountInfoKind::WriteFile };


    //let mut opts = ClientOptions::new();
    //opts.set_reconnect(ReconnectMethod::ReconnectAfter(Duration::from_secs(1)));
    //let mut client = opts.connect("127.0.0.1:1883", netopt).expect("Can't connect to server");
    println!("hello world");


    /*let mut a= 10;
    let client = mqtt_client::MqttClient::new(String::from("127.0.0.1:1883"),  &mut || {

        let manager = mqtt_mount_manager::MqttMountManager::new();
        manager.add_mqtt_mount(String::from("some path"), mqtt_mount_manager::MqttMountInfo{ topic: String::from("wow"), mount_type: mqtt_mount_manager::MqttMountInfoKind::CallScript });
        a = a + 1;

        println!("a is {}", a);

    });*/
    let listener = TcpListener::bind("127.0.0.1:8080").unwrap();


    let mut manager = mqtt_mount_manager::MqttMountManager::new();

    //let mut mqtt_client = MqttClient::start(client_options, Some(msg_callback)).expect("Coudn't start");


    for stream in listener.incoming() {
        match stream {
            Ok(stream) => {
                handle_client(stream, &mut manager);
            }
            Err(_) => {
                println!("Error");
            }
        }
    }


}
