# teamcity-monitor

Solution for monitoring realtime status of [JetBrains TeamCity](https://www.jetbrains.com/teamcity/) builds. Used in continuous integration team for quickly respond to any build fail.

Why another monitoring tool:

* Just for fun
* Written in Go (server-side)


# Installation

Download and install:

    % go get github.com/Gr1N/teamcity-monitor
    % cd $GOPATH/src/github.com/Gr1N/teamcity-monitor

(optional) Install [bee](http://beego.me/docs/install/bee.md) tool:

    % go get github.com/beego/bee

Install node dependencies:

    % npm install

Build assets:

    % node build

Build and run:

    % bee run

...or:

    % go build && ./teamcity-monitor


# Deployment

TBD


# License

*teamcity-monitor* is licensed under the MIT license. See the license file for details.
