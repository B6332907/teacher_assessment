package entity

import (
	//"fmt"
	//"time"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"

	"gorm.io/driver/sqlite"
)

var db *gorm.DB

func DB() *gorm.DB {

	return db

}

func SetupDatabase() {

	database, err := gorm.Open(sqlite.Open("sa-65.db"), &gorm.Config{})

	if err != nil {

		panic("failed to connect database")

	}

	// Migrate the schema

	database.AutoMigrate(
		&Student{},
		&Teacher{},
		&Teacher_assessment{},
		&Teaching_duration{},
		&Content_difficulty_level{},
		&Content_quality{},
	)

	db = database
	phone_1, err := bcrypt.GenerateFromPassword([]byte("0935463156"), 14)
	phone_2, err := bcrypt.GenerateFromPassword([]byte("0917271607"), 14)

	// student data
	db.Model(&Student{}).Create(&Student{
		Name:         "Khunjira Pantuket",
		User_student: "Student5001",
		College_year: 3,
		Gpx:          3.00,
		//Faculty_ID:    1004,
		Date_of_birth: "20/06/2001",
		Phone:         string(phone_1),
		Parent:        "Banjoong Puntuket",

		//Teacher_ID:    5001,
		//Officer_ID:    0101,
	})
	db.Model(&Student{}).Create(&Student{
		Name:         "Natiluk Srisakkwa",
		User_student: "Student5002",
		College_year: 3,
		Gpx:          3.00,
		//Faculty_ID:    1004,
		Date_of_birth: "26/07/2001",
		Phone:         string(phone_2),
		Parent:        "Sawat Puntuket",

		//Teacher_ID:    5001,
		//Officer_ID:    0102,
	})
	// teacher_data data
	db.Model(&Teacher{}).Create(&Teacher{
		//Faculty_id:                1004,
		Level: "Mr.",
		Name:  "Somchay Tunsamai",
		Email: "Somchay@gmail.com",
		//Graduate_faculty_level_id: 2002,
		//Officer_id:                0001,
	})
	db.Model(&Teacher{}).Create(&Teacher{
		//Faculty_ID:                1001,
		Level: "Mrs.",
		Name:  "Somying Junpon",
		Email: "Somying@gmail.com",
		//Graduate_faculty_level_id: 2002,
		//Officer_id:                0001,
	})

	//Teaching_duration data

	duration_one := Teaching_duration{
		Description: "longer than necessary",
	}
	db.Model(&Teaching_duration{}).Create(&duration_one)

	duration_two := Teaching_duration{
		Description: "Take the time to study",
	}
	db.Model(&Teaching_duration{}).Create(&duration_two)

	duration_three := Teaching_duration{
		Description: "is most suitable for learning",
	}
	db.Model(&Teaching_duration{}).Create(&duration_three)

	//content_quality data

	quality_one := Content_quality{
		Description: "content should be update",
	}
	db.Model(&Content_quality{}).Create(&quality_one)

	quality_two := Content_quality{
		Description: "content is not difficult and not easy",
	}
	db.Model(&Content_quality{}).Create(&quality_two)

	quality_three := Content_quality{
		Description: "is suitable for the course",
	}
	db.Model(&Content_quality{}).Create(&quality_three)

	quality_four := Content_quality{
		Description: "it is most appropriate for the quality of the content",
	}
	db.Model(&Content_quality{}).Create(&quality_four)

	//content difficulty level data

	difficulty_one := Content_difficulty_level{
		Description: "amend",
	}
	db.Model(&Content_difficulty_level{}).Create(&difficulty_one)

	difficulty_two := Content_difficulty_level{
		Description: "Fair",
	}
	db.Model(&Content_difficulty_level{}).Create(&difficulty_two)

	difficulty_three := Content_difficulty_level{
		Description: "appropriate",
	}
	db.Model(&Content_difficulty_level{}).Create(&difficulty_three)

	difficulty_four := Content_difficulty_level{
		Description: "easy to understand",
	}
	db.Model(&Content_difficulty_level{}).Create(&difficulty_four)

	var Khunjira Student
	var Natiluk Student
	db.Raw("SELECT * FROM Students WHERE name = ? ", "Khunjira Pantuket").Scan(&Khunjira)
	db.Raw("SELECT * FROM Students WHERE name = ? ", "Natiluk Srisakkwa").Scan(&Natiluk)

	var Somchay Teacher
	var Somying Teacher
	db.Raw("SELECT * FROM Teachers WHERE Name = ? ", "Somchay Tunsamai").Scan(&Somchay)
	db.Raw("SELECT * FROM Teachers WHERE Name = ? ", "Somying Junpon").Scan(&Somying)

}
