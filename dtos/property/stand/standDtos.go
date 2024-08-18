package stand

type StandCreationDTO struct {
	ManagerId          int                                  `json:"userId"`
	Price              int                                  `json:"price"`
	SizeNumber         int                                  `json:"sizeNumber"`
	AreaHasElectricity bool                                 `json:"areaHasElectricity"`
	IsServiced         bool                                 `json:"isServiced"`
	IsNegotiable       bool                                 `json:"isNegotiable"`
	Status             string                               `json:"status"`
	Level              string                               `json:"level"`
	SizeDimensions     string                               `json:"sizeDimensions"`
	Type               string                               `json:"type"`
	ProfilePicture     ManagerProfilePictureCreationDTO     `json:"profilePicture"`
	Contacts           []managerModels.ManagerContactNumber `json:"contacts"`
}
