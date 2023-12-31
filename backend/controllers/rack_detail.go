package controllers

import (
	"database/sql"
	"log"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/sessions"

	"github.com/segre5458/webapp_le4db/backend/models"
)

func Rack(ctx *gin.Context, DB *sql.DB) {
	unitNumber := ctx.Param("unitNumber")

	session := sessions.Default(ctx)
	userRole := session.Get("role")

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
	FROM network_device nd
	JOIN network_device_placement ndp ON ndp.port_mac_address = ANY(nd.port_mac_addresses)
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
		"Role": userRole,
	})
}
