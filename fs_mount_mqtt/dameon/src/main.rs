

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

extern crate mqtt3;

use std::net::{TcpListener, TcpStream};
use std::thread;
use std::io::Read;
use std::io::Write;
use std::str;


mod mqtt_client;
mod mqtt_mount_manager;




fn handle_client(mut stream: TcpStream) {
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
                let string = str::from_utf8(&data).unwrap();
                println!("as string: {}", string);
                stream.write(data).unwrap();
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

    for stream in listener.incoming() {
        match stream {
            Ok(stream) => {
                thread::spawn(move || {
                    handle_client(stream);
                });
            }
            Err(_) => {
                println!("Error");
            }
        }
    }


}
