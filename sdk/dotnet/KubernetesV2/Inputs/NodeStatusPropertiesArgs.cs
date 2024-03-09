// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.DigitalOceanNative.KubernetesV2.Inputs
{

    /// <summary>
    /// An object containing a `state` attribute whose value is set to a string indicating the current status of the node.
    /// </summary>
    public sealed class NodeStatusPropertiesArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// A string indicating the current status of the node.
        /// </summary>
        [Input("state")]
        public Input<Pulumi.DigitalOceanNative.KubernetesV2.NodeStatusPropertiesState>? State { get; set; }

        public NodeStatusPropertiesArgs()
        {
        }
        public static new NodeStatusPropertiesArgs Empty => new NodeStatusPropertiesArgs();
    }
}
