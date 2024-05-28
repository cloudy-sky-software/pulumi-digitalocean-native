# coding=utf-8
# *** WARNING: this file was generated by pulumigen. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from ... import _utilities

__all__ = ['KubernetesRunClusterLintArgs', 'KubernetesRunClusterLint']

@pulumi.input_type
class KubernetesRunClusterLintArgs:
    def __init__(__self__, *,
                 cluster_id: Optional[pulumi.Input[str]] = None,
                 exclude_checks: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 exclude_groups: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 include_checks: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 include_groups: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None):
        """
        The set of arguments for constructing a KubernetesRunClusterLint resource.
        :param pulumi.Input[str] cluster_id: A unique ID that can be used to reference a Kubernetes cluster.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] exclude_checks: An array of checks that will be run when clusterlint executes checks.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] exclude_groups: An array of check groups that will be omitted when clusterlint executes checks.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] include_checks: An array of checks that will be run when clusterlint executes checks.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] include_groups: An array of check groups that will be run when clusterlint executes checks.
        """
        if cluster_id is not None:
            pulumi.set(__self__, "cluster_id", cluster_id)
        if exclude_checks is not None:
            pulumi.set(__self__, "exclude_checks", exclude_checks)
        if exclude_groups is not None:
            pulumi.set(__self__, "exclude_groups", exclude_groups)
        if include_checks is not None:
            pulumi.set(__self__, "include_checks", include_checks)
        if include_groups is not None:
            pulumi.set(__self__, "include_groups", include_groups)

    @property
    @pulumi.getter(name="clusterId")
    def cluster_id(self) -> Optional[pulumi.Input[str]]:
        """
        A unique ID that can be used to reference a Kubernetes cluster.
        """
        return pulumi.get(self, "cluster_id")

    @cluster_id.setter
    def cluster_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "cluster_id", value)

    @property
    @pulumi.getter(name="excludeChecks")
    def exclude_checks(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        An array of checks that will be run when clusterlint executes checks.
        """
        return pulumi.get(self, "exclude_checks")

    @exclude_checks.setter
    def exclude_checks(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "exclude_checks", value)

    @property
    @pulumi.getter(name="excludeGroups")
    def exclude_groups(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        An array of check groups that will be omitted when clusterlint executes checks.
        """
        return pulumi.get(self, "exclude_groups")

    @exclude_groups.setter
    def exclude_groups(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "exclude_groups", value)

    @property
    @pulumi.getter(name="includeChecks")
    def include_checks(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        An array of checks that will be run when clusterlint executes checks.
        """
        return pulumi.get(self, "include_checks")

    @include_checks.setter
    def include_checks(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "include_checks", value)

    @property
    @pulumi.getter(name="includeGroups")
    def include_groups(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        An array of check groups that will be run when clusterlint executes checks.
        """
        return pulumi.get(self, "include_groups")

    @include_groups.setter
    def include_groups(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "include_groups", value)


class KubernetesRunClusterLint(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 cluster_id: Optional[pulumi.Input[str]] = None,
                 exclude_checks: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 exclude_groups: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 include_checks: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 include_groups: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 __props__=None):
        """
        Create a KubernetesRunClusterLint resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] cluster_id: A unique ID that can be used to reference a Kubernetes cluster.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] exclude_checks: An array of checks that will be run when clusterlint executes checks.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] exclude_groups: An array of check groups that will be omitted when clusterlint executes checks.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] include_checks: An array of checks that will be run when clusterlint executes checks.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] include_groups: An array of check groups that will be run when clusterlint executes checks.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: Optional[KubernetesRunClusterLintArgs] = None,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Create a KubernetesRunClusterLint resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param KubernetesRunClusterLintArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(KubernetesRunClusterLintArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 cluster_id: Optional[pulumi.Input[str]] = None,
                 exclude_checks: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 exclude_groups: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 include_checks: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 include_groups: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = KubernetesRunClusterLintArgs.__new__(KubernetesRunClusterLintArgs)

            __props__.__dict__["cluster_id"] = cluster_id
            __props__.__dict__["exclude_checks"] = exclude_checks
            __props__.__dict__["exclude_groups"] = exclude_groups
            __props__.__dict__["include_checks"] = include_checks
            __props__.__dict__["include_groups"] = include_groups
            __props__.__dict__["run_id"] = None
        super(KubernetesRunClusterLint, __self__).__init__(
            'digitalocean-native:kubernetes/v2:KubernetesRunClusterLint',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'KubernetesRunClusterLint':
        """
        Get an existing KubernetesRunClusterLint resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = KubernetesRunClusterLintArgs.__new__(KubernetesRunClusterLintArgs)

        __props__.__dict__["exclude_checks"] = None
        __props__.__dict__["exclude_groups"] = None
        __props__.__dict__["include_checks"] = None
        __props__.__dict__["include_groups"] = None
        __props__.__dict__["run_id"] = None
        return KubernetesRunClusterLint(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="excludeChecks")
    def exclude_checks(self) -> pulumi.Output[Optional[Sequence[str]]]:
        """
        An array of checks that will be run when clusterlint executes checks.
        """
        return pulumi.get(self, "exclude_checks")

    @property
    @pulumi.getter(name="excludeGroups")
    def exclude_groups(self) -> pulumi.Output[Optional[Sequence[str]]]:
        """
        An array of check groups that will be omitted when clusterlint executes checks.
        """
        return pulumi.get(self, "exclude_groups")

    @property
    @pulumi.getter(name="includeChecks")
    def include_checks(self) -> pulumi.Output[Optional[Sequence[str]]]:
        """
        An array of checks that will be run when clusterlint executes checks.
        """
        return pulumi.get(self, "include_checks")

    @property
    @pulumi.getter(name="includeGroups")
    def include_groups(self) -> pulumi.Output[Optional[Sequence[str]]]:
        """
        An array of check groups that will be run when clusterlint executes checks.
        """
        return pulumi.get(self, "include_groups")

    @property
    @pulumi.getter(name="runId")
    def run_id(self) -> pulumi.Output[Optional[str]]:
        """
        ID of the clusterlint run that can be used later to fetch the diagnostics.
        """
        return pulumi.get(self, "run_id")
