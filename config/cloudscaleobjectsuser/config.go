package cloudscaleobjectsuser

import (
	"fmt"

	"github.com/crossplane/terrajet/pkg/config"
)

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("cloudscale_objects_user", func(r *config.Resource) {
		r.ExternalName = config.IdentifierFromProvider
		r.Sensitive.AdditionalConnectionDetailsFn = func(attr map[string]interface{}) (map[string][]byte, error) {
			return extractMapFromInterfaceSlice(attr["keys"]), nil
		}
	})
}

func extractMapFromInterfaceSlice(t interface{}) map[string][]byte {
	listOfMaps := map[string][]byte{}

	// test if this is a slice
	if interfaceValue, ok := t.([]interface{}); ok {
		// loop over the slice
		for index, potentialMap := range interfaceValue {
			// check if entry is a map
			if realMap, ok := potentialMap.(map[string]interface{}); ok {
				// convert map[string]interface{} to map[string]string
				for key, value := range realMap {
					// check if value is a string
					if valueString, ok := value.(string); ok {
						listOfMaps[fmt.Sprintf("%s_%d", key, index)] = []byte(valueString)
					}
				}
			}
		}
	}

	return listOfMaps
}
