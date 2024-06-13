# Unitary tests

## Mode-S

| Downlink Format | Description                              | UT |
|-----------------|------------------------------------------|----|
| DF-0            | Short air-air surveillance (ACAS)        | ✓  |
| DF-4            | Surveillance, altitude reply             | ✓  |
| DF-5            | Surveillance, identify reply             | ✓  |
| DF-11           | All-call reply                           | ✓  |
| DF-16           | Long air-air surveillance (ACAS)         | ✓  |
| DF-17           | Extended squitter (ADSB)                 | ✓  |
| DF-18           | Extended squitter/non transponder (ADSB) | ✓  |
| DF-19           | Military extended squitter               | ✓  |
| DF-20           | Comm-B altitude reply                    | ✓  |
| DF-21           | Comm-B identify reply                    | ✓  |
| DF-24           | Comm-D (ELM)                             | ✓  |

## BDS

| BDS | Description                         | UT |
|-----|-------------------------------------|----|
| 0,0 | No message available                | ✓  |
| 0,5 | (ADSB) Airborne position            |    | 
| 0,6 | (ADSB) Surface position             |    |
| 0,7 | Status                              |    |
| 0,8 | (ADSB) Aircraft identification      | ✓  |
| 0,9 | (ADSB) Airborne velocity            | ✓  |
| 1,0 | Data link capability report         |    |
| 1,7 | Common usage GICB capability report |    |
| 2,0 | Aircraft identification             | ✓  |
| 3,0 | ACAS resolution advisory            |    |
| 4,0 | Selected vertical intention         | ✓  |
| 4,4 | Meteorological routine air report   | ✓  |
| 4,5 | Meteorological hazard report        | ✓  |
| 5,0 | Track and turn report               | ✓  |
| 6,0 | Heading and speed report            | ✓  |
| 6,1 | (ADSB) Emergency report             | ✓  |
| 6,2 | (ADSB) Target and status            | ✓  |
| 6,5 | (ADSB) Operational status           | ✓  |

## ADSB

| BDS | Description             | Formats        | UT |
|-----|-------------------------|----------------|----|
|     | No position information | 0              | ✓  |  
| 0,5 | Airborne position       | 9-18 and 20-22 |    |  
| 0,6 | Surface position        | 5-8            |    | 
| 0,8 | Aircraft identification | 1-4            |    | 
| 0,9 | Airborne velocity       | 19             |    | 
| 6,1 | Emergency report        | 28             |    | 
| 6,2 | Target and status       | 29             |    |  
| 6,5 | Operational status      | 31             |    | 

## Comm-B

TBD

## ACAS

| Code | Description         | UT |
|------|---------------------|----|
| 3,0  | Resolution Advisory | ✓  |
