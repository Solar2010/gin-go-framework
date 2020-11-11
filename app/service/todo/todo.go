package todo

import "goskeleton/app/model"

type Service struct {

}

func CreateTodoFactory() *Service {

	return &Service{}
}

func (s *Service) Submit(title string, content string, mark string, deadline int64, visitList string, shareList string, isPrivate int, userId int64)  bool{

	return model.CreateTodoFactory("").Add(title, content, mark, deadline, visitList, shareList, isPrivate, userId)
}