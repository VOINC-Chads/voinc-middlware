# Main.cpp
# Main/master of the volunteer computing architecture
# Volunteers provide registration information and then main distributes jobs from workers
# to volunteers. Due to bad actors, it is the responsibility of a consensus of responses to be received
# to be sent back to the worker.

import argparse
import zmq
from kazoo.client import KazooClient
from Middlewares.MainMW import MainMW
import logging
from flask import Flask, request
import time
import requests

class Main():


    def __init__(self, logger):

        self.logger = logger
        self.volunteers = {}
        self.pending = {}
        self.mw_obj = None
        self.name = None

        self.numVolunteers = None
        self.numOccupied = None

        self.code = None


    def update_code(self, code):

        try:

            self.code = code

        except Exception as e:
            raise e
    def configure(self, args):

        try:

            self.numOccupied = 0
            self.numVolunteers = 0
            self.name = args.name

            self.mw_obj = MainMW(self.logger)
            self.mw_obj.configure(args)



        except Exception as e:
            raise e


    def register_volunteer(self, information, id):

        try:

            name = information.info.id
            port = information.info.port
            addr = information.info.addr
            capacity = information.info.capacity

            if str(addr) + ":" + str(port) in self.volunteers:
                self.logger.info("Already registered worker")

            else:
                information = [name, port, addr, capacity]
                self.volunteers[str(addr) + ":" + str(port)] = information
                self.logger.info()
                print("Looking at new worker")

        except Exception as e:
            raise e


    def driver(self):

        try:


            # Set the upcall handler from the middleware
            self.mw_obj.set_upcall_handle(self)

            # start the event loop that just continuously runs and checks for requests

            self.mw_obj.event_loop(timeout=None)



        except Exception as e:
            raise e

def parseCmdLineArgs():


    parser = argparse.ArgumentParser(description="Main/Master Application")

    parser.add_argument("-n", "--name", default="main", help="Name assigned to master")
    parser.add_argument("-a", "--addr", default="localhost", help="Address process is running on")
    parser.add_argument("-p", "--port", default="5000", help="Port process is running on")
    parser.add_argument("-z", "--zkaddr", default="localhost", help="Address zookeeper is running on")
    parser.add_argument("-o", "--zkport", default="2181", help="Port zookeeper is running on")
    parser.add_argument("-l", "--loglevel", type=int, default=logging.INFO,
                        choices=[logging.DEBUG, logging.INFO, logging.WARNING, logging.ERROR, logging.CRITICAL],
                        help="logging level, choices 10,20,30,40,50: default 20=logging.INFO")

    return parser.parse_args()


def main():
    try:
        # obtain a system wide logger and initialize it to debug level to begin with
        logging.debug("Main - acquire a child logger and then log messages in the child")
        logger = logging.getLogger("Main")

        # first parse the arguments
        logger.debug("Main: parse command line arguments")
        args = parseCmdLineArgs()

        # reset the log level to as specified
        logger.debug("Main: resetting log level to {}".format(args.loglevel))
        logger.setLevel(args.loglevel)
        logger.debug("Main: effective log level is {}".format(logger.getEffectiveLevel()))

        # Obtain a publisher application
        logger.debug("Main: obtain the object")
        disc_app = Main(logger)

        # configure the object
        disc_app.configure(args)

        # now invoke the driver program
        disc_app.driver()

    except Exception as e:
        logger.error("Exception caught in main - {}".format(e))
        return

if __name__ == "__main__":

  # set underlying default logging capabilities
  logging.basicConfig (level=logging.DEBUG,
                       format='%(asctime)s - %(name)s - %(levelname)s - %(message)s')


  main()
