package entities

type Match struct {
	SeasonID              uint64                `json:"seasonId"`
	QueueID               uint64                `json:"queueId"`
	MatchId               uint64                `json:"gameId"`
	ParticipantIdentities []participantIdentity `json:"participantIdentity"`
	GameVersion           string                `json:"gameVersion"`
	PlatformID            string                `json:"platformId"`
	GameMode              string                `json:"gameMode"`
	MapID                 uint                  `json:"mapId"`
	GameType              string                `json:"gameType"`
	Teams                 []Team                `json:"teams"`
	Participants          []Participant         `json:"participants"`
	GameDuration          uint64                `json:"gameDuration"`
	GameCreation          uint64                `json:"gameCreation"`
}

type MatchPlayer struct {
	AccountID         uint64 `json:"accountId"`
	SummonerID        uint64 `json:"summonerId"`
	SummonerName      string `json:"summonerName"`
	ProfileIcon       uint64 `json:"profileIcon"`
	PlatformID        string `json:"platformId"`
	CurrentAccountID  uint64 `json:"currentAccountId"`
	MatchHistoryURI   string `json:"matchHistoryUri"`
	CurrentPlatformID string `json:"currentPlatformId"`
}

type participantIdentity struct {
	Player        MatchPlayer `json:"player"`
	ParticipantID uint        `json:"participantId"`
}

type Ban struct {
	PickTurn   uint `json:"pickTurn"`
	ChampionID uint `json:"championId"`
}

type Team struct {
	FirstDragon          bool   `json:"firstDragon"`
	Bans                 []Ban  `json:"bans"`
	FirstInhibitor       bool   `json:"firstInhibitor"`
	Win                  string `json:"win"`
	FirstRiftHerald      bool   `json:"firstRiftHerald"`
	FirstBaron           bool   `json:"firstBaron"`
	BaronKills           uint   `json:"baronKills"`
	RiftHeraldKills      uint   `json:"riftHeraldKills"`
	FirstBlood           bool   `json:"firstBlood"`
	TeamID               uint   `json:"teamId"`
	FirstTower           bool   `json:"firstTower"`
	VilemawKills         uint   `json:"vilemawKills"`
	InhibitorKills       uint   `json:"inhibitorKills"`
	TowerKills           uint   `json:"towerKills"`
	DominionVictoryScore uint   `json:"dominionVictoryScore"`
	DragonKills          uint   `json:"dragonKills"`
}

type Participant struct {
	ParticipantID             uint                `json:"participantId"`
	TeamID                    uint                `json:"teamId"`
	ChampionID                uint                `json:"championId"`
	Spell1ID                  uint                `json:"spell1Id"`
	Spell2ID                  uint                `json:"spell2Id"`
	Stats                     ParticipantStats    `json:"stats"`
	TimeLine                  ParticipantTimeLine `json:"timeline"`
	HighestAchievedSeasonTier string              `json:"highestAchievedSeasonTier"`
}

type Delta struct {
	Zero10 uint64 `json:"0-10"`
	One020 uint64 `json:"10-20"`
	Two030 uint64 `json:"20-30"`
}

type ParticipantTimeLine struct {
	Lane                        string `json:"lane"`
	Role                        string `json:"role"`
	ParticipantID               uint   `json:"participantId"`
	CsDiffPerMinDeltas          Delta  `json:"csDiffPerMinDeltas"`
	GoldPerMinDeltas            Delta  `json:"goldPerMinDeltas"`
	XpDiffPerMinDeltas          Delta  `json:"xpDiffPerMinDeltas"`
	CreepsPerMinDeltas          Delta  `json:"creepsPerMinDeltas"`
	XpPerMinDeltas              Delta  `json:"xpPerMinDeltas"`
	DamageTakenDiffPerMinDeltas Delta  `json:"damageTakenDiffPerMinDeltas"`
	DamageTakenPerMinDeltas     Delta  `json:"damageTakenPerMinDeltas"`
}

