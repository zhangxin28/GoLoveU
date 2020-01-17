package tests

import (
	"database/sql"
	"goloveu/utils/db"
	"testing"
)

func TestQueryParams(t *testing.T) {
	type User struct {
		db.GormModel
		Username    sql.NullString `gorm:"size:32;unique;" json:"username" form:"username"`
		Email       sql.NullString `gorm:"size:128;unique;" json:"email" form:"email"`
		Nickname    string         `gorm:"size:16;" json:"nickname" form:"nickname"`
		Avatar      string         `gorm:"type:text" json:"avatar" form:"avatar"`
		Password    string         `gorm:"size:512" json:"password" form:"password"`
		Status      int            `gorm:"index:idx_status;not null" json:"status" form:"status"`
		Roles       string         `gorm:"type:text" json:"roles" form:"roles"`
		Type        int            `gorm:"not null" json:"type" form:"type"`
		Description string         `gorm:"type:text" json:"description" form:"description"`
		CreateTime  int64          `json:"createTime" form:"createTime"`
		UpdateTime  int64          `json:"updateTime" form:"updateTime"`
	}
	type args struct {
		selectors  []string
		parameters map[string][]string
	}

	tests := []struct {
		name       string
		args       args
		myUsername string
	}{
		// TODO: Add test cases.
		{
			name:       "test for query paramters with Username,Email",
			myUsername: "Simon",
			args: args{
				selectors: []string{"Username", "Email"},
				parameters: map[string][]string{
					"Username": []string{"Simon", "Bob"},
					"Email":    []string{"simon@testgo.com"},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			newSQLCnd := db.NewSQLCnd(tt.args.selectors...).
				Where("1 = ?", 1).
				In("Username", tt.args.parameters["Username"]).
				Eq("Email", tt.args.parameters["Email"])

			user := &User{}
			newSQLCnd.Build(db.DB()).Find(&user)

			if user.Username.Valid && user.Username.String != tt.myUsername {
				t.Errorf("FindUser() Username = %v, want %v", user.Username.String, tt.myUsername)
			}
		})
	}
}
