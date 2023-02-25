package driver

type Driver struct {
	config Config
}

type Config struct {
	DriverName    string
	VendorVersion string
}

func NewHostPathDriver(cfg Config) (*Driver, error) {
	n := &Driver{
		config: cfg,
	}
	return n, nil
}
