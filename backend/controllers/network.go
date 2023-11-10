package controllers

import (
	"database/sql"
	"log"
	"strings"

	"github.com/gin-gonic/gin"

	"github.com/segre5458/webapp_le4db/backend/models"
)

func AddNetworkDevice(ctx *gin.Context, DB *sql.DB) {
	unitNumber := ctx.Param("unitNumber")
	deviceName := ctx.PostForm("deviceName")
	portCount := ctx.PostForm("portCount")
	ipAddressesStr := ctx.PostForm("ipAddresses")
	role := ctx.PostForm("role")
	portMacAddressesStr := ctx.PostForm("portMacAddresses")

	ipAddressesSlice := strings.Split(ipAddressesStr, ",")
	portMacAddressesSlice := strings.Split(portMacAddressesStr, ",")
	ipAddresses := "{" + strings.Join(ipAddressesSlice, ",") + "}"
	portMacAddresses := "{" + strings.Join(portMacAddressesSlice, ",") + "}"

	sql := "INSERT INTO NETWORK_DEVICE (device_name, port_count, ip_addresses, role, port_mac_addresses) VALUES ($1, $2, $3, $4, $5)"
	_, err := DB.Exec(sql, deviceName, portCount, ipAddresses, role, portMacAddresses)
	if err != nil {
		log.Fatal(err)
	}

	sql = "INSERT INTO NETWORK_DEVICE_PLACEMENT (port_mac_address, rack_unit_number) VALUES ($1, $2)"
	for _, portMacAddress := range portMacAddressesSlice {
		_, err = DB.Exec(sql, portMacAddress, unitNumber)
		if err != nil {
			log.Fatal(err)
		}
	}

	ctx.Redirect(302, "/rack/"+unitNumber)
}

func SearchNetworkDevice(ctx *gin.Context, DB *sql.DB) {
	deviceName := ctx.Query("deviceName")

	sql := `
	SELECT DISTINCT nd.*
	FROM NETWORK_DEVICE nd
	JOIN NETWORK_DEVICE_PLACEMENT ndp ON nd.port_mac_addresses @> ARRAY[ndp.port_mac_address]
	WHERE nd.device_name = $1;
	`
	var networkDevices []models.NetworkDevice
	rows, err := DB.Query(sql, deviceName)
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
		"networkDevices": networkDevices,
	})
}

func DeleteNetworkDevice(ctx *gin.Context, DB *sql.DB) {
	unitNumber := ctx.Param("unitNumber")
	deviceName := ctx.PostForm("deviceName")

	sql := "DELETE FROM NETWORK_DEVICE WHERE device_name = $1"
	_, err := DB.Exec(sql, deviceName)
	if err != nil {
		log.Fatal(err)
	}

	ctx.Redirect(302, "/rack/"+unitNumber)
}

func EditNetworkDevice(ctx *gin.Context, DB *sql.DB) {
	deviceName := ctx.Param("deviceName")

	sql := "SELECT * FROM NETWORK_DEVICE WHERE device_name = $1"
	var networkDevices []models.NetworkDevice
	rows, err := DB.Query(sql, deviceName)
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

	ctx.HTML(200, "edit_rack.html", gin.H{
		"networkDevices": networkDevices,
	})
}

func UpdateNetworkDevice(ctx *gin.Context, DB *sql.DB) {
	unitNumber := ctx.Param("unitNumber")
	deviceName := ctx.Param("deviceName")
	portCount := ctx.PostForm("portCount")
	ipAddressesStr := ctx.PostForm("ipAddresses")
	role := ctx.PostForm("role")
	portMacAddressesStr := ctx.PostForm("portMacAddresses")

	ipAddressesSlice := strings.Split(ipAddressesStr, ",")
	portMacAddressesSlice := strings.Split(portMacAddressesStr, ",")
	ipAddresses := "{" + strings.Join(ipAddressesSlice, ",") + "}"
	portMacAddresses := "{" + strings.Join(portMacAddressesSlice, ",") + "}"

	sql := "UPDATE NETWORK_DEVICE SET port_count = $1, ip_addresses = $2, role = $3, port_mac_addresses = $4 WHERE device_name = $5"
	_, err := DB.Exec(sql, portCount, ipAddresses, role, portMacAddresses, deviceName)
	if err != nil {
		log.Fatal(err)
	}

	sql = "DELETE FROM NETWORK_DEVICE_PLACEMENT WHERE rack_unit_number = $1"
	_, err = DB.Exec(sql, unitNumber)
	if err != nil {
		log.Fatal(err)
	}

	sql = "INSERT INTO NETWORK_DEVICE_PLACEMENT (port_mac_address, rack_unit_number) VALUES ($1, $2)"
	for _, portMacAddress := range portMacAddressesSlice {
		_, err = DB.Exec(sql, portMacAddress, unitNumber)
		if err != nil {
			log.Fatal(err)
		}
	}

	ctx.Redirect(302, "/rack/"+unitNumber)
}
