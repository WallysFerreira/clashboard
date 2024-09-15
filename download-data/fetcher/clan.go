package fetcher

import (
  "log"
  "strings"
  "net/http"
  "fmt"
  "encoding/json"
)

type Clan struct {
  Tag string `json:"tag"`
}

type ClanItems struct {
  Clans []Clan `json:"items"`
}

type Member struct {
  Tag string `json:"tag"`
  Name string `json:"name"`
  ExpLevel int `json:"expLevel"`
}

type MemberItems struct {
  Members []Member `json:"items"`
}

func (c Clan) getRelevantPlayerInformation() {
  for _, member := range c.GetMembers() {
    var playerInfo RelevantPlayer
    responsePlayer := member.GetPlayer()

    playerInfo.Name = member.Name
    playerInfo.Level = member.ExpLevel
    playerInfo.Trophies = responsePlayer.Trophies
    playerInfo.BattlesPlayed = responsePlayer.BattleCount

    /*
    log.Println("Player name: ", playerInfo.Name)
    log.Println("Player level: ", playerInfo.Level)
    log.Println("Player trophies: ", playerInfo.Trophies)
    log.Println("Player battles: ", playerInfo.BattlesPlayed)
    */
  }
}

func (c Clan) GetMembers() []Member {
  htmlTransformedTag := strings.Replace(c.Tag, "#", "%23", 1)

  response := MakeRequest(
    http.MethodGet,
    fmt.Sprintf("https://api.clashroyale.com/v1/clans/%s/members", htmlTransformedTag),
  )
  defer response.Body.Close()

  var memberItems MemberItems
  err := json.NewDecoder(response.Body).Decode(&memberItems)
  if err != nil {
    log.Fatal(err)
  }

  return memberItems.Members
}

func (m Member) GetPlayer() ResponsePlayer {
  htmlTransformedTag := strings.Replace(m.Tag, "#", "%23", 1)

  response := MakeRequest(
    http.MethodGet,
    fmt.Sprintf("https://api.clashroyale.com/v1/players/%s", htmlTransformedTag),
  )
  defer response.Body.Close()

  var player ResponsePlayer
  err := json.NewDecoder(response.Body).Decode(&player)
  if err != nil {
    log.Fatal(err)
  }

  return player
}

func (m Member) GetBattles() []ResponseBattle {
  htmlTransformedTag := strings.Replace(m.Tag, "#", "%23", 1)

  response := MakeRequest(
    http.MethodGet,
    fmt.Sprintf("https://api.clashroyale.com/v1/players/%s/battlelog", htmlTransformedTag),
  )
  defer response.Body.Close()

  var responseBattleList []ResponseBattle
  err := json.NewDecoder(response.Body).Decode(&responseBattleList)
  if err != nil {
    log.Fatal(err)
  }

  return responseBattleList;
}

