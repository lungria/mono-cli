# mono-cli
CLI tool that exports mono statements in csv format and prints it to stdout.
Works only with personal API key.

### Limitations
Due to monobank API limitations tool can only load statements for 31 days period in one request and only send one request per 60 seconds. So if you launched a tool and it's still running and nothing printed - just wait. 

### Configuration
Tool can be configured using following environment variables
```
MONO_APIKEY - required. Your personal API key, see https://api.monobank.ua/

MONO_STARTDATE - this tool will import statements occured after this date. Expected format: unix timestamp in UTC. If not set default value is monobank launch date (15 Nov 2017)

MONO_ENDDATE - this tool will import statements occured before this date. Expected format: unix timestamp in UTC. If not set default value is current date time
```
