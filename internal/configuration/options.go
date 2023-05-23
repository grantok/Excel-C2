package configuration

type options struct {
	tenantId     string
	clientId     string
	clientSecret string
	driveId      string
	sheetId      string
	debug        bool
}

var command options

func SetOptions(
	tenantId string,
	clientId string,
	clientSecret string,
	driveId string,
	sheetId string,
	debug bool) {
	command.tenantId = tenantId
	command.clientId = clientId
	command.clientSecret = clientSecret
	command.driveId = driveId
	command.sheetId = sheetId
	command.debug = debug
}

func GetOptionsTenantId() string {
	return command.tenantId
}

func GetOptionsClientId() string {
	return command.clientId
}

func GetOptionsClientSecret() string {
	return command.clientSecret
}

func GetOptionsDriveId() string {
	return command.driveId
}

func GetOptionsSheetId() string {
	return command.sheetId
}

func GetOptionsDebug() bool {
	return command.debug
}
