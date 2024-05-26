// Code generated by pulumigen DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v2

import (
	"fmt"

	"github.com/blang/semver"
	"github.com/cloudy-sky-software/pulumi-digitalocean-native/sdk/go/dgtlocn/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type module struct {
	version semver.Version
}

func (m *module) Version() semver.Version {
	return m.version
}

func (m *module) Construct(ctx *pulumi.Context, name, typ, urn string) (r pulumi.Resource, err error) {
	switch typ {
	case "digitalocean-native:domains/v2:A":
		r = &A{}
	case "digitalocean-native:domains/v2:Aaaa":
		r = &Aaaa{}
	case "digitalocean-native:domains/v2:Caa":
		r = &Caa{}
	case "digitalocean-native:domains/v2:Cname":
		r = &Cname{}
	case "digitalocean-native:domains/v2:Domain":
		r = &Domain{}
	case "digitalocean-native:domains/v2:Mx":
		r = &Mx{}
	case "digitalocean-native:domains/v2:Ns":
		r = &Ns{}
	case "digitalocean-native:domains/v2:Soa":
		r = &Soa{}
	case "digitalocean-native:domains/v2:Srv":
		r = &Srv{}
	case "digitalocean-native:domains/v2:Txt":
		r = &Txt{}
	default:
		return nil, fmt.Errorf("unknown resource type: %s", typ)
	}

	err = ctx.RegisterResource(typ, name, nil, r, pulumi.URN_(urn))
	return
}

func init() {
	version, err := internal.PkgVersion()
	if err != nil {
		version = semver.Version{Major: 1}
	}
	pulumi.RegisterResourceModule(
		"digitalocean-native",
		"domains/v2",
		&module{version},
	)
}
