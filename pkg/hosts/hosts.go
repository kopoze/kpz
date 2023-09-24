package hosts

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"time"

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
