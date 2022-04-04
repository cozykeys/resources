package unmarshal

import (
	"kb/pkg/models"

	"github.com/beevik/etree"
)

/*
   "Stack": {
       "Attributes": [
           "Name",
           "Rotation",
           "Orientation",
           "XOffset",
           "YOffset"
       ],
       "Children": [
           "Children"
       ]
   },
*/

func unmarshalStack(e *etree.Element) (*models.Stack, error) {
	// TODO: Implement this
	return nil, &unimplementedElementError{
		elementPath: getElementPath(e),
	}
}
