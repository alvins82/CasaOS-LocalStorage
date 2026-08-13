# Changelog

All notable changes to CasaOS LocalStorage are documented here.

## [0.4.24] - 2026-08-13

### Added

- Create the standard `Documents`, `Downloads`, `Gallery`, and `Media` directories in `/DATA` when an external merged storage is created and they are missing ([CasaOS-LocalStorage #7](https://github.com/alvins82/CasaOS-LocalStorage/pull/7)).

### Changed

- Reuse the default-directory helper during startup and merged-storage recovery while preserving the special system `AppData` compatibility mount.

### Verification

- Focused default-directory tests pass.
- Linux cross-compilation succeeds for the service and root package.

## [0.4.23] - 2026-08-13

### Added

- Add a protected `PUT /v1/storage/rename` endpoint for ext2/ext3/ext4 volumes and keep system storage protected from renaming ([CasaOS-LocalStorage #6](https://github.com/alvins82/CasaOS-LocalStorage/pull/6)).

### Fixed

- Read the filesystem label directly with `blkid` when `lsblk` has not refreshed udev data yet, so a successful rename is reflected immediately in Storage Manager ([CasaOS-LocalStorage #6](https://github.com/alvins82/CasaOS-LocalStorage/pull/6)).

### Verification

- `go generate ./...`
- `GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build ./...`
- Invalid-device API validation probe completed successfully.

## [0.4.22] - 2026-08-13

### Changed

- Traverse nested `lsblk` trees so filesystems under partitions and LVM logical volumes are represented accurately.
- Preserve the physical parent disk model and path for each storage entry ([CasaOS-LocalStorage #5](https://github.com/alvins82/CasaOS-LocalStorage/pull/5)).

### Fixed

- Report used and available space from the mounted filesystem rather than the allocated physical disk, with coverage for nested mounts and logical volumes ([CasaOS-LocalStorage #5](https://github.com/alvins82/CasaOS-LocalStorage/pull/5)).

### Verification

- `GOOS=linux GOARCH=amd64 go build ./...`
- `GOOS=linux GOARCH=amd64 go test -c ./service`
