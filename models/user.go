package models

type KalturaUser struct {
	AdminTags                string `json:"adminTags"`
	AllowedPartnerIds        string `json:"allowedPartnerIds"`
	AllowedPartnerPackages   string `json:"allowedPartnerPackages"`
	City                     string `json:"city"`
	Country                  string `json:"country"`
	CreatedAt                int    `json:"createdAt"`
	DeletedAt                int    `json:"deletedAt"`
	Description              string `json:"description"`
	Email                    string `json:"email"`
	FullName                 string `json:"fullName"`
	Id                       string `json:"id"`
	IndexedPartnerDataInt    int    `json:"indexedPartnerDataInt"`
	IndexedPartnerDataString string `json:"indexedPartnerDataString"`
	Language                 string `json:"language"`
	LastLoginTime            int    `json:"lastLoginTime"`
	ObjectType               string `json:"objectType"`
	PartnerData              string `json:"partnerData"`
	PartnerId                int    `json:"partnerId"`
	ScreenName               string `json:"screenName"`
	State                    string `json:"state"`
	Status                   int    `json:"status"`
	StatusUpdatedAt          int    `json:"statusUpdatedAt"`
	StorageSize              int    `json:"storageSize"`
	Tags                     string `json:"tags"`
	ThumbnailUrl             string `json:"thumbnailUrl"`
	UpdatedAt                int    `json:"updatedAt"`
	UserMode                 int    `json:"userMode"`
	Zip                      string `json:"zip"`
	AttendanceInfo           string `json:"attendanceInfo"`
	Company                  string `json:"company"`
	DateOfBirth              int    `json:"dateOfBirth"`
	EncryptedSeed            string `json:"encryptedSeed"`
	ExternalId               string `json:"externalId"`
	FirstName                string `json:"firstName"`
}
