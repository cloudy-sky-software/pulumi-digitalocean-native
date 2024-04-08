package provider

import (
	"context"
	"fmt"
	"net/http"
	"os"

	"github.com/getkin/kin-openapi/openapi3"

	"github.com/pkg/errors"

	"github.com/pulumi/pulumi/pkg/v3/resource/provider"

	"github.com/pulumi/pulumi/sdk/v3/go/common/resource"
	"github.com/pulumi/pulumi/sdk/v3/go/common/util/logging"
	pulumirpc "github.com/pulumi/pulumi/sdk/v3/proto/go"

	fwCallback "github.com/cloudy-sky-software/pulumi-provider-framework/callback"
	fwRest "github.com/cloudy-sky-software/pulumi-provider-framework/rest"
)

type digitalOceanProvider struct {
	name    string
	version string

	apiKey string
}

var (
	handler  *fwRest.Provider
	callback fwCallback.ProviderCallback
)

func makeProvider(host *provider.HostClient, name, version string, pulumiSchemaBytes, openapiDocBytes, metadataBytes []byte) (pulumirpc.ResourceProviderServer, error) {
	p := &digitalOceanProvider{
		name:    name,
		version: version,
	}

	callback = p
	rp, err := fwRest.MakeProvider(host, name, version, pulumiSchemaBytes, openapiDocBytes, metadataBytes, callback)

	handler = rp.(*fwRest.Provider)

	return rp, err
}

func (p *digitalOceanProvider) GetAuthorizationHeader() string {
	return fmt.Sprintf("%s %s", authSchemePrefix, p.apiKey)
}

func (p *digitalOceanProvider) OnPreInvoke(_ context.Context, _ *pulumirpc.InvokeRequest, _ *http.Request) error {
	return nil
}

func (p *digitalOceanProvider) OnPostInvoke(_ context.Context, _ *pulumirpc.InvokeRequest, outputs interface{}) (map[string]interface{}, error) {
	return outputs.(map[string]interface{}), nil
}

// OnConfigure is called by the provider framework when Pulumi calls Configure on
// the resource provider server.
func (p *digitalOceanProvider) OnConfigure(_ context.Context, req *pulumirpc.ConfigureRequest) (*pulumirpc.ConfigureResponse, error) {
	apiKey, ok := req.GetVariables()["digitalocean-native:config:apiKey"]
	if !ok {
		// Check if it's set as an env var.
		envVarNames := handler.GetSchemaSpec().Provider.InputProperties["apiKey"].DefaultInfo.Environment
		for _, n := range envVarNames {
			v := os.Getenv(n)
			if v != "" {
				apiKey = v
			}
		}

		// Return an error if the API key is still empty.
		if apiKey == "" {
			return nil, errors.New("api key is required")
		}
	}

	logging.V(3).Info("Configuring DigitalOcean API key")
	p.apiKey = apiKey

	return &pulumirpc.ConfigureResponse{
		AcceptSecrets: true,
	}, nil
}

// OnDiff checks what impacts a hypothetical update will have on the resource's properties.
func (p *digitalOceanProvider) OnDiff(_ context.Context, _ *pulumirpc.DiffRequest, _ string, _ *resource.ObjectDiff, _ *openapi3.MediaType) (*pulumirpc.DiffResponse, error) {
	return nil, nil
}

func (p *digitalOceanProvider) OnPreCreate(_ context.Context, _ *pulumirpc.CreateRequest, _ *http.Request) error {
	return nil
}

// OnPostCreate allocates a new instance of the provided resource and returns its unique ID afterwards.
func (p *digitalOceanProvider) OnPostCreate(_ context.Context, _ *pulumirpc.CreateRequest, outputs interface{}) (map[string]interface{}, error) {
	return outputs.(map[string]interface{}), nil
}

func (p *digitalOceanProvider) OnPreRead(_ context.Context, _ *pulumirpc.ReadRequest, _ *http.Request) error {
	return nil
}

func (p *digitalOceanProvider) OnPostRead(_ context.Context, _ *pulumirpc.ReadRequest, outputs map[string]interface{}) (map[string]interface{}, error) {
	return outputs, nil
}

func (p *digitalOceanProvider) OnPreUpdate(_ context.Context, _ *pulumirpc.UpdateRequest, _ *http.Request) error {
	return nil
}

func (p *digitalOceanProvider) OnPostUpdate(_ context.Context, _ *pulumirpc.UpdateRequest, _ http.Request, outputs interface{}) (map[string]interface{}, error) {
	return outputs.(map[string]interface{}), nil
}

func (p *digitalOceanProvider) OnPreDelete(_ context.Context, _ *pulumirpc.DeleteRequest, _ *http.Request) error {
	return nil
}

func (p *digitalOceanProvider) OnPostDelete(_ context.Context, _ *pulumirpc.DeleteRequest) error {
	return nil
}
