package controllers

import (
	"database/sql"
	"log"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"

	"github.com/segre5458/webapp_le4db/backend/models"
)

func Index(ctx *gin.Context, DB *sql.DB) {
	sql := "SELECT * FROM RACK"
	session := sessions.Default(ctx)
	userRole := session.Get("role")
	var racks []models.Rack
	rows, err := DB.Query(sql)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		var rack models.Rack
		err := rows.Scan(&rack.Height, &rack.UnitNumber, &rack.Width, &rack.Specification, &rack.Depth)
		if err != nil {
			log.Fatal(err)
		}
		racks = append(racks, rack)
	}

	ctx.HTML(200, "index.html", gin.H{
		"racks": racks,
		"Role": userRole,
	})
}

func AddRack(ctx *gin.Context, DB *sql.DB) {
	height := ctx.PostForm("height")
	unitNumber := ctx.PostForm("unitNumber")
	width := ctx.PostForm("width")
	specification := ctx.PostForm("specification")
	depth := ctx.PostForm("depth")

	sql := "INSERT INTO RACK (height, unit_number, width, specification, depth) VALUES ($1, $2, $3, $4, $5)"
	_, err := DB.Exec(sql, height, unitNumber, width, specification, depth)
	if err != nil {
		log.Fatal(err)
	}

	ctx.Redirect(302, "/")
}

func SearchRack(ctx *gin.Context, DB *sql.DB) {
	unitNumber := ctx.Query("unitNumber")

	sql := "SELECT * FROM RACK WHERE unit_number = $1"
	var racks []models.Rack
	rows, err := DB.Query(sql, unitNumber)
	defer rows.Close()
	if err != nil {
		log.Fatal(err)
	}

	session := sessions.Default(ctx)
	userRole := session.Get("role")

	for rows.Next() {
		var rack models.Rack
		err := rows.Scan(&rack.Height, &rack.UnitNumber, &rack.Width, &rack.Specification, &rack.Depth)
		if err != nil {
			log.Fatal(err)
		}
		racks = append(racks, rack)
	}
	ctx.HTML(200, "index.html", gin.H{
		"racks": racks,
		"Role": userRole,
	})
}

func DeleteRack(ctx *gin.Context, DB *sql.DB) {
	unitNumber := ctx.PostForm("unitNumber")

	sql := "DELETE FROM RACK WHERE unit_number = $1"
	_, err := DB.Exec(sql, unitNumber)
	if err != nil {
		log.Fatal(err)
	}

	ctx.Redirect(302, "/")
}

func EditRack(ctx *gin.Context, DB *sql.DB) {
	unitNumber := ctx.Param("unitNumber")
	var rack models.Rack

	sql := "SELECT * FROM RACK WHERE unit_number = $1"
	row := DB.QueryRow(sql, unitNumber)
	err := row.Scan(&rack.Height, &rack.UnitNumber, &rack.Width, &rack.Specification, &rack.Depth)
	if err != nil {
		log.Fatal(err)
	}

	ctx.HTML(200, "edit_index.html", gin.H{
		"rack": rack,
	})
}

func UpdateRack(ctx *gin.Context, DB *sql.DB) {
	height := ctx.PostForm("height")
	unitNumber := ctx.PostForm("unitNumber")
	width := ctx.PostForm("width")
	specification := ctx.PostForm("specification")
	depth := ctx.PostForm("depth")

	sql := "UPDATE RACK SET height = $1, width = $2, specification = $3, depth = $4 WHERE unit_number = $5"
	_, err := DB.Exec(sql, height, width, specification, depth, unitNumber)
	if err != nil {
		log.Fatal(err)
	}

	ctx.Redirect(302, "/")
}
