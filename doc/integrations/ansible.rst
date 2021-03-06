
.. _rs_ansible:

Dynamic Ansible Inventory using Profiles
~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

The following instructions show how to map Ansible Playbooks via
Digital Rebar with no Ansible specific configuration required.

The instructions are generic and could be adapted to run on any Ansible run.

Prereqs
-------

Before starting this process, a Digital Rebar Provision (DRP) server is required, along with the ability to provision machines.  These machines could be VMs, Packet servers or physical servers in a data center.  DRPCLI and Ansible must also be installed on the system.

Root ssh access to the systems is required for the script to work.  Make sure that the correct SSH keys have been installed on the target systems.  Review :ref:`rs_add_ssh` for details.

At this time, testing is on Centos 7 only using root as the login.  This documentation assumes provisioning has completed and the machines are ready for installation - there is no workflow automation to move from discovery or sledgehammer to the target o/s documented here.

Digital Rebar Provision Ansible Configuration
---------------------------------------------

The Integrations Ansible drpmachines.py script can be used to create a dynamic inventory from a Digital Rebar Endpoint.

The hosts inventory list defaults to all machines or can be restricted by setting a `ansible=[value]` Param.

Group membership managed by directly mapping Profiles into Ansible Groups.  If machines have been assigned a profile then it will be included in the Group hosts list.  Params in Profiles will be presents as Group vars.  *There is no additional mapping required.*

Note: Ansible dynamic inventory requires JSON output instead of YAML and the format is slightly different.


Optionally, parent groups can be configured by adding the `ansible/children` Param to any profile.  The Param is a simple list of groups to be listed in the Groups children.


DRP Environment Variables
-------------------------

The Ansible inventory assumes that you have correctly set the standard DRP environment variables for access: 

  * RS_ENDPOINT for the endpoint [default "https://127.0.0.1:8092"]
  * RS_KEY for authentication [default "rocketskates:r0cketsk8ts"]

You can also define specialized behavior for the Ansible inventory
 
  * RS_ANSIBLE will filter machines for the matching Ansible [default "all_machines" disables filter]
  * RS_HOST_ADDRESS see below
  * RS_ANSIBLE_USER see below

Ansible Dynamic Inventory from Digital Rebar Provision
------------------------------------------------------

Be certain to export the `RS_ENDPOINT` and `RS_KEY` to match the DRP endpoint information because the DRP dynamic Ansible inventory script relies on these values being set.

Optionally, you may limit the machines using the `ansible=[key]` Param by set `RS_ANSIBLE` to match the [key] value assigned.  The default is to ignore this value and use all machines.

For this example, please ensure that *jq* is installed.

Download the `drpmachines.py` inventory script to the local system to a convenient location and make it executable.  You can test the script by simply running it.  The script will output JSON in the required Ansible format.

  ::

    curl -s https://raw.githubusercontent.com/digitalrebar/provision/v4/integrations/ansible/drpmachines.py -o drpmachines.py
    chmod +x drpmachines.py
    ./drpmachines.py | jq

.. note:: If you run into problems using the drpmachines.py script try enabling the --debug flag for verbose logging.


In order to test the Ansible integration, use the ping command.  If everything is working, all the machines in the system should receive and respond to the ping command.

  ::

    ansible all -i drpmachines.py -m ping

.. note:: You may want to set `export ANSIBLE_HOST_KEY_CHECKING=False` to bypass the SSH key validation

Parent Children Groups
----------------------

Version 4.1 has a simple mapping for Parent groups: use the Profile.Profiles value to list the children groups of the desired parent group.  This is a direct 1 to 1 mapping with Ansible.

Use non-root Login
------------------

By default, the internal `ansible_user` will be set to `root`.  You can override this by setting the `RS_ANSIBLE_USER` value.

  ::

    RS_ANSIBLE_USER="username"


.. _rs_ansible_aws:

Use Alternative Host Address
----------------------------

By default, the internal Machine.Address value is used.  If this address does not work (e.g. Cloud IPs) then you can specify a parameter as the source of the IP address using the `RS_HOST_ADDRESS` value.

  ::

    export RS_HOST_ADDRESS="cloud/public-ipv4"

Summary
-------

Now that these steps are completed, the Digital Rebar Provision dynamic inventory script can be used in any number of ways.


Example of JSON output
----------------------

Here is an example of the output generated for the three node K3s install in AWS.  Please note that the expected JSON does not match Ansible documentation!

  ::

    {
      "node": {
        "hosts": [
          "ip-172-31-28-34.us-west-2.compute.internal",
          "ip-172-31-30-174.us-west-2.compute.internal"
        ],
        "vars": {}
      },
      "all": {
        "hosts": [
          "ip-172-31-28-34.us-west-2.compute.internal",
          "ip-172-31-22-153.us-west-2.compute.internal",
          "ip-172-31-30-174.us-west-2.compute.internal"
        ]
      },
      "_meta": {
        "rebar_profile": "all_machines",
        "rebar_user": "rocketskates",
        "hostvars": {
          "ip-172-31-30-174.us-west-2.compute.internal": {
            "cloud/public-ipv4": "18.236.144.191",
            "cloud/provider": "AWS",
            "detected-bios-mode": "legacy-bios",
            "rebar_uuid": "9646b873-0ecf-4cbe-94eb-c1deb2e20167",
            "ansible_user": "centos",
            "cloud/placement/availability-zone": "us-west-2b",
            "cloud/public-hostname": "ec2-18-236-144-191.us-west-2.compute.amazonaws.com",
            "cloud/instance-type": "t2.xlarge",
            "cloud/instance-id": "i-0c0ef4821c536246f",
            "ansible_host": "18.236.144.191"
          },
          "ip-172-31-28-34.us-west-2.compute.internal": {
            "cloud/public-ipv4": "34.222.134.226",
            "cloud/provider": "AWS",
            "detected-bios-mode": "legacy-bios",
            "rebar_uuid": "245468dc-b61b-471d-ac90-127165a51cb3",
            "ansible_user": "centos",
            "cloud/placement/availability-zone": "us-west-2b",
            "cloud/public-hostname": "ec2-34-222-134-226.us-west-2.compute.amazonaws.com",
            "cloud/instance-type": "t2.xlarge",
            "cloud/instance-id": "i-040fd5ebdcc3908b3",
            "ansible_host": "34.222.134.226"
          },
          "ip-172-31-22-153.us-west-2.compute.internal": {
            "cloud/public-ipv4": "34.221.97.235",
            "cloud/provider": "AWS",
            "detected-bios-mode": "legacy-bios",
            "rebar_uuid": "5556adcf-46d4-41e4-9e6c-c379f5edb743",
            "ansible_user": "centos",
            "cloud/placement/availability-zone": "us-west-2b",
            "cloud/public-hostname": "ec2-34-221-97-235.us-west-2.compute.amazonaws.com",
            "cloud/instance-type": "t2.xlarge",
            "cloud/instance-id": "i-01db8e5ccc14d98c0",
            "ansible_host": "34.221.97.235"
          }
        },
        "rebar_url": "https://34.222.216.7:8092"
      },
      "k3s-cluster": {
        "hosts": [],
        "children": [
          "master",
          "node"
        ],
        "vars": {}
      },
      "master": {
        "hosts": [
          "ip-172-31-22-153.us-west-2.compute.internal"
        ],
        "vars": {}
      }
    }

For reference only, the machines have been deleted.