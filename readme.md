# xianrail_exporter

Xianrail open data exporter for Prometheus. Currently, collects data from
the [拥挤度](https://www.xianrail.com/static/xianApp/index.html) website.

## Roadmap & Status

- [ ] Xianrail web API access
    - [x] passenger alarm api
    - [ ] thermogram api
- [ ] Prometheus collector
- [ ] Data processing (consider moving to <https://github.com/mitchellh/mapstructure>)
    - [ ] station data
        - [ ] dynamically get station data from xianrail website source code
    - [ ] passenger alram data
        - [x] total passengers
        - [ ] station passenger in
        - [ ] station passenger turn
    - [ ] thermogram data
- [ ] Data persistence
- [ ] Logger (consider using <https://github.com/sirupsen/logrus>)
- [ ] Test (consider using <https://github.com/onsi/ginkgo>)