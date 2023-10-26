CREATE TABLE rack (
    height integer NOT NULL,
    unit_number integer NOT NULL CHECK (unit_number >= 1 AND unit_number <= 100),
    width integer NOT NULL,
    specification integer CHECK (specification IN (19, 21, NULL)),
    depth integer NOT NULL
);

CREATE TABLE server (
    server_name character varying NOT NULL,
    ip_address character varying NOT NULL,
    os character varying NOT NULL,
    cpu character varying NOT NULL,
    memory integer NOT NULL,
    storage integer NOT NULL,
    nic_mac_address character varying NOT NULL
);

CREATE TABLE server_placement (
    nic_mac_address character varying NOT NULL,
    rack_unit_number integer NOT NULL
);

ALTER TABLE server_placement
    ADD CONSTRAINT placement_unique_unit_number UNIQUE (rack_unit_number);

CREATE TABLE network_device (
    device_name character varying NOT NULL,
    port_count integer NOT NULL,
    ip_addresses inet[] NOT NULL,
    role character varying NOT NULL,
    port_mac_addresses character varying[] NOT NULL
);

CREATE TABLE network_device_placement (
    port_mac_address character varying NOT NULL,
    rack_unit_number integer NOT NULL
);

ALTER TABLE network_device_placement
    ADD CONSTRAINT placement_unique_unit_number UNIQUE (rack_unit_number);

CREATE TABLE DCuser (
    uid character varying NOT NULL,
    name character varying NOT NULL
);

CREATE TABLE charge (
    uid character varying NOT NULL,
    server_or_network_device_name character varying NOT NULL
);

CREATE TABLE connection (
    rack_number integer NOT NULL,
    note character varying NOT NULL,
    server_or_network_device_name character varying NOT NULL
);
