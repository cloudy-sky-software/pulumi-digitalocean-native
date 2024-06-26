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

    [OutputType]
    public sealed class Credentials
    {
        /// <summary>
        /// A base64 encoding of bytes representing the certificate authority data for accessing the cluster.
        /// </summary>
        public readonly string? CertificateAuthorityData;
        /// <summary>
        /// A base64 encoding of bytes representing the x509 client
        /// certificate data for access the cluster. This is only returned for clusters
        /// without support for token-based authentication.
        /// 
        /// Newly created Kubernetes clusters do not return credentials using
        /// certificate-based authentication. For additional information,
        /// [see here](https://www.digitalocean.com/docs/kubernetes/how-to/connect-to-cluster/#authenticate).
        /// </summary>
        public readonly string? ClientCertificateData;
        /// <summary>
        /// A base64 encoding of bytes representing the x509 client key
        /// data for access the cluster. This is only returned for clusters without
        /// support for token-based authentication.
        /// 
        /// Newly created Kubernetes clusters do not return credentials using
        /// certificate-based authentication. For additional information,
        /// [see here](https://www.digitalocean.com/docs/kubernetes/how-to/connect-to-cluster/#authenticate).
        /// </summary>
        public readonly string? ClientKeyData;
        /// <summary>
        /// A time value given in ISO8601 combined date and time format that represents when the access token expires.
        /// </summary>
        public readonly string? ExpiresAt;
        /// <summary>
        /// The URL used to access the cluster API server.
        /// </summary>
        public readonly string? Server;
        /// <summary>
        /// An access token used to authenticate with the cluster. This is only returned for clusters with support for token-based authentication.
        /// </summary>
        public readonly string? Token;

        [OutputConstructor]
        private Credentials(
            string? certificateAuthorityData,

            string? clientCertificateData,

            string? clientKeyData,

            string? expiresAt,

            string? server,

            string? token)
        {
            CertificateAuthorityData = certificateAuthorityData;
            ClientCertificateData = clientCertificateData;
            ClientKeyData = clientKeyData;
            ExpiresAt = expiresAt;
            Server = server;
            Token = token;
        }
    }
}
