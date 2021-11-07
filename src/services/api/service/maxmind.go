package service

import (
	"net"

	"github.com/oschwald/geoip2-golang"
)

type Maxmind struct {
	reader *geoip2.Reader
}

type MaxmindOptions struct {
	Path  string
	Debug bool
}

func newMaxmind(options *MaxmindOptions) (*Maxmind, error) {
	var path string
	if options.Debug {
		path = "./share/GeoLite2-City.mmdb"
	} else {
		path = options.Path
	}

	reader, err := geoip2.Open(path)
	if err != nil {
		return nil, err
	}

	mx := &Maxmind{reader: reader}
	return mx, nil
}

func (m *Maxmind) Lookup(ip string) (country, region string, err error) {
	record, err := m.reader.City(net.ParseIP(ip))
	if err != nil {
		return "", "", err
	}

	country = record.Country.IsoCode
	if len(record.Subdivisions) > 0 {
		region = record.Subdivisions[len(record.Subdivisions)-1].IsoCode
	}
	return country, region, nil
}

func (m *Maxmind) Drop() error {
	return m.reader.Close()
}

func (m *Maxmind) DropMsg() string {
	return "close maxmind database"
}
