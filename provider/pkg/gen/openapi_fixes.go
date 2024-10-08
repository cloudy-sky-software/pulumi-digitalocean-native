package gen

import (
	"strings"

	"github.com/getkin/kin-openapi/openapi3"

	"github.com/pulumi/pulumi/sdk/v3/go/common/util/contract"
)

func fixCreateUptimeCheckRequest(openAPIDoc *openapi3.T) {
	pathItem := openAPIDoc.Paths.Find("/v2/uptime/checks")
	contract.Assertf(pathItem != nil, "Expected to find request path /v2/uptime/checks")

	reqSchema := pathItem.Post.RequestBody.Value.Content.Get("application/json")

	// `method` is an invalid required property as it is not even
	// a property in the request schema.
	reqSchema.Schema.Value.Required = []string{"name", "target", "regions", "type", "enabled"}
}

func fixReservedIPActionUnassignType(openAPIDoc *openapi3.T) {
	schemaRef, ok := openAPIDoc.Components.Schemas["reserved_ip_action_unassign"]
	contract.Assertf(ok, "Expected to find reserved_ip_action_unassign type")

	intSchema := openapi3.NewIntegerSchema()
	intSchema.Description = "The ID of the Droplet that the reserved IP will be unassigned from."
	schemaRef.Value.AllOf[1].Value.Properties = map[string]*openapi3.SchemaRef{
		"droplet_id": intSchema.NewRef(),
	}
	// Overwrite the list of required properties for this type,
	// so that only `droplet_id` is required.
	schemaRef.Value.Required = []string{"droplet_id"}
}

func fixReservedIPType(openAPIDoc *openapi3.T) {
	schemaRef, ok := openAPIDoc.Components.Schemas["reserved_ip"]
	contract.Assertf(ok, "Expected to find reserved_ip type")

	// The droplet property of this type can be null if the floating IP
	// is not attached to any droplet. The way this is represented in the
	// spec causes validation issues. So let's simplify it.
	dropletProp := schemaRef.Value.Properties["droplet"]
	dropletProp.Ref = "#/components/schemas/droplet"
	dropletProp.Value.AnyOf = nil
	dropletProp.Value.Example = nil
	dropletSchema, ok := openAPIDoc.Components.Schemas["droplet"]
	contract.Assertf(ok, "droplet schema type not found. required by reserved_ip type")
	dropletProp.Value = dropletSchema.Value
}

// fixObjectPropertyType adds `type: object` for any type that
// does not have a value defined for the `type` property.
func fixObjectPropertyType(objectType *openapi3.SchemaRef) {
	if (objectType.Value.Type == nil || len(objectType.Value.Type.Slice()) == 0) && len(objectType.Value.Properties) > 0 {
		objectType.Value.Type = &openapi3.Types{"object"}
	}

	for _, allOfTy := range objectType.Value.AllOf {
		fixObjectPropertyType(allOfTy)
	}
}

func fixResponseObjectPropertyType(responseRef *openapi3.ResponseRef) {
	responseContent := responseRef.Value.Content.Get("application/json")
	if responseContent == nil {
		return
	}

	if (responseContent.Schema.Value.Type == nil || len(responseContent.Schema.Value.Type.Slice()) == 0) && len(responseContent.Schema.Value.Properties) > 0 {
		responseContent.Schema.Value.Type = &openapi3.Types{"object"}
	}

	for _, allOfTy := range responseContent.Schema.Value.AllOf {
		fixObjectPropertyType(allOfTy)
	}
}

func fixFloatingIPType(openAPIDoc *openapi3.T) {
	floatingIP, ok := openAPIDoc.Components.Schemas["floating_ip"]
	contract.Assertf(ok, "Expected to find floating_ip type")

	// The droplet property of this type can be null if the floating IP
	// is not attached to any droplet. The way this is represented in the
	// spec causes validation issues. So let's simplify it.
	dropletProp := floatingIP.Value.Properties["droplet"]
	dropletProp.Ref = "#/components/schemas/droplet"
	dropletProp.Value.AnyOf = nil
	dropletProp.Value.Example = nil
	dropletSchema, ok := openAPIDoc.Components.Schemas["droplet"]
	contract.Assertf(ok, "droplet schema type not found. required by floating_ip schema type")
	dropletProp.Value = dropletSchema.Value
}

// fixDatabaseConfigType fixes the database_config type to use
// a OneOf schema (with a discriminator) instead of an AnyOf.
func fixDatabaseConfigType(openAPIDoc *openapi3.T) {
	databaseConfig, ok := openAPIDoc.Components.Schemas["database_config"]
	contract.Assertf(ok, "Expected to find database_config type")

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
	contract.Assertf(ok, "Expected to find the response schema for database_config")

	resp.Value.Content.Get("application/json").Schema.Ref = "#/components/schemas/database_config"
	resp.Value.Content.Get("application/json").Schema.Value = databaseConfig.Value
}

func fixPageLinksType(openAPIDoc *openapi3.T) {
	schemaRef, ok := openAPIDoc.Components.Schemas["page_links"]
	contract.Assertf(ok, "Expected to find page_links type")

	pagesProp, ok := schemaRef.Value.Properties["pages"]
	contract.Assertf(ok, "Expected to find 'pages' prop in page_links type")

	pagesProp.Value.Properties = map[string]*openapi3.SchemaRef{
		"first": openapi3.NewStringSchema().NewRef(),
		"last":  openapi3.NewStringSchema().NewRef(),
		"next":  openapi3.NewStringSchema().NewRef(),
		"prev":  openapi3.NewStringSchema().NewRef(),
	}
	pagesProp.Value.AnyOf = nil
}

