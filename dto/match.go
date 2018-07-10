package dto

import "encoding/json"

// MatchDTO MatchDTO
type MatchDTO struct {
	SeasonID              uint64                   `json:"seasonId"`
	QueueID               uint64                   `json:"queueId"`
	MatchID               uint64                   `json:"gameId"`
	ParticipantIdentities []ParticipantIdentityDTO `json:"participants"`
	GameVersion           string                   `json:"gameVersion"`
	PlatformID            string                   `json:"platformId"`
	GameMode              string                   `json:"gameMode"`
	MapID                 uint                     `json:"mapId"`
	GameType              string                   `json:"gameType"`
	Teams                 []TeamStatsDTO           `json:"teams"`
	Participants          []ParticipantDTO         `json:"participants"`
	GameDuration          uint64                   `json:"gameDuration"`
	GameCreation          uint64                   `json:"gameCreation"`
}

// rawMatchDTO rawMatchDTO
type rawMatchDTO struct {
	SeasonID              uint64                   `json:"seasonId"`
	QueueID               uint64                   `json:"queueId"`
	MatchID               uint64                   `json:"gameId"`
	ParticipantIdentities []ParticipantIdentityDTO `json:"ParticipantIdentityDTO"`
	GameVersion           string                   `json:"gameVersion"`
	PlatformID            string                   `json:"platformId"`
	GameMode              string                   `json:"gameMode"`
	MapID                 uint                     `json:"mapId"`
	GameType              string                   `json:"gameType"`
	Teams                 []TeamStatsDTO           `json:"teams"`
	Participants          []rawParticipantDTO      `json:"participants"`
	GameDuration          uint64                   `json:"gameDuration"`
	GameCreation          uint64                   `json:"gameCreation"`
}

// PlayerDTO PlayerDTO
type PlayerDTO struct {
	AccountID         uint64 `json:"accountId"`
	SummonerID        uint64 `json:"summonerId"`
	SummonerName      string `json:"summonerName"`
	ProfileIcon       uint64 `json:"profileIcon"`
	PlatformID        string `json:"platformId"`
	CurrentAccountID  uint64 `json:"currentAccountId"`
	MatchHistoryURI   string `json:"matchHistoryUri"`
	CurrentPlatformID string `json:"currentPlatformId"`
}

// ParticipantIdentityDTO ParticipantIdentityDTO
type ParticipantIdentityDTO struct {
	Player        PlayerDTO `json:"player"`
	ParticipantID uint      `json:"participantId"`
}

// TeamBansDTO TeamBansDTO
type TeamBansDTO struct {
	PickTurn   uint `json:"pickTurn"`
	ChampionID uint `json:"championId"`
}

// TeamStatsDTO TeamStatsDTO
type TeamStatsDTO struct {
	FirstDragon          bool          `json:"firstDragon"`
	Bans                 []TeamBansDTO `json:"bans"`
	FirstInhibitor       bool          `json:"firstInhibitor"`
	Win                  string        `json:"win"`
	FirstRiftHerald      bool          `json:"firstRiftHerald"`
	FirstBaron           bool          `json:"firstBaron"`
	BaronKills           uint          `json:"baronKills"`
	RiftHeraldKills      uint          `json:"riftHeraldKills"`
	FirstBlood           bool          `json:"firstBlood"`
	TeamID               uint          `json:"teamId"`
	FirstTower           bool          `json:"firstTower"`
	VilemawKills         uint          `json:"vilemawKills"`
	InhibitorKills       uint          `json:"inhibitorKills"`
	TowerKills           uint          `json:"towerKills"`
	DominionVictoryScore uint          `json:"dominionVictoryScore"`
	DragonKills          uint          `json:"dragonKills"`
}

