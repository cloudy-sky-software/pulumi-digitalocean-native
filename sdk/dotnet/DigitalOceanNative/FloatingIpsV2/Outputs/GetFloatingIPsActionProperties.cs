// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace CloudySkySoftware.Pulumi.DigitalOceanNative.FloatingIpsV2.Outputs
{

    [OutputType]
    public sealed class GetFloatingIPsActionProperties
    {
        public readonly Outputs.GetFloatingIPsActionPropertiesAction? Action;

        [OutputConstructor]
        private GetFloatingIPsActionProperties(Outputs.GetFloatingIPsActionPropertiesAction? action)
        {
            Action = action;
        }
    }
}
