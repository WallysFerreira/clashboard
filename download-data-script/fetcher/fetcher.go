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

func GetClans() []Clan {
  response := MakeRequest(http.MethodGet, "https://api.clashroyale.com/v1/clans?minMembers=10&limit=100")
  defer response.Body.Close()

  var clanItems ClanItems
  err := json.NewDecoder(response.Body).Decode(&clanItems)
  if err != nil {
    log.Fatal(err)
  }

  return clanItems.Clans;
}

