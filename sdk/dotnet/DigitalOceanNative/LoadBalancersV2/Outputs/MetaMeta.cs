// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace CloudySkySoftware.Pulumi.DigitalOceanNative.LoadBalancersV2.Outputs
{

    [OutputType]
    public sealed class MetaMeta
    {
        /// <summary>
        /// Number of objects returned by the request.
        /// </summary>
        public readonly int? Total;

        [OutputConstructor]
        private MetaMeta(int? total)
        {
            Total = total;
        }
    }
}
