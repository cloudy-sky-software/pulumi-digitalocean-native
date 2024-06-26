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
    public sealed class AppsDeploymentProgressStep
    {
        public readonly string? ComponentName;
        public readonly string? EndedAt;
        /// <summary>
        /// The base of a human-readable description of the step intended to be combined with the component name for presentation. For example:
        /// 
        /// `message_base` = "Building service"
        /// `component_name` = "api"
        /// </summary>
        public readonly string? MessageBase;
        public readonly string? Name;
        public readonly Outputs.AppsDeploymentProgressStepReason? Reason;
        public readonly string? StartedAt;
        public readonly CloudySkySoftware.Pulumi.DigitalOceanNative.AppsV2.AppsDeploymentProgressStepStatus? Status;
        public readonly ImmutableArray<object> Steps;

        [OutputConstructor]
        private AppsDeploymentProgressStep(
            string? componentName,

            string? endedAt,

            string? messageBase,

            string? name,

            Outputs.AppsDeploymentProgressStepReason? reason,

            string? startedAt,

            CloudySkySoftware.Pulumi.DigitalOceanNative.AppsV2.AppsDeploymentProgressStepStatus? status,

            ImmutableArray<object> steps)
        {
            ComponentName = componentName;
            EndedAt = endedAt;
            MessageBase = messageBase;
            Name = name;
            Reason = reason;
            StartedAt = startedAt;
            Status = status;
            Steps = steps;
        }
    }
}
