package models

/*
{
  "adminTags": "",
  "allowedPartnerIds": "",
  "allowedPartnerPackages": "",
  "city": "",
  "country": "",
  "createdAt": 0,
  "deletedAt": 0,
  "description": "",
  "email": "",
  "fullName": "",
  "id": "",
  "indexedPartnerDataInt": 0,
  "indexedPartnerDataString": "",
  "language": "aa",
  "lastLoginTime": 0,
  "objectType": "KalturaUser",
  "partnerData": "",
  "partnerId": 0,
  "screenName": "",
  "state": "",
  "status": 0,
  "statusUpdatedAt": 0,
  "storageSize": 0,
  "tags": "",
  "thumbnailUrl": "",
  "updatedAt": 0,
  "userMode": 0,
  "zip": "",
  "attendanceInfo": "",
  "company": "",
  "dateOfBirth": 0,
  "encryptedSeed": "",
  "externalId": "",
  "firstName": "",
  "gender": 0,
  "isAccountOwner": true,
  "isAdmin": true,
  "isGuest": true,
  "isSsoExcluded": true,
  "ksPrivileges": "",
  "lastName": "",
  "loginEnabled": true,
  "password": "",
  "registrationInfo": "",
  "roleIds": "",
  "roleNames": "",
  "title": "",
  "type": 0
}
*/
type KalturaUser struct {
	AdminTags string  `json:"adminTags"`
	AllowedPartnerIds string  `json:"allowedPartnerIds"`
	Email string  `json:"email"`
}