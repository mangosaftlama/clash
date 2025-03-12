package clash

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
)

type PlayerAchievementProgress struct {
	Stars int `json:"stars"`
	Value int `json:"value"`
	// JsonLocalizedName
	Name   string `json:"name"`
	Target int    `json:"target"`
	// JsonLocalizedName
	Info string `json:"info"`
	// JsonLocalizedName
	CompletionInfo string `json:"completionInfo"`
	// Enum: [ HOME_VILLAGE, BUILDER_BASE, CLAN_CAPITAL ]
	Village string `json:"village"`
}

type PlayerAchievementProgressList []PlayerAchievementProgress

type PlayerClan struct {
	Tag       string `json:"tag"`
	ClanLevel int    `json:"clanLevel"`
	Name      string `json:"name"`
	BadgeUrls any    `json:"badgeUrls"`
}

type PlayerHouse struct {
	Elements PlayerHouseElementList `json:"elements"`
}

type PlayerHouseElement struct {
	Id int `json:"id"`
	// Enum: [ GROUND, ROOF, FOOT, DECO ]
	Type string `json:"type"`
}

type PlayerHouseElementList []PlayerHouseElement

type PlayerItemLevel struct {
	Level    int    `json:"level"`
	Name     string `json:"name"`
	MaxLevel int    `json:"maxLevel"`
	// Enum: [ HOME_VILLAGE, BUILDER_BASE, CLAN_CAPITAL ]
	Village            string `json:"village"`
	SuperTroopIsActive bool   `json:"superTroopIsActive"`
	Equipment          any    `json:"equipment"`
}

type PlayerItemLevelList []PlayerItemLevel

type PlayerLegendStatistics struct{}

type PlayerRole string

type Player struct {
	playerTag string
	// League
	League struct {
		// JsonLocalizedName
		Name     string `json:"name"`
		Id       int    `json:"id"`
		IconUrls any    `json:"iconUrls"`
	} `json:"league"`
	// BuilderBaseLeague
	BuilderBaseLeague struct {
		// JsonLocalizedName
		Name string `json:"name"`
		Id   int    `json:"id"`
	} `json:"builderBaseLeague"`
	// PlayerClan
	Clan PlayerClan `json:"clan"`
	// Enum: [ NOT_MEMBER, MEMBER, LEADER, ADMIN, COLEADER ]
	Role string `json:"role"`
	// Enum: [ OUT, IN ]
	WarPreference       string                 `json:"warPreference"`
	AttackWins          int                    `json:"attackWins"`
	TownHallLevel       int                    `json:"townHallLevel"`
	TownHallWeaponLevel int                    `json:"townHallWeaponLevel"`
	LegendStatistics    PlayerLegendStatistics `json:"legendStatistics"`
	Troops              PlayerItemLevelList    `json:"troops"`
	Heroes              PlayerItemLevelList    `json:"heroes"`
	HeroEquipment       PlayerItemLevelList    `json:"heroEquipment"`
	Spells              PlayerItemLevelList    `json:"spells"`
	// LabelList
	Labels []struct {
		// Label
		Name     string `json:"name"`
		Id       int    `json:"id"`
		IconUrls any    `json:"iconUrls"`
	} `json:"labels"`
	Tag                      string                        `json:"tag"`
	Name                     string                        `json:"name"`
	ExpLevel                 int                           `json:"expLevel"`
	Trophies                 int                           `json:"trophies"`
	BestTrophies             int                           `json:"bestTrophies"`
	Donations                int                           `json:"donations"`
	DonationsReceived        int                           `json:"donationsReceived"`
	BuilderHallLevel         int                           `json:"builderHallLevel"`
	BuilderBaseTrophies      int                           `json:"builderBaseTrophies"`
	BestBuilderBaseTrophies  int                           `json:"bestBuilderBaseTrophies"`
	WarStars                 int                           `json:"warStars"`
	Achievements             PlayerAchievementProgressList `json:"achievements"`
	ClanCapitalContributions int                           `json:"clanCapitalContributions"`
	PlayerHouse              PlayerHouse                   `json:"playerHouse"`
}

type VerifyTokenResponse struct {
	Tag    string `json:"tag"`
	Token  string `json:"token"`
	Status string `json:"status"`
}

func (c Client) GetPlayer(playerTag string) (*Player, *ClientError, error) {
	url := "https://api.clashofclans.com/v1/players/" + playerTag

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, nil, err
	}
	req.Header = c.DefaultHeader()
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, nil, err
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, nil, err
	}

	switch resp.StatusCode {
	case 200:
		resp := &Player{}
		err = json.Unmarshal(body, resp)
		return resp, nil, err
	case 400, 403, 404, 429, 500, 503:
		clientErr := &ClientError{}
		err = json.Unmarshal(body, clientErr)
		return nil, clientErr, err
	}
	return nil, nil, nil
}

func (c Client) VerifyPlayerToken(playerTag string, token string) (*VerifyTokenResponse, *ClientError, error) {
	url := "https://api.clashofclans.com/v1/players/" + playerTag + "/verifytoken"

	bodyBuffer := bytes.NewBuffer(make([]byte, 0))

	encoder := json.NewEncoder(bodyBuffer)
	err := encoder.Encode(token)
	if err != nil {
		return nil, nil, err
	}

	req, err := http.NewRequest("POST", url, bodyBuffer)
	if err != nil {
		return nil, nil, err
	}
	req.Header = c.DefaultHeader()
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, nil, err
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, nil, err
	}

	switch resp.StatusCode {
	case 200:
		resp := &VerifyTokenResponse{}
		err = json.Unmarshal(body, resp)
		return resp, nil, err
	case 400, 403, 404, 429, 500, 503:
		clientErr := &ClientError{}
		err = json.Unmarshal(body, clientErr)
		return nil, clientErr, err
	}
	return nil, nil, nil
}
