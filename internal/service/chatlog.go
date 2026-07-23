package service

func (svc *Service) ServiceStoreToChatLogDB(name string, text string) error {
	return svc.sqlDB.StoreToChatLogDB(name, text)
}

func (svc *Service) ServiceReadFromChatLogDB(val int64) (string, error) {
	return svc.sqlDB.ReadFromChatLogDB(val)
}
