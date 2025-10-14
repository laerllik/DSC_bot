package telegramclientapi

// type TelegramClientConfigError struct {
// 	message string
// }

// func (e *TelegramClientConfigError) Error() string {
// 	return fmt.Sprintf("Went wrong while telegram config configure. %s", e.message)
// }

type TelegramClientNotConfiguredError struct{}

func (e *TelegramClientNotConfiguredError) Error() string {
	return "TelegramClientNotConfiguredError. Telegram client didn`t configure"
}