// The POST operation uses a oneOf definition in the
// request body schema to let the user create either
// one or many droplets. But the oneOf definition
// does not have a discriminator. Instead, we'll
// just change the definition to let the user
// create a single droplet. It doesn't make
// sense in IaC to create multiple resources
// from a single definition.
func fixCreateDropletsRequest(openAPIDoc *openapi3.T) {
	pathItem := openAPIDoc.Paths.Find("/v2/droplets")
	contract.Assertf(pathItem != nil, "Expected to find request path /v2/droplets")

	singleCreate := openAPIDoc.Components.Schemas["droplet_single_create"]

	pathItem.Post.OperationID = "droplet_create"
	reqSchema := pathItem.Post.RequestBody.Value.Content.Get("application/json")
	reqSchema.Schema = openapi3.NewSchemaRef("#/components/schemas/droplet_single_create", singleCreate.Value)
	reqSchema.Schema.Value.OneOf = nil
}

func fixDeleteDropletOperations(openAPIDoc *openapi3.T) {
	pathItem := openAPIDoc.Paths.Find("/v2/droplets")
	contract.Assertf(pathItem != nil, "Expected to find request path /v2/droplets")

	pathItem.Delete = nil

	pathItem = openAPIDoc.Paths.Find("/v2/droplets/{droplet_id}")
	contract.Assertf(pathItem != nil, "Expected to find request path /v2/droplets/{droplet_id}")
	pathItem.Delete.OperationID = "droplet_delete"
}

func fixCreateVolumeRequest(openAPIDoc *openapi3.T) {
	pathItem := openAPIDoc.Paths.Find("/v2/volumes")
	contract.Assertf(pathItem != nil, "Expected to find request path /v2/volumes")

	schemaRefs := pathItem.Post.RequestBody.Value.Content.Get("application/json").Schema.Value.AnyOf
	pathItem.Post.RequestBody.Value.Content.Get("application/json").Schema.Value.OneOf = schemaRefs
	volumesDiscriminator := &openapi3.Discriminator{
		Mapping: map[string]string{
			"ext4": "#/components/schemas/volumes_ext4",
			"xfs":  "#/components/schemas/volumes_xfs",
		},
		PropertyName: "filesystem_type",
	}
	pathItem.Post.RequestBody.Value.Content.Get("application/json").Schema.Value.Discriminator = volumesDiscriminator
	pathItem.Post.RequestBody.Value.Content.Get("application/json").Schema.Value.AnyOf = nil
}

func fixFileSystemLabelProperty(openAPIDoc *openapi3.T) {
	schemas := []string{"volumes_ext4", "volumes_xfs"}
	for _, schemaName := range schemas {
		schemaTypeRef := openAPIDoc.Components.Schemas[schemaName]
		for _, schemaRef := range schemaTypeRef.Value.AllOf {
			if len(schemaRef.Value.Properties) == 0 {
				continue
			}

			schemaRef.Value.Properties["filesystem_label"] = openapi3.NewSchemaRef("", openapi3.NewStringSchema())
		}
	}
}

func fixCreateCustomDomainRequest(openAPIDoc *openapi3.T) {
	// The POST operation uses an anyOf definition with a
	// discriminator. While it's technically correct, none
	// of the other operations on a domain record use the
	// discriminator leading to a situation where we are
	// not able to tie the other operations to the
	// corresponding type token properly.
	//
	// Moreover, it seems that the reasoning for using
	// the anyOf definition is to require separate properties
	// in the same underlying object that they all point to.
	// While we'd lose the per-record type resource,
	// we are ensuring that the custom domain record
	// resource is fully linked with all operations.
	//
	// We'll revisit this based on user feedback.
	recordsPathItem := openAPIDoc.Paths.Find("/v2/domains/{domain_name}/records")
	contract.Assertf(recordsPathItem != nil, "Expected to find request path /v2/domains/{domain_name}/records")

	domainRecord := openAPIDoc.Components.Schemas["domain_record"]
	recordsPathItem.Post.RequestBody.Value.Content.Get("application/json").Schema = openapi3.NewSchemaRef("#/components/schemas/domain_record", domainRecord.Value)
}

func FixOpenAPIDoc(openAPIDoc *openapi3.T) {
	regionSchema, ok := openAPIDoc.Components.Schemas["region"]
	contract.Assertf(ok, "Expected to find a schema for region type")

	regionSchema.Value.Properties["features"].Value.Type = &openapi3.Types{"array"}
	regionSchema.Value.Properties["sizes"].Value.Type = &openapi3.Types{"array"}

	for _, ty := range openAPIDoc.Components.Schemas {
		fixObjectPropertyType(ty)
	}

	for _, ty := range openAPIDoc.Components.Responses {
		fixResponseObjectPropertyType(ty)
	}

	fixDatabaseConfigType(openAPIDoc)
	fixFloatingIPType(openAPIDoc)
	fixReservedIPActionUnassignType(openAPIDoc)
	fixReservedIPType(openAPIDoc)
	fixPageLinksType(openAPIDoc)
	fixCreateUptimeCheckRequest(openAPIDoc)
	fixCreateDropletsRequest(openAPIDoc)
	fixDeleteDropletOperations(openAPIDoc)
	fixCreateVolumeRequest(openAPIDoc)
	fixFileSystemLabelProperty(openAPIDoc)
	fixCreateCustomDomainRequest(openAPIDoc)
}
