# coding=utf-8
# *** WARNING: this file was generated by pulumigen. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from ... import _utilities
from . import outputs
from ._enums import *
from ._inputs import *

__all__ = ['LoadBalancerArgs', 'LoadBalancer']

@pulumi.input_type
class LoadBalancerArgs:
    def __init__(__self__, *,
                 forwarding_rules: pulumi.Input[Sequence[pulumi.Input['ForwardingRuleArgs']]],
                 algorithm: Optional[pulumi.Input['LoadBalancerBaseAlgorithm']] = None,
                 created_at: Optional[pulumi.Input[str]] = None,
                 disable_lets_encrypt_dns_records: Optional[pulumi.Input[bool]] = None,
                 enable_backend_keepalive: Optional[pulumi.Input[bool]] = None,
                 enable_proxy_protocol: Optional[pulumi.Input[bool]] = None,
                 firewall: Optional[pulumi.Input['LbFirewallArgs']] = None,
                 health_check: Optional[pulumi.Input['HealthCheckArgs']] = None,
                 http_idle_timeout_seconds: Optional[pulumi.Input[int]] = None,
                 ip: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 project_id: Optional[pulumi.Input[str]] = None,
                 redirect_http_to_https: Optional[pulumi.Input[bool]] = None,
                 region: Optional[pulumi.Input['LoadBalancerPropertiesRegionEnum']] = None,
                 size: Optional[pulumi.Input['LoadBalancerBaseSize']] = None,
                 size_unit: Optional[pulumi.Input[int]] = None,
                 status: Optional[pulumi.Input['LoadBalancerBaseStatus']] = None,
                 sticky_sessions: Optional[pulumi.Input['StickySessionsArgs']] = None,
                 vpc_uuid: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a LoadBalancer resource.
        :param pulumi.Input[Sequence[pulumi.Input['ForwardingRuleArgs']]] forwarding_rules: An array of objects specifying the forwarding rules for a load balancer.
        :param pulumi.Input['LoadBalancerBaseAlgorithm'] algorithm: This field has been deprecated. You can no longer specify an algorithm for load balancers.
        :param pulumi.Input[str] created_at: A time value given in ISO8601 combined date and time format that represents when the load balancer was created.
        :param pulumi.Input[bool] disable_lets_encrypt_dns_records: A boolean value indicating whether to disable automatic DNS record creation for Let's Encrypt certificates that are added to the load balancer.
        :param pulumi.Input[bool] enable_backend_keepalive: A boolean value indicating whether HTTP keepalive connections are maintained to target Droplets.
        :param pulumi.Input[bool] enable_proxy_protocol: A boolean value indicating whether PROXY Protocol is in use.
        :param pulumi.Input['LbFirewallArgs'] firewall: An object specifying allow and deny rules to control traffic to the load balancer.
        :param pulumi.Input['HealthCheckArgs'] health_check: An object specifying health check settings for the load balancer.
        :param pulumi.Input[int] http_idle_timeout_seconds: An integer value which configures the idle timeout for HTTP requests to the target droplets.
        :param pulumi.Input[str] ip: An attribute containing the public-facing IP address of the load balancer.
        :param pulumi.Input[str] name: A human-readable name for a load balancer instance.
        :param pulumi.Input[str] project_id: The ID of the project that the load balancer is associated with. If no ID is provided at creation, the load balancer associates with the user's default project. If an invalid project ID is provided, the load balancer will not be created.
        :param pulumi.Input[bool] redirect_http_to_https: A boolean value indicating whether HTTP requests to the load balancer on port 80 will be redirected to HTTPS on port 443.
        :param pulumi.Input['LoadBalancerPropertiesRegionEnum'] region: The slug identifier for the region where the resource will initially be  available.
        :param pulumi.Input['LoadBalancerBaseSize'] size: This field has been replaced by the `size_unit` field for all regions except in AMS2, NYC2, and SFO1. Each available load balancer size now equates to the load balancer having a set number of nodes.
               * `lb-small` = 1 node
               * `lb-medium` = 3 nodes
               * `lb-large` = 6 nodes
               
               You can resize load balancers after creation up to once per hour. You cannot resize a load balancer within the first hour of its creation.
        :param pulumi.Input[int] size_unit: How many nodes the load balancer contains. Each additional node increases the load balancer's ability to manage more connections. Load balancers can be scaled up or down, and you can change the number of nodes after creation up to once per hour. This field is currently not available in the AMS2, NYC2, or SFO1 regions. Use the `size` field to scale load balancers that reside in these regions.
        :param pulumi.Input['LoadBalancerBaseStatus'] status: A status string indicating the current state of the load balancer. This can be `new`, `active`, or `errored`.
        :param pulumi.Input['StickySessionsArgs'] sticky_sessions: An object specifying sticky sessions settings for the load balancer.
        :param pulumi.Input[str] vpc_uuid: A string specifying the UUID of the VPC to which the load balancer is assigned.
        """
        pulumi.set(__self__, "forwarding_rules", forwarding_rules)
        if algorithm is None:
            algorithm = 'round_robin'
        if algorithm is not None:
            pulumi.set(__self__, "algorithm", algorithm)
        if created_at is not None:
            pulumi.set(__self__, "created_at", created_at)
        if disable_lets_encrypt_dns_records is None:
            disable_lets_encrypt_dns_records = False
        if disable_lets_encrypt_dns_records is not None:
            pulumi.set(__self__, "disable_lets_encrypt_dns_records", disable_lets_encrypt_dns_records)
        if enable_backend_keepalive is None:
            enable_backend_keepalive = False
        if enable_backend_keepalive is not None:
            pulumi.set(__self__, "enable_backend_keepalive", enable_backend_keepalive)
        if enable_proxy_protocol is None:
            enable_proxy_protocol = False
        if enable_proxy_protocol is not None:
            pulumi.set(__self__, "enable_proxy_protocol", enable_proxy_protocol)
        if firewall is not None:
            pulumi.set(__self__, "firewall", firewall)
        if health_check is not None:
            pulumi.set(__self__, "health_check", health_check)
        if http_idle_timeout_seconds is None:
            http_idle_timeout_seconds = 60
        if http_idle_timeout_seconds is not None:
            pulumi.set(__self__, "http_idle_timeout_seconds", http_idle_timeout_seconds)
        if ip is not None:
            pulumi.set(__self__, "ip", ip)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if project_id is not None:
            pulumi.set(__self__, "project_id", project_id)
        if redirect_http_to_https is None:
            redirect_http_to_https = False
        if redirect_http_to_https is not None:
            pulumi.set(__self__, "redirect_http_to_https", redirect_http_to_https)
        if region is not None:
            pulumi.set(__self__, "region", region)
        if size is None:
            size = 'lb-small'
        if size is not None:
            pulumi.set(__self__, "size", size)
        if size_unit is None:
            size_unit = 1
        if size_unit is not None:
            pulumi.set(__self__, "size_unit", size_unit)
        if status is not None:
            pulumi.set(__self__, "status", status)
        if sticky_sessions is not None:
            pulumi.set(__self__, "sticky_sessions", sticky_sessions)
        if vpc_uuid is not None:
            pulumi.set(__self__, "vpc_uuid", vpc_uuid)

    @property
    @pulumi.getter(name="forwardingRules")
    def forwarding_rules(self) -> pulumi.Input[Sequence[pulumi.Input['ForwardingRuleArgs']]]:
        """
        An array of objects specifying the forwarding rules for a load balancer.
        """
        return pulumi.get(self, "forwarding_rules")

    @forwarding_rules.setter
    def forwarding_rules(self, value: pulumi.Input[Sequence[pulumi.Input['ForwardingRuleArgs']]]):
        pulumi.set(self, "forwarding_rules", value)

    @property
    @pulumi.getter
    def algorithm(self) -> Optional[pulumi.Input['LoadBalancerBaseAlgorithm']]:
        """
        This field has been deprecated. You can no longer specify an algorithm for load balancers.
        """
        return pulumi.get(self, "algorithm")

    @algorithm.setter
    def algorithm(self, value: Optional[pulumi.Input['LoadBalancerBaseAlgorithm']]):
        pulumi.set(self, "algorithm", value)

    @property
    @pulumi.getter(name="createdAt")
    def created_at(self) -> Optional[pulumi.Input[str]]:
        """
        A time value given in ISO8601 combined date and time format that represents when the load balancer was created.
        """
        return pulumi.get(self, "created_at")

    @created_at.setter
    def created_at(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "created_at", value)

    @property
    @pulumi.getter(name="disableLetsEncryptDnsRecords")
    def disable_lets_encrypt_dns_records(self) -> Optional[pulumi.Input[bool]]:
        """
        A boolean value indicating whether to disable automatic DNS record creation for Let's Encrypt certificates that are added to the load balancer.
        """
        return pulumi.get(self, "disable_lets_encrypt_dns_records")

    @disable_lets_encrypt_dns_records.setter
    def disable_lets_encrypt_dns_records(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "disable_lets_encrypt_dns_records", value)

    @property
    @pulumi.getter(name="enableBackendKeepalive")
    def enable_backend_keepalive(self) -> Optional[pulumi.Input[bool]]:
        """
        A boolean value indicating whether HTTP keepalive connections are maintained to target Droplets.
        """
        return pulumi.get(self, "enable_backend_keepalive")

    @enable_backend_keepalive.setter
    def enable_backend_keepalive(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "enable_backend_keepalive", value)

    @property
    @pulumi.getter(name="enableProxyProtocol")
    def enable_proxy_protocol(self) -> Optional[pulumi.Input[bool]]:
        """
        A boolean value indicating whether PROXY Protocol is in use.
        """
        return pulumi.get(self, "enable_proxy_protocol")

    @enable_proxy_protocol.setter
    def enable_proxy_protocol(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "enable_proxy_protocol", value)

    @property
    @pulumi.getter
    def firewall(self) -> Optional[pulumi.Input['LbFirewallArgs']]:
        """
        An object specifying allow and deny rules to control traffic to the load balancer.
        """
        return pulumi.get(self, "firewall")

    @firewall.setter
    def firewall(self, value: Optional[pulumi.Input['LbFirewallArgs']]):
        pulumi.set(self, "firewall", value)

    @property
    @pulumi.getter(name="healthCheck")
    def health_check(self) -> Optional[pulumi.Input['HealthCheckArgs']]:
        """
        An object specifying health check settings for the load balancer.
        """
        return pulumi.get(self, "health_check")

    @health_check.setter
    def health_check(self, value: Optional[pulumi.Input['HealthCheckArgs']]):
        pulumi.set(self, "health_check", value)

    @property
    @pulumi.getter(name="httpIdleTimeoutSeconds")
    def http_idle_timeout_seconds(self) -> Optional[pulumi.Input[int]]:
        """
        An integer value which configures the idle timeout for HTTP requests to the target droplets.
        """
        return pulumi.get(self, "http_idle_timeout_seconds")

    @http_idle_timeout_seconds.setter
    def http_idle_timeout_seconds(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "http_idle_timeout_seconds", value)

    @property
    @pulumi.getter
    def ip(self) -> Optional[pulumi.Input[str]]:
        """
        An attribute containing the public-facing IP address of the load balancer.
        """
        return pulumi.get(self, "ip")

    @ip.setter
    def ip(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "ip", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        A human-readable name for a load balancer instance.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="projectId")
    def project_id(self) -> Optional[pulumi.Input[str]]:
        """
        The ID of the project that the load balancer is associated with. If no ID is provided at creation, the load balancer associates with the user's default project. If an invalid project ID is provided, the load balancer will not be created.
        """
        return pulumi.get(self, "project_id")

    @project_id.setter
    def project_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "project_id", value)

    @property
    @pulumi.getter(name="redirectHttpToHttps")
    def redirect_http_to_https(self) -> Optional[pulumi.Input[bool]]:
        """
        A boolean value indicating whether HTTP requests to the load balancer on port 80 will be redirected to HTTPS on port 443.
        """
        return pulumi.get(self, "redirect_http_to_https")

    @redirect_http_to_https.setter
    def redirect_http_to_https(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "redirect_http_to_https", value)

    @property
    @pulumi.getter
    def region(self) -> Optional[pulumi.Input['LoadBalancerPropertiesRegionEnum']]:
        """
        The slug identifier for the region where the resource will initially be  available.
        """
        return pulumi.get(self, "region")

    @region.setter
    def region(self, value: Optional[pulumi.Input['LoadBalancerPropertiesRegionEnum']]):
        pulumi.set(self, "region", value)

    @property
    @pulumi.getter
    def size(self) -> Optional[pulumi.Input['LoadBalancerBaseSize']]:
        """
        This field has been replaced by the `size_unit` field for all regions except in AMS2, NYC2, and SFO1. Each available load balancer size now equates to the load balancer having a set number of nodes.
        * `lb-small` = 1 node
        * `lb-medium` = 3 nodes
        * `lb-large` = 6 nodes

        You can resize load balancers after creation up to once per hour. You cannot resize a load balancer within the first hour of its creation.
        """
        return pulumi.get(self, "size")

    @size.setter
    def size(self, value: Optional[pulumi.Input['LoadBalancerBaseSize']]):
        pulumi.set(self, "size", value)

    @property
    @pulumi.getter(name="sizeUnit")
    def size_unit(self) -> Optional[pulumi.Input[int]]:
        """
        How many nodes the load balancer contains. Each additional node increases the load balancer's ability to manage more connections. Load balancers can be scaled up or down, and you can change the number of nodes after creation up to once per hour. This field is currently not available in the AMS2, NYC2, or SFO1 regions. Use the `size` field to scale load balancers that reside in these regions.
        """
        return pulumi.get(self, "size_unit")

    @size_unit.setter
    def size_unit(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "size_unit", value)

    @property
    @pulumi.getter
    def status(self) -> Optional[pulumi.Input['LoadBalancerBaseStatus']]:
        """
        A status string indicating the current state of the load balancer. This can be `new`, `active`, or `errored`.
        """
        return pulumi.get(self, "status")

    @status.setter
    def status(self, value: Optional[pulumi.Input['LoadBalancerBaseStatus']]):
        pulumi.set(self, "status", value)

    @property
    @pulumi.getter(name="stickySessions")
    def sticky_sessions(self) -> Optional[pulumi.Input['StickySessionsArgs']]:
        """
        An object specifying sticky sessions settings for the load balancer.
        """
        return pulumi.get(self, "sticky_sessions")

    @sticky_sessions.setter
    def sticky_sessions(self, value: Optional[pulumi.Input['StickySessionsArgs']]):
        pulumi.set(self, "sticky_sessions", value)

    @property
    @pulumi.getter(name="vpcUuid")
    def vpc_uuid(self) -> Optional[pulumi.Input[str]]:
        """
        A string specifying the UUID of the VPC to which the load balancer is assigned.
        """
        return pulumi.get(self, "vpc_uuid")

    @vpc_uuid.setter
    def vpc_uuid(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "vpc_uuid", value)


class LoadBalancer(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 algorithm: Optional[pulumi.Input['LoadBalancerBaseAlgorithm']] = None,
                 created_at: Optional[pulumi.Input[str]] = None,
                 disable_lets_encrypt_dns_records: Optional[pulumi.Input[bool]] = None,
                 enable_backend_keepalive: Optional[pulumi.Input[bool]] = None,
                 enable_proxy_protocol: Optional[pulumi.Input[bool]] = None,
                 firewall: Optional[pulumi.Input[pulumi.InputType['LbFirewallArgs']]] = None,
                 forwarding_rules: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['ForwardingRuleArgs']]]]] = None,
                 health_check: Optional[pulumi.Input[pulumi.InputType['HealthCheckArgs']]] = None,
                 http_idle_timeout_seconds: Optional[pulumi.Input[int]] = None,
                 ip: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 project_id: Optional[pulumi.Input[str]] = None,
                 redirect_http_to_https: Optional[pulumi.Input[bool]] = None,
                 region: Optional[pulumi.Input['LoadBalancerPropertiesRegionEnum']] = None,
                 size: Optional[pulumi.Input['LoadBalancerBaseSize']] = None,
                 size_unit: Optional[pulumi.Input[int]] = None,
                 status: Optional[pulumi.Input['LoadBalancerBaseStatus']] = None,
                 sticky_sessions: Optional[pulumi.Input[pulumi.InputType['StickySessionsArgs']]] = None,
                 vpc_uuid: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Create a LoadBalancer resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input['LoadBalancerBaseAlgorithm'] algorithm: This field has been deprecated. You can no longer specify an algorithm for load balancers.
        :param pulumi.Input[str] created_at: A time value given in ISO8601 combined date and time format that represents when the load balancer was created.
        :param pulumi.Input[bool] disable_lets_encrypt_dns_records: A boolean value indicating whether to disable automatic DNS record creation for Let's Encrypt certificates that are added to the load balancer.
        :param pulumi.Input[bool] enable_backend_keepalive: A boolean value indicating whether HTTP keepalive connections are maintained to target Droplets.
        :param pulumi.Input[bool] enable_proxy_protocol: A boolean value indicating whether PROXY Protocol is in use.
        :param pulumi.Input[pulumi.InputType['LbFirewallArgs']] firewall: An object specifying allow and deny rules to control traffic to the load balancer.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['ForwardingRuleArgs']]]] forwarding_rules: An array of objects specifying the forwarding rules for a load balancer.
        :param pulumi.Input[pulumi.InputType['HealthCheckArgs']] health_check: An object specifying health check settings for the load balancer.
        :param pulumi.Input[int] http_idle_timeout_seconds: An integer value which configures the idle timeout for HTTP requests to the target droplets.
        :param pulumi.Input[str] ip: An attribute containing the public-facing IP address of the load balancer.
        :param pulumi.Input[str] name: A human-readable name for a load balancer instance.
        :param pulumi.Input[str] project_id: The ID of the project that the load balancer is associated with. If no ID is provided at creation, the load balancer associates with the user's default project. If an invalid project ID is provided, the load balancer will not be created.
        :param pulumi.Input[bool] redirect_http_to_https: A boolean value indicating whether HTTP requests to the load balancer on port 80 will be redirected to HTTPS on port 443.
        :param pulumi.Input['LoadBalancerPropertiesRegionEnum'] region: The slug identifier for the region where the resource will initially be  available.
        :param pulumi.Input['LoadBalancerBaseSize'] size: This field has been replaced by the `size_unit` field for all regions except in AMS2, NYC2, and SFO1. Each available load balancer size now equates to the load balancer having a set number of nodes.
               * `lb-small` = 1 node
               * `lb-medium` = 3 nodes
               * `lb-large` = 6 nodes
               
               You can resize load balancers after creation up to once per hour. You cannot resize a load balancer within the first hour of its creation.
        :param pulumi.Input[int] size_unit: How many nodes the load balancer contains. Each additional node increases the load balancer's ability to manage more connections. Load balancers can be scaled up or down, and you can change the number of nodes after creation up to once per hour. This field is currently not available in the AMS2, NYC2, or SFO1 regions. Use the `size` field to scale load balancers that reside in these regions.
        :param pulumi.Input['LoadBalancerBaseStatus'] status: A status string indicating the current state of the load balancer. This can be `new`, `active`, or `errored`.
        :param pulumi.Input[pulumi.InputType['StickySessionsArgs']] sticky_sessions: An object specifying sticky sessions settings for the load balancer.
        :param pulumi.Input[str] vpc_uuid: A string specifying the UUID of the VPC to which the load balancer is assigned.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: LoadBalancerArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Create a LoadBalancer resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param LoadBalancerArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(LoadBalancerArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 algorithm: Optional[pulumi.Input['LoadBalancerBaseAlgorithm']] = None,
                 created_at: Optional[pulumi.Input[str]] = None,
                 disable_lets_encrypt_dns_records: Optional[pulumi.Input[bool]] = None,
                 enable_backend_keepalive: Optional[pulumi.Input[bool]] = None,
                 enable_proxy_protocol: Optional[pulumi.Input[bool]] = None,
                 firewall: Optional[pulumi.Input[pulumi.InputType['LbFirewallArgs']]] = None,
                 forwarding_rules: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['ForwardingRuleArgs']]]]] = None,
                 health_check: Optional[pulumi.Input[pulumi.InputType['HealthCheckArgs']]] = None,
                 http_idle_timeout_seconds: Optional[pulumi.Input[int]] = None,
                 ip: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 project_id: Optional[pulumi.Input[str]] = None,
                 redirect_http_to_https: Optional[pulumi.Input[bool]] = None,
                 region: Optional[pulumi.Input['LoadBalancerPropertiesRegionEnum']] = None,
                 size: Optional[pulumi.Input['LoadBalancerBaseSize']] = None,
                 size_unit: Optional[pulumi.Input[int]] = None,
                 status: Optional[pulumi.Input['LoadBalancerBaseStatus']] = None,
                 sticky_sessions: Optional[pulumi.Input[pulumi.InputType['StickySessionsArgs']]] = None,
                 vpc_uuid: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = LoadBalancerArgs.__new__(LoadBalancerArgs)

            if algorithm is None:
                algorithm = 'round_robin'
            __props__.__dict__["algorithm"] = algorithm
            __props__.__dict__["created_at"] = created_at
            if disable_lets_encrypt_dns_records is None:
                disable_lets_encrypt_dns_records = False
            __props__.__dict__["disable_lets_encrypt_dns_records"] = disable_lets_encrypt_dns_records
            if enable_backend_keepalive is None:
                enable_backend_keepalive = False
            __props__.__dict__["enable_backend_keepalive"] = enable_backend_keepalive
            if enable_proxy_protocol is None:
                enable_proxy_protocol = False
            __props__.__dict__["enable_proxy_protocol"] = enable_proxy_protocol
            __props__.__dict__["firewall"] = firewall
            if forwarding_rules is None and not opts.urn:
                raise TypeError("Missing required property 'forwarding_rules'")
            __props__.__dict__["forwarding_rules"] = forwarding_rules
            __props__.__dict__["health_check"] = health_check
            if http_idle_timeout_seconds is None:
                http_idle_timeout_seconds = 60
            __props__.__dict__["http_idle_timeout_seconds"] = http_idle_timeout_seconds
            __props__.__dict__["ip"] = ip
            __props__.__dict__["name"] = name
            __props__.__dict__["project_id"] = project_id
            if redirect_http_to_https is None:
                redirect_http_to_https = False
            __props__.__dict__["redirect_http_to_https"] = redirect_http_to_https
            __props__.__dict__["region"] = region
            if size is None:
                size = 'lb-small'
            __props__.__dict__["size"] = size
            if size_unit is None:
                size_unit = 1
            __props__.__dict__["size_unit"] = size_unit
            __props__.__dict__["status"] = status
            __props__.__dict__["sticky_sessions"] = sticky_sessions
            __props__.__dict__["vpc_uuid"] = vpc_uuid
            __props__.__dict__["load_balancer"] = None
        super(LoadBalancer, __self__).__init__(
            'digitalocean-native:load_balancers/v2:LoadBalancer',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'LoadBalancer':
        """
        Get an existing LoadBalancer resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = LoadBalancerArgs.__new__(LoadBalancerArgs)

        __props__.__dict__["algorithm"] = None
        __props__.__dict__["created_at"] = None
        __props__.__dict__["disable_lets_encrypt_dns_records"] = None
        __props__.__dict__["enable_backend_keepalive"] = None
        __props__.__dict__["enable_proxy_protocol"] = None
        __props__.__dict__["firewall"] = None
        __props__.__dict__["forwarding_rules"] = None
        __props__.__dict__["health_check"] = None
        __props__.__dict__["http_idle_timeout_seconds"] = None
        __props__.__dict__["ip"] = None
        __props__.__dict__["load_balancer"] = None
        __props__.__dict__["name"] = None
        __props__.__dict__["project_id"] = None
        __props__.__dict__["redirect_http_to_https"] = None
        __props__.__dict__["region"] = None
        __props__.__dict__["size"] = None
        __props__.__dict__["size_unit"] = None
        __props__.__dict__["status"] = None
        __props__.__dict__["sticky_sessions"] = None
        __props__.__dict__["vpc_uuid"] = None
        return LoadBalancer(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def algorithm(self) -> pulumi.Output[Optional['LoadBalancerBaseAlgorithm']]:
        """
        This field has been deprecated. You can no longer specify an algorithm for load balancers.
        """
        return pulumi.get(self, "algorithm")

    @property
    @pulumi.getter(name="createdAt")
    def created_at(self) -> pulumi.Output[Optional[str]]:
        """
        A time value given in ISO8601 combined date and time format that represents when the load balancer was created.
        """
        return pulumi.get(self, "created_at")

    @property
    @pulumi.getter(name="disableLetsEncryptDnsRecords")
    def disable_lets_encrypt_dns_records(self) -> pulumi.Output[Optional[bool]]:
        """
        A boolean value indicating whether to disable automatic DNS record creation for Let's Encrypt certificates that are added to the load balancer.
        """
        return pulumi.get(self, "disable_lets_encrypt_dns_records")

    @property
    @pulumi.getter(name="enableBackendKeepalive")
    def enable_backend_keepalive(self) -> pulumi.Output[Optional[bool]]:
        """
        A boolean value indicating whether HTTP keepalive connections are maintained to target Droplets.
        """
        return pulumi.get(self, "enable_backend_keepalive")

    @property
    @pulumi.getter(name="enableProxyProtocol")
    def enable_proxy_protocol(self) -> pulumi.Output[Optional[bool]]:
        """
        A boolean value indicating whether PROXY Protocol is in use.
        """
        return pulumi.get(self, "enable_proxy_protocol")

    @property
    @pulumi.getter
    def firewall(self) -> pulumi.Output[Optional['outputs.LbFirewall']]:
        """
        An object specifying allow and deny rules to control traffic to the load balancer.
        """
        return pulumi.get(self, "firewall")

    @property
    @pulumi.getter(name="forwardingRules")
    def forwarding_rules(self) -> pulumi.Output[Optional[Sequence['outputs.ForwardingRule']]]:
        """
        An array of objects specifying the forwarding rules for a load balancer.
        """
        return pulumi.get(self, "forwarding_rules")

    @property
    @pulumi.getter(name="healthCheck")
    def health_check(self) -> pulumi.Output[Optional['outputs.HealthCheck']]:
        """
        An object specifying health check settings for the load balancer.
        """
        return pulumi.get(self, "health_check")

    @property
    @pulumi.getter(name="httpIdleTimeoutSeconds")
    def http_idle_timeout_seconds(self) -> pulumi.Output[Optional[int]]:
        """
        An integer value which configures the idle timeout for HTTP requests to the target droplets.
        """
        return pulumi.get(self, "http_idle_timeout_seconds")

    @property
    @pulumi.getter
    def ip(self) -> pulumi.Output[Optional[str]]:
        """
        An attribute containing the public-facing IP address of the load balancer.
        """
        return pulumi.get(self, "ip")

    @property
    @pulumi.getter(name="loadBalancer")
    def load_balancer(self) -> pulumi.Output[Optional['outputs.LoadBalancer']]:
        return pulumi.get(self, "load_balancer")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[Optional[str]]:
        """
        A human-readable name for a load balancer instance.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="projectId")
    def project_id(self) -> pulumi.Output[Optional[str]]:
        """
        The ID of the project that the load balancer is associated with. If no ID is provided at creation, the load balancer associates with the user's default project. If an invalid project ID is provided, the load balancer will not be created.
        """
        return pulumi.get(self, "project_id")

    @property
    @pulumi.getter(name="redirectHttpToHttps")
    def redirect_http_to_https(self) -> pulumi.Output[Optional[bool]]:
        """
        A boolean value indicating whether HTTP requests to the load balancer on port 80 will be redirected to HTTPS on port 443.
        """
        return pulumi.get(self, "redirect_http_to_https")

    @property
    @pulumi.getter
    def region(self) -> pulumi.Output[Optional['LoadBalancerPropertiesRegionEnum']]:
        """
        The slug identifier for the region where the resource will initially be  available.
        """
        return pulumi.get(self, "region")

    @property
    @pulumi.getter
    def size(self) -> pulumi.Output[Optional['LoadBalancerBaseSize']]:
        """
        This field has been replaced by the `size_unit` field for all regions except in AMS2, NYC2, and SFO1. Each available load balancer size now equates to the load balancer having a set number of nodes.
        * `lb-small` = 1 node
        * `lb-medium` = 3 nodes
        * `lb-large` = 6 nodes

        You can resize load balancers after creation up to once per hour. You cannot resize a load balancer within the first hour of its creation.
        """
        return pulumi.get(self, "size")

    @property
    @pulumi.getter(name="sizeUnit")
    def size_unit(self) -> pulumi.Output[Optional[int]]:
        """
        How many nodes the load balancer contains. Each additional node increases the load balancer's ability to manage more connections. Load balancers can be scaled up or down, and you can change the number of nodes after creation up to once per hour. This field is currently not available in the AMS2, NYC2, or SFO1 regions. Use the `size` field to scale load balancers that reside in these regions.
        """
        return pulumi.get(self, "size_unit")

    @property
    @pulumi.getter
    def status(self) -> pulumi.Output[Optional['LoadBalancerBaseStatus']]:
        """
        A status string indicating the current state of the load balancer. This can be `new`, `active`, or `errored`.
        """
        return pulumi.get(self, "status")

    @property
    @pulumi.getter(name="stickySessions")
    def sticky_sessions(self) -> pulumi.Output[Optional['outputs.StickySessions']]:
        """
        An object specifying sticky sessions settings for the load balancer.
        """
        return pulumi.get(self, "sticky_sessions")

    @property
    @pulumi.getter(name="vpcUuid")
    def vpc_uuid(self) -> pulumi.Output[Optional[str]]:
        """
        A string specifying the UUID of the VPC to which the load balancer is assigned.
        """
        return pulumi.get(self, "vpc_uuid")

