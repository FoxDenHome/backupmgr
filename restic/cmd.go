package restic

func (r *Repo) Backup() error {
	return r.RunWait("backup", "backup")
}

func (r *Repo) Prune() error {
	return r.RunWait("forget", "prune")
}

func (r *Repo) Mount(target string) error {
	return r.RunWait("mount", "mount", target)
}
