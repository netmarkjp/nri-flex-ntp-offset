# Build

```sh
go build -buildvcs=false -trimpath -ldflags "-w -s -extldflags '-static'" -tags "netgo,osusergo" -o nri-flex-ntp-offset
```

# New Relic Infrastructure Integration Setting

```yaml
# /etc/newrelic-infra/integrations.d/ntp-offset.yml
integrations:
  - name: nri-flex
    config:
      name: ntpOffset
      apis:
        - name: ntpOffset
          commands:
            - run: /opt/nri-flex-ntp-offset
```

# Graph/Dashboard NRQL Example

```sql
SELECT latest(`clock.offset.ms`) FROM ntpOffsetSample FACET fullHostname TIMESERIES
```

# Alert Example

```sql
SELECT abs(latest(`clock.offset.ms`)) FROM ntpOffsetSample
```
