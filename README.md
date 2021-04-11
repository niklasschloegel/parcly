![Parcly Image](resources/parcly_image.jpg)

# Parcly

Parcly is a CLI app for tracking parcels.

This app is built with [cobra](https://github.com/spf13/cobra) and [Tracktry API](https://www.tracktry.com/).

## Getting started

### Installation

#### macOS Homebrew

```bash
brew tap niklasschloegel/parcly; brew install niklasschloegel/parcly/parcly
```

#### With go install

```bash
go install github.com/niklasschloegel/parcly@latest
```

### Setting the API key

To use this app, you need to create a free account at [Tracktry](https://www.tracktry.com).
Your API key can be found under Settings/API.

The API Key can be provided in three ways:

1. As a flag:

```bash
parcly <noun> <command> --tracktrykey <key>
```

2. As an environment variable:

```bash
export PARCLY_TRACKTRYKEY=<key>
```

3. In a config file:

default config file is $HOME/.parcly.yaml and should contain:

```yaml
tracktrykey: <key>
```

When you want to use another location, you can
specify the location with another flag:

```bash
parcly ... --config <filepath>
```

4. Through the config set command:

```bash
parcly config set --tracktry key
```

## Example usage

You first have to add a tracking item:

```bash
parcly tracking add <trackingNr> --carrier dhl-germany
```

The tracking item now gets saved to Tracktry and can get requested afterwards with

```bash
parcly tracking list
```

For more information you can just type

```bash
parcly help
```
