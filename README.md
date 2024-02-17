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
<a href="https://www.dwd.de/"><img src="https://www.dwd.de/SharedDocs/bilder/DE/logos/dwd/dwd_logo_258x69.png" alt="Deutscher Wetterdienst"></a>
The Data is provided via the [BrightSky API](https://brightsky.dev/).

