# hmcli

Tool and Nagios/Icinga check plugin for Homematic/Raspberrymatic based on XMLAPI AddOn


[![Go Report Card](https://goreportcard.com/badge/github.com/tommi2day/hmcli)](https://goreportcard.com/report/github.com/tommi2day/hmcli)
![CI](https://github.com/tommi2day/hmcli/actions/workflows/main.yml/badge.svg)
[![codecov](https://codecov.io/gh/Tommi2Day/hmcli/branch/main/graph/badge.svg?token=3EBK75VLC8)](https://codecov.io/gh/Tommi2Day/hmcli)
![GitHub release (latest SemVer)](https://img.shields.io/github/v/release/tommi2day/hmcli)



## Setup
- install XMLAPI-Addon on your Homematic/Raspberrymatic (see https://github.com/homematic-community/XML-API#installation)
- create a token for the XMLAPI-Addon using `tokenregister.cgi` and notice the value (see https://github.com/homematic-community/XML-API#authentication) 
- put the binary in your nagios/icinga plugin directory
- you may create a config file `hmcli.yaml` (usually in the same directory,$HOME/.hmcli or /etc/nagios-plugins/config or provide full path via --config flag) 
with the content  in [hmcli.yaml](hmcli_config.yaml) and modify for your site. You may provide all thes flags also on the command line
- test the plugin with `hmcli version --debug` to see if the correct config file has been taken  and `hmcli device list` to see if your CCU access is working. 
It should return a list of your devices. You may enable --debug to see details

## Configuration
- create the your command and service definitions as usual. For Icinga2 see special hints in [Icinga2.md](Icinga2.md) and Icinga Director configuration basket [Icinga2_basket.json](Icinga2_basket.json)
- to define a service, you have usually provide the ise_id (--id flag) of your devices, datapoints or system variables. You can get them with `hmcli device list` or `hmcli datapoint list` or `hmcli sysvar list`

## hmcli tool reference
```bash
>hmcli --help
Query Tool and Nagios/Icinga check plugin for Homematic/Raspberrymatic with XMLAPI

Usage:
  hmcli [command]

Available Commands:
  completion    Generate the autocompletion script for the specified shell
  datapoint     command related to datapoints
  device        command related to devices
  help          Help about any command
  notifications check hmWarnThreshold notifications
  rssi          list rssi values of deviceList
  sysvar        command related to system variables
  value         command related to device master values
  version       version print version string

Flags:
  -C, --config string    config file name
  -c, --crit string      critical level
  -d, --debug            verbose debug output
  -h, --help             help for hmcli
  -T, --show-threshold   show threshold section
  -t, --token string     Homematic XMLAPI Token
      --unit-test        redirect output for unit tests
  -u, --url string       Homematic URL (default "http://ccu")
  -w, --warn string      warning level


#-------------------
>hmcli datapoint
command related to datapoints

Usage:
  hmcli datapoint [command]

Available Commands:
  check       check a single datapoint
  list        List all data points for a device

#-------------------
hmcli datapoint list --help
List all data points for a device

Usage:
  hmcli datapoint list [flags]

Flags:
  -h, --help           help for list
  -m, --match string   list data points with name matching regexp
#-------------------
>hmcli  datapoint check --help
select a single datapoint by name or id and check its value
You can set -w or -c to set the warning or critical threshold on numeric values

Usage:
  hmcli datapoint check [flags]

Flags:
  -h, --help           help for check
  -i, --id string      datapoint id (one) to query
  -m, --match string   datapoint value must match this regexp
  -n, --name string    datapoint name(one) to query

#-------------------
>hmcli  devices list --help
list all devices or a specify one

Usage:
  hmcli device list [flags]

Flags:
  -a, --address string   device addresses to query, comma separated
  -h, --help             help for list
  -i, --id string        device ids to query, comma separated
      --internal         also list internal channels
  -n, --name string      device name to query, comma separated

#-------------------

>hmcli mastervalue
command related to device master values

Usage:
  hmcli value [command]

Aliases:
  value, values, mastervalue, mastervalues

Available Commands:
  check       check a single device Value
  list        List all  device mastervalues or a specify one

Flags:
  -h, --help   help for value
#-------------------
>hmcli  mastervalue list --help
List for a given device all mastervalues  or a named value. 
Address or id must be given

Usage:
  hmcli value list [flags]

Flags:
  -a, --address string   device addresses to query, comma separated
  -h, --help             help for list
  -i, --id string        device ids to query, comma separated
  -n, --name string      requested mastervalue names, comma separated

#-------------------
>hmcli  mastervalue check --help
this retrives a given device mastervalue and may check if it contains a string (-m)
You can set -w or -c to set the warning or critical threshold on numeric values

Usage:
  hmcli value check [flags]

Flags:
  -a, --address string   device address (only one) to query, alternative to id
  -h, --help             help for check
  -i, --id string        device id (one) to query, alternative to address
  -m, --match string     returned master value must match this regexp
  -n, --name string      requested master value (not device) name ( only one)
#-------------------
>hmcli notifications --help
List all current notifications. 
You can set -w or -c to set the warning or critical threshold on notification count

Usage:
  hmcli notifications [flags]

Flags:
  -h, --help   help for notifications
  -I, --ignore string   regexp to ignore notifications

#-------------------
>hmcli rssi --help
List available rssi values of devices

Usage:
  hmcli rssi [flags]

Flags:
  -h, --help   help for rssi

#-------------------
>hmcli sysvar --help
command related to system variables

Usage:
  hmcli sysvar [command]

Available Commands:
  check       check a single system variable
  list        List all system variables
#-------------------
hmcli.exe sysvar list --help
List all system variables

Usage:
  hmcli sysvar list [flags]

Flags:
  -h, --help   help for list
#-------------------
>hmcli sysvar check --help
select a single system variable by name or id and check its value

Usage:
  hmcli sysvar check [flags]

Flags:
  -h, --help           help for check
  -i, --id string      variable id (one) to query
  -m, --match string   value must match this regexp
  -n, --name string    value name(one) to query

```

## hmcli examples
```bash
# list datapoints of a device (using regexp). omit --match to see all
>hmcli datapoint list --match "000955699D3D84"
Device: Bewegungsmelder Garage Channel: Bewegungsmelder Garage:0 Datapoint: HmIP-RF.000955699D3D84:0.CONFIG_PENDING(4742) = true 
Device: Bewegungsmelder Garage Channel: Bewegungsmelder Garage:0 Datapoint: HmIP-RF.000955699D3D84:0.DUTY_CYCLE(4746) = false 
Device: Bewegungsmelder Garage Channel: Bewegungsmelder Garage:0 Datapoint: HmIP-RF.000955699D3D84:0.ERROR_CODE(4747) = 0
Device: Bewegungsmelder Garage Channel: Bewegungsmelder Garage:0 Datapoint: HmIP-RF.000955699D3D84:0.LOW_BAT(4748) = false
....
Device: Bewegungsmelder Garage Channel: Bewegungsmelder Garage Datapoint: HmIP-RF.000955699D3D84:1.CURRENT_ILLUMINATION(7799) = 0.140000
Device: Bewegungsmelder Garage Channel: Bewegungsmelder Garage Datapoint: HmIP-RF.000955699D3D84:1.CURRENT_ILLUMINATION_STATUS(7800) = 0
Device: Bewegungsmelder Garage Channel: Bewegungsmelder Garage Datapoint: HmIP-RF.000955699D3D84:1.ILLUMINATION_STATUS(7801) = 0

# check a single datapoint (using exact name) and set warning level if value is not "true"
>hmcli datapoint check --name HmIP-RF.000955699D3D84:0.LOW_BAT --match "true"
CRITICAL: HmIP-RF.000955699D3D84:0.LOW_BAT returned value 'false', does not match 'true' 

Datapoint:HmIP-RF.000955699D3D84:0.LOW_BAT ID:4748 Value:=false
 | 'time'=847ms;;;;

# list devices by name (omit --name to see all)
>hmcli device list --name "Bewegungsmelder Garage"
ID:4740, Name: Bewegungsmelder Garage, Address: 000955699D3D84, Type: HmIP-SMO

# list master values of a device by id (id or name required)
>hmcli mastervalues list --id="4740"
Device (ID=4740): Bewegungsmelder Garage, Type: HmIP-SMO
  ARR_TIMEOUT=10
  CYCLIC_BIDI_INFO_MSG_DISCARD_FACTOR=1
  CYCLIC_BIDI_INFO_MSG_DISCARD_VALUE=57
  CYCLIC_INFO_MSG=1
  CYCLIC_INFO_MSG_DIS=1
  CYCLIC_INFO_MSG_DIS_UNCHANGED=20
  CYCLIC_INFO_MSG_OVERDUE_THRESHOLD=2
  DISABLE_MSG_TO_AC=0
  DUTYCYCLE_LIMIT=180
  ENABLE_ROUTING=1
  LOCAL_RESET_DISABLED=0
  LOW_BAT_LIMIT=2.200000

# check a single master value (usinng id) and set warning level if value is  >=9
>xheck_hm.exe mastervalues check --id 4740 --name ARR_TIMEOUT -w 9
ARR_TIMEOUT=10 


Device:4740 ID:Bewegungsmelder Garage ValueName:ARR_TIMEOUT=10
 | 'Bewegungsmelder Garage(4740).ARR_TIMEOUT'=10;9;;; 'time'=156ms;;;;
 
# check notifications and warn if there is one or more pending
>hmcli.exe notifications -w 1
1 notifications pending 

CONFIG_PENDING: HmIP-RF.000955699D3D84:0.CONFIG_PENDING(Bewegungsmelder Garage) since 2024-02-17T18:25:31+01:00

 | 'notifications'=1;1;;; 'time'=2029ms;;;;

# check notifications as before, but ignore sticky and config pending notifications
>hmcli.exe notifications -I 'STICKY|PENDING' -w 1
0 notifications pending | 'notifications'=0;1;;; 'time'=1853ms;;;;

# list rssi values of devices
>hmcli.exe rssi
Address:BidCoS-RF rx:65536 tx: -56
Address:KEQ1070683 rx:65536 tx: -73
Address:MEQ0272433 rx:65536 tx: -58

# list system variables and find id of a variable
>hmcli.exe sysvar list
ID:7931, Name: DutyCycle-Alarm, Value: nicht ausgelöst, INFO: DutyCycle 98% (CCU)
ID:7794, Name: WatchDog-Alarm, Value: nicht ausgelöst, INFO: Unclean shutdown or system crash identified
ID:1233, Name: Alarmzone 1, Value: ausgelöst, INFO: Alarmmeldung Alarmzone 1, Since:2024-02-17 19:15:22
ID:950, Name: Anwesenheit, Value: anwesend, INFO: Anwesenheit, Since:2024-01-31 21:33:04
ID:6548, Name: DutyCycle, Value: 36.000000 %, INFO: DutyCycle CCU, Since:2024-02-17 21:54:00

# check HM DutyCyle variable and set warning level if value is >=30
# DutyCycle needs to be published by removing the "internal" flag in CCU Systemvariables settings
>hmcli.exe sysvar check --id 6548 -w 30
DutyCycle=34.000000 (%) 

ID:6548, Name: DutyCycle, Value: 34.000000 %, INFO: DutyCycle CCU, Since:2024-02-17 21:58:00
 | 'DutyCycle(6548).'=34.000000;30;;; 'time'=229ms;;;;

```


## see also
- [RaspberryMatic](https://github.com/jens-maus/RaspberryMatic)
- [XMLAPI-Addon](https://github.com/homematic-community/XML-API)
- [gomodules/hmlib](https://pkg.go.dev/github.com/tommi2day/gomodules/hmlib)
- [go-nagios](https://github.com/atc0005/go-nagios)
- [Threshold Format](https://nagios-plugins.org/doc/guidelines.html#THRESHOLDFORMAT)
- [Icinga2](https://icinga.com/docs/icinga-2/latest/doc/03-monitoring-basics/)
- [Icinga2 Director](https://github.com/Icinga/icingaweb2-module-director)
- [Icinga2 configuration](Icinga2.md)
- [Icinga2 Director Basket](Icinga2_basket.json)










