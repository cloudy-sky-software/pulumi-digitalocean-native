// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.DigitalOceanNative.FunctionsV2.Outputs
{

    [OutputType]
    public sealed class TriggerInfoScheduledRunsProperties
    {
        /// <summary>
        /// Indicates last run time. null value indicates trigger not run yet.
        /// </summary>
        public readonly string? LastRunAt;
        /// <summary>
        /// Indicates next run time. null value indicates trigger will not run.
        /// </summary>
        public readonly string? NextRunAt;

        [OutputConstructor]
        private TriggerInfoScheduledRunsProperties(
            string? lastRunAt,

            string? nextRunAt)
        {
            LastRunAt = lastRunAt;
            NextRunAt = nextRunAt;
        }
    }
}
