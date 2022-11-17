package mail

func SendEmail() func(Mail) error {
	return func(ml Mail) error {
		return nil
	}
}
