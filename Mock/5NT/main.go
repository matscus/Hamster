package main

import (
	"encoding/json"
	"flag"
	"io/ioutil"
	"log"
	"math/rand"
	"os"
	"strings"
	"sync"
	"time"

	"github.com/matscus/mq-golang-jms20/jms20subset"
	"github.com/matscus/mq-golang-jms20/mqjms"
)

var (
	scenario  string
	Config    config
	endCh     chan bool
	methodMap sync.Map 
)

type config struct {
	Host     string `json:"Host"`
	Port     int    `json:"Port"`
	Manager  string `json:"Manager"`
	Cannel   string `json:"Cannel"`
	UserName string `json:"UserName"`
	Password string `json:"Password"`
	Data     []struct {
		Operation struct {
			Name     string `json:"name"`
			Queuein  string `json:"queuein"`
			Queueout string `json:"queueout"`
		} `json:"Operation"`
	} `json:"Data"`
}

type Receiver struct {
	Connect    mqjms.ConnectionFactoryImpl
	Operations []operation
}
type operation struct {
	Name     string
	Queuein  string
	Queueout string
}
type IntRange struct {
	min, max int
}

func main() {
	flag.StringVar(&scenario, "scenario", "./config/config.json", "path to config file")
	flag.Parse()
	readConfig(scenario)
	readJSON()
	receiver := Receiver{}
	receiver.Init()
	receiver.Run()
	<-endCh
	//readJSON()
}

func (r *Receiver) Init() {
	r.Connect = mqjms.ConnectionFactoryImpl{
		QMName:      Config.Manager,
		Hostname:    Config.Host,
		PortNumber:  Config.Port,
		ChannelName: Config.Cannel,
		UserName:    Config.UserName,
		Password:    Config.Password,
	}
	var o operation
	for i := 0; i < len(Config.Data); i++ {
		o.Name = Config.Data[i].Operation.Name
		o.Queuein = Config.Data[i].Operation.Queuein
		o.Queueout = Config.Data[i].Operation.Queueout
		r.Operations = append(r.Operations, o)
	}
}

func (r *Receiver) Run() {
	for i := 0; i < len(r.Operations); i++ {
		go receiver(r.Connect, r.Operations[i])
	}
}
func receiver(cf mqjms.ConnectionFactoryImpl, ops operation) {
	context, ctxErr := cf.CreateTLSContext("/home/matscus/IBM/key")
	if ctxErr != nil {
		log.Println("Create Receiver context: ", ctxErr)
	}
	if context != nil {
		defer context.Close()
	}
	queuein := context.CreateQueue(ops.Queuein)
	queueout := context.CreateQueue(ops.Queueout)
	consumer, conErr := context.CreateConsumer(queuein)
	if conErr != nil {
		log.Println("Create Receiver consumer: ", ctxErr)
	}
	if consumer != nil {
		defer consumer.Close()
	}
	for {
		msg, jmsErr := consumer.ReceiveNoWait()
		if jmsErr != nil {
			log.Println("Receive receiver msg: ", jmsErr)
		}
		if msg != nil {
			r := rand.New(rand.NewSource(55))
			ir := IntRange{300, 600}
			time.Sleep(time.Duration(ir.NextRandom(r)) * time.Millisecond)
			switch msg := msg.(type) {
			case jms20subset.TextMessage:
				method, err := msg.GetStringProperty("esfl_methodName")
				checkError(err, "Get string property")
				pt, _ := methodMap.Load(method)
				if pt != nil {
					msg.SetText(pt.(string))
					context.CreateProducer().Send(queueout, msg)
				} else {
					log.Println("not correct metod for jms header")
				}

			default:
				jmsErr = jms20subset.CreateJMSException(
					"Received message is not a TextMessage", "MQJMS6068", nil)
			}
		}
	}
}

func readConfig(configpath string) {
	fl, err := os.Open(configpath)
	checkError(err, "error open config:")
	defer fl.Close()
	err = json.NewDecoder(fl).Decode(&Config)
	checkError(err, "error decode config:")
}
func checkError(err error, str string) {
	if err != nil {
		log.Println(str, err)
	}
}
func checkJMSError(err jms20subset.JMSException, str string) {
	if err != nil {
		log.Println(str, err)
	}
}
func (ir *IntRange) NextRandom(r *rand.Rand) int {
	return r.Intn(ir.max-ir.min+1) + ir.min
}

func readJSON() {
	files, err := ioutil.ReadDir("./json")
	if err != nil {
		log.Fatal(err)
	}

	for _, f := range files {
		file, err := ioutil.ReadFile("./json/" + f.Name())
		checkError(err, "Open json error: ")
		body := string(file)
		r := strings.Replace(f.Name(), ".json", "", 1)
		methodMap.Store(r, body)
	}
}
