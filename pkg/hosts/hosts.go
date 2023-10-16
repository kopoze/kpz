package hosts

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"time"

	"github.com/kopoze/kpz/pkg/app"
	"github.com/kopoze/kpz/pkg/config"
	"github.com/txn2/txeh"
)

func InitDomain() {
	conf := config.LoadConfig()
	log.Printf("Add %s to hosts", conf.Kopoze.Domain)

	hosts, err := txeh.NewHostsDefault()
	if err != nil {
		log.Panic(err)
	}

	backupHosts(hosts)
	hosts.AddHost("127.0.0.1", conf.Kopoze.Domain)
	err = hosts.Save()
	if err != nil {
		log.Panic(err)
	}
}

func SetDomain() {
	var old_domain string
	conf := config.LoadConfig()
	curr_domain := conf.Kopoze.Domain
	old_domain, err := config.GetOldDomain()
	if err != nil && !os.IsNotExist(err) {
		log.Panic(err)
	}
	hosts, err := txeh.NewHostsDefault()
	if err != nil {
		log.Fatal(err)
	}

	if os.IsNotExist(err) || old_domain == "" {
		config.SetOldDomain(curr_domain)
	}

	if curr_domain != old_domain {
		log.Println("Updating existing domain")
		backupHosts(hosts)

		hosts.RemoveHost(old_domain)
		hosts.AddHost("127.0.0.1", curr_domain)

		var apps []app.App
		app.ConnectDB()
		app.DB.Find(&apps)
		for _, currApp := range apps {
			hosts.RemoveHost(fmt.Sprintf("%s.%s", currApp.Subdomain, old_domain))
			hosts.AddHost("127.0.0.1", fmt.Sprintf("%s.%s", currApp.Subdomain, curr_domain))
		}

		config.SetOldDomain(curr_domain)
		// TODO: Update nginx config
		// TODO: Dynamic reload server
		// TODO: Add redirect from old_domain to new_domain
		log.Printf("Updating hosts successfully.\n\nTo finalize the configuration update, you need to: \n1. manually change your nginx config file value: \n\tfrom\t'*.%s' \n\tto\t'*.%s'\n2. Restart server", old_domain, curr_domain)
	}
}

func AddSubdomain(subdomain string) {
	log.Printf("Adding %s to hosts", subdomain)
	conf := config.LoadConfig()
	hosts, err := txeh.NewHostsDefault()
	if err != nil {
		log.Panic(err)
	}

	backupHosts(hosts)
	hosts.AddHost("127.0.0.1", fmt.Sprintf("%s.%s", subdomain, conf.Kopoze.Domain))
	err = hosts.Save()
	if err != nil {
		log.Panic(err)
	}
}

func RemoveSubdomain(subdomain string) {
	log.Printf("Removing %s from hosts", subdomain)
	conf := config.LoadConfig()
	hosts, err := txeh.NewHostsDefault()
	if err != nil {
		log.Panic(err)
	}

	backupHosts(hosts)

	hosts.RemoveHost(fmt.Sprintf("%s.%s", subdomain, conf.Kopoze.Domain))
	err = hosts.Save()
	if err != nil {
		log.Panic(err)
	}
}

func backupHosts(h *txeh.Hosts) {
	backupHostPath := filepath.Join(config.GetConfigPath(), "hosts", "backup")
	if _, err := os.Stat(backupHostPath); os.IsNotExist(err) {
		os.MkdirAll(backupHostPath, 0755)
	}

	file := filepath.Join(backupHostPath, buildFilename())
	h.SaveAs(file)
}

func buildFilename() string {
	curr_date := time.Now()
	formated_date := fmt.Sprintf("%d-%02d-%d_%d-%d-%d", curr_date.Year(), curr_date.Month(), curr_date.Day(), curr_date.Hour(), curr_date.Minute(), curr_date.Second())
	return fmt.Sprintf("hosts_%s", formated_date)
}
