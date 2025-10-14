package telegramclientapi

type telegramClient struct {
	configured bool
}

func (tc *telegramClient) Configure() {
	tc.configured = true
}

func (tc *telegramClient) GetChats() error {
	if !tc.configured {
		return &TelegramClientNotConfiguredError{}
	}

	return nil
}

var TelegramClientApi = telegramClient{}
