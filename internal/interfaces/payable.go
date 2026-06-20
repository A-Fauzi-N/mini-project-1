package interfaces


// mendefinisikan kontrak kerja
// untuk entitas yang memiliki gaji
type Payable interface{
	CalculatedPay() float64
}