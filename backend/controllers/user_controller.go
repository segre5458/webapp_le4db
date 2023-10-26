package controllers

import (
	"database/sql"
	"log"

	"github.com/gin-gonic/gin"

	"github.com/segre5458/webapp_le4db/backend/models"
)

func Index(ctx *gin.Context, DB *sql.DB) {
	sql := "SELECT * FROM RACK"
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
