// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.DigitalOceanNative.ImagesV2.Outputs
{

    [OutputType]
    public sealed class GetImagesProperties
    {
        public readonly Outputs.Image Image;

        [OutputConstructor]
        private GetImagesProperties(Outputs.Image image)
        {
            Image = image;
        }
    }
}
