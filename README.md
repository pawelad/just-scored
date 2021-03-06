# Just Scored!
![AWS Lambda](https://img.shields.io/badge/AWS-Lambda-ff9900.svg)
[![Build status](https://img.shields.io/circleci/project/github/pawelad/just-scored.svg)][circleci]
[![Test coverage](https://img.shields.io/coveralls/pawelad/just-scored.svg)][coveralls]
[![Release](https://img.shields.io/github/release/pawelad/just-scored.svg)][github latest release]
[![License](https://img.shields.io/github/license/pawelad/just-scored.svg)][license]

Ever wanted to be notified in Slack when somebody scored a goal in the World
Cup? No? Well, that's also fine - this was made mostly to play around with Go
and Lambda, so no harm done. But in case you *are* interested - read on!

<p align="center" >
    <img src="https://cdn.rawgit.com/pawelad/just-scored/62a25ee9/screenshot.png" alt="Slack message">
</p>

## Architecture
The project is made of two AWS Lambda functions:
- `goal-checker`, which runs every minute, checks for scored goals in
  currently played matches and saves them to a DynamoDB table
- `goal-notifier`, which is triggered on DynamoDB table item creation
  and sends a notification to the configured Slack webhook(s)

The third piece of it all is `worldcup` - a very simple API wrapper I made for
http://worldcup.sfg.io/. I may put it in a separate repository after I
implement all endpoints, but I decided to leave it here at the moment.

## Running it yourself
Given the serverless nature of this project and usage of the awesome
[Serverless framework][serverless], the *entire* deploy comes up to running
`serverless deploy` - 🎉.

If you never used it, then I'd recommend at least skimming through its
[AWS docs][serverless aws docs], but the only thing you *need* to set up
locally are the [AWS credentials][serverless aws credentials] and the Slack
webhook URL exported as a `SLACK_WEBHOOK_URLS` environment variable
(it supports multiple comma separated URLs).

So, all in all, it should look similar to:

```shell
$ # Setup
$ npm install serverless -g
$ serverless config credentials --provider aws --key FOO --secret BAR
$ export SLACK_WEBHOOK_URLS='https://hooks.slack.com/services/...'
$ git clone https://github.com/pawelad/just-scored && cd just-scored
$ make build 
$ # Deployment
$ serverless deploy
```

And if you want to go _all_ the way, you can fork this repository and plug
it into [CircleCI][circleci] - it will use the existing config I built, which
implements a full CI / CD pipeline. It runs tests on each push and deploys the
app on each version tag.

## Contributions
Package source code is available at [GitHub][github].

Feel free to use, ask, fork, star, report bugs, fix them, suggest enhancements,
add functionality and point out any mistakes.

## Authors
Developed and maintained by [Paweł Adamczak][pawelad].

Released under [MIT License][license].


[circleci]: https://circleci.com/gh/pawelad/just-scored
[coveralls]: https://coveralls.io/github/pawelad/just-scored
[github]: https://github.com/pawelad/just-scored
[github latest release]: https://github.com/pawelad/just-scored/releases/latest
[license]: https://github.com/pawelad/just-scored/blob/master/LICENSE
[pawelad]: https://github.com/pawelad
[serverless]: https://serverless.com/
[serverless aws credentials]: https://serverless.com/framework/docs/providers/aws/guide/credentials/
[serverless aws docs]: https://serverless.com/framework/docs/providers/aws/guide/quick-start/
