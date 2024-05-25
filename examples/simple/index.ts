import * as pulumi from "@pulumi/pulumi";
import { droplets } from "@cloudyskysoftware/pulumi-digitalocean-native";

const droplet = new droplets.v2.Droplets("my-vm", {
  names: ["my-vm"],
  image: "ubuntu-20-04-x64",
  size: "s-1vcpu-1gb",
});
