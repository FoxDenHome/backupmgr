package restic

func (r *Repo) Backup() error {
	err := r.RunWait("backup")
	if err != nil {
		return err
	}

	return nil
}

func (r *Repo) Prune() error {
	err := r.RunWait("prune")
	if err != nil {
		return err
	}

	return nil
}

func (r *Repo) Mount(target string) error {
	err := r.RunWait("mount", target)
	if err != nil {
		return err
	}

	return nil
}
