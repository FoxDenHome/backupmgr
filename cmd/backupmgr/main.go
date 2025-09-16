package main

import (
	"flag"
	"log"
	"os"

	"github.com/FoxDenHome/backupmgr/restic"
	"github.com/FoxDenHome/backupmgr/util"
)

func main() {
	log.Printf("backupmgr version %s (git rev %s)", util.GetVersion(), util.GetGitRev())

	configFile := os.Getenv("BACKUPMGR_CONFIG")
	if configFile == "" {
		configFile = "/etc/backupmgr/config.json"
	}

	config, err := restic.LoadConfig(configFile)
	if err != nil {
		log.Fatalf("Failed to load config %s: %v", configFile, err)
	}

	targetRepo := flag.String("repo", "*", "Name of the repository to use (* for all)")
	cmdMode := flag.String("mode", "help", "Mode to run in (backup, mount, help)")
	flag.Parse()

	switch *cmdMode {
	case "backup":
		forTargetRepo(true, *targetRepo, config, func(name string, repo *restic.Repo) {
			log.Printf("Backing up repo %s", name)
			err := repo.Backup()
			if err != nil {
				log.Printf("Failed to back up repo %s: %v", name, err)
			}
		})
	case "mount":
		forTargetRepo(false, *targetRepo, config, func(name string, repo *restic.Repo) {
			log.Printf("Mounting repo %s", name)
			err := repo.Mount(flag.Arg(0))
			if err != nil {
				log.Printf("Failed to mount repo %s: %v", name, err)
			}
		})
	case "help":
		flag.Usage()
		return
	default:
		log.Fatalf("Unknown mode: %s", *cmdMode)
	}

	log.Printf("All done!")
}

func forTargetRepo(allowMulti bool, target string, config *restic.Config, fn func(name string, repo *restic.Repo)) {
	if target == "*" {
		for name, repo := range config.Repos {
			fn(name, repo)
			if !allowMulti {
				break
			}
		}
	} else {
		repo, ok := config.Repos[target]
		if !ok {
			log.Fatalf("No such repo: %s", target)
		}
		fn(target, repo)
	}
}
