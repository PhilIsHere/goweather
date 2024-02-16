# goweather

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

Please note that you need to cd into the directory where the binary is located.

## My Antivirus Program detects a Trojan!
This seems to be a common false positive especially with Windows Machines.
Don't take my word for it. Here's the official Notice from the Golang team: https://go.dev/doc/faq#virus

```
Why does my virus-scanning software think my Go distribution or compiled binary is infected?

This is a common occurrence, especially on Windows machines, and is almost always a false positive.
Commercial virus scanning programs are often confused by the structure of Go binaries,
which they don't see as often as those compiled from other languages.

If you've just installed the Go distribution and the system reports it is infected,
that's certainly a mistake. To be really thorough, you can verify the
download by comparing the checksum with those on the downloads page.

In any case, if you believe the report is in error, please report
a bug to the supplier of your virus scanner.
Maybe in time virus scanners can learn to understand Go programs."

```
https://go.dev/doc/faq#virus

## Aknowledgements and Privacy Policies
The privacy policy of OpenStreetMap can be found here: https://osmfoundation.org/wiki/Privacy_Policy

Goweather uses the raw data from the [Deutscher
Wetterdienst](https://www.dwd.de/).
The Data is provided via the [BrightSky API](https://brightsky.dev/).

<a href="https://www.dwd.de/"><img src="docs/img/dwd.svg" alt="Deutscher Wetterdienst" height="100"></a>

