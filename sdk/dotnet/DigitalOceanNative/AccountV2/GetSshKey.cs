// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace CloudySkySoftware.Pulumi.DigitalOceanNative.AccountV2
{
    public static class GetSshKey
    {
        public static Task<Outputs.GetSshKeyProperties> InvokeAsync(GetSshKeyArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<Outputs.GetSshKeyProperties>("digitalocean-native:account/v2:getSshKey", args ?? new GetSshKeyArgs(), options.WithDefaults());

        public static Output<Outputs.GetSshKeyProperties> Invoke(GetSshKeyInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<Outputs.GetSshKeyProperties>("digitalocean-native:account/v2:getSshKey", args ?? new GetSshKeyInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetSshKeyArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Either the ID or the fingerprint of an existing SSH key.
        /// </summary>
        [Input("sshKeyIdentifier", required: true)]
        public string SshKeyIdentifier { get; set; } = null!;

        public GetSshKeyArgs()
        {
        }
        public static new GetSshKeyArgs Empty => new GetSshKeyArgs();
    }

    public sealed class GetSshKeyInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Either the ID or the fingerprint of an existing SSH key.
        /// </summary>
        [Input("sshKeyIdentifier", required: true)]
        public Input<string> SshKeyIdentifier { get; set; } = null!;

        public GetSshKeyInvokeArgs()
        {
        }
        public static new GetSshKeyInvokeArgs Empty => new GetSshKeyInvokeArgs();
    }
}
