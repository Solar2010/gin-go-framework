package model

type Model struct {
	*BaseModel
	CreatedTime int64
	UpdatedTime int64
	DeletedTime int64
}

func (m Model) Add(table string, uniques map[string] string, args ...interface{})  {

}