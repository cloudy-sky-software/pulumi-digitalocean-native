// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.DigitalOceanNative.DatabasesV2.Outputs
{

    [OutputType]
    public sealed class GetDatabasesUserProperties
    {
        public readonly Outputs.DatabaseUser User;

        [OutputConstructor]
        private GetDatabasesUserProperties(Outputs.DatabaseUser user)
        {
            User = user;
        }
    }
}