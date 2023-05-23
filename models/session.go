package models

/*
	{
	  "expiry": 0,
	  "ks": "",
	  "partnerId": 0,
	  "privileges": "",
	  "sessionType": 0,
	  "userId": ""
	}
*/
type KalturaSessionInfo struct {
	Expiry      int    `json:"expiry"`
	Ks          string `json:"ks"`
	PartnerId   int    `json:"partnerId"`
	Privileges  string `json:"privileges"`
	SessionType int    `json:"sessionType"`
	UserId      string `json:"userId"`
}

/*
	{
	  "ks": "",
	  "objectType": "KalturaStartWidgetSessionResponse",
	  "partnerId": 0,
	  "userId": ""
	}
*/
type KalturaSessionResponse struct {
	Ks         string `json:"ks"`
	ObjectType string `json:"objectType"`
	PartnerId  int    `json:"partnerId"`
	UserId     string `json:"userId"`
}
