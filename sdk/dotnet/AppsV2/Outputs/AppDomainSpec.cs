// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.DigitalOceanNative.AppsV2.Outputs
{

    [OutputType]
    public sealed class AppDomainSpec
    {
        /// <summary>
        /// The hostname for the domain
        /// </summary>
        public readonly string Domain;
        /// <summary>
        /// The minimum version of TLS a client application can use to access resources for the domain.  Must be one of the following values wrapped within quotations: `"1.2"` or `"1.3"`.
        /// </summary>
        public readonly Pulumi.DigitalOceanNative.AppsV2.AppDomainSpecMinimumTlsVersion? MinimumTlsVersion;
        /// <summary>
        /// - DEFAULT: The default `.ondigitalocean.app` domain assigned to this app
        /// - PRIMARY: The primary domain for this app that is displayed as the default in the control panel, used in bindable environment variables, and any other places that reference an app's live URL. Only one domain may be set as primary.
        /// - ALIAS: A non-primary domain
        /// </summary>
        public readonly Pulumi.DigitalOceanNative.AppsV2.AppDomainSpecType? Type;
        /// <summary>
        /// Indicates whether the domain includes all sub-domains, in addition to the given domain
        /// </summary>
        public readonly bool? Wildcard;
        /// <summary>
        /// Optional. If the domain uses DigitalOcean DNS and you would like App
        /// Platform to automatically manage it for you, set this to the name of the
        /// domain on your account.
        /// 
        /// For example, If the domain you are adding is `app.domain.com`, the zone
        /// could be `domain.com`.
        /// </summary>
        public readonly string? Zone;

        [OutputConstructor]
        private AppDomainSpec(
            string domain,

            Pulumi.DigitalOceanNative.AppsV2.AppDomainSpecMinimumTlsVersion? minimumTlsVersion,

            Pulumi.DigitalOceanNative.AppsV2.AppDomainSpecType? type,

            bool? wildcard,

            string? zone)
        {
            Domain = domain;
            MinimumTlsVersion = minimumTlsVersion;
            Type = type;
            Wildcard = wildcard;
            Zone = zone;
        }
    }
}
