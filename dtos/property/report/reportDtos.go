package report

type PropertyReportCreationDto struct {
	PropertyId int    `json:"propertyId"`
	ManagerId  int    `json:"managerId"`
	Report     string `json:"report"`
}

type PropertyReportUpdateAndResponseDto struct {
	Id         int    `json:"id"`
	PropertyId int    `json:"propertyId"`
	ManagerId  int    `json:"managerId"`
	Report     string `json:"report"`
}
