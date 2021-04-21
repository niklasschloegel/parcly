/*
Copyright © 2021 Niklas Schlögel <niklasschloegel@gmail.com>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package config

const (
	BasePath          = "https://api.tracktry.com/v1"
	TrackTryConfigKey = "tracktrykey"
)

var TrackingStatuses = []string{"pending", "notfound", "transit", "pickup",
	"delivered", "undelivered", "exception", "expired"}
var TracktryApiKey string
var ConfigFilePath string
var ErrorMsg = `API Key is missing.
You need to provide an API Key from Tracktry.
Sign up at https://www.tracktry.com/signup-en.html,
and copy the API key under 'Settings'.

The API Key can be provided in three ways:

1) As a flag:
	parcly <noun> <command> --tracktrykey <key>

2) As an environment variable:
	export PARCLY_TRACKTRYKEY=<key>

3) In a config file:
	default config file is $HOME/.parcly.yaml and should contain:
	tracktrykey: <key>

	When you want to use another location, you can
	specify the location with another flag:
	parcly ... --config <filepath>

4) Through the config set command:
	parcly config set --tracktry key
	For more information see
	parcly help config set
			`
