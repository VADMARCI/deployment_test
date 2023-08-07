package nats

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"runtime/debug"
	"time"

	"github.com/nats-io/nats.go"
	"github.com/pepusz/go_redirect/messaging"
	"github.com/pepusz/go_redirect/utils"
)

type Gateway struct {
	NC            *nats.Conn
	JS            nats.JetStreamContext
	Subscriptions []*nats.Subscription
}

func NewGateway() Gateway {
	g := Gateway{}
	g.connect()
	return g
}

func (g *Gateway) connect() error {
	servers := utils.GetEnvString("NATS_SERVERS")
	devMode := utils.GetEnvBool("DEV")

	if devMode {
		servers = "nats://ns1:4222, nats://ns2:4222, nats://ns3:4222"
	}
	if servers == "" {
		servers = "nats://nats:4222"
	}

	nc, err := nats.Connect(servers,
		nats.Timeout(5*time.Second),
		nats.DisconnectErrHandler(func(conn *nats.Conn, err error) {
			log.Println("NATS disconnected:", err)
		}),
		nats.ReconnectHandler(func(conn *nats.Conn) {
			log.Println("NATS reconnected to:", conn.ConnectedUrl())
		}),
		nats.ErrorHandler(func(conn *nats.Conn, _ *nats.Subscription, err error) {
			log.Println("NATS error:", err)
		}),
		nats.MaxReconnects(100),
		nats.ReconnectWait(3*time.Second),
		nats.RetryOnFailedConnect(true))
	if err != nil {
		return fmt.Errorf("Got an error on Connect with Secure Options: %s\n", err)
	}
	g.NC = nc
	js, err := nc.JetStream()
	g.JS = js
	return nil
}

func (g *Gateway) IsConnected() bool {
	return g.NC.IsConnected()
}

func (g *Gateway) Publish(subject messaging.Subject, data interface{}) (*nats.PubAck, error) {
	dataJSON, _ := json.Marshal(data)
	log.Print("sending data on nats: ", subject)
	return g.JS.Publish(string(subject), dataJSON)
}

func (g *Gateway) CreateStream(stream string, subjects []messaging.Subject) {
	subjectsString := []string{}
	for _, subject := range subjects {
		subjectsString = append(subjectsString, string(subject))
	}
	info, _ := g.JS.StreamInfo(stream)

	if info == nil {
		_, err := g.JS.AddStream(&nats.StreamConfig{
			Name:     stream,
			Subjects: subjectsString,
			Replicas: len(g.NC.DiscoveredServers()),
		})
		if err != nil {
			log.Println("INFO CREATE ERROR", err)
		}
	} else {
		_, err := g.JS.UpdateStream(&nats.StreamConfig{
			Name:     stream,
			Subjects: subjectsString,
			Replicas: len(g.NC.DiscoveredServers()),
		})
		if err != nil {
			log.Println("INFO UPDATE ERROR", err)
		}
	}
}

func (g *Gateway) QueueSubscribe(subSubjectName messaging.Subject, queue string, durable string, streamName string, handler nats.MsgHandler) {
	_, err := g.checkStreamPresence(streamName, 0)
	if err != nil {
		log.Println("Cannot start because", err)
		return
	}
	sub, err := g.JS.QueueSubscribe(string(subSubjectName), queue, handler,
		nats.MaxDeliver(1),
		nats.AckExplicit(),
		nats.DeliverNew(),
		nats.Durable(durable),
		nats.BindStream(streamName),
	)

	if err != nil && (err != nats.ErrTimeout && err != context.DeadlineExceeded) {
		log.Println("FATAL", subSubjectName, queue, ":", err)
	}
	g.Subscriptions = append(g.Subscriptions, sub)
}
func (g *Gateway) Subscribe(subSubjectName messaging.Subject, durable string, streamName string, handler nats.MsgHandler) {
	_, err := g.checkStreamPresence(streamName, 0)
	if err != nil {
		log.Println("Cannot start because", err)
		return
	}
	sub, err := g.JS.Subscribe(string(subSubjectName), handler,
		nats.MaxDeliver(1),
		nats.AckExplicit(),
		nats.DeliverNew(),
		nats.Durable(durable),
		nats.BindStream(streamName),
	)

	if err != nil && (err != nats.ErrTimeout && err != context.DeadlineExceeded) {
		log.Println("FATAL", subSubjectName, durable, ":", err)
	}
	g.Subscriptions = append(g.Subscriptions, sub)
}

func (g *Gateway) checkStreamPresence(streamName string, counter int) (int, error) {
	info, _ := g.JS.StreamInfo(streamName)
	if counter > 100 {
		return -1, fmt.Errorf("Timeout error after 10 try")
	}
	if info == nil {
		log.Println("NO stream so sleep: ", streamName)
		time.Sleep(300 * time.Millisecond)
		counter++
		g.checkStreamPresence(streamName, counter)
		return counter, nil
	}
	return -1, nil
}

func (g *Gateway) Stop() {
	for _, subscription := range g.Subscriptions {
		subscription.Unsubscribe()
	}
	g.NC.Close()
}

func RecoverFunc(subject string) {
	if r := recover(); r != nil {
		log.Printf("%s - Recovering from panic in %s error is: %v \n", time.Now().Format("2006-01-02 15:04:05.000000"), subject, r)
		log.Printf("Stack trace: \n%s", debug.Stack())
	}
}
