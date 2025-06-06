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
    public static class ListProjectsResourcesDefault
    {
        public static Task<Outputs.ListProjectsResourcesDefaultItems> InvokeAsync(ListProjectsResourcesDefaultArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<Outputs.ListProjectsResourcesDefaultItems>("digitalocean-native:projects/v2:listProjectsResourcesDefault", args ?? new ListProjectsResourcesDefaultArgs(), options.WithDefaults());

        public static Output<Outputs.ListProjectsResourcesDefaultItems> Invoke(InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<Outputs.ListProjectsResourcesDefaultItems>("digitalocean-native:projects/v2:listProjectsResourcesDefault", InvokeArgs.Empty, options.WithDefaults());

        public static Output<Outputs.ListProjectsResourcesDefaultItems> Invoke(InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<Outputs.ListProjectsResourcesDefaultItems>("digitalocean-native:projects/v2:listProjectsResourcesDefault", InvokeArgs.Empty, options.WithDefaults());
    }


    public sealed class ListProjectsResourcesDefaultArgs : global::Pulumi.InvokeArgs
    {
        public ListProjectsResourcesDefaultArgs()
        {
        }
        public static new ListProjectsResourcesDefaultArgs Empty => new ListProjectsResourcesDefaultArgs();
    }
}
