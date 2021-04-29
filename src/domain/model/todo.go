package model

//Todo is TodoModel
type Todo struct {
	ID        int    `db:"id" json:"id"`                        //TaskのID
	Task      string `db:"task" json:"task"`                    //Task自体
	LimitDate string `gorm:"column:limitDate" json:"limitDate"` //Taskの完了期限
	Status    bool   `db:"status" json:"status"`                //Taskの状態(0=未済,1=済)
}
