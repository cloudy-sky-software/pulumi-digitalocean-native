// Copyright 2022, Cloudy Sky Software LLC.

package provider

import (
	"context"
	"encoding/json"
	"os"
	"path/filepath"
	"testing"

	"github.com/cloudy-sky-software/pulumi-provider-framework/openapi"
	"github.com/cloudy-sky-software/pulumi-provider-framework/state"

	"github.com/stretchr/testify/assert"

	"github.com/pulumi/pulumi/sdk/v3/go/common/resource"
	"github.com/pulumi/pulumi/sdk/v3/go/common/resource/plugin"
	pulumirpc "github.com/pulumi/pulumi/sdk/v3/proto/go"
)

const testCreateJSONPayload = `{
    "image":"ubuntu-20-04-x64",
	"names":["my-droplet"],
	"size":"s-1vcpu-1gb"
}
`

func readFileFromProviderResourceDir(t *testing.T, filename string) []byte {
	t.Helper()

	b, err := os.ReadFile(filepath.Join("..", "..", "cmd", "pulumi-resource-digitalocean-native", filename))
	if err != nil {
		t.Fatalf("Failed reading openapi.yml: %v", err)
	}

	return b
}

func makeTestProvider(ctx context.Context, t *testing.T) pulumirpc.ResourceProviderServer {
	t.Helper()

	openapiBytes := readFileFromProviderResourceDir(t, "openapi_generated.yml")
	d := openapi.GetOpenAPISpec(openapiBytes)
	d.Servers[0].URL = "http://localhost:8080"
	openapiBytes, _ = d.MarshalJSON()

	pSchemaBytes := readFileFromProviderResourceDir(t, "schema.json")
	metadataBytes := readFileFromProviderResourceDir(t, "metadata.json")

	p, err := makeProvider(nil, "", "", pSchemaBytes, openapiBytes, metadataBytes)

	if err != nil {
		t.Fatalf("Could not create a provider instance: %v", err)
	}

	_, err = p.Configure(ctx, &pulumirpc.ConfigureRequest{
		Variables: map[string]string{"digitalocean-native:config:apiKey": "fakeapikey"},
	})

	if err != nil {
		t.Fatalf("Error configuring the provider: %v", err)
	}

	return p
}

func TestDiff(t *testing.T) {
	ctx := context.Background()

	p := makeTestProvider(ctx, t)

	outputs := make(map[string]interface{})
	outputs["size"] = "s-1vcpu-1gb"
	oldsStruct, _ := plugin.MarshalProperties(state.GetResourceState(outputs, resource.NewPropertyMapFromMap(outputs)), state.DefaultMarshalOpts)

	news := make(map[string]interface{})
	news["size"] = "s-1vcpu-2gb"
	newsStruct, _ := plugin.MarshalProperties(resource.NewPropertyMapFromMap(news), state.DefaultMarshalOpts)

	resp, err := p.Diff(ctx, &pulumirpc.DiffRequest{Id: "", Urn: "urn:pulumi:some-stack::some-project::digitalocean-native:droplets/v2:Droplets::someResourceName", Olds: oldsStruct, News: newsStruct})
	assert.Nil(t, err)
	assert.Equal(t, pulumirpc.DiffResponse_DIFF_SOME, resp.Changes)
	assert.NotEmpty(t, resp.Diffs)
	assert.Len(t, resp.Diffs, 1)
	assert.NotEmpty(t, resp.Replaces)
}

func TestCreate(t *testing.T) {
	ctx := context.Background()

	var inputs map[string]interface{}
	if err := json.Unmarshal([]byte(testCreateJSONPayload), &inputs); err != nil {
		t.Fatal("Failed to unmarshal test payload")
	}

	p := makeTestProvider(ctx, t)

	inputProperties, _ := plugin.MarshalProperties(resource.NewPropertyMapFromMap(inputs), state.DefaultMarshalOpts)

	_, err := p.Create(ctx, &pulumirpc.CreateRequest{
		Urn:        "urn:pulumi:dev::digitalocean-native-ts::digitalocean-native:droplets/v2:Droplets::webservice",
		Properties: inputProperties,
	})

	assert.NotNil(t, err)
	// For now just assume that if we got to the point of making the request, we are good to go.
	assert.Contains(t, err.Error(), "connect: connection refused")
}
