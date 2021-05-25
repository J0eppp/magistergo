package magistergo

type Teacher struct {
	ID int64 `json:"Id"`
	Name string `json:"Naam"`
	TeacherCode string `json:"Docentcode"`
}

type Subject struct {
	ID int64 `json:"Id"`
	Name string `json:"Naam"`
}

type Classroom struct {
	Name string `json:"Naam"`
}

type Appointment struct {
	ID int64 `json:"Id"`
	// Skip "Links", is that interesting to know?
	Start string `json:"Start"`
	End string `json:"Einde"`
	StartHour int8 `json:"LesuurVan"`
	EndHour int8 `json:"LesuurTotMet"`
	TakesWholeDay bool `json:"DuurtHeleDag"`
	Description string `json:"Omschrijving"`
	Location string `json:"Locatie"`
	Status int64 `json:"Status"`
	Type int64 `json:"Type"`
	IsOnline bool `json:"IsOnlineDeelname"`
	DisplayType int64 `json:"WeergaveType"`
	Content string `json:"Inhoud"`
	InfoType int64 `json:"InfoType"`
	Notes string `json:"Aantekeningen"`
	Finished bool `json:"Afgerond"`
	RepeatStatus int64 `json:"HerhaalStatus"`
	Repeat string `json:"Herhaling"`
	Subjects []Subject `json:"Vakken"`
	Teachers []Teacher `json:"Docenten"`
	Classrooms []Classroom `json:"Lokalen"`
	Groups string `json:"Groepen"`
	AssignmentID int64 `json:"OpdrachtId"`
	HasAttachements bool `json:"HeeftBijlagen"`
	Attachments string `json:"Bijlagen"`
}