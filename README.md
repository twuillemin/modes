# Introduction
ModeS is a Mode-S and ADSB decoder written in Go. The project is a pure Go library, without dependency, that allows to
quickly deformat and access the content of Mode-S message.

A small example application is provided that can be used to dump the content of a message. A more complete usage
of the library is done in the sister project ModeS Viewer (https://github.com/twuillemin/modes-viewer). 

The main goal of this project is exactitude, code clarity and completeness (if achievable).

# Warning
This project is experimental and has bugs and issues. It should by no means be used in a context where one will have
to rely (in the broadest meaning) on the provided information.

# Format Supported
## Mode-S
| Downlink Format | Description                              |
|-----------------|------------------------------------------| 
| DF-0            | Short air-air surveillance (ACAS)        | 
| DF-4            | Surveillance, altitude reply             | 
| DF-5            | Surveillance, identify reply             |
| DF-11           | All-call reply                           |
| DF-16           | Long air-air surveillance (ACAS)         |
| DF-17           | Extended squitter (ADSB)                 |
| DF-18           | Extended squitter/non transponder (ADSB) |
| DF-19           | Military extended squitter               |
| DF-20           | Comm-B altitude reply                    |
| DF-21           | Comm-B identify reply                    |
| DF-24           | Comm-D (ELM)                             |

## ADSB
| Code | BDS | Description        | V0 | V1 | V2 |
|------|-----|--------------------|----|----|----| 
| 0    | ?   | No Position        |    |    |    |
| 1    | 0,8 | Aircraft Id        | OK | OK | OK |
| 2    | 0,8 | Aircraft Id        | OK | OK | OK |
| 3    | 0,8 | Aircraft Id        | OK | OK | OK |
| 4    | 0,8 | Aircraft Id        | OK | OK | OK |
| 5    | 0,6 | Surface position   | OK | OK | OK |
| 6    | 0,6 | Surface position   | OK | OK | OK |
| 7    | 0,6 | Surface position   | OK | OK | OK |
| 8    | 0,6 | Surface position   | OK | OK | OK |
| 9    | 0,5 | Airborne position  | OK | OK | OK |
| 10   | 0,5 | Airborne position  | OK | OK | OK |
| 14   | 0,5 | Airborne position  | OK | OK | OK |
| 12   | 0,5 | Airborne position  | OK | OK | OK |
| 13   | 0,5 | Airborne position  | OK | OK | OK |
| 14   | 0,5 | Airborne position  | OK | OK | OK |
| 15   | 0,5 | Airborne position  | OK | OK | OK |
| 16   | 0,5 | Airborne position  | OK | OK | OK |
| 17   | 0,5 | Airborne position  | OK | OK | OK |
| 18   | 0,5 | Airborne position  | OK | OK | OK |
| 19   | 0,9 | Airborne velocity  | OK | OK | OK |
| 20   | 0,5 | Airborne position  | OK | OK | OK |
| 21   | 0,5 | Airborne position  | OK | OK | OK |
| 22   | 0,5 | Airborne position  | OK | OK | OK |
| 23   |     | Reserved           |    |    |    |
| 24   |     | Reserved           |    |    |    |
| 25   |     | Reserved           |    |    |    |
| 26   |     | Reserved           |    |    |    |
| 27   |     | Reserved           |    |    |    |
| 28   | 6,1 | Emergency report   | OK | OK | OK |
| 29   | 6,2 | Target and status  | _  | OK | OK |
| 30   |     | Reserved           |    |    |    |
| 31   | 6,5 | Operational status | OK | OK | OK |

## ACAS
| Code | Description         |
|------|---------------------|
| 3,0  | Resolution Advisory | 

# Usage
## Library
A full example using the library is given in the file _main.go_ which goes from a single "line" of data 
received to outputting the detailed content.

A simple workflow can be:

```go
package main

import (
    "encoding/hex"
    "fmt"
    "github.com/twuillemin/modes/pkg/bds/adsb"
    adsbReader "github.com/twuillemin/modes/pkg/bds/reader"
    modeSMessages "github.com/twuillemin/modes/pkg/modes/messages"
    modeSReader "github.com/twuillemin/modes/pkg/modes/reader"
)

func main() {
	// Convert the string to its hexadecimal value
	binaryData, _ := hex.DecodeString("8D40768DEA3AB864013C088209CA")
	// Read to a Mode-S message if possible
	messageModeS, _ := modeSReader.ReadMessage(binaryData)
	// If the message has ADSB data
	if messageModeS.GetDownLinkFormat() == 17 {
		// Convert the message to its real type
		messageDF17 := messageModeS.(*modeSMessages.MessageDF17)
		// Read the ADSB content
		messageADSB, _, _ := adsbReader.ReadADSBMessage(
			adsb.ReaderLevel2,
			false,
			false,
			messageDF17.MessageExtendedSquitter)
		fmt.Print(messageADSB.ToString())
	}
}
```

which should hopefully print a report like:
```
Message:                                       29 - Target state and status information (BDS 6,2) [ADSB V2]
Subtype:                                       1 - Subtype 1
Selected Altitude Type:                        0 - MCP/FCU (Mode Control Panel / Flight Control Unit)
Selected Altitude :                            30016 feet
Barometric Pressure Setting (minus 800 mbars): 213.60000000000002 feet
Selected Heading :                             0 degrees
Navigation Accuracy Category - Position:       9 - EPU < 30 m - e.g. GPS (SA off)
Navigation Integrity Category - Baro:          1 - barometric altitude based on a Gilham input that has been cross-checked or on a non Gilham input
Source Integrity ReaderLevel:                  1 - SIL <= 1 * 10^-7 per flight hour or per sample
Source Integrity ReaderLevel Supplement:       0 - Probability of exceeding NIC radius of containment is based on "per hour"
Status of MCP/FPU Mode Bits:                   0 - no information provided for MCP/FCU mode bits
Autopilot Engaged:                             No information provided from MCP/FCU
VNAV Mode Engaged:                             No information provided from MCP/FCU
Altitude Hold Mode Engaged:                    No information provided from MCP/FCU
Approach Mode Engaged:                         No information provided from MCP/FCU
LNAV Mode Engaged:                             No information provided from MCP/FCU
TCAS / ACAS Operational :                      1 - TCAS/ACAS is operational
```

All the values printed are in fact extracted in their own structure, using either numerical value or enumeration when
applicable. The _ToString()_ function is present on most if not all the attributes and just builds a readable 
representation of the data otherwise accessible by the structure of the object.

## Application
For extracting a single line content
```bash
go run cmd/main.go 8D40768DEA3AB864013C088209CA
```

For reading a file with multiple lines
```bash
go run cmd/main.go -file .\example\example2.txt
```

# Sources
The main sources used are:
 * ICAO, Annex 10, Volume IV (July 2014): Surveillance and Collision Avoidance System
 * ICAO, Doc 9871, AN/460: Technical Provisions for Mode S Services and Extended Squitter
 
I am always open to add new valid sources of information. In particular, information about ACAS messages are welcome.

# Versions
 * v0.4.0: 
    * Mode ADSBSpy and "current state" of plane into a separate project
 * v0.3.0:
    * Clean relation Format, SubType Version for ADSB messages
    * Bug fixes on some fields not well fetched from messages
    * Add unitary test on all main classes
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