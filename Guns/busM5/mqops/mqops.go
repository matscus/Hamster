package mqops

import (
	"log"
	"math"
	"math/rand"
	"net/http"
	"sync"
	"time"

	"../asserts"
	"../errors"

	"github.com/matscus/mq-golang-jms20/jms20subset"
	"github.com/matscus/mq-golang-jms20/mqjms"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var (
	letters   = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890!@#$%^&*()")
	timestamp sync.Map
	//MQOps global slice operation
	MQOps        []*Operation
	responseTime = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "mqops_response_time",
			Help: "Operation response time",
		},
		[]string{"metod", "operation"},
	)
	produserCounter = promauto.NewCounter(prometheus.CounterOpts{
		Name: "mqops_produser_count",
		Help: "produser_count_message",
	})
	consumerCounter = promauto.NewCounter(prometheus.CounterOpts{
		Name: "mqops_consumer_count",
		Help: "consumer_count_message",
	})
)

// Operation - struct for test operation
type Operation struct {
	Name           string
	Produser       int
	Consumer       int
	Receiver       int
	Step           int
	Rumpup         int
	Throughput     int
	Queuein        string
	Queueout       string
	Connect        mqjms.ConnectionFactoryImpl
	PoolCh         chan string
	StringProperty map[string]string
	Body           string
}

//Action - interface for test operation
type Action interface {
	InitMQOps()
	RunProduser()
	RunConsumer()
	RunReceiver()
}

func init() {
	prometheus.MustRegister(responseTime)
	http.Handle("/metrics", promhttp.Handler())
	log.Println("Listenport: 9900")
	go http.ListenAndServe(":9900", nil)
}

//New - return new struct Operation
func New() *Operation {
	return new(Operation)
}

//NewCh - return new Ch
func NewCh(i int) chan string {
	ch := make(chan string, i)
	return ch
}

//InitMQOps - init operation struct
func (o *Operation) InitMQOps() {
	operationsCnt := len(MQConf.Data.Operation)
	for i := 0; i < operationsCnt; i++ {
		if MQConf.Data.Operation[i].Data.Name == o.Name {
			o.Connect = mqjms.ConnectionFactoryImpl{
				QMName:      MQConf.Data.Operation[i].Data.Manager,
				Hostname:    MQConf.Data.Operation[i].Data.Host,
				PortNumber:  MQConf.Data.Operation[i].Data.Port,
				ChannelName: MQConf.Data.Operation[i].Data.Cannel,
				UserName:    MQConf.Data.Operation[i].Data.UserName,
				Password:    MQConf.Data.Operation[i].Data.Password,
			}
			o.Queuein = MQConf.Data.Operation[i].Data.Queuein
			o.Queueout = MQConf.Data.Operation[i].Data.Queueout
			if !MQConf.Data.Operation[i].Data.Defaultparams {
				o.Produser = getRoundCountOps(MQConf.Data.Defaultproduser, operationsCnt)
				o.Consumer = getRoundCountOps(MQConf.Data.Defaultconsumer, operationsCnt)
				o.Receiver = getRoundCountOps(MQConf.Data.Defaultreceiver, operationsCnt)
				o.Step = MQConf.Data.Defaultstep
				o.Rumpup = MQConf.Data.Defaultrumpup
				o.Throughput = MQConf.Data.Defaultthroughput
			} else {
				o.Produser = MQConf.Data.Operation[i].Data.Produser
				o.Consumer = MQConf.Data.Operation[i].Data.Consumer
				o.Receiver = MQConf.Data.Operation[i].Data.Receiver
				o.Step = MQConf.Data.Operation[i].Data.Step
				o.Rumpup = MQConf.Data.Operation[i].Data.Rumpup
				o.Throughput = MQConf.Data.Operation[i].Data.Throughput
			}
			break
		}
	}
}

//InitMQOps - init operation struct
func InitMQOps(o Action) {
	o.InitMQOps()
}

//RunProduser - create produser and run operation send to queue
func (o *Operation) RunProduser() {
	s := float64(o.Produser) / float64(o.Step)
	lenStep := int(math.Round(s))
	if lenStep == 0 {
		lenStep = 1
	}
	stepDuration := (time.Duration(o.Rumpup / lenStep)) * time.Second
	users := o.Produser / lenStep
	for i := 0; i < lenStep; i++ {
		for i := 0; i <= users; i++ {
			go producer(o.Connect, o.Queuein, o.Throughput, o.PoolCh, o.StringProperty)
		}
		time.Sleep(stepDuration)
	}
}

