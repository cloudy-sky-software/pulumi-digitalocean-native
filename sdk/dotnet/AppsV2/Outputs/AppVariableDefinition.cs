// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.DigitalOceanNative.AppsV2.Outputs
{

    [OutputType]
    public sealed class AppVariableDefinition
    {
        /// <summary>
        /// The variable name
        /// </summary>
        public readonly string Key;
        /// <summary>
        /// - RUN_TIME: Made available only at run-time
        /// - BUILD_TIME: Made available only at build-time
        /// - RUN_AND_BUILD_TIME: Made available at both build and run-time
        /// </summary>
        public readonly Pulumi.DigitalOceanNative.AppsV2.AppVariableDefinitionScope? Scope;
        /// <summary>
        /// - GENERAL: A plain-text environment variable
        /// - SECRET: A secret encrypted environment variable
        /// </summary>
        public readonly Pulumi.DigitalOceanNative.AppsV2.AppVariableDefinitionType? Type;
        /// <summary>
        /// The value. If the type is `SECRET`, the value will be encrypted on first submission. On following submissions, the encrypted value should be used.
        /// </summary>
        public readonly string? Value;

        [OutputConstructor]
        private AppVariableDefinition(
            string key,

            Pulumi.DigitalOceanNative.AppsV2.AppVariableDefinitionScope? scope,

            Pulumi.DigitalOceanNative.AppsV2.AppVariableDefinitionType? type,

            string? value)
        {
            Key = key;
            Scope = scope;
            Type = type;
            Value = value;
        }
    }
}