// ParticipantDTO ParticipantDTO
type ParticipantDTO struct {
	ParticipantID             uint                   `json:"participantId"`
	TeamID                    uint                   `json:"teamId"`
	ChampionID                uint                   `json:"championId"`
	PrimarySpellID            uint                   `json:"spell1Id"`
	SecondarySpellID          uint                   `json:"spell2Id"`
	Stats                     ParticipantStatsDTO    `json:"stats"`
	TimeLine                  ParticipantTimeLineDTO `json:"timeline"`
	HighestAchievedSeasonTier string                 `json:"highestAchievedSeasonTier"`
}

type rawParticipantDTO struct {
	ParticipantID             uint                   `json:"participantId"`
	TeamID                    uint                   `json:"teamId"`
	ChampionID                uint                   `json:"championId"`
	PrimarySpellID            uint                   `json:"spell1Id"`
	SecondarySpellID          uint                   `json:"spell2Id"`
	Stats                     rawParticipantStatsDTO `json:"stats"`
	TimeLine                  ParticipantTimeLineDTO `json:"timeline"`
	HighestAchievedSeasonTier string                 `json:"highestAchievedSeasonTier"`
}

// DeltaDTO DeltaDTO
type DeltaDTO struct {
	Zero10 uint64 `json:"0-10"`
	One020 uint64 `json:"10-20"`
	Two030 uint64 `json:"20-30"`
}

// ParticipantTimeLineDTO ParticipantTimeLineDTO
type ParticipantTimeLineDTO struct {
	Lane                        string   `json:"lane"`
	Role                        string   `json:"role"`
	ParticipantID               uint     `json:"participantId"`
	CsDiffPerMinDeltas          DeltaDTO `json:"csDiffPerMinDeltas"`
	GoldPerMinDeltas            DeltaDTO `json:"goldPerMinDeltas"`
	XpDiffPerMinDeltas          DeltaDTO `json:"xpDiffPerMinDeltas"`
	CreepsPerMinDeltas          DeltaDTO `json:"creepsPerMinDeltas"`
	XpPerMinDeltas              DeltaDTO `json:"xpPerMinDeltas"`
	DamageTakenDiffPerMinDeltas DeltaDTO `json:"damageTakenDiffPerMinDeltas"`
	DamageTakenPerMinDeltas     DeltaDTO `json:"damageTakenPerMinDeltas"`
}

// ParticipantKDADTO ParticipantKDADTO
type ParticipantKDADTO struct {
	Kills               uint `json:"kills"`
	Deaths              uint `json:"deaths"`
	Assists             uint `json:"assists"`
	DoubleKills         uint `json:"doubleKills"`
	TripleKills         uint `json:"tripleKills"`
	QuadraKills         uint `json:"quadraKills"`
	PentaKills          uint `json:"pentaKills"`
	HexaKills           uint `json:"unrealKills"`
	LargestMultiKill    uint `json:"largestMultiKill"`
	LargestKillingSpree uint `json:"largestKillingSpree"`
	FirstBloodKill      bool `json:"firstBloodKill"`
	FirstBloodAssist    bool `json:"firstBloodAssist"`
	KillingSprees       uint `json:"killingSprees"`
}

// ParticipantGoldDTO ParticipantGoldDTO
type ParticipantGoldDTO struct {
	Earned uint `json:"goldEarned"`
	Spent  uint `json:"goldSpent"`
}

// ParticipantDamageDTO ParticipantDamageDTO
type ParticipantDamageDTO struct {
	Dealt            ParticipantDamageDealtDTO            `json:"ParticipantDamageDealtDTO"`
	Taken            ParticipantDamageTakenDTO            `json:"ParticipantDamageTakenDTO"`
	Mitigated        ParticipantDamageMitigatedDTO        `json:"ParticipantDamageMitigatedDTO"`
	DealtToChampions ParticipantDamageDealtToChampionsDTO `json:"ParticipantDamageDealtToChampionsDTO"`
}

// ParticipantDamageMitigatedDTO ParticipantDamageMitigatedDTO
type ParticipantDamageMitigatedDTO struct {
	Total uint `json:"damageSelfMitigated"`
}

