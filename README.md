## Build

```
go get github.com/randomtask1155/logpas
```

## Usage

```
Usage of logpas:
  -f string
    	What log component are you reading.  Defaults assumes log formatter is '{}'.
Available components include: (default "default")
  -l string
    	Specify path to logfile to parse.  If not defined the default is stdin
```

## formats

| Format  | Description  |   
|---|---|
| Lager  | [lager](https://github.com/cloudfoundry/lager) converts the epoch timestamp into human readable form  | 



## Example STDIN

original line

```
{"timestamp":"1513880255.460468292","source":"tps-listener","message":"tps-listener.lrp-stats.fetching-container-metrics","log_level":1,"data":{"log-guid":"12345","process-guid":"12345","session":"1059"}}
```

parsed

```
:> cat samples/tps-lister.txt | logpas -f lager
2017-12-21 12:17:35.460468292 -0600 CST:tps-listener:tps-listener.lrp-stats.fetching-container-metrics:1: {"log-guid":"12345","process-guid":"12345","session":"1059"}
```
