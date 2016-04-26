go-websocket-sample
===================

To see the sample in action, simply run:

.. sourcecode:: text

    $ go run <path/to/server.go>

This starts up the server; it immediately reports which port it's listening on
(the port can be changed with the ``-port`` flag).

Then open a browser and visit ``http://localhost:<portnum>`` to see the HTML
page that talks with the server using websockets. The server also spins up
a `net/trace <https://godoc.org/golang.org/x/net/trace>`__ debugging page on
``/debug/requests``.

License
=======

This code is in the public domain. See the ``LICENSE`` file for more details.
