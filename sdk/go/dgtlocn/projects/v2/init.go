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
	case "digitalocean-native:projects/v2:Project":
		r = &Project{}
	case "digitalocean-native:projects/v2:ProjectsAssignResource":
		r = &ProjectsAssignResource{}
	case "digitalocean-native:projects/v2:ProjectsAssignResourcesDefault":
		r = &ProjectsAssignResourcesDefault{}
	case "digitalocean-native:projects/v2:ProjectsDefault":
		r = &ProjectsDefault{}
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
		"projects/v2",
		&module{version},
	)
}
