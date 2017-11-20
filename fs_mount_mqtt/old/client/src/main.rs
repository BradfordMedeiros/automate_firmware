#[macro_use]
extern crate serde_derive;

extern crate serde;
extern crate serde_json;
extern crate getopts;

use std::io::prelude::*;
use std::net::TcpStream;
use std::result;
use getopts::Options;
use std::str;
use serde_json::{Value, Error};

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

fn write_stream(command: &str) -> Result<String, std::io::Error>{
    let mut stream = TcpStream::connect("127.0.0.1:8080").unwrap();

    let to_write = command.as_bytes();
    let _ = stream.write(&to_write); // ignore the Result

    let mut read = [0; 1028];
    match stream.read(&mut read) {
        Ok(n) => {
            let data = &read[0..n];
            Ok(String::from(str::from_utf8(&data).unwrap()))
        }
        Err(err) => Err(err),
    }
}

fn print_usage(program: &str, opts: Options){
    let description = format!("Usage: {} FILE [options]", program);
    println!("{}", opts.usage(&description));
}

fn reset(){
    let resetCommand = SimpleCommand { action_type: String::from("reset")};
    let response = write_stream(&serde_json::to_string(&resetCommand).unwrap()).unwrap();
    println!("response is {}", response);
}

fn list(){
    let listCommand = SimpleCommand { action_type: String::from("list")};
    let response = write_stream(&serde_json::to_string(&listCommand).unwrap()).unwrap();
    println!("{}", response);
}

fn add_topic_with_path(topic: String, file_path: String) {
    // maybe write json?
    // { action: 'subscribe', topic, value }
    let addTopicWithPath = AddTopicFileCommand {
        action_type: String::from("add_topic_path"),
        topic,
        file_path,
    };

    let response = write_stream(&serde_json::to_string(&addTopicWithPath).unwrap()).unwrap();
    println!("{}", response);
}

fn add_topic_with_script(topic: String, script_path: String){
    let addTopicWithScript = AddTopicScriptCommand {
        action_type: String::from("add_topic_script"),
        topic,
        script_path,
    };

    let response = write_stream(&serde_json::to_string(&addTopicWithScript).unwrap()).unwrap();
    println!("{}", response);
}

fn delete() {
    let deleteCommand = DeleteTopicCommand { action_type: String::from("delete"), id: 1 };
    let response = write_stream(&serde_json::to_string(&deleteCommand).unwrap()).unwrap();
    println!("{}", response);
}


fn main(){
    let args: Vec<String> = std::env::args().collect();
    let program_name = args[0].clone();
    let mut opts = Options::new();
    opts.optflag("h", "help", "prints help information (this message)");
    opts.optflag("r", "reset", "resets the daemon");
    opts.optflag("l", "list", "list current topic subscriptions");
    opts.optflag("d", "delete", "delete a subscription by id (find id with list");
    opts.optopt("t", "topic", "mqtt topic to subscribe to data from (must also use p or s option)",  "TOPIC");
    opts.optopt("p", "path", "path to write file to when mqtt data is received (cannot be used with script)", "PATH");
    opts.optopt("s", "script", "script or executable to call when a topic is received (cannot be used with path", "SCRIPT");

    let matches = match opts.parse(&args[1..]){
        Ok(m) => { m }
        Err(f) => { panic!(f.to_string()) }
    };

    if matches.opt_present("h"){
        print_usage(&program_name, opts);
    }else if matches.opt_present("r"){
        reset();
    }else if matches.opt_present("l"){
        list();
    }else if matches.opt_present("t") {

        let topic = matches.opt_str("t").unwrap();
        match matches.opt_str("p") {
            Some(file_path) => {
                add_topic_with_path(String::from(topic), file_path);
            },
            None => {
                match matches.opt_str("s") {
                    Some(script_path) => {
                        add_topic_with_script(String::from(topic),script_path);
                    },
                    None => {
                        print_usage(&program_name, opts);
                    }
                }
            }
        }
    }else if matches.opt_present("d"){
        delete();
    }else{
        print_usage(&program_name, opts);
    }
}
