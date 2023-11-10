package controllers

import (
	"database/sql"
	"log"
	"strings"

	"github.com/gin-gonic/gin"

	"github.com/segre5458/webapp_le4db/backend/models"
)

func Rack(ctx *gin.Context, DB *sql.DB) {
	unitNumber := ctx.Param("unitNumber")

	sql := `
		SELECT *
		FROM SERVER
		JOIN SERVER_PLACEMENT ON SERVER.nic_mac_address = SERVER_PLACEMENT.nic_mac_address
		WHERE SERVER_PLACEMENT.rack_unit_number = $1
	`
	var servers []models.Server
	rows, err := DB.Query(sql, unitNumber)
	defer rows.Close()
	if err != nil {
		log.Fatal(err)
	}

	for rows.Next() {
		var server models.Server
		var serverPlacement models.ServerPlacement
		err := rows.Scan(&server.ServerName, &server.IpAddress, &server.Os, &server.Cpu, &server.Memory, &server.Storage, &server.NicMacAddress, &serverPlacement.NicMacAddress, &serverPlacement.RackUnitNumber)
		if err != nil {
			log.Fatal(err)
		}
		servers = append(servers, server)
	}

	sql = `
	SELECT DISTINCT nd.*
	FROM NETWORK_DEVICE nd
	JOIN NETWORK_DEVICE_PLACEMENT ndp ON nd.port_mac_addresses @> ARRAY[ndp.port_mac_address]
	WHERE ndp.rack_unit_number = $1;
	`
	var networkDevices []models.NetworkDevice
	rows, err = DB.Query(sql, unitNumber)
	defer rows.Close()
	if err != nil {
		log.Fatal(err)
	}

	for rows.Next() {
		var networkDevice models.NetworkDevice
		var IpAddressStr string
		var portMacAddressStr string

		err := rows.Scan(&networkDevice.DeviceName, &networkDevice.PortCount, &IpAddressStr, &networkDevice.Role, &portMacAddressStr)
		if err != nil {
			log.Fatal(err)
		}
		networkDevice.IpAddress = strings.Split(IpAddressStr, ",")
		networkDevice.PortMacAddress = strings.Split(portMacAddressStr, ",")
		
		networkDevices = append(networkDevices, networkDevice)
	}

	ctx.HTML(200, "rack.html", gin.H{
		"unitNumber": unitNumber,
		"servers": servers,
		"networkDevices": networkDevices,
	})
}
