package poke


type Pokemon struct {
	Abilities              []Abilities   `json:"abilities"`
	BaseExperience         int           `json:"base_experience"`
	Forms                  []Forms       `json:"forms"`
	GameIndices            []GameIndices `json:"game_indices"`
	Height                 int           `json:"height"`
	HeldItems              []HeldItems   `json:"held_items"`
	ID                     int           `json:"id"`
	IsDefault              bool          `json:"is_default"`
	LocationAreaEncounters string        `json:"location_area_encounters"`
	Moves                  []Moves       `json:"moves"`
	Name                   string        `json:"name"`
	Order                  int           `json:"order"`
	Species                Species       `json:"species"`
	Sprites                Sprites       `json:"sprites"`
	Stats                  []Stats       `json:"stats"`
	Types                  []Types       `json:"types"`
	Weight                 int           `json:"weight"`
}
type Ability struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}
type Abilities struct {
	Ability  Ability `json:"ability"`
	IsHidden bool    `json:"is_hidden"`
	Slot     int     `json:"slot"`
}
type Forms struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}
type Version struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}
type GameIndices struct {
	GameIndex int     `json:"game_index"`
	Version   Version `json:"version"`
}
type Item struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}
type VersionDetails struct {
	Rarity  int     `json:"rarity"`
	Version Version `json:"version"`
}
type HeldItems struct {
	Item           Item             `json:"item"`
	VersionDetails []VersionDetails `json:"version_details"`
}
type Move struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}
type MoveLearnMethod struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}
type VersionGroup struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}
type VersionGroupDetails struct {
	LevelLearnedAt  int             `json:"level_learned_at"`
	MoveLearnMethod MoveLearnMethod `json:"move_learn_method"`
	VersionGroup    VersionGroup    `json:"version_group"`
}
type Moves struct {
	Move                Move                  `json:"move"`
	VersionGroupDetails []VersionGroupDetails `json:"version_group_details"`
}
type Species struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}
type DreamWorld struct {
	FrontDefault string      `json:"front_default"`
	FrontFemale  interface{} `json:"front_female"`
}
type OfficialArtwork struct {
	FrontDefault string `json:"front_default"`
}
type Other struct {
	DreamWorld      DreamWorld      `json:"dream_world"`
	OfficialArtwork OfficialArtwork `json:"official-artwork"`
}
type RedBlue struct {
	BackDefault  string `json:"back_default"`
	BackGray     string `json:"back_gray"`
	FrontDefault string `json:"front_default"`
	FrontGray    string `json:"front_gray"`
}
type Yellow struct {
	BackDefault  string `json:"back_default"`
	BackGray     string `json:"back_gray"`
	FrontDefault string `json:"front_default"`
	FrontGray    string `json:"front_gray"`
}
type GenerationI struct {
	RedBlue RedBlue `json:"red-blue"`
	Yellow  Yellow  `json:"yellow"`
}
type Crystal struct {
	BackDefault  string `json:"back_default"`
	BackShiny    string `json:"back_shiny"`
	FrontDefault string `json:"front_default"`
	FrontShiny   string `json:"front_shiny"`
}
type Gold struct {
	BackDefault  string `json:"back_default"`
	BackShiny    string `json:"back_shiny"`
	FrontDefault string `json:"front_default"`
	FrontShiny   string `json:"front_shiny"`
}
type Silver struct {
	BackDefault  string `json:"back_default"`
	BackShiny    string `json:"back_shiny"`
	FrontDefault string `json:"front_default"`
	FrontShiny   string `json:"front_shiny"`
}
type GenerationIi struct {
	Crystal Crystal `json:"crystal"`
	Gold    Gold    `json:"gold"`
	Silver  Silver  `json:"silver"`
}
type Emerald struct {
	FrontDefault string `json:"front_default"`
	FrontShiny   string `json:"front_shiny"`
}
type FireredLeafgreen struct {
	BackDefault  string `json:"back_default"`
	BackShiny    string `json:"back_shiny"`
	FrontDefault string `json:"front_default"`
	FrontShiny   string `json:"front_shiny"`
}
type RubySapphire struct {
	BackDefault  string `json:"back_default"`
	BackShiny    string `json:"back_shiny"`
	FrontDefault string `json:"front_default"`
	FrontShiny   string `json:"front_shiny"`
}
type GenerationIii struct {
	Emerald          Emerald          `json:"emerald"`
	FireredLeafgreen FireredLeafgreen `json:"firered-leafgreen"`
	RubySapphire     RubySapphire     `json:"ruby-sapphire"`
}
type DiamondPearl struct {
	BackDefault      string      `json:"back_default"`
	BackFemale       interface{} `json:"back_female"`
	BackShiny        string      `json:"back_shiny"`
	BackShinyFemale  interface{} `json:"back_shiny_female"`
	FrontDefault     string      `json:"front_default"`
	FrontFemale      interface{} `json:"front_female"`
	FrontShiny       string      `json:"front_shiny"`
	FrontShinyFemale interface{} `json:"front_shiny_female"`
}
type HeartgoldSoulsilver struct {
	BackDefault      string      `json:"back_default"`
	BackFemale       interface{} `json:"back_female"`
	BackShiny        string      `json:"back_shiny"`
	BackShinyFemale  interface{} `json:"back_shiny_female"`
	FrontDefault     string      `json:"front_default"`
	FrontFemale      interface{} `json:"front_female"`
	FrontShiny       string      `json:"front_shiny"`
	FrontShinyFemale interface{} `json:"front_shiny_female"`
}
type Platinum struct {
	BackDefault      string      `json:"back_default"`
	BackFemale       interface{} `json:"back_female"`
	BackShiny        string      `json:"back_shiny"`
	BackShinyFemale  interface{} `json:"back_shiny_female"`
	FrontDefault     string      `json:"front_default"`
	FrontFemale      interface{} `json:"front_female"`
	FrontShiny       string      `json:"front_shiny"`
	FrontShinyFemale interface{} `json:"front_shiny_female"`
}
type GenerationIv struct {
	DiamondPearl        DiamondPearl        `json:"diamond-pearl"`
	HeartgoldSoulsilver HeartgoldSoulsilver `json:"heartgold-soulsilver"`
	Platinum            Platinum            `json:"platinum"`
}
type Animated struct {
	BackDefault      string      `json:"back_default"`
	BackFemale       interface{} `json:"back_female"`
	BackShiny        string      `json:"back_shiny"`
	BackShinyFemale  interface{} `json:"back_shiny_female"`
	FrontDefault     string      `json:"front_default"`
	FrontFemale      interface{} `json:"front_female"`
	FrontShiny       string      `json:"front_shiny"`
	FrontShinyFemale interface{} `json:"front_shiny_female"`
}
type BlackWhite struct {
	Animated         Animated    `json:"animated"`
	BackDefault      string      `json:"back_default"`
	BackFemale       interface{} `json:"back_female"`
	BackShiny        string      `json:"back_shiny"`
	BackShinyFemale  interface{} `json:"back_shiny_female"`
	FrontDefault     string      `json:"front_default"`
	FrontFemale      interface{} `json:"front_female"`
	FrontShiny       string      `json:"front_shiny"`
	FrontShinyFemale interface{} `json:"front_shiny_female"`
}
type GenerationV struct {
	BlackWhite BlackWhite `json:"black-white"`
}
type OmegarubyAlphasapphire struct {
	FrontDefault     string      `json:"front_default"`
	FrontFemale      interface{} `json:"front_female"`
	FrontShiny       string      `json:"front_shiny"`
	FrontShinyFemale interface{} `json:"front_shiny_female"`
}
type XY struct {
	FrontDefault     string      `json:"front_default"`
	FrontFemale      interface{} `json:"front_female"`
	FrontShiny       string      `json:"front_shiny"`
	FrontShinyFemale interface{} `json:"front_shiny_female"`
}
type GenerationVi struct {
	OmegarubyAlphasapphire OmegarubyAlphasapphire `json:"omegaruby-alphasapphire"`
	XY                     XY                     `json:"x-y"`
}
type Icons struct {
	FrontDefault string      `json:"front_default"`
	FrontFemale  interface{} `json:"front_female"`
}
type UltraSunUltraMoon struct {
	FrontDefault     string      `json:"front_default"`
	FrontFemale      interface{} `json:"front_female"`
	FrontShiny       string      `json:"front_shiny"`
	FrontShinyFemale interface{} `json:"front_shiny_female"`
}
type GenerationVii struct {
	Icons             Icons             `json:"icons"`
	UltraSunUltraMoon UltraSunUltraMoon `json:"ultra-sun-ultra-moon"`
}
type GenerationViii struct {
	Icons Icons `json:"icons"`
}
type Versions struct {
	GenerationI    GenerationI    `json:"generation-i"`
	GenerationIi   GenerationIi   `json:"generation-ii"`
	GenerationIii  GenerationIii  `json:"generation-iii"`
	GenerationIv   GenerationIv   `json:"generation-iv"`
	GenerationV    GenerationV    `json:"generation-v"`
	GenerationVi   GenerationVi   `json:"generation-vi"`
	GenerationVii  GenerationVii  `json:"generation-vii"`
	GenerationViii GenerationViii `json:"generation-viii"`
}
type Sprites struct {
	BackDefault      string      `json:"back_default"`
	BackFemale       interface{} `json:"back_female"`
	BackShiny        string      `json:"back_shiny"`
	BackShinyFemale  interface{} `json:"back_shiny_female"`
	FrontDefault     string      `json:"front_default"`
	FrontFemale      interface{} `json:"front_female"`
	FrontShiny       string      `json:"front_shiny"`
	FrontShinyFemale interface{} `json:"front_shiny_female"`
	Other            Other       `json:"other"`
	Versions         Versions    `json:"versions"`
}
type Stat struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}
type Stats struct {
	BaseStat int  `json:"base_stat"`
	Effort   int  `json:"effort"`
	Stat     Stat `json:"stat"`
}
type Type struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}
type Types struct {
	Slot int  `json:"slot"`
	Type Type `json:"type"`
}

// response body methods

func (p *Pokemon) IsResponseBody() {}
