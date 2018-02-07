package controllers

import "ideaparLog/Repo"

type ControllerBase struct {
}

var GoalRepo = new(Repo.GoalRepo)
var PvRepo = new(Repo.PvRepo)
