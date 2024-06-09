// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace CloudySkySoftware.Pulumi.DigitalOceanNative.VolumesV2.Outputs
{

    [OutputType]
    public sealed class VolumeFull
    {
        /// <summary>
        /// A time value given in ISO8601 combined date and time format that represents when the block storage volume was created.
        /// </summary>
        public readonly string? CreatedAt;
        /// <summary>
        /// An optional free-form text field to describe a block storage volume.
        /// </summary>
        public readonly string? Description;
        /// <summary>
        /// An array containing the IDs of the Droplets the volume is attached to. Note that at this time, a volume can only be attached to a single Droplet.
        /// </summary>
        public readonly ImmutableArray<int> DropletIds;
        /// <summary>
        /// The label currently applied to the filesystem.
        /// </summary>
        public readonly string? FilesystemLabel;
        /// <summary>
        /// The type of filesystem currently in-use on the volume.
        /// </summary>
        public readonly string? FilesystemType;
        /// <summary>
        /// The unique identifier for the block storage volume.
        /// </summary>
        public readonly string? Id;
        /// <summary>
        /// A human-readable name for the block storage volume. Must be lowercase and be composed only of numbers, letters and "-", up to a limit of 64 characters. The name must begin with a letter.
        /// </summary>
        public readonly string? Name;
        public readonly Outputs.VolumeFullPropertiesRegion? Region;
        /// <summary>
        /// The size of the block storage volume in GiB (1024^3). This field does not apply  when creating a volume from a snapshot.
        /// </summary>
        public readonly int? SizeGigabytes;
        /// <summary>
        /// A flat array of tag names as strings to be applied to the resource. Tag names may be for either existing or new tags.
        /// </summary>
        public readonly ImmutableArray<string> Tags;

        [OutputConstructor]
        private VolumeFull(
            string? createdAt,

            string? description,

            ImmutableArray<int> dropletIds,

            string? filesystemLabel,

            string? filesystemType,

            string? id,

            string? name,

            Outputs.VolumeFullPropertiesRegion? region,

            int? sizeGigabytes,

            ImmutableArray<string> tags)
        {
            CreatedAt = createdAt;
            Description = description;
            DropletIds = dropletIds;
            FilesystemLabel = filesystemLabel;
            FilesystemType = filesystemType;
            Id = id;
            Name = name;
            Region = region;
            SizeGigabytes = sizeGigabytes;
            Tags = tags;
        }
    }
}