package datatype

import "time"

// TimelineEntry timeline entry data type
type TimelineEntry struct {
	BlockID                    ID
	UserID                     ID
	UserName                   string
	UserProfileImageURL        string
	AssetID                    ID
	AssetName                  string
	AssetSymbol                string
	OracleID                   ID
	OracleName                 string
	Text                       string
	EthereumTransactionAddress string
	Status                     int
	YtVideoID                  string
	FavoritesCount             int
	DidUserLike                bool
	CreatedAt                  time.Time
	CreatedAtHuman             string
	Images                     []string
	DidUserLikeTopic           bool
	Reactions                  []PostEmoji
}

// PostEmoji post emoji type
type PostEmoji struct {
	Users     []ID   `json:"Users"`
	EmojiID   ID     `json:"EmojiID"`
	EmojiName string `json:"EmojiName"`
	EmojiLogo string `json:"EmojiLogo"`
}
