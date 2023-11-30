package controllers

import (
	"database/sql"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/sessions"

	"github.com/segre5458/webapp_le4db/backend/models"
)

func AddServer(ctx *gin.Context, DB *sql.DB) {
	unitNumber := ctx.Param("unitNumber")
	serverName := ctx.PostForm("serverName")
	ipAddress := ctx.PostForm("ipAddress")
	os := ctx.PostForm("os")
	cpu := ctx.PostForm("cpu")
	memory := ctx.PostForm("memory")
	storage := ctx.PostForm("storage")
	nicMacAddress := ctx.PostForm("nicMacAddress")

	sql := "INSERT INTO SERVER (server_name, ip_address, os, cpu, memory, storage, nic_mac_address) VALUES ($1, $2, $3, $4, $5, $6, $7)"
	_, err := DB.Exec(sql, serverName, ipAddress, os, cpu, memory, storage, nicMacAddress)
	if err != nil {
		log.Fatal(err)
	}

	sql = "INSERT INTO SERVER_PLACEMENT (nic_mac_address, rack_unit_number) VALUES ($1, $2)"
	_, err = DB.Exec(sql, nicMacAddress, unitNumber)
	if err != nil {
		log.Fatal(err)
	}

	ctx.Redirect(302, "/rack/"+unitNumber)
}

func SearchServer(ctx *gin.Context, DB *sql.DB) {
	serverName := ctx.Query("serverName")

	session := sessions.Default(ctx)
	userRole := session.Get("role")

	sql := `
	SELECT DISTINCT s.*
	FROM SERVER s
	JOIN SERVER_PLACEMENT sp ON s.nic_mac_address = sp.nic_mac_address
	WHERE s.server_name = $1;
	`
	var servers []models.Server
	rows, err := DB.Query(sql, serverName)
	defer rows.Close()
	if err != nil {
		log.Fatal(err)
	}

	for rows.Next() {
		var server models.Server
		err := rows.Scan(&server.ServerName, &server.IpAddress, &server.Os, &server.Cpu, &server.Memory, &server.Storage, &server.NicMacAddress)
		if err != nil {
			log.Fatal(err)
		}
		servers = append(servers, server)
	}

	ctx.HTML(200, "rack.html", gin.H{
		"servers": servers,
		"Role": userRole,
	})
}

func DeleteServer(ctx *gin.Context, DB *sql.DB) {
	unitNumber := ctx.Param("unitNumber")
	nicMacAddress := ctx.PostForm("nicMacAddress")

	sql := "DELETE FROM SERVER WHERE nic_mac_address = $1"
	_, err := DB.Exec(sql, nicMacAddress)
	if err != nil {
		log.Fatal(err)
	}

	sql = "DELETE FROM SERVER_PLACEMENT WHERE nic_mac_address = $1"
	_, err = DB.Exec(sql, nicMacAddress)
	if err != nil {
		log.Fatal(err)
	}

	ctx.Redirect(302, "/rack/"+unitNumber)
}

func EditServer(ctx *gin.Context, DB *sql.DB) {
	serverName := ctx.Param("serverName")
	unitNumber := ctx.Param("unitNumber")

	sql := "SELECT * FROM SERVER WHERE server_name = $1"
	var servers []models.Server
	rows, err := DB.Query(sql, serverName)
	defer rows.Close()
	if err != nil {
		log.Fatal(err)
	}

	for rows.Next() {
		var server models.Server
		err := rows.Scan(&server.ServerName, &server.IpAddress, &server.Os, &server.Cpu, &server.Memory, &server.Storage, &server.NicMacAddress)
		if err != nil {
			log.Fatal(err)
		}
		servers = append(servers, server)
	}

	ctx.HTML(200, "edit_server.html", gin.H{
		"server": servers[0],
		"unitNumber": unitNumber,
	})
}

func UpdateServer(ctx *gin.Context, DB *sql.DB) {
	unitNumber := ctx.Param("unitNumber")
	serverName := ctx.Param("serverName")
	ipAddress := ctx.PostForm("ipAddress")
	os := ctx.PostForm("os")
	cpu := ctx.PostForm("cpu")
	memory := ctx.PostForm("memory")
	storage := ctx.PostForm("storage")
	nicMacAddress := ctx.PostForm("nicMacAddress")

	sql := "UPDATE SERVER SET ip_address = $1, os = $2, cpu = $3, memory = $4, storage = $5, nic_mac_address = $6 WHERE server_name = $7"
	_, err := DB.Exec(sql, ipAddress, os, cpu, memory, storage, nicMacAddress, serverName)
	if err != nil {
		log.Fatal(err)
	}

	sql = "UPDATE SERVER_PLACEMENT SET nic_mac_address = $1 WHERE rack_unit_number = $2"
	_, err = DB.Exec(sql, nicMacAddress, unitNumber)
	if err != nil {
		log.Fatal(err)
	}

	ctx.Redirect(302, "/rack/"+unitNumber)
}
