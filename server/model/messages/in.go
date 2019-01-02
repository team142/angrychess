package messages

//BaseMessage for all incoming messages
type BaseMessage struct {
	Msg string `json:"msg"`
}

//MessageNick describes a nickname change
type MessageNick struct {
	Nick string `json:"nick"`
}

//MessageJoinGame for joining a game by id
type MessageJoinGame struct {
	ID string `json:"id"`
}

//MessageMove describes a movement
type MessageMove struct {
	PieceID string `json:"pieceId"`
	FromX   int    `json:"fx"`
	FromY   int    `json:"fy"`
	ToX     int    `json:"tx"`
	ToY     int    `json:"ty"`
}

//MessagePlace describes placing a piece
type MessagePlace struct {
	ID  string `json:"id"`
	ToX int    `json:"tx"`
	ToY int    `json:"ty"`
}

type MessageChangeSeat struct {
	Seat int `json:"seat"`
}
