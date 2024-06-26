// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace CloudySkySoftware.Pulumi.DigitalOceanNative.AppsV2
{
    [DigitalOceanNativeResourceType("digitalocean-native:apps/v2:AppsValidateAppSpec")]
    public partial class AppsValidateAppSpec : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The monthly cost of the proposed app in USD using the next pricing plan tier. For example, if you propose an app that uses the Basic tier, the `app_tier_upgrade_cost` field displays the monthly cost of the app if it were to use the Professional tier. If the proposed app already uses the most expensive tier, the field is empty.
        /// </summary>
        [Output("appCost")]
        public Output<int?> AppCost { get; private set; } = null!;

        /// <summary>
        /// An optional ID of an existing app. If set, the spec will be treated as a proposed update to the specified app. The existing app is not modified using this method.
        /// </summary>
        [Output("appId")]
        public Output<string?> AppId { get; private set; } = null!;

        /// <summary>
        /// Indicates whether the app is a static app.
        /// </summary>
        [Output("appIsStatic")]
        public Output<bool?> AppIsStatic { get; private set; } = null!;

        /// <summary>
        /// Indicates whether the app name is available.
        /// </summary>
        [Output("appNameAvailable")]
        public Output<bool?> AppNameAvailable { get; private set; } = null!;

        /// <summary>
        /// The suggested name if the proposed app name is unavailable.
        /// </summary>
        [Output("appNameSuggestion")]
        public Output<string?> AppNameSuggestion { get; private set; } = null!;

        /// <summary>
        /// The monthly cost of the proposed app in USD using the previous pricing plan tier. For example, if you propose an app that uses the Professional tier, the `app_tier_downgrade_cost` field displays the monthly cost of the app if it were to use the Basic tier. If the proposed app already uses the lest expensive tier, the field is empty.
        /// </summary>
        [Output("appTierDowngradeCost")]
        public Output<int?> AppTierDowngradeCost { get; private set; } = null!;

        /// <summary>
        /// The maximum number of free static apps the account can have. We will charge you for any additional static apps.
        /// </summary>
        [Output("existingStaticApps")]
        public Output<string?> ExistingStaticApps { get; private set; } = null!;

        /// <summary>
        /// The desired configuration of an application.
        /// </summary>
        [Output("spec")]
        public Output<Outputs.AppSpec> Spec { get; private set; } = null!;


        /// <summary>
        /// Create a AppsValidateAppSpec resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public AppsValidateAppSpec(string name, AppsValidateAppSpecArgs args, CustomResourceOptions? options = null)
            : base("digitalocean-native:apps/v2:AppsValidateAppSpec", name, args ?? new AppsValidateAppSpecArgs(), MakeResourceOptions(options, ""))
        {
        }

        private AppsValidateAppSpec(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("digitalocean-native:apps/v2:AppsValidateAppSpec", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                PluginDownloadURL = "github://api.github.com/cloudy-sky-software/pulumi-digitalocean-native",
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing AppsValidateAppSpec resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static AppsValidateAppSpec Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new AppsValidateAppSpec(name, id, options);
        }
    }

    public sealed class AppsValidateAppSpecArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// An optional ID of an existing app. If set, the spec will be treated as a proposed update to the specified app. The existing app is not modified using this method.
        /// </summary>
        [Input("appId")]
        public Input<string>? AppId { get; set; }

        /// <summary>
        /// The desired configuration of an application.
        /// </summary>
        [Input("spec", required: true)]
        public Input<Inputs.AppSpecArgs> Spec { get; set; } = null!;

        public AppsValidateAppSpecArgs()
        {
        }
        public static new AppsValidateAppSpecArgs Empty => new AppsValidateAppSpecArgs();
    }
}
