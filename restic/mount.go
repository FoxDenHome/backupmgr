package restic

func (r *Repo) Mount(target string) error {
	err := r.RunWait("mount", target)
	if err != nil {
		return err
	}

	return nil
}
