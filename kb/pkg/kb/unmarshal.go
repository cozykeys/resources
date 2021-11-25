package kb

import (
	"encoding/json"
)

// TODO: Explanation; this is the only exported method related to unmarshaling.
// User of the package is expected to unmarshal a la:
//
//     kb := &kbutil.Keyboard{}
//     if err := json.Unmarshal(bytes, kb); err != nil { ... }
//
// Internally, all unmarshaling is done by handle from generic maps/slices. The
// primary reason for this is that the schema is somewhat complicated and
// requires overriding some of the unmarshal functions anyways to support:
// - Sum types
// - Required fields
// As an added benefit, we don't have to unmarshal anything twice so it ends up
// being a bit more performant, not that that matters much.
func (k *Keyboard) UnmarshalJSON(bytes []byte) error {
	kbMap := map[string]interface{}{}
	if err := json.Unmarshal(bytes, &kbMap); err != nil {
		return err
	}

	if err := k.fromMap(kbMap); err != nil {
		return err
	}

	return nil
}
