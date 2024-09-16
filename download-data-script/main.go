package main

import (
  "log"
  "context"

  "download-data-script/fetcher"
  "download-data-script/db"
  "github.com/joho/godotenv"
)

func main() {
  err := godotenv.Load()
  if err != nil {
    log.Fatal("Error loading .env file")
  }
  client :=  db.Connect()
  collection := client.Database("db").Collection("players")


  c := make(chan fetcher.RelevantPlayer)
  getRelevantPlayer(c)

  relevantPlayer := <-c

  _, err = collection.InsertOne(context.Background(), relevantPlayer)
  if err != nil {
    log.Fatal(err)
  }
}

func makeRelevantPlayerFrom(member fetcher.Member) fetcher.RelevantPlayer {
  var playerInfo fetcher.RelevantPlayer

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
    var battleInfo fetcher.RelevantBattle

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

  return playerInfo
}

func getRelevantPlayer(c chan fetcher.RelevantPlayer) {
  for _, clan := range fetcher.GetClans() {
    for _, member := range clan.GetMembers() {
      c <- makeRelevantPlayerFrom(member)
    }
  }

}
