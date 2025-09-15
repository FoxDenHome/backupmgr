package restic

func (r *Repo) Backup() error {
	err := r.RunWait("backup")
	if err != nil {
		return err
	}

	return nil
}
