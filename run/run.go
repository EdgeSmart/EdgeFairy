package run

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"net"
	"net/http"
	"os"
	"strings"

	MQTT "github.com/eclipse/paho.mqtt.golang"
)

type topicItem struct {
	topic   string
	qos     byte
	handler MQTT.MessageHandler
}

type ProxyStruct struct {
	ResponseTopic string
	Data          []byte
}

var (
	mqttClient MQTT.Client
	clientID   string
	mqttServe  chan os.Signal
	topicList  = []topicItem{
		topicItem{topic: "%s/DockerAPI", qos: 0, handler: dockerAPIHandler},
		topicItem{topic: "gateway_broadcast", qos: 0, handler: broadcastHandler},
	}
)

func Run(clusterKey string, clusterToken string, server string) {
	mqttServe = make(chan os.Signal)
	hostname, err := os.Hostname()
	if err != nil {
		return
	}
	clientID = fmt.Sprintf("gateway/%s/%s", clusterKey, hostname)
	opts := MQTT.NewClientOptions().AddBroker(server)
	opts.SetClientID(clientID)
	opts.SetUsername(clientID)
	opts.SetPassword(clusterToken)

	mqttClient = MQTT.NewClient(opts)
	if token := mqttClient.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}

	opts.SetDefaultPublishHandler(defaultHandler)
	for i := range topicList {
		topic := topicList[i].topic
		if strings.Count(topic, "%s") == 1 {
			topic = fmt.Sprintf(topic, clusterKey)
		}
		if token := mqttClient.Subscribe(topic, topicList[i].qos, topicList[i].handler); token.Wait() && token.Error() != nil {
			fmt.Println(token.Error())
			os.Exit(1)
		}
	}

	// publish register message
	token := mqttClient.Publish("Manager/gateway_register", 0, false, "")
	token.Wait()

	quitSignal := <-mqttServe
	fmt.Println("Quit", quitSignal)
}

func defaultHandler(client MQTT.Client, msg MQTT.Message) {
	fmt.Printf("TOPIC: %s\n", msg.Topic())
	fmt.Printf("MSG: %s\n", msg.Payload())
}

func dockerAPIHandler(client MQTT.Client, msg MQTT.Message) {
	// var GolngGob bytes.Buffer
	payload := bytes.NewReader(msg.Payload())
	dec := gob.NewDecoder(payload)
	proxyData := ProxyStruct{}
	gob.Register(http.NoBody)
	err := dec.Decode(&proxyData)
	if err != nil {
		fmt.Println(err)
		return
	}
	addr, err := net.ResolveUnixAddr("unix", "/var/run/docker.sock")
	if err != nil {
		panic("Cannot resolve unix addr: " + err.Error())
	}
	c, err := net.DialUnix("unix", nil, addr)
	if err != nil {
		panic("DialUnix failed.")
	}

	_, err = c.Write(proxyData.Data)
	if err != nil {
		panic("Writes failed.")
	}

	buf := make([]byte, 5120)
	realLen, err := c.Read(buf)
	realData := buf[0:realLen]
	if err != nil {
		panic("Read: " + err.Error())
	}
	token := mqttClient.Publish(proxyData.ResponseTopic, 0, false, realData)
	token.Wait()
}

func broadcastHandler(client MQTT.Client, msg MQTT.Message) {
	fmt.Printf("TOPIC: %s\n", msg.Topic())
	fmt.Printf("MSG: %s\n", msg.Payload())
}
