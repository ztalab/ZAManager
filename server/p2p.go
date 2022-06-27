package server

import (
	"time"

	"github.com/ztalab/ZAManager/pkg/schema"
	"github.com/ztalab/ZAManager/pkg/util/json"

	"github.com/sirupsen/logrus"
	"github.com/ztalab/ZAManager/pkg/confer"
	"github.com/ztalab/ZAManager/pkg/logger"
	"github.com/ztalab/ZAManager/pkg/p2p"
)

func runP2P() error {
	cfg := confer.GlobalConfig()
	// Create a new P2PHost
	p2phost := p2p.NewP2P(cfg.P2P.ServiceDiscoveryID)
	logger.Infof("Completed P2P Setup")
	// Connect to peers with the chosen discovery method
	switch cfg.P2P.ServiceDiscoveryMode {
	case "announce":
		p2phost.AnnounceConnect()
	case "advertise":
		p2phost.AdvertiseConnect()
	default:
		p2phost.AdvertiseConnect()
	}
	logger.Infof("Connected to Service Peers")
	// Join the chat room
	pubsub, err := p2p.JoinPubSub(p2phost, "server_provider", cfg.P2P.ServiceMetadataTopic)
	if err != nil {
		logger.Errorf(nil, "Join PubSub Error: %v", err)
		return err
	}
	logrus.Infof("Successfully joined [%s] P2P channel.", cfg.P2P.ServiceMetadataTopic)
	go startEventHandler(pubsub)
	return nil
}

func startEventHandler(ps *p2p.PubSub) {
	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()

	for {
		select {
		case msg := <-ps.Inbound:
			// Print the received messages to the message box
			logger.Infof("Received message:%s", msg)
		case <-ticker.C:
			// publish
			ps.Outbound <- getServerInfo()
		}
	}
}

func getServerInfo() string {
	result := schema.ServerInfo{
		PeerId: "I'm peer id",
		Addr:   "server.zta.com",
		Port:   5091,
		MetaData: schema.MetaData{
			Ip:   "127.0.0.1",
			Loc:  "nanjing",
			Colo: "China",
		},
		GasPrice: 1000000,
		Type:     schema.ServerTypeProvider,
	}
	return json.MarshalToString(result)
}
