# Introduction
ModeS is a Mode-S and ADSB decoder written in Go. The project is mainly a library that can be quickly plugged to 
deformat the messages.

An example application is provided that can can connect either to an AirSpy device (via ADSB-Spy) or can 
read from a text file.

The main goal of this project is exactitude, code clarity and completeness (if achievable).

# Warning
This project is experimental and has bugs and issues. It should by no means be used in a context where one will have
to rely (in the broadest meaning) on the provided information.

# Format Supported
## Mode-S
| Downlink Format | Description |
| --- | --- | 
| DF-0 | Short air-air surveillance (ACAS) | 
| DF-4 | Surveillance, altitude reply | 
| DF-5 | Surveillance, identify reply |
| DF-11 | All-call reply |
| DF-16 | Long air-air surveillance (ACAS) |
| DF-17 | Extended squitter (ADSB) |
| DF-18 | Extended squitter/non transponder (ADSB) |
| DF-19 | Military extended squitter |
| DF-20 | Comm-B altitude reply |
| DF-21 | Comm-B identify reply |
| DF-24 | Comm-D (ELM) |

## ADSB
| Code |  BDS | Description  |         V0   |   V1   |   V2 |
| --- | --- | --- | --- | --- | --- | 
|  0   |   ?  |  No Position
|  1   |  0,8 |  Aircraft Id         | OK | OK | OK |
|  2   |  0,8 |  Aircraft Id         | OK | OK | OK |
|  3   |  0,8 |  Aircraft Id         | OK | OK | OK |
|  4   |  0,8 |  Aircraft Id         | OK | OK | OK |
|  5   |  0,6 |  Surface position    | OK | OK | OK |
|  6   |  0,6 |  Surface position    | OK | OK | OK |
|  7   |  0,6 |  Surface position    | OK | OK | OK |
|  8   |  0,6 |  Surface position    | OK | OK | OK |
|  9   |  0,5 |  Airborne position   | OK | OK | OK |
| 10   |  0,5 |  Airborne position   | OK | OK | OK |
| 14   |  0,5 |  Airborne position   | OK | OK | OK |
| 12   |  0,5 |  Airborne position   | OK | OK | OK |
| 13   |  0,5 |  Airborne position   | OK | OK | OK |
| 14   |  0,5 |  Airborne position   | OK | OK | OK |
| 15   |  0,5 |  Airborne position   | OK | OK | OK |
| 16   |  0,5 |  Airborne position   | OK | OK | OK |
| 17   |  0,5 |  Airborne position   | OK | OK | OK |
| 18   |  0,5 |  Airborne position   | OK | OK | OK |
| 19   |  0,9 |  Airborne velocity   | OK | OK | OK |
| 20   |  0,5 |  Airborne position   | OK | OK | OK |
| 21   |  0,5 |  Airborne position   | OK | OK | OK |
| 22   |  0,5 |  Airborne position   | OK | OK | OK |
| 23   |      | Reserved             |    |    |    |
| 24   |      | Reserved             |    |    |    |
| 25   |      |  Reserved            |    |    |    |
| 26   |      |  Reserved            |    |    |    |
| 27   |      |  Reserved            |    |    |    |
| 28   |  6,1 |  Emergency report    | OK | OK | OK |
| 29   |  6,2 |  Target and status   | _  | OK | OK |
| 30   |      |  Reserved            |    |    |    |
| 31   |  6,5 |  Operational status  | OK | OK | OK |

## ACAS
| Code |  Description |
| --- | --- |
| 3,0  | Resolution Advisory | 

# Usage
## Library
A full example using the library is given in the file _message_processor.go_ which goes from a single "line" of data 
received to outputting the detailed content.

A simple workflow can be:

```go
// Split the line into its various fields
messageADSBSpy, _ := adsbspy.ReadLine(str)
// Convert to a Mode-S message if possible
messageModeS, err := modeSReader.ReadMessage(messageADSBSpy.Message)
// If the message is an ADSB message
if messageModeS.GetDownLinkFormat() == 17 {
    messageDF17 := messageModeS.(*modeSMessages.MessageDF17)
    messageADSB, _, _ := adsbReader.ReadADSBMessage(adsb.Level2, false, false, messageDF17.MessageExtendedSquitter.Data)
}
```

##Application
For reading a file
```bash
go run cmd/main.go -file .\example\example2.txt
```

For connecting to an ADSBSpy server
```bash
go run cmd/main.go -adsb_spy_server localhost -adsb_spy_port 47806
```

# Sources
The main sources used are:
 * ICAO - Annex 10 - Volume IV (July 2014): Surveillance and Collision Avoidance System
 * ICAO Doc 9871 AN/460: Technical Provisions for Mode S Services and Extended Squitter
 
I am always open to add new valid sources of information. In particular, information about ACAS messages are welcome.

# Versions

 * v0.2.0: 
    * First public version
 * v0.1.0: First version

# License

Copyright 2019 Thomas Wuillemin  <thomas.wuillemin@gmail.com>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this project or its content except in compliance with the License.
You may obtain a copy of the License at

http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.