netapp_exporter
---

This exporter will expose metrics about a netapp ontap NAS storage system, which can be scraped by prometheus. it utilized [go-netapp](https://github.com/pepabo/go-netapp) as underlying lib.



example configure set as [example](./scripts/netapp_exporter.yml)
```yaml
credentials:
    10.36.48.39:
      group: BSU
      username: admin
      password: pass
      debug: false
```

here `group` which is used to confrom to the `netapp-harvest` group filter.



then start netapp_exporter via 
```sh
netapp_exporter --config.file=netapp_exporter.yml
```

then we can get the metrics via 
```
curl http://<netapp-export host>:9609/netapp?target=10.36.48.39

```

## prometheus job conf
add netapp-exporter job conif as following
```yaml
  - job_name: 'netapp-exporter'

    # metrics_path defaults to '/metrics'
    metrics_path: /netapp


    # scheme defaults to 'http'.

    static_configs:
    - targets:
       - 10.36.48.39
    relabel_configs:
      - source_labels: [__address__]
        target_label: __param_target
      - source_labels: [__param_target]
        target_label: instance
      - target_label: __address__
        replacement: localhost:9609  ### the address of the netapp-exporter address
````