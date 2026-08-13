# Changelog

All notable changes to CasaOS LocalStorage are documented here.

## [0.4.22] - 2026-08-13

### Changed

- Traverse nested `lsblk` trees so filesystems under partitions and LVM logical volumes are represented accurately.
- Preserve the physical parent disk model and path for each storage entry ([CasaOS-LocalStorage #5](https://github.com/alvins82/CasaOS-LocalStorage/pull/5)).

### Fixed

- Report used and available space from the mounted filesystem rather than the allocated physical disk, with coverage for nested mounts and logical volumes ([CasaOS-LocalStorage #5](https://github.com/alvins82/CasaOS-LocalStorage/pull/5)).

### Verification

- `GOOS=linux GOARCH=amd64 go build ./...`
- `GOOS=linux GOARCH=amd64 go test -c ./service`
