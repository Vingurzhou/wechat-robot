package main

import (
	"log"
	"os"
	"path/filepath"

	"github.com/Vingurzhou/wechat-robot/db"

	"gorm.io/gen"
	"gorm.io/gorm"
)

func readSQLFile(filePath string) (string, error) {
	content, err := os.ReadFile(filePath)
	if err != nil {
		return "", err
	}
	return string(content), nil
}

func executeSQLFile(db *gorm.DB, filePath string) error {
	sqlContent, err := readSQLFile(filePath)
	if err != nil {
		return err
	}

	// 执行 SQL 语句
	return db.Exec(sqlContent).Error
}

func main() {
	db := db.NewDatabase()

	err := executeSQLFile(db.GormDB, "cmd/gen/wechat_robot_db.sql")
	if err != nil {
		log.Fatal(err)
	}

	log.Println("SQL file executed successfully")

	projectRoot, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	g := gen.NewGenerator(gen.Config{
		OutPath: filepath.Join(projectRoot, "query"),
		Mode:    gen.WithoutContext | gen.WithDefaultQuery | gen.WithQueryInterface,
	})

	g.UseDB(db.GormDB)
	g.ApplyBasic(
		g.GenerateModelAs("group", "Group"),
		g.GenerateModelAs("user", "User"),
		g.GenerateModelAs("msg", "Msg"),
	)
	g.Execute()

	log.Println("GORM GEN executed successfully")

}
