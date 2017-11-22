package main


type subscription struct {
	uuid string
	path string
	is_script bool	// if not script is file path, how to do enums?
}

type mqtt_manager struct {
	topic_subscriptions map[string] subscription
}

func New_mqtt_manager () mqtt_manager {
	//subscriptions := make([]subscription, 0)
	subscriptions := make(map[string] subscription )
	return mqtt_manager{ topic_subscriptions: subscriptions }
}

func (manager mqtt_manager) add_script_subscription(topic string, file_path string) {
	_, ok := manager.topic_subscriptions[topic]
	if ok {

	}else{

	}

}

func (manager mqtt_manager) add_file_subscription(topic string, file_path string){

}

func (manager mqtt_manager) remove_subscription(uuid string){

}

func (manager mqtt_manager) reset(){

}

func (manager mqtt_manager) list_subscription() string {
	return "hello"
}

func (manager mqtt_manager) on_mqtt_message() {

}

