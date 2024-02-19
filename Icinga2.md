# Icinga2 configuration

These are simple examples how to use the check_hm plugin with Icinga2.

## Definitions

You need to define
- data fields
- the command definition
- the service template
- the service definitions

for your convenience, a icinga director basket is available [here](Icinga2_basket.json)

### commands
```conf
object CheckCommand "hm" {
    import "plugin-check-command"
    command = [ PluginDir + "/check_hm" ]
    arguments += {
        "(no key)" = {
            description = "check_hm command"
            order = 1
            required = true
            skip_key = true
            value = "$hm_command$"
        }
        "(no key.2)" = {
            description = "check_hm Subcommand"
            order = 2
            skip_key = true
            value = "$hm_subcommand$"
        }
        "(no key.3)" = {
            description = "Extra params on command line"
            order = -1
            skip_key = true
            value = "$hm_extra_params$"
        }
        "-C" = {
            description = "full name of configuration file"
            value = "$hm_config$"
        }
        "-I" = {
            description = "regexp to ignore notifications"
            value = "$hm_ignore$"
        }
        "-a" = {
            description = "HM Address to lookup"
            value = "$hm_address$"
        }
        "-c" = {
            description = "Crtical value string (nagios syntax)"
            value = "$hm_critical$"
        }
        "-i" = {
            description = "XMLAPI ise_id of device/datapoint/value"
            value = "$hm_id$"
        }
        "-n" = {
            description = "HM Name (target depends on command)"
            value = "$hm_name$"
        }
        "-t" = {
            description = "HM Token"
            value = "$hm_token$"
        }
        "-u" = {
            description = "HM Base URL (e.g. http://ccu/)"
            value = "$hm_url$"
        }
        "-w" = {
            description = "Warning Level (nagios syntax)"
            value = "$hm_warning$"
        }
    }
}  
```

### Service template

```conf
template Service "Homematic" {
    import "service-default"

    check_command = "hm"
    check_interval = 15m
    enable_active_checks = true
    enable_perfdata = true
    command_endpoint = null
    vars.hm_token = "<token>"
    vars.hm_url = "http://ccu"
}
```
 
check CCU Duty Cycle

```conf
object Service "Homematic DutyCycle" {
    host_name = "ccu"
    import "Homematic"

    vars.hm_command = "sysvar"
    vars.hm_critical = "80"
    vars.hm_name = "DutyCycle"
    vars.hm_subcommand = "check"
    vars.hm_warning = "50"
}
```

Check a outdoor datapoint and set warning level if value is below 0

```conf
object Service "Outdoor Temperature" {
    host_name = "ccu"
    import "Homematic"

    vars.hm_command = "datapoint"
    vars.hm_id = "6471"
    vars.hm_subcommand = "check"
    vars.hm_warning = "0:"
}
```

Check notifications and ignore sticky and config pending notifications

```conf
object Service "Homematic Notifications" {
    host_name = "ccu"
    import "Homematic"

    vars.hm_command = "notifications"
    vars.hm_ignore = "STICKY_UNREACH|CONFIG_PENDING"
    vars.hm_warning = "1"
}
```
