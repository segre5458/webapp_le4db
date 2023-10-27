package models

type Rack struct {
    Height int
    UnitNumber int
    Width int
    Specification int
    Depth int
}

type Server struct {
    ServerName string
    IpAddress string
    Os string
    Cpu string
    Memory int
    Storage int
    NicMacAddress string
}

type ServerPlacement struct {
    NicMacAddress string
    RackUnitNumber int
}

type NetworkDevice struct {
    DeviceName string
    PortCount int
    IpAddress []string
    Role string
    PortMacAddress []string
}

type NetworkDevicePlacement struct {
    PortMacAddress string
    RackUnitNumber int
}

type DCUser struct {
    Uid string
	Name string
	Password string
	Email string
}

type Table struct {
    Uid string
    ServerOrNetworkDeviceName string
}

type Connection struct {
    RackNumber int
    Note string
    ServerOrNetworkDeviceName string
}
