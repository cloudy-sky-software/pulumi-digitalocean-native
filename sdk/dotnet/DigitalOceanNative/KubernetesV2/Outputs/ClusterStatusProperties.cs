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

    /// <summary>
    /// An object containing a `state` attribute whose value is set to a string indicating the current status of the cluster.
    /// </summary>
    [OutputType]
    public sealed class ClusterStatusProperties
    {
        /// <summary>
        /// An optional message providing additional information about the current cluster state.
        /// </summary>
        public readonly string? Message;
        /// <summary>
        /// A string indicating the current status of the cluster.
        /// </summary>
        public readonly CloudySkySoftware.Pulumi.DigitalOceanNative.KubernetesV2.ClusterStatusPropertiesState? State;

        [OutputConstructor]
        private ClusterStatusProperties(
            string? message,

            CloudySkySoftware.Pulumi.DigitalOceanNative.KubernetesV2.ClusterStatusPropertiesState? state)
        {
            Message = message;
            State = state;
        }
    }
}