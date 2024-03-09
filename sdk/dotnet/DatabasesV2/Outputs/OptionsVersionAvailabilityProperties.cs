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
    public sealed class OptionsVersionAvailabilityProperties
    {
        /// <summary>
        /// An array of objects, each indicating the version end-of-life, end-of-availability for various database engines
        /// </summary>
        public readonly ImmutableArray<Outputs.DatabaseVersionAvailability> Mongodb;
        /// <summary>
        /// An array of objects, each indicating the version end-of-life, end-of-availability for various database engines
        /// </summary>
        public readonly ImmutableArray<Outputs.DatabaseVersionAvailability> Mysql;
        /// <summary>
        /// An array of objects, each indicating the version end-of-life, end-of-availability for various database engines
        /// </summary>
        public readonly ImmutableArray<Outputs.DatabaseVersionAvailability> Pg;
        /// <summary>
        /// An array of objects, each indicating the version end-of-life, end-of-availability for various database engines
        /// </summary>
        public readonly ImmutableArray<Outputs.DatabaseVersionAvailability> Redis;

        [OutputConstructor]
        private OptionsVersionAvailabilityProperties(
            ImmutableArray<Outputs.DatabaseVersionAvailability> mongodb,

            ImmutableArray<Outputs.DatabaseVersionAvailability> mysql,

            ImmutableArray<Outputs.DatabaseVersionAvailability> pg,

            ImmutableArray<Outputs.DatabaseVersionAvailability> redis)
        {
            Mongodb = mongodb;
            Mysql = mysql;
            Pg = pg;
            Redis = redis;
        }
    }
}
