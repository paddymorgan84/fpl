<h1 align="center"> FPL </h1> <br>
<p align="center">
  <img alt="GitPoint" title="GitPoint" src="img/header.png" width="450">
</p>

<p align="center">
  A Go tool for Fantasy Premier League
</p>

<!-- START doctoc generated TOC please keep comment here to allow auto update -->
<!-- DON'T EDIT THIS SECTION, INSTEAD RE-RUN doctoc TO UPDATE -->
## Table of Contents

- [Introduction](#introduction)
- [Pre-requisites](#pre-requisites)
- [Installation](#installation)
- [Configuration](#configuration)
  - [Configuring your team ID](#configuring-your-team-id)
- [Acknowledgements](#acknowledgements)

<!-- END doctoc generated TOC please keep comment here to allow auto update -->

## Introduction

![Go](https://github.com/paddymorgan84/fpl/workflows/fpl/badge.svg)
[![Go Report Card](https://goreportcard.com/badge/github.com/paddymorgan84/fpl?style=flat-square)](https://goreportcard.com/report/github.com/paddymorgan84/fpl)

`fpl` is a CLI tool written in Go that pulls information from the Fantasy Premier League API and presents it in a readable format.

```bash
A CLI tool for retrieving FPL data

Usage:
  fpl [command]

Available Commands:
  details     Returns details of manager for current season, e.g. league standings, cash in the bank, overall points etc
  fixtures    Get the fixtures for a specific gameweek
  help        Help about any command
  history     Returns history for a managers current and past seasons
  points      Get the points for a specified gameweek (defaults to latest active gameweek)

Flags:
      --config string     config file (default is $HOME/.fpl.yaml)
  -g, --gameweek string   The gameweek you wish to see the fixtures for
  -h, --help              help for fpl
  -t, --toggle            Help message for toggle

Use "fpl [command] --help" for more information about a command.
```

## Pre-requisites

- [Go version 1.16 or higher](https://golang.org/dl/)

## Installation

Installing the tool is very straightforward:

```bash
git clone https://github.com/paddymorgan84/fpl.git
cd fpl
go install
```

## Configuration

`fpl` accepts a config file called `.fpl.yaml`, by default it will search your home directory.

### Configuring your team ID

Setting a default team id prevents the need for you to provide it when you want to use `fpl` for your own team.

```bash
team-id: 1327470
```

You can get your FPL team id via the Gameweek History page, linked on your FPL squad page. Simply click that link and check the URL displayed in your browser. You should see something like this:

`http://fantasy.premierleague.com/entry/xxxx/history/`

## Acknowledgements

- Shoutout to [Frenzel Timothy](https://medium.com/@frenzelts) for his [how-to guide](https://medium.com/@frenzelts/fantasy-premier-league-api-endpoints-a-detailed-guide-acbd5598eb19) on using the Fantasy Premier League API 👏
- I've used [Cobra](https://github.com/spf13/cobra) to help me create the scaffolding for a go CLI tool. I would highly recommend it 👍