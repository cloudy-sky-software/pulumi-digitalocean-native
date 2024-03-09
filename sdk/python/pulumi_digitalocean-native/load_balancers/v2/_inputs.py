# coding=utf-8
# *** WARNING: this file was generated by pulumigen. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from ... import _utilities
from ._enums import *

__all__ = [
    'ForwardingRuleArgs',
]

@pulumi.input_type
class ForwardingRuleArgs:
    def __init__(__self__, *,
                 entry_port: pulumi.Input[int],
                 entry_protocol: pulumi.Input['ForwardingRuleEntryProtocol'],
                 target_port: pulumi.Input[int],
                 target_protocol: pulumi.Input['ForwardingRuleTargetProtocol'],
                 certificate_id: Optional[pulumi.Input[str]] = None,
                 tls_passthrough: Optional[pulumi.Input[bool]] = None):
        """
        An object specifying a forwarding rule for a load balancer.
        :param pulumi.Input[int] entry_port: An integer representing the port on which the load balancer instance will listen.
        :param pulumi.Input['ForwardingRuleEntryProtocol'] entry_protocol: The protocol used for traffic to the load balancer. The possible values are: `http`, `https`, `http2`, `http3`, `tcp`, or `udp`. If you set the  `entry_protocol` to `udp`, the `target_protocol` must be set to `udp`.  When using UDP, the load balancer requires that you set up a health  check with a port that uses TCP, HTTP, or HTTPS to work properly.
        :param pulumi.Input[int] target_port: An integer representing the port on the backend Droplets to which the load balancer will send traffic.
        :param pulumi.Input['ForwardingRuleTargetProtocol'] target_protocol: The protocol used for traffic from the load balancer to the backend Droplets. The possible values are: `http`, `https`, `http2`, `tcp`, or `udp`. If you set the `target_protocol` to `udp`, the `entry_protocol` must be set to  `udp`. When using UDP, the load balancer requires that you set up a health  check with a port that uses TCP, HTTP, or HTTPS to work properly.
        :param pulumi.Input[str] certificate_id: The ID of the TLS certificate used for SSL termination if enabled.
        :param pulumi.Input[bool] tls_passthrough: A boolean value indicating whether SSL encrypted traffic will be passed through to the backend Droplets.
        """
        pulumi.set(__self__, "entry_port", entry_port)
        pulumi.set(__self__, "entry_protocol", entry_protocol)
        pulumi.set(__self__, "target_port", target_port)
        pulumi.set(__self__, "target_protocol", target_protocol)
        if certificate_id is not None:
            pulumi.set(__self__, "certificate_id", certificate_id)
        if tls_passthrough is not None:
            pulumi.set(__self__, "tls_passthrough", tls_passthrough)

    @property
    @pulumi.getter(name="entryPort")
    def entry_port(self) -> pulumi.Input[int]:
        """
        An integer representing the port on which the load balancer instance will listen.
        """
        return pulumi.get(self, "entry_port")

    @entry_port.setter
    def entry_port(self, value: pulumi.Input[int]):
        pulumi.set(self, "entry_port", value)

    @property
    @pulumi.getter(name="entryProtocol")
    def entry_protocol(self) -> pulumi.Input['ForwardingRuleEntryProtocol']:
        """
        The protocol used for traffic to the load balancer. The possible values are: `http`, `https`, `http2`, `http3`, `tcp`, or `udp`. If you set the  `entry_protocol` to `udp`, the `target_protocol` must be set to `udp`.  When using UDP, the load balancer requires that you set up a health  check with a port that uses TCP, HTTP, or HTTPS to work properly.
        """
        return pulumi.get(self, "entry_protocol")

    @entry_protocol.setter
    def entry_protocol(self, value: pulumi.Input['ForwardingRuleEntryProtocol']):
        pulumi.set(self, "entry_protocol", value)

    @property
    @pulumi.getter(name="targetPort")
    def target_port(self) -> pulumi.Input[int]:
        """
        An integer representing the port on the backend Droplets to which the load balancer will send traffic.
        """
        return pulumi.get(self, "target_port")

    @target_port.setter
    def target_port(self, value: pulumi.Input[int]):
        pulumi.set(self, "target_port", value)

    @property
    @pulumi.getter(name="targetProtocol")
    def target_protocol(self) -> pulumi.Input['ForwardingRuleTargetProtocol']:
        """
        The protocol used for traffic from the load balancer to the backend Droplets. The possible values are: `http`, `https`, `http2`, `tcp`, or `udp`. If you set the `target_protocol` to `udp`, the `entry_protocol` must be set to  `udp`. When using UDP, the load balancer requires that you set up a health  check with a port that uses TCP, HTTP, or HTTPS to work properly.
        """
        return pulumi.get(self, "target_protocol")

    @target_protocol.setter
    def target_protocol(self, value: pulumi.Input['ForwardingRuleTargetProtocol']):
        pulumi.set(self, "target_protocol", value)

    @property
    @pulumi.getter(name="certificateId")
    def certificate_id(self) -> Optional[pulumi.Input[str]]:
        """
        The ID of the TLS certificate used for SSL termination if enabled.
        """
        return pulumi.get(self, "certificate_id")

    @certificate_id.setter
    def certificate_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "certificate_id", value)

    @property
    @pulumi.getter(name="tlsPassthrough")
    def tls_passthrough(self) -> Optional[pulumi.Input[bool]]:
        """
        A boolean value indicating whether SSL encrypted traffic will be passed through to the backend Droplets.
        """
        return pulumi.get(self, "tls_passthrough")

    @tls_passthrough.setter
    def tls_passthrough(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "tls_passthrough", value)

