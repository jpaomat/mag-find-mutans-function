package credentialsdb

type SecretData struct {
	Hostname string `json:"host"`
	UserName string `json:"username"`
	Password string `json:"password"`
	Database string `json:"dbInstanceIdentifier"`
	Port     int    `json:"port"`
}
