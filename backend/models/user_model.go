package models

type Rack struct {
    height int
    unit_number int
    width int
    specification int
    depth int
}

type Server struct {
    server_name string
    ip_address string
    os string
    cpu string
    memory int
    storage int
    nic_mac_address string
}

type ServerPlacement struct {
    nic_mac_address string
    rack_unit_number int
}

type NetworkDevice struct {
    device_name string
    port_count int
    ip_address []string
    role string
    port_mac_address []string
}

type DCUser struct {
    uid string
    name string
}

type Table struct {
    uid string
    server_or_network_device_name string
}

type Connection struct {
    rack_number int
    note string
    server_or_network_device_name string
}
