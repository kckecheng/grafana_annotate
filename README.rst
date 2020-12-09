Grafana Annotate
=================

About
------

I'm using Grafana for end to end monitoring during product qualifications. Each qualification task consists of several use cases which need to be aligned with the monitoring results. Simply speaking, I need to mark the start and end timestamps of each use case on Grafana panels. It is not convenient to do the work manually, hence this CLI tool is created.

Usage
------

::

  git clone https://github.com/kckecheng/grafana_annotate.git
  cd grafana_annotate
  go build
  ./grafana_annotate --help

Issues
-------

- Not compatible with Grafana 7.0+ due to SDK issues.
