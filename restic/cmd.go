package restic

import (
	"encoding/json"
	"errors"
	"io"
	"os"
	"os/exec"
)

func (r *Repo) setup(cfg *Config) {
	r.subArgs = make(map[string][]string)
	for cmd, args := range cfg.Args {
		if _, ok := r.subArgs[cmd]; !ok {
			r.subArgs[cmd] = append([]string{"-r", r.URI}, args...)
		}
	}
}

func (r *Repo) makeCmd(json bool, cmd string, args ...string) *exec.Cmd {
	preArgs := r.subArgs[cmd]
	if preArgs == nil {
		preArgs = r.subArgs["default"]
	}

	cmdArgs := []string{cmd}
	if json {
		cmdArgs = append(cmdArgs, "--json")
	}

	cmdArgs = append(cmdArgs, preArgs...)
	cmdArgs = append(cmdArgs, args...)

	eCmd := exec.Command("restic", cmdArgs...)
	eCmd.Env = append(os.Environ(), "RESTIC_PASSWORD="+r.Password)
	return eCmd
}

func (r *Repo) RunJSON(cmd string, args ...string) error {
	eCmd := r.makeCmd(true, cmd, args...)
	stdout, err := eCmd.StdoutPipe()
	if err != nil {
		return err
	}
	eCmd.Stderr = eCmd.Stdout

	if err := eCmd.Start(); err != nil {
		return err
	}

	decoder := json.NewDecoder(stdout)
	go func() {
		for {
			decErr := r.handleNextMessage(decoder)
			if decErr == nil {
				continue
			}
			if errors.Is(decErr, io.EOF) {
				break
			}

			// log.Printf("Error handling restic message: %v", decErr)
		}
	}()

	if err := eCmd.Wait(); err != nil {
		return err
	}

	return nil
}

func (r *Repo) RunWait(cmd string, args ...string) error {
	eCmd := r.makeCmd(false, cmd, args...)
	eCmd.Stdout = os.Stdout
	eCmd.Stderr = os.Stderr
	return eCmd.Run()
}
