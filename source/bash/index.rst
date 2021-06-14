=====
Bash
=====


.. toctree::
   :maxdepth: 2
   :caption: Contents:


Base
====

How to run Bash script
-----------------------

Shell scripts usually starts with `shebang`, which specifies interpreter.

.. code:: bash

   #!/bin/bash

 
There are two way to run script:

1. run script as cli argument

    .. code:: bash

       bash script.sh


2. make script executable(need shebang)

      .. code:: bash
      
         chmod 755 script.sh  # or chmod a+x script.sh
         ./script.sh
