extern crate mqtt3;

use std::env;
use std::net::TcpStream;
use std::io::{Read, Write, BufReader, BufWriter};
use std::process::exit;
use mqtt3::{MqttRead, MqttWrite, Packet, Connect, Publish, Subscribe, Protocol, QoS, PacketIdentifier};
use std::sync::Arc;

pub struct MqttClient {
    stream: TcpStream,
    writer: BufWriter<TcpStream>,
    reader: BufReader<TcpStream>,

}

impl MqttClient {
    pub fn new(address: String) -> MqttClient {
        let mut stream = TcpStream::connect(address.as_str()).unwrap();
        let mut writer = BufWriter::new(stream.try_clone().unwrap());
        let mut reader = BufReader::new(stream.try_clone().unwrap());

        let connect_packet = Packet::Connect(Box::new(Connect {
            protocol: Protocol::MQTT(4),
            keep_alive: 30,
            client_id: "rust-mq-example-sub".to_owned(),
            clean_session: true,
            last_will: None,
            username: None,
            password: None
        }));

        writer.write_packet(&connect_packet);
        MqttClient {  stream, writer, reader }
    }
    pub fn subscribe_to_topic (&mut self, topic: String)  {
        // SUBSCRIBE
        let subscribePacket = Packet::Subscribe(Box::new(Subscribe {
            pid: PacketIdentifier(260),
            topics: vec![
                mqtt3::SubscribeTopic { topic_path: topic, qos: QoS::ExactlyOnce }
            ]
        }));
        self.writer.write_packet(&subscribePacket);
        self.writer.flush();
    }
}
