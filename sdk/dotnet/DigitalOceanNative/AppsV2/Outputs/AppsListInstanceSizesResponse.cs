// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace CloudySkySoftware.Pulumi.DigitalOceanNative.AppsV2.Outputs
{

    [OutputType]
    public sealed class AppsListInstanceSizesResponse
    {
        public readonly double? DiscountPercent;
        public readonly ImmutableArray<Outputs.AppsInstanceSize> InstanceSizes;

        [OutputConstructor]
        private AppsListInstanceSizesResponse(
            double? discountPercent,

            ImmutableArray<Outputs.AppsInstanceSize> instanceSizes)
        {
            DiscountPercent = discountPercent;
            InstanceSizes = instanceSizes;
        }
    }
}
