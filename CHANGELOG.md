# Changelog
All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog][keepachangelog] and this project
adheres to [Semantic Versioning][semver].

## [Unreleased][unreleased]

### Fixed
- Use `DYNAMODB_TABLE` environment variable when connecting to DynamoDB table.

## [v0.0.2][v0.0.2] - 2018-06-20
### Fixed
- Added DynamoDB definition and appropiate IAM to Serverless config file. 

## [v0.0.1][v0.0.1] - 2018-06-19
### Added
- Initial release.
- Added `worldcup` package.
- Added `goal-checker` Lambda function.


[keepachangelog]: http://keepachangelog.com/en/1.0.0/
[semver]: http://semver.org/spec/v2.0.0.html
[unreleased]: https://github.com/pawelad/just-scored/compare/v0.0.2...HEAD
[v0.0.1]: https://github.com/pawelad/just-scored/releases/tag/v0.0.1
[v0.0.2]: https://github.com/pawelad/just-scored/releases/tag/v0.0.2
