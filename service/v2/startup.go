package v2

// InitializeDataLayout restores configured mergerfs mounts before creating
// default directories. The mount point must be empty while CreateMerge runs;
// creating directories concurrently can make a persisted merge fail to come
// back after a restart or package upgrade.
func InitializeDataLayout(mergerFSEnabled bool, restoreMerge func(), ensureDirectories func()) {
	if mergerFSEnabled {
		restoreMerge()
		return
	}

	ensureDirectories()
}
