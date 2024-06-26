// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace CloudySkySoftware.Pulumi.DigitalOceanNative.DropletsV2.Outputs
{

    [OutputType]
    public sealed class DropletSnapshot
    {
        /// <summary>
        /// A time value given in ISO8601 combined date and time format that represents when the snapshot was created.
        /// </summary>
        public readonly string CreatedAt;
        /// <summary>
        /// The minimum size in GB required for a volume or Droplet to use this snapshot.
        /// </summary>
        public readonly int MinDiskSize;
        /// <summary>
        /// A human-readable name for the snapshot.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// An array of the regions that the snapshot is available in. The regions are represented by their identifying slug values.
        /// </summary>
        public readonly ImmutableArray<string> Regions;
        /// <summary>
        /// The billable size of the snapshot in gigabytes.
        /// </summary>
        public readonly double SizeGigabytes;
        /// <summary>
        /// Describes the kind of image. It may be one of `snapshot` or `backup`. This specifies whether an image is a user-generated Droplet snapshot or automatically created Droplet backup.
        /// </summary>
        public readonly CloudySkySoftware.Pulumi.DigitalOceanNative.DropletsV2.DropletSnapshotPropertiesType Type;

        [OutputConstructor]
        private DropletSnapshot(
            string createdAt,

            int minDiskSize,

            string name,

            ImmutableArray<string> regions,

            double sizeGigabytes,

            CloudySkySoftware.Pulumi.DigitalOceanNative.DropletsV2.DropletSnapshotPropertiesType type)
        {
            CreatedAt = createdAt;
            MinDiskSize = minDiskSize;
            Name = name;
            Regions = regions;
            SizeGigabytes = sizeGigabytes;
            Type = type;
        }
    }
}
