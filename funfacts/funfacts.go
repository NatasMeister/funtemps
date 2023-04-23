package funfacts

/**
  Implementer funfacts-funksjon:
    GetFunFacts(about string) []string
      hvor about kan ha en av tre testverdier, -
        sun, luna eller terra
  Sett inn alle Funfucts i en struktur
  type FunFacts struct {
      Sun []string
      Luna []string
      Terra []string
  }
*/

type FunFacts struct {
  Sun   []string
  Luna  []string
  Terra []string
}

func GetFunFacts(about string) []string {
  facts := FunFacts{
      Sun: []string{
          "Sola er ei stjerne.",
          "Sola er nesten 5 milliarder år gammel.",
      },
      Luna: []string{
          "Månen går i bane rundt jorda.",
          "Månen er faktisk ikke lagd av ost.",
      },
      Terra: []string{
          "Vi bor på jorda.",
          "Jorda er kul og sånt.",
      },
  }

  switch about {
  case "sun":
      return facts.Sun
  case "luna":
      return facts.Luna
  case "terra":
      return facts.Terra
  default:
      return []string{"Unknown fact category."}
  }
}