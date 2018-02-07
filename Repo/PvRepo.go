package Repo

import (
	"ideaparLog/model"
	"ideaparLog/database"
	"log"
	"fmt"
)

const (
	createSql     = "insert into pv(date, url, referrerUrl, browserUuId, userAgent, userId, userIp) values (?, ?, ?, ?, ?, ?, ?)"
	getLastRowSql = "select * from pv order by id desc limit 1"
)

type PvRepo struct {
}

func (repo *PvRepo) Create(pv *model.Pv) {
	stmt, err := database.DBCon.Prepare(createSql)
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	stmt.Exec(pv.Date, pv.Url, pv.ReferrerUrl, pv.BrowserUuId, pv.UserAgent, pv.UserId, pv.UserIp)
}

func (repo *PvRepo) List(query string) ([]model.Pv, error) {
	listSql := "select * from pv " + query
	rows, err := database.DBCon.Query(listSql)
	defer rows.Close()
	var pvs = []model.Pv{}

	if err != nil {
		fmt.Println("errrrr",err)
		return pvs, err
	}

	for rows.Next() {
		pv := model.Pv{}
		err = rows.Scan(&pv.Id, &pv.Date, &pv.Url, &pv.ReferrerUrl, &pv.BrowserUuId, &pv.UserAgent, &pv.UserId, &pv.UserIp)
		pvs = append(pvs, pv)
	}

	return pvs, err
}

func (repo *PvRepo) GetLastRow() model.Pv {
	pv := model.Pv{}
	result := database.DBCon.QueryRow(getLastRowSql)
	err := result.Scan(&pv.Id, &pv.Date, &pv.Url, &pv.ReferrerUrl, &pv.BrowserUuId, &pv.UserAgent, &pv.UserId, &pv.UserIp)
	if err != nil {
		fmt.Println(err)
	}

	return pv
}
