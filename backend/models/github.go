package models

import (
	"gorm.io/gorm"
)

type GithubAppConnection struct {
	gorm.Model
	GithubId                  int64 // app id
	ClientID                  string
	ClientSecretEncrypted     string
	WebhookSecretEncrypted    string
	PrivateKeyEncrypted       string
	PrivateKeyBase64Encrypted string
	Org                       string
	Name                      string
	GithubAppUrl              string
	OrganisationID            uint
	Organisation              Organisation
}

type GithubAppInstallStatus int

const (
	GithubAppInstallActive  GithubAppInstallStatus = 1
	GithubAppInstallDeleted GithubAppInstallStatus = 2
)

type GithubAppInstallation struct {
	gorm.Model
	GithubInstallationId int64
	GithubAppId          int64
	AccountId            int
	Login                string
	Repo                 string
	Status               GithubAppInstallStatus
}

type GithubAppInstallationLinkStatus int8

const (
	GithubAppInstallationLinkActive   GithubAppInstallationLinkStatus = 1
	GithubAppInstallationLinkInactive GithubAppInstallationLinkStatus = 2
)

// GithubAppInstallationLink links GitHub App installation Id to Digger's organisation Id
type GithubAppInstallationLink struct {
	gorm.Model
	GithubInstallationId int64 `gorm:"index:idx_github_installation_org"`
	OrganisationId       uint  `gorm:"index:idx_github_installation_org"`
	Organisation         *Organisation
	Status               GithubAppInstallationLinkStatus
}
