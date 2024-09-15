package fetcher

type RelevantCard struct {
  Name string `json:"name"`
}

type RelevantBattle struct {
  Won bool
  BattleTime string
  Deck []RelevantCard
  OppDeck []RelevantCard
  TowersDestroyed int
  OppTowersDestroyed int
  TrophiesOnStart int
  OppTrophiesOnStart int
}

type RelevantPlayer struct {
  Name string
  BattlesPlayed int
  Level int
  Trophies int
  Battles []RelevantBattle
}

type ResponsePlayer struct {
  Trophies int `json:"trophies"`
  BattleCount int `json:"battleCount"`
}

type BattleData struct {
  TowersDestroyed int `json:"crowns"`
  TrophiesOnStart int `json:"startingTrophies"`
  Deck []RelevantCard `json:"cards"`
}

type ResponseBattle struct {
  Time string `json:"battleTime"`
  BattleData []BattleData `json:"team"`
  OppBattleData []BattleData `json:"opponent"`
}

type ResponseBattleList struct {
  Battles []ResponseBattle
}


