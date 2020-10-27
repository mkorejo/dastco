## DAST CLI
### BurpSuite Enterprise
```

```

### WebInspect
```
docker run mkorejo/dastco wi
time="2020-10-27T17:36:06Z" level=warning msg="No configuration file specified."
Usage:
  dastco wi [flags]
  dastco wi [command]

Available Commands:
  list          List scans
  resume        Resume a stopped scan
  retest        Start a scan retest
  retest-status Get the status of a scan retest
  status        Get the status of a scan
  stop          Stop a running scan

Flags:
  -h, --help              help for wi
  -p, --password string   Password for WebInspect username
      --scan-id string    Scan ID
  -U, --url string        WebInspect API URL
  -u, --username string   WebInspect username

Global Flags:
  -c, --config string   Configuration file (defaults to ~/.dastco)
  -k, --insecure        Trust all SSL certificates

Use "dastco wi [command] --help" for more information about a command.
```
