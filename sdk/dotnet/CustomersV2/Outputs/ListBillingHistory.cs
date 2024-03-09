// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.DigitalOceanNative.CustomersV2.Outputs
{

    [OutputType]
    public sealed class ListBillingHistory
    {
        public readonly ImmutableArray<Outputs.BillingHistory> BillingHistory;
        public readonly Outputs.PageLinks? Links;
        /// <summary>
        /// Information about the response itself.
        /// </summary>
        public readonly Outputs.MetaProperties Meta;

        [OutputConstructor]
        private ListBillingHistory(
            ImmutableArray<Outputs.BillingHistory> billingHistory,

            Outputs.PageLinks? links,

            Outputs.MetaProperties meta)
        {
            BillingHistory = billingHistory;
            Links = links;
            Meta = meta;
        }
    }
}
