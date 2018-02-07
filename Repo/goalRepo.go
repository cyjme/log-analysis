package Repo

import (
	"ideaparLog/model"
	"ideaparLog/database"
	"log"
)

const (
	list   = "select * from goal"
	create = "insert into goal(name, `desc`, dataUrl) values (?, ?, ?)"
	update = "update goal set name=?, `desc`=?, dataUrl=? where id = ?"
	delete = "delete from goal where id = ?"
)

type GoalRepo struct {
}

func (repo *GoalRepo) List() []model.Goal {
	stmt, err := database.DBCon.Prepare(list)
	if err != nil {
		log.Fatal(err)
	}

	defer stmt.Close()

	rows, err := stmt.Query()
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	result := []model.Goal{}

	for rows.Next() {
		goal := model.Goal{}
		err := rows.Scan(&goal.Id, &goal.Name, &goal.Desc, &goal.DataUrl)
		if err != nil {
			log.Fatal(err)
		}

		result = append(result, goal)
	}

	return result
}

func (repo *GoalRepo) Create(goal *model.Goal) {
	stmt, err := database.DBCon.Prepare(create)
	if err != nil {
		log.Fatal(err)
	}

	defer stmt.Close()

	stmt.Exec(goal.Name, goal.Desc, goal.DataUrl)
}

func (repo *GoalRepo) Update(goal *model.Goal) {
	stmt, err := database.DBCon.Prepare(update)
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	stmt.Exec(goal.Name, goal.Desc, goal.DataUrl, goal.Id)
}

func (repo *GoalRepo) Delete(id int) {
	stmt, err := database.DBCon.Prepare(delete)

	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	stmt.Exec(id)
}
