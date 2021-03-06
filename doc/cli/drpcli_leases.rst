drpcli leases
-------------

Access CLI commands relating to leases

Synopsis
~~~~~~~~

Access CLI commands relating to leases

Options
~~~~~~~

::

     -h, --help   help for leases

Options inherited from parent commands
~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

::

     -c, --catalog string          The catalog file to use to get product information (default "https://repo.rackn.io")
     -d, --debug                   Whether the CLI should run in debug mode
     -D, --download-proxy string   HTTP Proxy to use for downloading catalog and content
     -E, --endpoint string         The Digital Rebar Provision API endpoint to talk to (default "https://127.0.0.1:8092")
     -f, --force                   When needed, attempt to force the operation - used on some update/patch calls
     -F, --format string           The serialization we expect for output.  Can be "json" or "yaml" or "text" or "table" (default "json")
     -H, --no-header               Should header be shown in "text" or "table" mode
     -x, --noToken                 Do not use token auth or token cache
     -P, --password string         password of the Digital Rebar Provision user (default "r0cketsk8ts")
     -J, --print-fields string     The fields of the object to display in "text" or "table" mode. Comma separated
     -r, --ref string              A reference object for update commands that can be a file name, yaml, or json blob
     -T, --token string            token of the Digital Rebar Provision access
     -t, --trace string            The log level API requests should be logged at on the server side
     -Z, --traceToken string       A token that individual traced requests should report in the server logs
     -j, --truncate-length int     Truncate columns at this length (default 40)
     -U, --username string         Name of the Digital Rebar Provision user to talk to (default "rocketskates")

SEE ALSO
~~~~~~~~

-  `drpcli <drpcli.html>`__ - A CLI application for interacting with the
   DigitalRebar Provision API
-  `drpcli leases action <drpcli_leases_action.html>`__ - Display the
   action for this lease
-  `drpcli leases actions <drpcli_leases_actions.html>`__ - Display
   actions for this lease
-  `drpcli leases destroy <drpcli_leases_destroy.html>`__ - Destroy
   lease by id
-  `drpcli leases exists <drpcli_leases_exists.html>`__ - See if a
   leases exists by id
-  `drpcli leases indexes <drpcli_leases_indexes.html>`__ - Get indexes
   for leases
-  `drpcli leases list <drpcli_leases_list.html>`__ - List all leases
-  `drpcli leases meta <drpcli_leases_meta.html>`__ - Gets metadata for
   the lease
-  `drpcli leases runaction <drpcli_leases_runaction.html>`__ - Run
   action on object from plugin
-  `drpcli leases show <drpcli_leases_show.html>`__ - Show a single
   leases by id
-  `drpcli leases wait <drpcli_leases_wait.html>`__ - Wait for a lease’s
   field to become a value within a number of seconds