// ParticipantDamageDealtDTO ParticipantDamageDealtDTO
type ParticipantDamageDealtDTO struct {
	Total                 uint `json:"totalDamageDealt"`
	Magic                 uint `json:"magicDamageDealt"`
	True                  uint `json:"trueDamageDealt"`
	Physical              uint `json:"physicalDamageDealt"`
	ToTurrets             uint `json:"damageDealtToTurrets"`
	ToObjectives          uint `json:"damageDealtToObjectives"`
	LargestCriticalStrike uint `json:"largestCriticalStrike"`
}

// ParticipantDamageDealtToChampionsDTO ParticipantDamageDealtToChampionsDTO
type ParticipantDamageDealtToChampionsDTO struct {
	Total    uint `json:"totalDamageDealtToChampions"`
	Magic    uint `json:"magicDamageDealtToChampions"`
	True     uint `json:"trueDamageDealtToChampions"`
	Physical uint `json:"physicalDamageDealtToChampions"`
}

// ParticipantDamageTakenDTO ParticipantDamageTakenDTO
type ParticipantDamageTakenDTO struct {
	Total    uint `json:"totalDamageTaken"`
	Magic    uint `json:"magicalDamageTaken"`
	True     uint `json:"trueDamageTaken"`
	Physical uint `json:"physicalDamageTaken"`
}

// ParticipantCSDTO ParticipantCSDTO
type ParticipantCSDTO struct {
	Lane   ParticipantLaneCSDTO   `json:"ParticipantLaneCSDTO"`
	Jungle ParticipantJungleCSDTO `json:"ParticipantJungleCSDTO"`
}

// ParticipantLaneCSDTO ParticipantLaneCSDTO
type ParticipantLaneCSDTO struct {
	Total uint `json:"totalMinionsKilled"`
}

// ParticipantJungleCSDTO ParticipantJungleCSDTO
type ParticipantJungleCSDTO struct {
	Total       uint `json:"neutralMinionsKilled"`
	OwnJungle   uint `json:"neutralMinionsKilledTeamJungle"`
	EnemyJungle uint `json:"neutralMinionsKilledEnemyJungle"`
}

// ParticipantObjectivesDTO ParticipantObjectivesDTO
type ParticipantObjectivesDTO struct {
	FirstTowerKill       bool `json:"firstTowerKill"`
	FirstInhibitorKill   bool `json:"firstInhibitorKill"`
	FirstTowerAssist     bool `json:"firstTowerAssist"`
	FirstInhibitorAssist bool `json:"firstInhibitorAssist"`
	TurretKills          uint `json:"turretKills"`
	InhibitorKills       uint `json:"inhibitorKills"`
}

// ParticipantScoreDTO ParticipantScoreDTO
type ParticipantScoreDTO struct {
	Rank      uint `json:"totalScoreRank"`
	Total     uint `json:"totalPlayerScore"`
	Combat    uint `json:"combatPlayerScore"`
	Objective uint `json:"objectivePlayerScore"`
	Vision    uint `json:"visionScore"`
}

// ParticipantHealingDTO ParticipantHealingDTO
type ParticipantHealingDTO struct {
	Allies uint `json:"totalUnitsHealed"`
	Self   uint `json:"totalHeal"`
}

// ParticipantCrowdControlDTO ParticipantCrowdControlDTO
type ParticipantCrowdControlDTO struct {
	Dealt uint `json:"totalTimeCrowdControlDealt"`
	Taken uint `json:"timeCCingOthers"`
}

// ParticipantWardsDTO ParticipantWardsDTO
type ParticipantWardsDTO struct {
	Placed uint `json:"wardsPlaced"`
	Bought uint `json:"visionWardsBoughtInGame"`
	Killed uint `json:"wardsKilled"`
}

