// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace CloudySkySoftware.Pulumi.DigitalOceanNative.KubernetesV2.Outputs
{

    [OutputType]
    public sealed class GetKubernetesAvailableUpgradeProperties
    {
        public readonly ImmutableArray<Outputs.KubernetesVersion> AvailableUpgradeVersions;

        [OutputConstructor]
        private GetKubernetesAvailableUpgradeProperties(ImmutableArray<Outputs.KubernetesVersion> availableUpgradeVersions)
        {
            AvailableUpgradeVersions = availableUpgradeVersions;
        }
    }
}
