package restic

func (r *Repo) Backup() error {
	err := r.RunJSON("backup")
	if err != nil {
		return err
	}

	return nil
}
