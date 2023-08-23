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

# Development

```sh
[[ -f "$(go env GOPATH)/bin/task" ]] || go install github.com/go-task/task/v3/cmd/task@latest
$(go env GOPATH)/bin/task lint
```