// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.DigitalOceanNative.SizesV2.Outputs
{

    [OutputType]
    public sealed class PageLinks
    {
        public readonly Outputs.PageLinksPagesProperties? Pages;

        [OutputConstructor]
        private PageLinks(Outputs.PageLinksPagesProperties? pages)
        {
            Pages = pages;
        }
    }
}
