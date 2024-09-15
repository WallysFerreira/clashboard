package fetcher

import (
  "net/http"
  "fmt"
  "log"
  "encoding/json"
  "os"
)

func MakeRequest(method string, endpoint string) *http.Response {
  key := os.Getenv("CLASH_API_KEY")

  request, err := http.NewRequest(method, endpoint, nil)
  if err != nil {
    log.Fatal("Something went wrong when creating request")
  }

  request.Header.Set(
    "Authorization",
    fmt.Sprintf(
      "Bearer %s",
      key,
    ),
  )

  log.Print(request.URL)

  response, err := http.DefaultClient.Do(request)
  if err != nil {
    log.Fatal("Something went wrong when requesting")
  }
  if response.StatusCode != 200 {
    var t interface{}
    json.NewDecoder(response.Body).Decode(&t)
    log.Fatal(t)
  }

  return response
}

func GetRelevantInformation() {
  for _, clan := range getClans() {
    for _, member := range clan.GetMembers() {
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

      responseBattles := member.GetBattles()

      for _, battle := range responseBattles {
        var battleInfo RelevantBattle

        if (battle.BattleData[0].TowersDestroyed > battle.OppBattleData[0].TowersDestroyed) {
          battleInfo.Won = true
        } else {
          battleInfo.Won = false
        }

        battleInfo.BattleTime = battle.Time
        battleInfo.Deck = battle.BattleData[0].Deck
        battleInfo.OppDeck = battle.OppBattleData[0].Deck
        battleInfo.TowersDestroyed = battle.BattleData[0].TowersDestroyed
        battleInfo.OppTowersDestroyed = battle.OppBattleData[0].TowersDestroyed
        battleInfo.TrophiesOnStart = battle.BattleData[0].TrophiesOnStart
        battleInfo.OppTrophiesOnStart = battle.OppBattleData[0].TrophiesOnStart

        playerInfo.Battles = append(playerInfo.Battles, battleInfo)
      }

      /*
      if len(playerInfo.Battles) > 0 {
        log.Println(playerInfo.Battles[0].Deck[0].Name)
      }
      */
    }
  }
}

func getClans() []Clan {
  response := MakeRequest(http.MethodGet, "https://api.clashroyale.com/v1/clans?maxMembers=20&limit=10")
  defer response.Body.Close()

  var clanItems ClanItems
  err := json.NewDecoder(response.Body).Decode(&clanItems)
  if err != nil {
    log.Fatal(err)
  }

  return clanItems.Clans;
}

