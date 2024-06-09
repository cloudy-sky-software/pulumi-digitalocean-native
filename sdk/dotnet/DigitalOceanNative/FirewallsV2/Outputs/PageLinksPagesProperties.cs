// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace CloudySkySoftware.Pulumi.DigitalOceanNative.FirewallsV2.Outputs
{

    [OutputType]
    public sealed class PageLinksPagesProperties
    {
        public readonly string? First;
        public readonly string? Last;
        public readonly string? Next;
        public readonly string? Prev;

        [OutputConstructor]
        private PageLinksPagesProperties(
            string? first,

            string? last,

            string? next,

            string? prev)
        {
            First = first;
            Last = last;
            Next = next;
            Prev = prev;
        }
    }
}
