// (c) 2019 Max Wolter - ALL RIGHTS RESERVED

package website

type Owner struct {
	ID               string  `json:"id"`
	Nickname         *string `json:"nickname"`
	AvatarURL        *string `json:"avatarURL"`
	TwitterAvatarURL *string `json:"twitterAvatarURL"`
	Status           string  `json:"status"`
}
