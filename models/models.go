package models

type User struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
	UniqueID string `json:"unique_id"`
}

type Friendship struct {
	ID 	 int `json:"id"`
	UserID   int `json:"user_id"`
	FriendID int `json:"friend_id"`
	Status   string `json:"status"`
}

type Game struct {
	ID		 int `json:"id"`
	Player1ID int `json:"player1_id"`
	Player2ID int `json:"player2_id"`
	Status    string `json:"status"`
}