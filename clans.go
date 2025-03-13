package clash

import (
	"encoding/json"
	"io"
	"net/http"
	"net/url"
	"strconv"
)

type CapitalLeague struct {
	// JsonLocalizedName
	Name string `json:"name"`
	Id   int    `json:"id"`
}

type WarLeague struct {
	// JsonLocalizedName
	Name string `json:"name"`
	Id   int    `json:"id"`
}

type Clan struct {
	WarLeague     WarLeague      `json:"warLeague"`
	CapitalLeague CapitalLeague  `json:"capitalLeague"`
	MemberList    ClanMemberList `json:"memberList"`
	Tag           string         `json:"tag"`
	// Enum: [ UNKNOWN, ALWAYS, MORE_THAN_ONCE_PER_WEEK, ONCE_PER_WEEK, LESS_THAN_ONCE_PER_WEEK, NEVER, ANY ]
	WarFrequency                string    `json:"warFrequency"`
	ClanLevel                   int       `json:"clanLevel"`
	WarWinStreak                int       `json:"warWinStreak"`
	WarWins                     int       `json:"warWins"`
	WarTies                     int       `json:"warTies"`
	WarLosses                   int       `json:"warLosses"`
	ClanPoints                  int       `json:"clanPoints"`
	ClanBuilderBasePoints       int       `json:"clanBuilderBasePoints"`
	ClanCapitalPoints           int       `json:"clanCapitalPoints"`
	RequiredTrophies            int       `json:"requiredTrophies"`
	RequiredBuilderBaseTrophies int       `json:"requiredBuilderBaseTrophies"`
	RequiredTownhallLevel       int       `json:"requiredTownhallLevel"`
	IsFamilyFriendly            bool      `json:"isFamilyFriendly"`
	IsWarLogPublic              bool      `json:"isWarLogPublic"`
	ChatLanguage                Language  `json:"chatLanguage"`
	Labels                      LabelList `json:"labels"`
	Name                        string    `json:"name"`
	Location                    Location  `json:"location"`
	// Enum: [ OPEN, INVITE_ONLY, CLOSED ]
	Type        string      `json:"type"`
	Members     int         `json:"members"`
	Description string      `json:"description"`
	ClanCapital ClanCapital `json:"clanCapital"`
	BadgeUrls   any         `json:"badgeUrls"`
}

type ClanCapital struct {
	CapitalHallLevel int                  `json:"capitalHallLevel"`
	Districts        ClanDistrictDataList `json:"districts"`
}

type ClanDistrictData struct {
	// JsonLocalizedName
	Name              string `json:"name"`
	Id                int    `json:"id"`
	DistrictHallLevel int    `json:"districtHallLevel"`
}

type ClanDistrictDataList []ClanDistrictData

type ClanList struct {
	Clans []Clan `json:"items"`
}

type ClanMember struct {
	League            League            `json:"league"`
	BuilderBaseLeague BuilderBaseLeague `json:"builderBaseLeague"`
	Tag               string            `json:"tag"`
	Name              string            `json:"name"`
	// Enum: [ NOT_MEMBER, MEMBER, LEADER, ADMIN, COLEADER ]
	Role                string      `json:"role"`
	TownHallLevel       int         `json:"townHallLevel"`
	ExpLevel            int         `json:"expLevel"`
	ClanRank            int         `json:"clanRank"`
	PreviousClanRank    int         `json:"previousClanRank"`
	Donations           int         `json:"donations"`
	DonationsReceived   int         `json:"donationsReceived"`
	Trophies            int         `json:"trophies"`
	BuilderBaseTrophies int         `json:"builderBaseTrophies"`
	PlayerHouse         PlayerHouse `json:"playerHouse"`
}

type ClanMemberList []ClanMember

type ClanSearchQuery struct {
	query map[string]string
}

func NewClanSearchQuery() ClanSearchQuery {
	return ClanSearchQuery{query: make(map[string]string)}
}

func (q *ClanSearchQuery) SetName(name string) {
	q.query["name"] = name
}

func (q *ClanSearchQuery) SetWarFrequency(warFrequency string) {
	q.query["warFrequency"] = warFrequency
}

func (q *ClanSearchQuery) SetLocationId(locationId int) {
	q.query["locationId"] = strconv.Itoa(locationId)
}

func (q *ClanSearchQuery) SetMinMembers(minMembers int) {
	q.query["minMembers"] = strconv.Itoa(minMembers)
}

func (q *ClanSearchQuery) SetMaxMembers(maxMembers int) {
	q.query["maxMembers"] = strconv.Itoa(maxMembers)
}

func (q *ClanSearchQuery) SetMinClanPoints(minClanPoints int) {
	q.query["minClanPoints"] = strconv.Itoa(minClanPoints)
}

func (q *ClanSearchQuery) SetMinClanLevel(minClanLevel int) {
	q.query["minClanLevel"] = strconv.Itoa(minClanLevel)
}

func (q *ClanSearchQuery) SetLimit(limit int) {
	q.query["limit"] = strconv.Itoa(limit)
}

func (q *ClanSearchQuery) SetAfter(after string) {
	q.query["after"] = after
}

func (q *ClanSearchQuery) SetBefore(before string) {
	q.query["before"] = before
}

func (q *ClanSearchQuery) SetLabelIds(labelIds string) {
	q.query["labelIds"] = labelIds
}

func (c Client) GetClan(clanTag string) (*Clan, *ClientError, error) {
	url := "https://api.clashofclans.com/v1/clans/" + url.QueryEscape(clanTag)

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
		resp := &Clan{}
		err = json.Unmarshal(body, resp)
		return resp, nil, err
	case 400, 403, 404, 429, 500, 503:
		clientErr := &ClientError{}
		err = json.Unmarshal(body, clientErr)
		return nil, clientErr, err
	}
	return nil, nil, nil
}

func (c Client) SearchClans(query ClanSearchQuery) (*ClanList, *ClientError, error) {
	endpoint := "https://api.clashofclans.com/v1/clans"

	url, err := BuildQueryURL(endpoint, query.query)
	if err != nil {
		return nil, nil, err
	}

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
		resp := &ClanList{}
		err = json.Unmarshal(body, resp)
		return resp, nil, err
	case 400, 403, 404, 429, 500, 503:
		clientErr := &ClientError{}
		err = json.Unmarshal(body, clientErr)
		return nil, clientErr, err
	}
	return nil, nil, nil
}
