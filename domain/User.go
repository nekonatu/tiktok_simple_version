package domain

type User struct {
	// id、密码、随机盐字段在返回给用户时应屏蔽
	IsFollow       bool   `json:"is_follow" gorm:"-"`
	Id             int64  `json:"id" gorm:"primaryKey"`
	FollowCount    int64  `json:"follow_count,omitempty" gorm:"-"`
	FollowerCount  int64  `json:"follower_count,omitempty" gorm:"-"`
	TotalFavorited int64  `json:"total_favorited,omitempty"`
	FavoriteCount  int64  `json:"favorite_count,omitempty"`
	Salt           string `json:"-" gorm:"type:char(4)"`
	Name           string `json:"name" gorm:"type:varchar(32); index"`
	Pwd            string `json:"-" gorm:"type:char(60)"`
}

type UserLoginResponse struct {
	Response
	UserId int64  `json:"user_id"`
	Token  string `json:"token"`
}

type UserResponse struct {
	Response
	User User `json:"user"`
}

type UserFollowListResponse struct {
	Response
	UserFollowList []User
}
