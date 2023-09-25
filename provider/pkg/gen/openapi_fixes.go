package gen

import (
	"strings"

	"github.com/getkin/kin-openapi/openapi3"
)

func FixOpenAPIDoc(openAPIDoc *openapi3.T) {
	regionSchema, ok := openAPIDoc.Components.Schemas["region"]
	if !ok {
		panic("Expected to find a schema for region type")
	}

	regionSchema.Value.Properties["features"].Value.Type = "array"
	regionSchema.Value.Properties["sizes"].Value.Type = "array"

	// Add a title for droplet_action type.
	addTitleForType(openAPIDoc, "droplet_action", "DropletAction")

	// Add a title for image_action_base type.
	addTitleForType(openAPIDoc, "image_action_base", "ImageAction")

	for _, ty := range openAPIDoc.Components.Schemas {
		fixObjectPropertyType(ty)
	}

	// for _, ty := range openAPIDoc.Components.Responses {
	// 	fixResponseObjectPropertyType(ty)
	// }

	fixDatabaseConfigType(openAPIDoc)
	fixFloatingIpType(openAPIDoc)
	fixReservedIpActionUnassignType(openAPIDoc)
	fixReservedIpType(openAPIDoc)
	fixPageLinksType(openAPIDoc)
}

func addTitleForType(openAPIDoc *openapi3.T, typeName, title string) {
	theType, ok := openAPIDoc.Components.Schemas[typeName]
	if !ok {
		panic("Expected to find " + typeName + " type")
	}
	theType.Value.Title = title
}

func fixReservedIpActionUnassignType(openAPIDoc *openapi3.T) {
	schemaRef, ok := openAPIDoc.Components.Schemas["reserved_ip_action_unassign"]
	if !ok {
		panic("Expected to find reserved_ip_action_unassign type")
	}

	intSchema := openapi3.NewIntegerSchema()
	intSchema.Description = "The ID of the Droplet that the reserved IP will be unassigned from."
	schemaRef.Value.AllOf[1].Value.Properties = map[string]*openapi3.SchemaRef{
		"droplet_id": intSchema.NewRef(),
	}
	// Overwrite the list of required properties for this type,
	// so that only `droplet_id` is required.
	schemaRef.Value.Required = []string{"droplet_id"}
}

func fixReservedIpType(openAPIDoc *openapi3.T) {
	schemaRef, ok := openAPIDoc.Components.Schemas["reserved_ip"]
	if !ok {
		panic("Expected to find reserved_ip type")
	}

	// The droplet property of this type can be null if the floating IP
	// is not attached to any droplet. The way this is represented in the
	// spec causes validation issues. So let's simplify it.
	schemaRef.Value.Properties["droplet"].Ref = "#/components/schemas/droplet"
	schemaRef.Value.Properties["droplet"].Value.AnyOf = nil
	schemaRef.Value.Properties["droplet"].Value.Example = nil
}

func fixObjectPropertyType(objectType *openapi3.SchemaRef) {
	// Add `type: object` for any type that does not have the `type` property
	// of the schema set.
	if objectType.Value.Type != "" {
		return
	}

	if len(objectType.Value.Properties) > 0 {
		objectType.Value.Type = "object"
	}

	// if len(objectType.Value.AllOf) > 0 {
	// 	for _, allOfTy := range objectType.Value.AllOf {
	// 		fixObjectPropertyType(allOfTy)
	// 	}
	// }
}

func fixResponseObjectPropertyType(responseRef *openapi3.ResponseRef) {
	responseContent := responseRef.Value.Content.Get("application/json")
	if responseContent == nil {
		return
	}

	// Add `type: object` for any type that does not have the `type` property
	// of the schema set.
	if responseContent.Schema.Value.Type != "" {
		return
	}

	// if len(responseContent.Schema.Value.Properties) > 0 {
	// 	responseContent.Schema.Value.Type = "object"
	// }

	// if len(responseContent.Schema.Value.AllOf) > 0 {
	// 	for _, allOfTy := range responseContent.Schema.Value.AllOf {
	// 		fixObjectPropertyType(allOfTy)
	// 	}
	// }
}

func fixFloatingIpType(openAPIDoc *openapi3.T) {
	floatingIp, ok := openAPIDoc.Components.Schemas["floating_ip"]
	if !ok {
		panic("Expected to find floating_ip type")
	}

	// The droplet property of this type can be null if the floating IP
	// is not attached to any droplet. The way this is represented in the
	// spec causes validation issues. So let's simplify it.
	floatingIp.Value.Properties["droplet"].Ref = "#/components/schemas/droplet"
	floatingIp.Value.Properties["droplet"].Value.AnyOf = nil
	floatingIp.Value.Properties["droplet"].Value.Example = nil
}

// fixDatabaseConfigType fixes the database_config type to use
// a OneOf schema (with a discriminator) instead of an AnyOf.
func fixDatabaseConfigType(openAPIDoc *openapi3.T) {
	databaseConfig, ok := openAPIDoc.Components.Schemas["database_config"]
	if !ok {
		panic("Expected to find database_config type")
	}

	configProp := databaseConfig.Value.Properties["config"]
	types := make([]*openapi3.SchemaRef, 0, len(configProp.Value.AnyOf))
	for _, ty := range configProp.Value.AnyOf {
		types = append(types, ty)
		// Add a property to each of the sub-types.
		stringSchema := openapi3.NewStringSchema()
		ref := ty.Ref
		refParts := strings.SplitAfterN(ref, "/", 4)
		refName := refParts[3]
		stringSchema.Default = refName
		stringSchema.ReadOnly = true
		ty.Value.Properties["type"] = stringSchema.NewRef()
	}
	configProp.Value.OneOf = types
	configProp.Value.Discriminator = &openapi3.Discriminator{
		PropertyName: "type",
		Mapping: map[string]string{
			"mysql":    "#/components/schemas/mysql",
			"postgres": "#/components/schemas/postgres",
			"redis":    "#/components/schemas/redis",
		},
	}

	// Finally, remove the AnyOf schema.
	configProp.Value.AnyOf = nil

	// Use the database_config type as a ref for the response schema.
	resp, ok := openAPIDoc.Components.Responses["database_config"]
	if !ok {
		panic("Expected to find the response schema for database_config")
	}

	resp.Value.Content.Get("application/json").Schema.Ref = "#/components/schemas/database_config"
}

func fixPageLinksType(openAPIDoc *openapi3.T) {
	schemaRef, ok := openAPIDoc.Components.Schemas["page_links"]
	if !ok {
		panic("Expected to find page_links type")
	}

	pagesProp, ok := schemaRef.Value.Properties["pages"]
	if !ok {
		panic("Expected to find 'pages' prop in page_links type")
	}

	pagesProp.Value.Properties = map[string]*openapi3.SchemaRef{
		"first": openapi3.NewStringSchema().NewRef(),
		"last":  openapi3.NewStringSchema().NewRef(),
		"next":  openapi3.NewStringSchema().NewRef(),
		"prev":  openapi3.NewStringSchema().NewRef(),
	}
	pagesProp.Value.AnyOf = nil
}
