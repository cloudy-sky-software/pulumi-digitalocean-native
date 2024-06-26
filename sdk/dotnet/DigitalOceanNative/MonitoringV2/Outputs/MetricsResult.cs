// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace CloudySkySoftware.Pulumi.DigitalOceanNative.MonitoringV2.Outputs
{

    [OutputType]
    public sealed class MetricsResult
    {
        /// <summary>
        /// An object containing the metric labels.
        /// </summary>
        public readonly object Metric;
        public readonly ImmutableArray<ImmutableArray<Union<int, string>>> Values;

        [OutputConstructor]
        private MetricsResult(
            object metric,

            ImmutableArray<ImmutableArray<Union<int, string>>> values)
        {
            Metric = metric;
            Values = values;
        }
    }
}
