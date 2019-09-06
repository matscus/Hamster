package subset

import "github.com/matscus/Hamster/Package/Datapool/structs"

type Datapool interface {
	GetDefaultDatapool(lenpool int, c chan<- structs.DefaultDatapool) error
}
