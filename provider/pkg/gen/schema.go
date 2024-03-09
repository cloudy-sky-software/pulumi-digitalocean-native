// Copyright 2022, Cloudy Sky Software LLC.

package gen

import (
	"bytes"
	"encoding/json"

	"github.com/getkin/kin-openapi/openapi3"

	dotnetgen "github.com/pulumi/pulumi/pkg/v3/codegen/dotnet"
	nodejsgen "github.com/pulumi/pulumi/pkg/v3/codegen/nodejs"
	pschema "github.com/pulumi/pulumi/pkg/v3/codegen/schema"

	"github.com/pulumi/pulumi/sdk/v3/go/common/util/contract"

	openapigen "github.com/cloudy-sky-software/pulschema/pkg"

	"github.com/cloudy-sky-software/pulumi-digitalocean-native/provider/pkg/gen/examples"
)

const packageName = "digitalocean-native"

// PulumiSchema will generate a Pulumi schema for the given k8s schema.
func PulumiSchema(openapiDoc *openapi3.T) (pschema.PackageSpec, openapigen.ProviderMetadata, openapi3.T) {
	pkg := pschema.PackageSpec{
		Name:        packageName,
		Description: "A Pulumi package for creating and managing DigitalOcean resources.",
		DisplayName: "DigitalOcean",
		License:     "Apache-2.0",
		Keywords: []string{
			"pulumi",
			packageName,
			"category/cloud",
			"kind/native",
		},
		Homepage:   "https://cloudysky.software",
		Publisher:  "Cloudy Sky Software",
		Repository: "https://github.com/cloudy-sky-software/pulumi-digitalocean-native",

		Config: pschema.ConfigSpec{
			Variables: map[string]pschema.PropertySpec{
				"apiKey": {
					Description: "The API key",
					TypeSpec:    pschema.TypeSpec{Type: "string"},
					Language: map[string]pschema.RawMessage{
						"csharp": rawMessage(map[string]interface{}{
							"name": "ApiKey",
						}),
					},
					Secret: true,
				},
			},
		},

		Provider: pschema.ResourceSpec{
			ObjectTypeSpec: pschema.ObjectTypeSpec{
				Description: "The provider type for the DigitalOcean package.",
				Type:        "object",
			},
			InputProperties: map[string]pschema.PropertySpec{
				"apiKey": {
					DefaultInfo: &pschema.DefaultSpec{
						Environment: []string{
							"DIGITALOCEAN_NATIVE_APIKEY",
						},
					},
					Description: "The DigitalOcean API key.",
					TypeSpec:    pschema.TypeSpec{Type: "string"},
					Language: map[string]pschema.RawMessage{
						"csharp": rawMessage(map[string]interface{}{
							"name": "ApiKey",
						}),
					},
					Secret: true,
				},
			},
		},

		PluginDownloadURL: "github://api.github.com/cloudy-sky-software/pulumi-digitalocean-native",
		Types:             map[string]pschema.ComplexTypeSpec{},
		Resources:         map[string]pschema.ResourceSpec{},
		Functions:         map[string]pschema.FunctionSpec{},
		Language:          map[string]pschema.RawMessage{},
	}

	csharpNamespaces := map[string]string{
		"digitalocean-native": "DigitalOceanNative",
		// TODO: Is this needed?
		"": "Provider",
	}

	openAPICtx := &openapigen.OpenAPIContext{
		Doc: *openapiDoc,
		Pkg: &pkg,
		ExcludedPaths: []string{
			"/v2/customers/my/invoices/{invoice_uuid}",
			"/v2/customers/my/invoices/{invoice_uuid}/csv",
			"/v2/customers/my/invoices/{invoice_uuid}/pdf",
			"/v2/customers/my/invoices/{invoice_uuid}/summary",
			// pulschema does not support POST/PUT requests without a request body yet.
			"/v2/apps/{app_id}/rollback/commit",
			"/v2/apps/{app_id}/deployments/{deployment_id}/cancel",
			"/v2/apps/{app_id}/rollback/revert",
			"/v2/registry/{registry_name}/garbage-collection",
			"/v2/databases/{database_cluster_uuid}/replicas/{replica_name}/promote",
			"/v2/droplets/{droplet_id}/destroy_with_associated_resources/retry",
			//
			// pulschema does not support application/yaml response type yet.
			"/v2/kubernetes/clusters/{cluster_id}/kubeconfig",
			//
			// Similar to /v2/volumes/{volume_id}/actions which is more appropriate
			// since it is by volume ID.
			"/v2/volumes/actions",
		},
	}

	providerMetadata, updatedOpenAPIDoc, err := openAPICtx.GatherResourcesFromAPI(csharpNamespaces)
	if err != nil {
		contract.Failf("generating resources from OpenAPI spec: %v", err)
	}

	// Override some module names.
	csharpNamespaces["1-clicks/v2"] = "OneClicksV2"

	// Add examples to resources
	for k, v := range examples.ResourceExample {
		if r, ok := pkg.Resources[k]; ok {
			r.Description += "\n\n" + v
			pkg.Resources[k] = r
		}
	}

	pkg.Language["csharp"] = rawMessage(dotnetgen.CSharpPackageInfo{
		// TODO: What does this enable?
		// DictionaryConstructors: true,
		Namespaces: csharpNamespaces,
		PackageReferences: map[string]string{
			"Pulumi": "3.*",
		},
		RootNamespace: "Pulumi",
	})

	pkg.Language["go"] = rawMessage(map[string]interface{}{
		"importBasePath": "github.com/cloudy-sky-software/pulumi-digitalocean-native/sdk/go/dgtlocn",
	})

	pkg.Language["nodejs"] = rawMessage(nodejsgen.NodePackageInfo{
		PackageName: "@cloudyskysoftware/pulumi-digitalocean-native",
		ModuleToPackage: map[string]string{
			"1-clicks/v2": "oneclicks/v2",
		},
	})

	pkg.Language["python"] = rawMessage(map[string]interface{}{
		"packageName": "pulumi_digitalocean-native",
		"requires": map[string]string{
			"pulumi": ">=3.0.0,<4.0.0",
		},
	})

	metadata := openapigen.ProviderMetadata{
		ResourceCRUDMap:  providerMetadata.ResourceCRUDMap,
		AutoNameMap:      providerMetadata.AutoNameMap,
		SdkToApiNameMap:  providerMetadata.SdkToApiNameMap,
		PathParamNameMap: providerMetadata.PathParamNameMap,
	}
	return pkg, metadata, updatedOpenAPIDoc
}

func rawMessage(v interface{}) pschema.RawMessage {
	var out bytes.Buffer
	encoder := json.NewEncoder(&out)
	encoder.SetEscapeHTML(false)
	err := encoder.Encode(v)
	contract.Assert(err == nil)
	return out.Bytes()
}
