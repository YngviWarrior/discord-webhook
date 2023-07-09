package discordstructs

// ChannelUrl example: /1057831941514207282/AlQM4bvDA3DX1nFTiYwzMcZ0hKfgjWwlBndZd1g-FUblfglaLSFJmK9kZudQYiNJi2T9
type Notification struct {
	Channel    string `json:"channel"`
	ChannelUrl string `json:"channel_url"`
	Content    string `json:"content"`
}
