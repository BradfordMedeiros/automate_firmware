extern crate getopts;

use std::fs;
use std::fs::File;
use std::io::prelude::*;
use std::process::Command;
use std::ffi::OsStr;
use std::path::Path;
use getopts::Options;

fn print_usage(program: &str, opts: Options){
   let description = format!("Usage: {} FILE [options]", program);
   println!("{}", opts.usage(&description));
}

fn has_permission(user_hash_path:String, system_hash_path: String) -> bool {
   let mut file_user = File::open(user_hash_path).unwrap();
   let mut contents_user = String::new();
   let _ = file_user.read_to_string(&mut contents_user);

   let mut file_system = File::open(system_hash_path).unwrap();
   let mut contents_system = String::new();
   let _ = file_system.read_to_string(&mut contents_system);

   contents_user == contents_system
}

fn is_script(file_path: &String) -> bool {
   Path::new(file_path).extension()  == Some(OsStr::new("sh"))
}

fn execute_script(path: &str){
   Command::new("sh").arg("-c").arg(path).output().expect("failed to execute script");
}

fn get_scripts_to_execute(directory_path: String) -> Vec<String> {
   let path = Path::new(&directory_path);
   let is_directory = path.is_dir();
   let mut scripts: Vec<String> = Vec::new();

   if is_directory {
      let paths = fs::read_dir(&directory_path).unwrap();
      for path in paths {
         let script  = path.unwrap().path().display().to_string();
         if is_script(&script){
            scripts.push(script);
         }
      }
   }else{
      panic!("{} is not directory", directory_path);
   }
   scripts
}

fn main() {
   let args: Vec<String> = std::env::args().collect();
   let program_name = args[0].clone();

   let mut opts = Options::new();
   opts.optflag("h", "help", "prints help information (this message)");
   opts.optopt("s", "scriptpath", "the directory to from which to execute scripts at the top level", "PATH");
   opts.optopt("k", "key", "(not yet implemented) - optional file to match against before executing the scripts. uUsed for file hash- if unused will  be non-secure",  "KEY_PATH");
   opts.optopt("m", "sys", "(not yet implemented) - optional file to match against before executing script", "SYS_PATH");

   let matches = match opts.parse(&args[1..]){
      Ok(m) => { m }
      Err(f) => { panic!(f.to_string()) }
   };

   if matches.opt_present("h"){
      print_usage(&program_name, opts);
   }

   let key = matches.opt_str("k").unwrap();
   let sys = matches.opt_str("m").unwrap();

   let should_execute = has_permission(key, sys);

   if should_execute {
      let script_path = matches.opt_str("s");
      let scripts = get_scripts_to_execute(script_path.unwrap());

      for script in scripts.iter() {
         execute_script(&script);
      }
   }else{
      panic!("invalid keys");
   }
}
