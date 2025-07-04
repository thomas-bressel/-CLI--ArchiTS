# Changelog

All notable changes to this project will be documented in this file.

## [Unreleased]

## [1.8.1] - 2025-06-08
### Fixed
- Refacto command variable for the CLI into command.go file.
- Refecto type of structure intro a new folder called models. 

## [1.8.0] - 2025-06-06
### Added
- New command archi version to display versions of Linux, Go, Node and Archi TS

## [1.7.0] - 2025-06-05
### Added
- Generation of index.ts or index.js file depending on Express or vanilla Node.js choice.

## [1.6.3] - 2025-06-05
### Fixed
- Fixed npm.go file by adding condition for Express installation in package.json
- Added confirmation display for Express installation status in the installation report

## [1.6.2] - 2025-06-04
### Added
- Execution of npm init after configuration validation
- Execution of dependencies installation
- Execution of pnpm installation

### Fixed
- Added automatic Go module initialization (go mod init) if go.mod file is missing
- Added automatic dependencies installation via go mod tidy during make build
- Fixed errors encountered during first make install on freshly cloned repository
- Removed npm init execution (redundant since package.json is already generated)
- Added PayPal donation link for those who wish to support this project

*Thanks to Mahery Randrianirina for the helpful feedback! ;)*

## [1.5.2] - 2025-06-03
### Added
- Generate tsconfig.json file if TypeScript is selected
- Generate .gitignore file
- Generate jest.config file with different configuration for TS or JS
- Add new prompt to choose Express library installation

### Changed
- Updated package.json template and logic to add dev dependencies in correct order
- Added Makefile to automate installation and configuration

## [1.4.0] - 2025-06-02
### Added
- Generate package.json file with the selected language type
- Generate .env file template

## [1.3.0] - 2025-05-31
### Added
- Added prompt to choose language type between TypeScript and JavaScript
- Added 3rd architecture choice to the architecture prompt
- Create folder structure based on the selected architecture type

## [1.1.0] - 2025-05-30
### Added
- Added first prompt to enter project name with `./archi create` command

## [1.0.0] - 2025-05-30
### Added
- Started project development
- Set up initial project structure