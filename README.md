# tsc-spotpricecharger

Docker container to be added to an existing [Teslamate](https://docs.teslamate.org/) and
[TeslaSolarCharger](https://github.com/pkuehnel/TeslaSolarCharger). It can initiate car charging (via
TeslaSolarCharger) based on spot price (below a certain limit) up to defined state of charge (SoC) limit.

To use it just add it to the docker-compose.yml of your Teslamate installation. For example:

```
  tsc-spotpricecharger:
    image: andreasfaerber/tsc-spotpricecharger
    container_name: tsc-spotpricecharger
    restart: always
    depends_on:
      - teslasolarcharger
    environment:
      - TZ=Europe/Berlin
      - TSC_SPOT_TESLAMATEAPI_URL=http://teslamateapi:8080
      - TSC_SPOT_TSC_URL=http://teslasolarcharger:7190
      - TSC_SPOT_SPOTCHARGEPRICE=0.05
      - TSC_SPOT_CHECKINTERVAL=300
      - TSC_SPOT_STARTUPDELAY=180
      - TSC_SPOT_CARID=1
      - TSC_SPOT_CHARGESOCLIMIT=80
      - TSC_SPOT_FALLBACKCHARGELIMIT=80
      - TSC_SPOT_DEBUG=false
```

| Variable | Default               | Description           |
|----------|-----------------------|-----------------------|
| TSC_SPOT_TSC_URL | http://teslasolarcharger:7190 | TeslaSolarCharger URL |
| TSC_SPOT_TESLAMATEAPI_URL | http://teslamateapi:8080      | Teslamate API URL |     
| TSC_SPOT_SPOTCHARGEPRICE | 0.06                  | Spot price under which charging should charge. Raw spot price without any additional price components |
| TSC_SPOT_CHECKINTERVAL | 300                   | Interval in seconds to check for spot price and if charging should start / stop. |
| TSC_SPOT_CARID | 1                     | Car ID in TeslaSolarCharger to initiate charging for when spot price falls below TSC_SPOT_SPOTCHARGEPRICE |
| TSC_SPOT_CHARGESOCLIMIT | 80                    | Upper SoC limit to which the car should be charged which spot price is below TSC_SPOT_SPOTCHARGEPRICE | 
| TSC_SPOT_FALLBACKSOCLIMIT | 80                    | Fallback SoC limit when the initial SoC limit in TeslaSolarCharger has not been saved (usually across restarts) |
| TSC_SPOT_STARTUPDELAY | 180                   | Delay after startup to allow TeslaSolarCharger to properly start up | 
| TSC_SPOT_DEBUG | false | Log debug information |
| TSC_SPOT_DRYRUN | false | Do not issue API calls to TeslaSolarCharger |
| TZ | Europe/Berlin | Local timezone |

