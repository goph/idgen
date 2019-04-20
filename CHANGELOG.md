# Change Log


All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](http://keepachangelog.com/en/1.0.0/)
and this project adheres to [Semantic Versioning](http://semver.org/spec/v2.0.0.html).


## [Unreleased]

### Added

- List generator


## [0.3.0] - 2019-03-11

### Added

- UUID support (using [github.com/gofrs/uuid](https://github.com/gofrs/uuid))

### Changed

- Add functional options to ULID generator


## [0.2.0] - 2019-03-08

### Added

- New `Generator` interface

### Changed

- Renamed `New` method to `Generate`
- Renamed `Generator` interface to `SafeGenerator`


## 0.1.0 - 2019-02-11

- Initial release


[Unreleased]: https://github.com/goph/idgen/compare/v0.3.0...HEAD
[0.3.0]: https://github.com/goph/idgen/compare/v0.2.0...v0.3.0
[0.2.0]: https://github.com/goph/idgen/compare/v0.1.0...v0.2.0
