package dto

// TimeLine TimeLine
type TimeLine struct {
	Frames        []Frame `json:"frames"`
	FrameInterval int     `json:"frameInterval"`
}

// Position Position
type Position struct {
	Y int `json:"y"`
	X int `json:"x"`
}

// Frame Frame
type Frame struct {
	Timestamp         int               `json:"timestamp"`
	ParticipantFrames ParticipantFrames `json:"participantFrames"`
	Events            []interface{}     `json:"events"`
}

// ParticipantFrames ParticipantFrames
type ParticipantFrames struct {
	Participant1  ParticipantFrame `json:"1"`
	Participant2  ParticipantFrame `json:"2"`
	Participant3  ParticipantFrame `json:"3"`
	Participant4  ParticipantFrame `json:"4"`
	Participant5  ParticipantFrame `json:"5"`
	Participant6  ParticipantFrame `json:"6"`
	Participant7  ParticipantFrame `json:"7"`
	Participant8  ParticipantFrame `json:"8"`
	Participant9  ParticipantFrame `json:"9"`
	Participant10 ParticipantFrame `json:"10"`
}

// ParticipantFrame ParticipantFrame
type ParticipantFrame struct {
	TotalGold           uint     `json:"totalGold"`
	TeamScore           uint     `json:"teamScore"`
	ParticipantID       uint     `json:"participantId"`
	Level               uint     `json:"level"`
	CurrentGold         uint     `json:"currentGold"`
	MinionsKilled       uint     `json:"minionsKilled"`
	DominionScore       uint     `json:"dominionScore"`
	Position            Position `json:"position"`
	Xp                  uint     `json:"xp"`
	JungleMinionsKilled uint     `json:"jungleMinionsKilled"`
}
