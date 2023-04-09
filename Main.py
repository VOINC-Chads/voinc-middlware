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

class Main():


    def __init__(self, logger):

        self.logger = logger
        self.volunteers = {}
        self.pending = {}
        self.mw_obj = None
        self.name = None

        self.numVolunteers = None
        self.numOccupied = None



    def configure(self, args):

        try:

            self.numOccupied = 0
            self.numVolunteers = 0
            self.name = args.name

            self.mw_obj = MainMW(self.logger)
            self.mw_obj.configure(args)



        except Exception as e:
            raise e


    def register_volunteer(self, information):

        try:

            name = information.info.id
            port = information.info.port
            addr = information.info.addr
            capacity = information.info.capacity

            if name in self.volunteers:
                print("Already registered worker")
            else:
                information = [name, port, addr, capacity]
                self.volunteers[name] = information
                print("Looking at new worker")

        except Exception as e:
            raise e


def parseCmdLineArgs():


    parser = argparse.ArgumentParser(description="Main/Master Application")

    parser.add_argument("-n", "--name", default="main", help="Name assigned to master")
    parser.add_argument("-a", "--addr", default="localhost", help="Address process is running on")
    parser.add_argument("-p", "--port", default="5000", help="Port process is running on")
    parser.add_argument("-z", "--zkaddr", default="localhost", help="Address zookeeper is running on")
    parser.add_argument("-o", "--zkport", default="2181", help="Port zookeeper is running on")

    parser.add_argument()