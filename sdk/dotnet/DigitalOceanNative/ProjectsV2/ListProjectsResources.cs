// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace CloudySkySoftware.Pulumi.DigitalOceanNative.ProjectsV2
{
    public static class ListProjectsResources
    {
        public static Task<Outputs.ListProjectsResourcesItems> InvokeAsync(ListProjectsResourcesArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<Outputs.ListProjectsResourcesItems>("digitalocean-native:projects/v2:listProjectsResources", args ?? new ListProjectsResourcesArgs(), options.WithDefaults());

        public static Output<Outputs.ListProjectsResourcesItems> Invoke(ListProjectsResourcesInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<Outputs.ListProjectsResourcesItems>("digitalocean-native:projects/v2:listProjectsResources", args ?? new ListProjectsResourcesInvokeArgs(), options.WithDefaults());
    }


    public sealed class ListProjectsResourcesArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// A unique identifier for a project.
        /// </summary>
        [Input("projectId", required: true)]
        public string ProjectId { get; set; } = null!;

        public ListProjectsResourcesArgs()
        {
        }
        public static new ListProjectsResourcesArgs Empty => new ListProjectsResourcesArgs();
    }

    public sealed class ListProjectsResourcesInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// A unique identifier for a project.
        /// </summary>
        [Input("projectId", required: true)]
        public Input<string> ProjectId { get; set; } = null!;

        public ListProjectsResourcesInvokeArgs()
        {
        }
        public static new ListProjectsResourcesInvokeArgs Empty => new ListProjectsResourcesInvokeArgs();
    }
}
