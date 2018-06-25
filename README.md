# Just Scored!
Ever wanted to be notified in Slack when somebody scored a goal in the World
Cup? No? Well, that's also fine - this was made mostly to play around with Go
and Lambda so no harm done. But in case you *are* interested - read on!

## Architecture
The project is made of two AWS Lambda functions:
- `goal-checker`, which runs every minute, checks for scored goals in
  currently played match and saves them to a DynamoDB table
- `goal-notifier`, which is triggered on DynamoDB table item creation and
  sends a notification to configured Slack webhook(s)

The third piece of it all is `worlcup` - a very simple API wrapper I made for
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
$ # Deployment
$ serverless deploy
```

And if you want to go _all_ the way, you can fork this repository and plug
it into [CircleCI][circleci] using the existing CI / CD config I built - it
runs tests on each push and deploys the app on each version tag.

## Contributions
Package source code is available at [GitHub][github].

Feel free to use, ask, fork, star, report bugs, fix them, suggest enhancements,
add functionality and point out any mistakes.

## Authors
Developed and maintained by [Paweł Adamczak][pawelad].

Released under [MIT License][license].


[circleci]: https://circleci.com/
[github]: https://github.com/pawelad/just-scored
[license]: https://github.com/pawelad/just-scored/blob/master/LICENSE
[pawelad]: https://github.com/pawelad
[serverless]: https://serverless.com/
[serverless aws credentials]: https://serverless.com/framework/docs/providers/aws/guide/credentials/
[serverless aws docs]: https://serverless.com/framework/docs/providers/aws/guide/quick-start/