//RunConsumer - create  consumer and run operation read to queue
func (o *Operation) RunConsumer() {
	s := float64(o.Consumer) / float64(o.Step)
	lenStep := int(math.Round(s))
	if lenStep == 0 {
		lenStep = 1
	}
	stepDuration := (time.Duration(o.Rumpup / lenStep)) * time.Second
	users := o.Consumer / lenStep
	for i := 0; i < lenStep; i++ {
		for i := 0; i <= users; i++ {
			go consumer(o.Connect, o.Queueout, o.Name)
		}
		time.Sleep(stepDuration)
	}
}

//RunProduser - fot interface
func RunProduser(o Action) {
	o.RunProduser()
}

//RunConsumer - fot interface
func RunConsumer(o Action) {
	o.RunConsumer()
}

//RunResiever - fot interface
func RunResiever(o Action) {
	o.RunReceiver()
}

//RunReceiver - run receiver(stub)
func (o *Operation) RunReceiver() {

}
func producer(cf mqjms.ConnectionFactoryImpl, q string, t int, bodych chan string, stringProperty map[string]string) {
	throughput := convertThroughput(t)
	context, ctxErr := cf.CreateContext()
	errors.CheckJMSError(ctxErr, "Create context: ")
	if context != nil {
		defer context.Close()
	}
	queue := context.CreateQueue(q)
	for {
		t := time.Now()
		collerationID := getUUID(12)
		msg := context.CreateTextMessage()
		msg.SetJMSCorrelationID(collerationID)
		body := <-bodych
		msg.SetText(body)
		produser := context.CreateProducer()
		for k, v := range stringProperty {
			produser.SetStringProperty(k, v)
		}
		err := produser.Send(queue, msg)
		errors.CheckJMSError(err, "Create send error: ")
		produserCounter.Inc()
		pt := msg.GetJMSTimestamp()
		timestamp.Store(collerationID, pt)
		t2 := time.Since(t)
		if t2 < throughput {
			sleep := throughput - t2
			time.Sleep(sleep)
		}
	}
}
func consumer(cf mqjms.ConnectionFactoryImpl, q string, n string) {
	context, ctxErr := cf.CreateContext()
	errors.CheckJMSError(ctxErr, "Create consumer context: ")
	if context != nil {
		defer context.Close()
	}
	queue := context.CreateQueue(q)
	consumer, conErr := context.CreateConsumer(queue)
	errors.CheckJMSError(conErr, "Create consumer: ")
	if consumer != nil {
		defer consumer.Close()
	}
	for {
		t1 := time.Now()
		msg, jmsErr := consumer.ReceiveNoWait()
		errors.CheckJMSError(jmsErr, "Receive consumer msg: ")
		if msg != nil {
			id := msg.GetJMSCorrelationID()
			pt, ok := timestamp.Load(id)
			var produserTimestamp, consumerTimeStamp, res int64
			var response float64
			if ok {
				produserTimestamp = pt.(int64)
				consumerTimeStamp = msg.GetJMSTimestamp()
				res = (consumerTimeStamp - produserTimestamp)
				t := time.Duration(res) * time.Millisecond
				response = t.Seconds()
				timestamp.Delete(id)
			}
			switch msg := msg.(type) {
			case jms20subset.TextMessage:
				consumerCounter.Inc()
				var msgBodyStrPtr *string
				msgBodyStrPtr = msg.GetText()
				text := *msgBodyStrPtr
				go asserts.CheckAssert(text, n)
			default:
				jmsErr = jms20subset.CreateJMSException(
					"Received message is not a TextMessage", "MQJMS6068", nil)
			}
			t2 := time.Now()
			responseConsumer := t2.Sub(t1).Seconds()
			responseTime.With(prometheus.Labels{"metod": "totaltime", "operation": n}).Set(response)
			responseTime.With(prometheus.Labels{"metod": "consumer", "operation": n}).Set(responseConsumer)
		}
	}
}
func getUUID(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}
func convertThroughput(i int) time.Duration {
	return time.Duration(i) * time.Millisecond
}
func convertDurationToFloat64(d time.Duration) (res float64) {
	return float64(d * time.Millisecond)
}
func getRoundCountOps(i int, i2 int) int {
	return int(math.Round(float64(i) / float64(i2)))
}
