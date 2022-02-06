package main

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Todo struct {
	gorm.Model
	Id   int
	Text string
}

type Database struct {
	db *gorm.DB
}

func (d *Database) Connect() (err error) {
	var db *gorm.DB
	db, err = gorm.Open(sqlite.Open("todo.db"), &gorm.Config{})
	d.db = db
	return
}

func (d *Database) Migrate() {
	d.db.AutoMigrate(&Todo{})
}

func (d *Database) Create(text string) {
	d.db.Create(&Todo{Text: text})
}

func (d *Database) FetchTodos() (todos []Todo) {
	d.db.Find(&todos)
	return
}

func (d *Database) FetchTodo(id string) (todo Todo) {
	d.db.Find(&todo, id)
	return
}

func (d *Database) UpdateTodo(id string, text string) {
	var todo Todo
	d.db.Find(&todo, id)
	d.db.Model(&todo).Update("Text", text)
}

func (d *Database) DeleteTodo(id string) {
	var todo Todo
	d.db.Delete(&todo, id)
}

func init() {
	d := Database{}
	err := d.Connect()
	if err != nil {
		panic("failed to connect database")
	}
	d.Migrate()
}
