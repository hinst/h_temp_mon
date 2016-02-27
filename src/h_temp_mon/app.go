package h_temp_mon

type TApp struct {
	TempReader TTempReader
}

func NewApp() *TApp {
	return &TApp{}
}

func (this *TApp) Run() {
}
