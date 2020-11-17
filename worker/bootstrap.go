package worker

func StartWorker() {
	NewUserRegisterWorker().Start()

	select {}
}
