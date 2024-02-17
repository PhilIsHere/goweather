# goweather

[![PhilIsHere - goweather](https://img.shields.io/static/v1?label=PhilIsHere&message=goweather&color=blue&logo=github)](https://github.com/PhilIsHere/goweather "Go to GitHub repo")
[![stars - goweather](https://img.shields.io/github/stars/PhilIsHere/goweather?style=social)](https://github.com/PhilIsHere/goweather)
[![forks - goweather](https://img.shields.io/github/forks/PhilIsHere/goweather?style=social)](https://github.com/PhilIsHere/goweather)

[![Buildtest](https://github.com/PhilIsHere/goweather/workflows/Buildtest/badge.svg)](https://github.com/PhilIsHere/goweather/actions?query=workflow:"Buildtest")
[![GitHub release](https://img.shields.io/github/release/PhilIsHere/goweather?include_prereleases=&sort=semver&color=blue)](https://github.com/PhilIsHere/goweather/releases/)
[![License](https://img.shields.io/badge/License-GPL--3.0-blue)](https://github.com/PhilIsHere/goweather/blob/main/LICENSE)
[![issues - goweather](https://img.shields.io/github/issues/PhilIsHere/goweather)](https://github.com/PhilIsHere/goweather/issues)

goweather aims to provide a simple command line interface to get the current weather and forecast for a given location in Germany. The data is provided by the [Deutscher Wetterdienst](https://www.dwd.de/) via the [BrightSky API](https://brightsky.dev/).

The GPS coordinates are provided by the [OpenStreetMap](https://www.openstreetmap.org/) API.

The DWD ([Deutscher Wetterdienst](https://www.dwd.de/)), as Germany's
meteorological service, publishes a myriad of meteorological observations and
calculations as part of their [Open Data
program](https://www.dwd.de/DE/leistungen/opendata/opendata.html).

[**Bright Sky**](https://brightsky.dev/) is an open-source project aiming to
make some of the more popular data — in particular weather observations from
the DWD station network and weather forecasts from the MOSMIX model — available
in a free, simple JSON API.

## Usage
Download the binary from the release page and run it with the following command:
./goweather on Unix or .\goweather.exe on Windows

You can also just double click the binary.

## My Antivirus Program detects a Trojan!
This seems to be a common false positive especially on windows machines.
Don't take my word for it. Here's the official notice from the golang team: https://go.dev/doc/faq#virus

```
Why does my virus-scanning software think my Go
distribution or compiled binary is infected?

This is a common occurrence, especially on Windows machines,
and is almost always a false positive.
Commercial virus scanning programs are often confused by
the structure of Go binaries, which they don't see as often
as those compiled from other languages.

(...)

In any case, if you believe the report is in error, please report
a bug to the supplier of your virus scanner.
Maybe in time virus scanners can learn to understand Go programs."

```
Please rest assured, that the tool does not contain any malicious components.

## Aknowledgements and Privacy Policies
The privacy policy of OpenStreetMap can be found here: https://osmfoundation.org/wiki/Privacy_Policy

Goweather uses the raw data from the [Deutscher
Wetterdienst](https://www.dwd.de/).
The Data is provided via the [BrightSky API](https://brightsky.dev/).

<a href="https://www.dwd.de/"><img src="dwd.svg" alt="Deutscher Wetterdienst" height="100"></a>