// rawParticipantStatsDTO rawParticipantStatsDTO
type rawParticipantStatsDTO struct {
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

// ParticipantItemDTO ParticipantItemDTO
type ParticipantItemDTO struct {
	Slot   uint
	ItemID uint
}

// ParticipantTrinketDTO ParticipantTrinketDTO
type ParticipantTrinketDTO struct {
	TrinketID uint
}

// ParticipantStatsDTO ParticipantStatsDTO
type ParticipantStatsDTO struct {
	ParticipantID          uint
	ChampLevel             uint
	Win                    bool
	Gold                   ParticipantGoldDTO
	KDA                    ParticipantKDADTO
	Items                  []ParticipantItemDTO
	Trinket                ParticipantTrinketDTO
	RunePage               ParticipantRunePageDTO
	Damage                 ParticipantDamageDTO
	Objectives             ParticipantObjectivesDTO
	CS                     ParticipantCSDTO
	Healing                ParticipantHealingDTO
	CrowdControl           ParticipantCrowdControlDTO
	Score                  ParticipantScoreDTO
	Wards                  ParticipantWardsDTO
	LongestTimeSpentLiving uint
}

// ParticipantRunePageDTO ParticipantRunePageDTO
type ParticipantRunePageDTO struct {
	Primary   ParticipantPrimaryRunePathDTO
	Secondary ParticipantSecondaryRunePathDTO
}

// ParticipantPrimaryRunePathDTO ParticipantPrimaryRunePathDTO
type ParticipantPrimaryRunePathDTO struct {
	PathType     uint
	KeystoneRune ParticipantRuneDTO
	Rune1        ParticipantRuneDTO
	Rune2        ParticipantRuneDTO
	Rune3        ParticipantRuneDTO
}

// ParticipantSecondaryRunePathDTO ParticipantSecondaryRunePathDTO
type ParticipantSecondaryRunePathDTO struct {
	PathType uint
	Rune1    ParticipantRuneDTO
	Rune2    ParticipantRuneDTO
}

// ParticipantRuneDTO ParticipantRuneDTO
type ParticipantRuneDTO struct {
	RuneID    uint
	RuneStat1 uint
	RuneStat2 uint
	RuneStat3 uint
}

// NewMatch NewMatch
func NewMatch(data string) MatchDTO {
	var rawMatch rawMatchDTO
	var participants []ParticipantDTO
	var match MatchDTO

	json.Unmarshal([]byte(data), &rawMatch)

	match.SeasonID = rawMatch.SeasonID
	match.QueueID = rawMatch.QueueID
	match.MatchID = rawMatch.MatchID
	match.ParticipantIdentities = rawMatch.ParticipantIdentities
	match.GameVersion = rawMatch.GameVersion
	match.PlatformID = rawMatch.PlatformID
	match.GameMode = rawMatch.GameMode
	match.MapID = rawMatch.MapID
	match.GameType = rawMatch.GameType
	match.Teams = rawMatch.Teams
	match.GameDuration = rawMatch.GameDuration
	match.Participants = participants
	match.GameCreation = rawMatch.GameCreation

	for _, rawParticipant := range rawMatch.Participants {
		var participant ParticipantDTO

		participant.ParticipantID = rawParticipant.ParticipantID
		participant.TeamID = rawParticipant.TeamID
		participant.ChampionID = rawParticipant.ChampionID
		participant.PrimarySpellID = rawParticipant.PrimarySpellID
		participant.SecondarySpellID = rawParticipant.SecondarySpellID
		participant.TimeLine = rawParticipant.TimeLine
		participant.HighestAchievedSeasonTier = rawParticipant.HighestAchievedSeasonTier

		var rawStats = rawParticipant.Stats

		participant.Stats.ParticipantID = rawStats.ParticipantID
		participant.Stats.ChampLevel = rawStats.ChampLevel
		participant.Stats.Win = rawStats.Win

		participant.Stats.Gold.Earned = rawStats.GoldEarned
		participant.Stats.Gold.Spent = rawStats.GoldSpent

		participant.Stats.KDA.Kills = rawStats.Kills
		participant.Stats.KDA.Deaths = rawStats.Deaths
		participant.Stats.KDA.Assists = rawStats.Assists
		participant.Stats.KDA.DoubleKills = rawStats.DoubleKills
		participant.Stats.KDA.TripleKills = rawStats.TripleKills
		participant.Stats.KDA.QuadraKills = rawStats.QuadraKills
		participant.Stats.KDA.PentaKills = rawStats.PentaKills
		participant.Stats.KDA.HexaKills = rawStats.UnrealKills
		participant.Stats.KDA.LargestMultiKill = rawStats.LargestMultiKill
		participant.Stats.KDA.LargestKillingSpree = rawStats.LargestKillingSpree
		participant.Stats.KDA.FirstBloodKill = rawStats.FirstBloodKill
		participant.Stats.KDA.FirstBloodAssist = rawStats.FirstBloodAssist
		participant.Stats.KDA.KillingSprees = rawStats.KillingSprees

		participant.Stats.Items = append(participant.Stats.Items, ParticipantItemDTO{Slot: 1, ItemID: rawStats.Item0})
		participant.Stats.Items = append(participant.Stats.Items, ParticipantItemDTO{Slot: 2, ItemID: rawStats.Item1})
		participant.Stats.Items = append(participant.Stats.Items, ParticipantItemDTO{Slot: 3, ItemID: rawStats.Item2})
		participant.Stats.Items = append(participant.Stats.Items, ParticipantItemDTO{Slot: 4, ItemID: rawStats.Item3})
		participant.Stats.Items = append(participant.Stats.Items, ParticipantItemDTO{Slot: 5, ItemID: rawStats.Item4})
		participant.Stats.Items = append(participant.Stats.Items, ParticipantItemDTO{Slot: 6, ItemID: rawStats.Item5})

		participant.Stats.Trinket.TrinketID = rawStats.Item6

		participant.Stats.RunePage.Primary.KeystoneRune.RuneID = rawStats.Perk0
		participant.Stats.RunePage.Primary.KeystoneRune.RuneStat1 = rawStats.Perk0Var1
		participant.Stats.RunePage.Primary.KeystoneRune.RuneStat2 = rawStats.Perk0Var2
		participant.Stats.RunePage.Primary.KeystoneRune.RuneStat3 = rawStats.Perk0Var3

		participant.Stats.RunePage.Primary.Rune1.RuneID = rawStats.Perk1
		participant.Stats.RunePage.Primary.Rune1.RuneStat1 = rawStats.Perk1Var1
		participant.Stats.RunePage.Primary.Rune1.RuneStat2 = rawStats.Perk1Var2
		participant.Stats.RunePage.Primary.Rune1.RuneStat3 = rawStats.Perk1Var3

		participant.Stats.RunePage.Primary.Rune2.RuneID = rawStats.Perk2
		participant.Stats.RunePage.Primary.Rune2.RuneStat1 = rawStats.Perk2Var1
		participant.Stats.RunePage.Primary.Rune2.RuneStat2 = rawStats.Perk2Var2
		participant.Stats.RunePage.Primary.Rune2.RuneStat3 = rawStats.Perk2Var3

		participant.Stats.RunePage.Primary.Rune3.RuneID = rawStats.Perk3
		participant.Stats.RunePage.Primary.Rune3.RuneStat1 = rawStats.Perk3Var1
		participant.Stats.RunePage.Primary.Rune3.RuneStat2 = rawStats.Perk3Var2
		participant.Stats.RunePage.Primary.Rune3.RuneStat3 = rawStats.Perk3Var3

		participant.Stats.RunePage.Secondary.Rune1.RuneID = rawStats.Perk4
		participant.Stats.RunePage.Secondary.Rune1.RuneStat1 = rawStats.Perk4Var1
		participant.Stats.RunePage.Secondary.Rune1.RuneStat2 = rawStats.Perk4Var2
		participant.Stats.RunePage.Secondary.Rune1.RuneStat3 = rawStats.Perk4Var3

		participant.Stats.RunePage.Secondary.Rune2.RuneID = rawStats.Perk5
		participant.Stats.RunePage.Secondary.Rune2.RuneStat1 = rawStats.Perk5Var1
		participant.Stats.RunePage.Secondary.Rune2.RuneStat2 = rawStats.Perk5Var2
		participant.Stats.RunePage.Secondary.Rune2.RuneStat3 = rawStats.Perk5Var3

		participant.Stats.RunePage.Primary.PathType = rawStats.PerkPrimaryStyle
		participant.Stats.RunePage.Secondary.PathType = rawStats.PerkSubStyle

		participant.Stats.Damage.Mitigated.Total = rawStats.DamageSelfMitigated

		participant.Stats.Damage.Dealt.Total = rawStats.TotalDamageDealt
		participant.Stats.Damage.Dealt.Magic = rawStats.MagicDamageDealt
		participant.Stats.Damage.Dealt.True = rawStats.TrueDamageDealt
		participant.Stats.Damage.Dealt.Physical = rawStats.PhysicalDamageDealt
		participant.Stats.Damage.Dealt.LargestCriticalStrike = rawStats.LargestCriticalStrike

		participant.Stats.Damage.Taken.Total = rawStats.TotalDamageTaken
		participant.Stats.Damage.Taken.Magic = rawStats.MagicDamageTaken
		participant.Stats.Damage.Taken.True = rawStats.TrueDamageTaken
		participant.Stats.Damage.Taken.Physical = rawStats.PhysicalDamageTaken

		participant.Stats.Damage.DealtToChampions.Total = rawStats.TotalDamageDealtToChampions
		participant.Stats.Damage.DealtToChampions.Magic = rawStats.MagicDamageDealtToChampions
		participant.Stats.Damage.DealtToChampions.True = rawStats.TrueDamageDealtToChampions
		participant.Stats.Damage.DealtToChampions.Physical = rawStats.PhysicalDamageDealtToChampions

		participant.Stats.Damage.Dealt.ToTurrets = rawStats.DamageDealtToTurrets
		participant.Stats.Damage.Dealt.ToObjectives = rawStats.DamageDealtToObjectives

		participant.Stats.Objectives.TurretKills = rawStats.TurretKills
		participant.Stats.Objectives.InhibitorKills = rawStats.InhibitorKills
		participant.Stats.Objectives.FirstTowerKill = rawStats.FirstTowerKill
		participant.Stats.Objectives.FirstInhibitorKill = rawStats.FirstInhibitorKill
		participant.Stats.Objectives.FirstTowerAssist = rawStats.FirstTowerAssist
		participant.Stats.Objectives.FirstInhibitorAssist = rawStats.FirstInhibitorAssist

		participant.Stats.CS.Lane.Total = rawStats.TotalMinionsKilled
		participant.Stats.CS.Jungle.Total = rawStats.NeutralMinionsKilled
		participant.Stats.CS.Jungle.OwnJungle = rawStats.NeutralMinionsKilledTeamJungle
		participant.Stats.CS.Jungle.EnemyJungle = rawStats.NeutralMinionsKilledEnemyJungle

		participant.Stats.Healing.Self = rawStats.TotalHeal
		participant.Stats.Healing.Allies = rawStats.TotalUnitsHealed

		participant.Stats.Score.Rank = rawStats.TotalScoreRank
		participant.Stats.Score.Vision = rawStats.VisionScore
		participant.Stats.Score.Total = rawStats.TotalPlayerScore
		participant.Stats.Score.Combat = rawStats.CombatPlayerScore
		participant.Stats.Score.Objective = rawStats.ObjectivePlayerScore

		participant.Stats.CrowdControl.Dealt = rawStats.TotalTimeCrowdControlDealt
		participant.Stats.CrowdControl.Taken = rawStats.TimeCCingOthers

		participant.Stats.Wards.Killed = rawStats.WardsKilled
		participant.Stats.Wards.Placed = rawStats.WardsPlaced
		participant.Stats.Wards.Bought = rawStats.VisionWardsBoughtInGame

		participant.Stats.LongestTimeSpentLiving = rawStats.LongestTimeSpentLiving

		match.Participants = append(match.Participants, participant)
	}

	return match
}
