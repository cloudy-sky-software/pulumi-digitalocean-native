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
    public sealed class Size
    {
        /// <summary>
        /// This is a boolean value that represents whether new Droplets can be created with this size.
        /// </summary>
        public readonly bool Available;
        /// <summary>
        /// A string describing the class of Droplets created from this size. For example: Basic, General Purpose, CPU-Optimized, Memory-Optimized, or Storage-Optimized.
        /// </summary>
        public readonly string Description;
        /// <summary>
        /// The amount of disk space set aside for Droplets of this size. The value is represented in gigabytes.
        /// </summary>
        public readonly int Disk;
        /// <summary>
        /// The amount of RAM allocated to Droplets created of this size. The value is represented in megabytes.
        /// </summary>
        public readonly int Memory;
        /// <summary>
        /// This describes the price of the Droplet size as measured hourly. The value is measured in US dollars.
        /// </summary>
        public readonly double PriceHourly;
        /// <summary>
        /// This attribute describes the monthly cost of this Droplet size if the Droplet is kept for an entire month. The value is measured in US dollars.
        /// </summary>
        public readonly double PriceMonthly;
        /// <summary>
        /// An array containing the region slugs where this size is available for Droplet creates.
        /// </summary>
        public readonly ImmutableArray<string> Regions;
        /// <summary>
        /// A human-readable string that is used to uniquely identify each size.
        /// </summary>
        public readonly string Slug;
        /// <summary>
        /// The amount of transfer bandwidth that is available for Droplets created in this size. This only counts traffic on the public interface. The value is given in terabytes.
        /// </summary>
        public readonly double Transfer;
        /// <summary>
        /// The integer of number CPUs allocated to Droplets of this size.
        /// </summary>
        public readonly int Vcpus;

        [OutputConstructor]
        private Size(
            bool available,

            string description,

            int disk,

            int memory,

            double priceHourly,

            double priceMonthly,

            ImmutableArray<string> regions,

            string slug,

            double transfer,

            int vcpus)
        {
            Available = available;
            Description = description;
            Disk = disk;
            Memory = memory;
            PriceHourly = priceHourly;
            PriceMonthly = priceMonthly;
            Regions = regions;
            Slug = slug;
            Transfer = transfer;
            Vcpus = vcpus;
        }
    }
}
