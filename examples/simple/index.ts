import * as pulumi from "@pulumi/pulumi";
import { droplets } from "@cloudyskysoftware/pulumi-digitalocean-native";

const droplet = new droplets.v2.Droplet("my-vm", {
  name: "my-vm",
  image: "ubuntu-20-04-x64",
  size: "s-1vcpu-1gb",
});
