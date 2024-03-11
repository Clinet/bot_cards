package main

import (
	//std necessities
	"os"
	"os/signal"
	"syscall"

	//Logging
	"github.com/JoshuaDoes/logger"

	//Clinet bot framework
	"github.com/Clinet/clinet_bot"
	"github.com/Clinet/clinet_config"
	"github.com/Clinet/clinet_features"

	//Lore Deck's features
	"github.com/Clinet/clinet_features_cards"

	//Lore Deck's services
	"github.com/Clinet/clinet_services_discord"
)

var sock *bot.Bot
var log  *logger.Logger

func doBot() {
	log = logger.NewLogger("bot", verbosity)
	log.Trace("--- doBot() ---")

	//For some reason we don't automatically exit as planned when we return to main()
	defer os.Exit(0)

	log.Info("Loading features...")
	cfg, err := config.LoadConfig(featuresFile, config.ConfigTypeJSON)
	if err != nil {
		log.Error("Error loading features: ", err)
		return
	}

	if cfg.Features == nil || len(cfg.Features) == 0 {
		log.Error("No features found in configuration!")
		return
	}

	log.Debug("Syncing format of features file...")
	cfg.SaveTo(featuresFile, config.ConfigTypeJSON)

	//Initialize Lore Deck using the configuration provided
	log.Info("Initializing instance of Lore Deck...")
	sock = bot.NewBot(cfg)

	//Lore Deck is effectively online at this stage, so defer shutdown in case of errors below
	defer sock.Shutdown()

	//Register the features to handle commands
	log.Debug("Registering features...")
	log.Trace("- cards")
	logFatalError(sock.RegisterFeature(cards.Feature))

	log.Debug("Registering chat services...")
	log.Trace("- discord")
	logFatalError(sock.RegisterFeature(discord.Feature))

	if writeFeaturesTemplate {
		log.Debug("Updating features template...")
		templateCfg := config.NewConfig()
		templateCfg.Features = features.FM.Features
		templateCfg.SaveTo("features.template.json", config.ConfigTypeJSON)
	}

	log.Info("Lore Deck is now online!")

	log.Debug("Waiting for SIGINT syscall signal...")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT)
	<-sc

	log.Info("Good-bye!")
}

func logFatalError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

