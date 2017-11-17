extern crate getopts;

use std::io::prelude::*;
use std::net::TcpStream;
use std::result;
use getopts::Options;

fn write_stream(command: &str) -> Result<usize, std::io::Error>{
    let mut stream = TcpStream::connect("127.0.0.1:8080").unwrap();

    let to_write = command.as_bytes();
    let _ = stream.write(&to_write); // ignore the Result
    let response = stream.read(&mut [0; 128]); // ignore this too*/}
    return response;
}

fn print_usage(program: &str, opts: Options){
    let description = format!("Usage: {} FILE [options]", program);
    println!("{}", opts.usage(&description));
}

fn reset(){
    let response = write_stream(&"reset");
    match response {
        Ok(response) => {
            println!("response { }", response);
        }
        Err(_) => {
            println!("fuck");
        }
    }
}

fn list(){
    println!("list placeholder");
}

fn delete() {
    println!("delete placeholder");
}

fn add_topic_with_path() {
    println!("add topic with path placeholder");
}

fn add_topic_with_script(){
    println!("add topic with script placeholder");
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
    }else if matches.opt_present("t"){
        let topic = matches.opt_str("t").unwrap();
        let path = matches.opt_str("p");
        let script = matches.opt_str("s");


    }else if matches.opt_present("d"){
        delete();
    }else{
        print_usage(&program_name, opts);
    }
}