type ParticipantStats struct {
	ParticipantID                   uint `json:"participantId"`
	ChampLevel                      uint `json:"champLevel"`
	GoldEarned                      uint `json:"goldEarned"`
	GoldSpent                       uint `json:"goldSpent"`
	Win                             bool `json:"win"`
	Kills                           uint `json:"kills"`
	Deaths                          uint `json:"deaths"`
	Assists                         uint `json:"assists"`
	DoubleKills                     uint `json:"doubleKills"`
	TripleKills                     uint `json:"tripleKills"`
	QuadraKills                     uint `json:"quadraKills"`
	PentaKills                      uint `json:"pentaKills"`
	LargestMultiKill                uint `json:"largestMultiKill"`
	LargestKillingSpree             uint `json:"largestKillingSpree"`
	PlayerScore0                    uint `json:"playerScore0"`
	PlayerScore1                    uint `json:"playerScore1"`
	PlayerScore2                    uint `json:"playerScore2"`
	PlayerScore3                    uint `json:"playerScore3"`
	PlayerScore4                    uint `json:"playerScore4"`
	PlayerScore5                    uint `json:"playerScore5"`
	PlayerScore6                    uint `json:"playerScore6"`
	PlayerScore7                    uint `json:"playerScore7"`
	PlayerScore8                    uint `json:"playerScore8"`
	PlayerScore9                    uint `json:"playerScore9"`
	Item0                           uint `json:"item0"`
	Item1                           uint `json:"item1"`
	Item2                           uint `json:"item2"`
	Item3                           uint `json:"item3"`
	Item4                           uint `json:"item4"`
	Item5                           uint `json:"item5"`
	Item6                           uint `json:"item6"`
	Perk0                           uint `json:"perk0"`
	Perk0Var1                       uint `json:"perk0Var1"`
	Perk0Var2                       uint `json:"perk0Var2"`
	Perk0Var3                       uint `json:"perk0Var3"`
	Perk1                           uint `json:"perk1"`
	Perk1Var1                       uint `json:"perk1Var1"`
	Perk1Var2                       uint `json:"perk1Var2"`
	Perk1Var3                       uint `json:"perk1Var3"`
	Perk2                           uint `json:"perk2"`
	Perk2Var1                       uint `json:"perk2Var1"`
	Perk2Var2                       uint `json:"perk2Var2"`
	Perk2Var3                       uint `json:"perk2Var3"`
	Perk3                           uint `json:"perk3"`
	Perk3Var1                       uint `json:"perk3Var1"`
	Perk3Var2                       uint `json:"perk3Var2"`
	Perk3Var3                       uint `json:"perk3Var3"`
	Perk4                           uint `json:"perk4"`
	Perk4Var1                       uint `json:"perk4Var1"`
	Perk4Var2                       uint `json:"perk4Var2"`
	Perk4Var3                       uint `json:"perk4Var3"`
	Perk5                           uint `json:"perk5"`
	Perk5Var1                       uint `json:"perk5Var1"`
	Perk5Var2                       uint `json:"perk5Var2"`
	Perk5Var3                       uint `json:"perk5Var3"`
	PerkPrimaryStyle                uint `json:"perkPrimaryStyle"`
	PerkSubStyle                    uint `json:"perkSubStyle"`
	TotalDamageDealt                uint `json:"totalDamageDealt"`
	MagicDamageDealt                uint `json:"magicDamageDealt"`
	TrueDamageDealt                 uint `json:"trueDamageDealt"`
	PhysicalDamageDealt             uint `json:"physicalDamageDealt"`
	TotalDamageTaken                uint `json:"totalDamageTaken"`
	MagicDamageTaken                uint `json:"magicalDamageTaken"`
	TrueDamageTaken                 uint `json:"trueDamageTaken"`
	PhysicalDamageTaken             uint `json:"physicalDamageTaken"`
	TotalDamageDealtToChampions     uint `json:"totalDamageDealtToChampions"`
	MagicDamageDealtToChampions     uint `json:"magicDamageDealtToChampions"`
	TrueDamageDealtToChampions      uint `json:"trueDamageDealtToChampions"`
	PhysicalDamageDealtToChampions  uint `json:"physicalDamageDealtToChampions"`
	DamageDealtToTurrets            uint `json:"damageDealtToTurrets"`
	DamageDealtToObjectives         uint `json:"damageDealtToObjectives"`
	FirstBloodKill                  bool `json:"firstBloodKill"`
	FirstTowerKill                  bool `json:"firstTowerKill"`
	FirstInhibitorKill              bool `json:"firstInhibitorKill"`
	FirstBloodAssist                bool `json:"firstBloodAssist"`
	FirstTowerAssist                bool `json:"firstTowerAssist"`
	FirstInhibitorAssist            bool `json:"firstInhibitorAssist"`
	NeutralMinionsKilled            uint `json:"neutralMinionsKilled"`
	NeutralMinionsKilledTeamJungle  uint `json:"neutralMinionsKilledTeamJungle"`
	NeutralMinionsKilledEnemyJungle uint `json:"neutralMinionsKilledEnemyJungle"`
	TotalHeal                       uint `json:"totalHeal"`
	TotalScoreRank                  uint `json:"totalScoreRank"`
	VisionScore                     uint `json:"visionScore"`
	TotalPlayerScore                uint `json:"totalPlayerScore"`
	CombatPlayerScore               uint `json:"combatPlayerScore"`
	ObjectivePlayerScore            uint `json:"objectivePlayerScore"`
	TotalUnitsHealed                uint `json:"totalUnitsHealed"`
	TotalMinionsKilled              uint `json:"totalMinionsKilled"`
	TotalTimeCrowdControlDealt      uint `json:"totalTimeCrowdControlDealt"`
	TurretKills                     uint `json:"turretKills"`
	UnrealKills                     uint `json:"unrealKills"`
	InhibitorKills                  uint `json:"inhibitorKills"`
	WardsKilled                     uint `json:"wardsKilled"`
	LongestTimeSpentLiving          uint `json:"longestTimeSpentLiving"`
	LargestCriticalStrike           uint `json:"largestCriticalStrike"`
	DamageSelfMitigated             uint `json:"damageSelfMitigated"`
	SightWardsBoughtInGame          uint `json:"sightWardsBoughtInGame"`
	WardsPlaced                     uint `json:"wardsPlaced"`
	KillingSprees                   uint `json:"killingSprees"`
	VisionWardsBoughtInGame         uint `json:"visionWardsBoughtInGame"`
	TimeCCingOthers                 uint `json:"timeCCingOthers"`
}
