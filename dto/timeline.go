package dto

import "encoding/json"

// TimeLineDTO TimeLineDTO
type TimeLineDTO struct {
	Frames        []TimeLineFrameDTO
	FrameInterval int
}

type rawTimeLineDTO struct {
	Frames        []rawTimeLineFrameDTO `json:"frames"`
	FrameInterval int                   `json:"frameInterval"`
}

// PositionDTO PositionDTO
type PositionDTO struct {
	Y int `json:"y"`
	X int `json:"x"`
}

// TimeLineFrameDTO TimeLineFrameDTO
type TimeLineFrameDTO struct {
	Timestamp         int
	ParticipantFrames []TimeLineParticipantFrameDTO
	Events            []TimeLineMatchEventDTO
}

type rawTimeLineFrameDTO struct {
	Timestamp         int                             `json:"timestamp"`
	ParticipantFrames rawTimeLineParticipantFramesDTO `json:"participantFrames"`
	Events            []TimeLineMatchEventDTO         `json:"events"`
}

type rawTimeLineParticipantFramesDTO struct {
	Participant1  TimeLineParticipantFrameDTO `json:"1"`
	Participant2  TimeLineParticipantFrameDTO `json:"2"`
	Participant3  TimeLineParticipantFrameDTO `json:"3"`
	Participant4  TimeLineParticipantFrameDTO `json:"4"`
	Participant5  TimeLineParticipantFrameDTO `json:"5"`
	Participant6  TimeLineParticipantFrameDTO `json:"6"`
	Participant7  TimeLineParticipantFrameDTO `json:"7"`
	Participant8  TimeLineParticipantFrameDTO `json:"8"`
	Participant9  TimeLineParticipantFrameDTO `json:"9"`
	Participant10 TimeLineParticipantFrameDTO `json:"10"`
}

// TimeLineParticipantFrameDTO TimeLineParticipantFrameDTO
type TimeLineParticipantFrameDTO struct {
	TotalGold           uint        `json:"totalGold"`
	TeamScore           uint        `json:"teamScore"`
	ParticipantID       uint        `json:"participantId"`
	Level               uint        `json:"level"`
	CurrentGold         uint        `json:"currentGold"`
	MinionsKilled       uint        `json:"minionsKilled"`
	DominionScore       uint        `json:"dominionScore"`
	Position            PositionDTO `json:"position"`
	Xp                  uint        `json:"xp"`
	JungleMinionsKilled uint        `json:"jungleMinionsKilled"`
}

// TimeLineMatchEventDTO TimeLineMatchEventDTO
type TimeLineMatchEventDTO struct {
	EventType               string      `json:"eventType"`
	TowerType               string      `json:"towerType"`
	TeamID                  uint        `json:"teamId"`
	AscendedType            string      `json:"ascendedType"`
	KillerID                uint        `json:"killerId"`
	LevelUpType             string      `json:"levelUpType"`
	PointCaptured           string      `json:"pointCaptured"`
	AssistingParticipantIds []uint      `json:"assistingParticipantIds"`
	WardType                string      `json:"wardType"`
	MonsterType             string      `json:"monsterType"`
	Type                    string      `json:"type"`
	SkillSlot               uint        `json:"skillSlot"`
	VictimID                uint        `json:"victimId"`
	Timestamp               uint64      `json:"timestamp"`
	AfterID                 uint        `json:"afterId"`
	MonsterSubType          string      `json:"monsterSubType"`
	LaneType                string      `json:"laneType"`
	ItemID                  uint        `json:"itemId"`
	ParticipantID           uint        `json:"participantId"`
	BuildingType            string      `json:"buildingType"`
	CreatorID               uint        `json:"creatorId"`
	Position                PositionDTO `json:"position"`
	BeforeID                uint        `json:"beforeId"`
}

// NewTimeLine NewTimeLine
func NewTimeLine(data string) TimeLineDTO {
	var rawTimeLine rawTimeLineDTO
	var timeLine TimeLineDTO

	json.Unmarshal([]byte(data), &rawTimeLine)

	timeLine.FrameInterval = rawTimeLine.FrameInterval

	for _, rawFrame := range rawTimeLine.Frames {
		var frame TimeLineFrameDTO

		frame.Events = rawFrame.Events
		frame.Timestamp = rawFrame.Timestamp
		frame.ParticipantFrames = append(frame.ParticipantFrames, rawFrame.ParticipantFrames.Participant1)
		frame.ParticipantFrames = append(frame.ParticipantFrames, rawFrame.ParticipantFrames.Participant2)
		frame.ParticipantFrames = append(frame.ParticipantFrames, rawFrame.ParticipantFrames.Participant3)
		frame.ParticipantFrames = append(frame.ParticipantFrames, rawFrame.ParticipantFrames.Participant4)
		frame.ParticipantFrames = append(frame.ParticipantFrames, rawFrame.ParticipantFrames.Participant5)
		frame.ParticipantFrames = append(frame.ParticipantFrames, rawFrame.ParticipantFrames.Participant6)
		frame.ParticipantFrames = append(frame.ParticipantFrames, rawFrame.ParticipantFrames.Participant7)
		frame.ParticipantFrames = append(frame.ParticipantFrames, rawFrame.ParticipantFrames.Participant8)
		frame.ParticipantFrames = append(frame.ParticipantFrames, rawFrame.ParticipantFrames.Participant9)
		frame.ParticipantFrames = append(frame.ParticipantFrames, rawFrame.ParticipantFrames.Participant10)

		timeLine.Frames = append(timeLine.Frames, frame)
	}

	return timeLine
}
