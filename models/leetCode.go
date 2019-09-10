package models

// key-questionId,value-Question
var QuestionMap map[string]*Question

type Question struct {
	Id       string
	Target   string
	KeyWord  string
	Solution string
}