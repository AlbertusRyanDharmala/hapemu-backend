package service

import (
	"database/sql"
	"fmt"
	"hapemu/model"

	_ "github.com/lib/pq"
)

const (
	host     = "aws-0-ap-southeast-1.pooler.supabase.com"
	port     = 6543
	user     = "postgres.teprmsxuirxhmriekpgh"
	password = "testinkdatabasepassword"
	dbname   = "postgres"

	// host     = "aws-0-ap-southeast-1.pooler.supabase.com"
	// port     = 5432
	// user     = "postgres.yovcppevikilglvpktzq"
	// password = "HapemuPostgres123"
	// dbname   = "postgres"
)

// region get from database
func getSmartphoneList() []model.Smartphone {
	var smartphones []model.Smartphone
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		fmt.Println("failed on connecting to database " + err.Error())
		return smartphones
	}
	defer db.Close()

	sqlStatement := `SELECT name, "segmentPrice", processor,"dxomarkScore", battery, ram, storage FROM smartphones`
	rows, err := db.Query(sqlStatement)
	if err != nil {
		fmt.Println("failed on query " + err.Error())
		return smartphones
	}
	defer rows.Close()

	for rows.Next() {
		var smartphone model.Smartphone
		var name sql.NullString
		var segmentPrice sql.NullString
		var processor sql.NullString
		var dxomarkScore sql.NullInt64
		var battery sql.NullString
		var ram sql.NullString
		var storage sql.NullString

		err := rows.Scan(&name, &segmentPrice, &processor, &dxomarkScore, &battery, &ram, &storage)
		if err != nil {
			fmt.Println("Error when scanning to Go struct:", err)
			return smartphones
		}

		smartphone.Name = name.String
		if name.Valid {
			smartphone.Name = name.String
		} else {
			smartphone.Name = "" // Default value or handle appropriately
		}

		if segmentPrice.Valid {
			smartphone.SegmentPrice = segmentPrice.String
		} else {
			smartphone.SegmentPrice = "" // Default value or handle appropriately
		}

		if processor.Valid {
			smartphone.Processor = processor.String
		} else {
			smartphone.Processor = "" // Default value or handle appropriately
		}

		if dxomarkScore.Valid {
			smartphone.DxomarkScore = int(dxomarkScore.Int64)
		} else {
			smartphone.DxomarkScore = 0 // Default value or handle appropriately
		}

		if battery.Valid {
			smartphone.Battery = battery.String
		} else {
			smartphone.Battery = "" // Default value or handle appropriately
		}

		if ram.Valid {
			smartphone.Ram = ram.String
		} else {
			smartphone.Ram = "" // Default value or handle appropriately
		}

		if storage.Valid {
			smartphone.Storage = storage.String
		} else {
			smartphone.Storage = "" // Default value or handle appropriately
		}
		smartphones = append(smartphones, smartphone)
	}

	// for _, phone := range smartphones {
	// 	fmt.Println("Name: " + phone.Name)
	// 	fmt.Println("SegmentPrice: " + phone.SegmentPrice)
	// 	fmt.Println("Processor: " + phone.Processor)
	// 	fmt.Print("Dxomark Score: ")
	// 	fmt.Println(phone.DxomarkScore)
	// 	fmt.Println("Battery: " + phone.Battery)
	// 	fmt.Println("Ram: " + phone.Ram)
	// 	fmt.Println("Storage: " + phone.Storage)
	// 	fmt.Println()
	// }
	return smartphones
}

// region get from database
func getSmartphoneByName(phoneName string) model.Smartphone {
	var smartphone model.Smartphone
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		fmt.Println("failed on connecting to database " + err.Error())
		return smartphone
	}
	defer db.Close()

	sqlStatement := `SELECT name, "segmentPrice", processor,"dxomarkScore", battery, ram, storage FROM smartphones WHERE name = $1`
	rows, err := db.Query(sqlStatement, phoneName)
	if err != nil {
		fmt.Println("failed on query " + err.Error())
		return smartphone
	}
	defer rows.Close()
	var name sql.NullString
	var segmentPrice sql.NullString
	var processor sql.NullString
	var dxomarkScore sql.NullInt64
	var battery sql.NullString
	var ram sql.NullString
	var storage sql.NullString

	for rows.Next() {
		err := rows.Scan(&name, &segmentPrice, &processor, &dxomarkScore, &battery, &ram, &storage)
		if err != nil {
			fmt.Println("Error when scanning to Go struct:", err)
			return smartphone
		}

		smartphone.Name = name.String
		if name.Valid {
			smartphone.Name = name.String
		} else {
			smartphone.Name = "" // Default value or handle appropriately
		}

		if segmentPrice.Valid {
			smartphone.SegmentPrice = segmentPrice.String
		} else {
			smartphone.SegmentPrice = "" // Default value or handle appropriately
		}

		if processor.Valid {
			smartphone.Processor = processor.String
		} else {
			smartphone.Processor = "" // Default value or handle appropriately
		}

		if dxomarkScore.Valid {
			smartphone.DxomarkScore = int(dxomarkScore.Int64)
		} else {
			smartphone.DxomarkScore = 0 // Default value or handle appropriately
		}

		if battery.Valid {
			smartphone.Battery = battery.String
		} else {
			smartphone.Battery = "10000" // Default value or handle appropriately
		}

		if ram.Valid {
			smartphone.Ram = ram.String
		} else {
			smartphone.Ram = "10000" // Default value or handle appropriately
		}

		if storage.Valid {
			smartphone.Storage = storage.String
		} else {
			smartphone.Storage = "10000" // Default value or handle appropriately
		}
		fmt.Println(smartphone)
	}
	return smartphone
}
