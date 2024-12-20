// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace CloudySkySoftware.Pulumi.DigitalOceanNative.CertificatesV2
{
    public static class GetCertificate
    {
        public static Task<Outputs.GetCertificateProperties> InvokeAsync(GetCertificateArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<Outputs.GetCertificateProperties>("digitalocean-native:certificates/v2:getCertificate", args ?? new GetCertificateArgs(), options.WithDefaults());

        public static Output<Outputs.GetCertificateProperties> Invoke(GetCertificateInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<Outputs.GetCertificateProperties>("digitalocean-native:certificates/v2:getCertificate", args ?? new GetCertificateInvokeArgs(), options.WithDefaults());

        public static Output<Outputs.GetCertificateProperties> Invoke(GetCertificateInvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<Outputs.GetCertificateProperties>("digitalocean-native:certificates/v2:getCertificate", args ?? new GetCertificateInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetCertificateArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// A unique identifier for a certificate.
        /// </summary>
        [Input("certificateId", required: true)]
        public string CertificateId { get; set; } = null!;

        public GetCertificateArgs()
        {
        }
        public static new GetCertificateArgs Empty => new GetCertificateArgs();
    }

    public sealed class GetCertificateInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// A unique identifier for a certificate.
        /// </summary>
        [Input("certificateId", required: true)]
        public Input<string> CertificateId { get; set; } = null!;

        public GetCertificateInvokeArgs()
        {
        }
        public static new GetCertificateInvokeArgs Empty => new GetCertificateInvokeArgs();
    }
}
