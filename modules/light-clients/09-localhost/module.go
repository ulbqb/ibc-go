package localhost

import (
	"github.com/Finschia/ibc-go/v4/modules/light-clients/09-localhost/types"
)

// Name returns the IBC client name
func Name() string {
	return types.SubModuleName
}
