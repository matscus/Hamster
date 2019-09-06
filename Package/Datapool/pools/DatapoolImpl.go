package pools

import "github.com/matscus/Hamster/Package/Datapool/subset"

//Datapool - default struct for  datapool, not contains users data
type Datapool struct{}

//New - return new dapatool interface
func (p Datapool) New() subset.Datapool {
	var dp subset.Datapool
	dp = Datapool{}
	return dp
}
