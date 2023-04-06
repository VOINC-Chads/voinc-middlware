# Main.cpp
# Main/master of the volunteer computing architecture
# Volunteers provide registration information and then main distributes jobs from workers
# to volunteers. Due to bad actors, it is the responsibility of a consensus of responses to be received
# to be sent back to the worker.


import zmq
from kazoo.client import KazooClient

class Main():


    def __init__(self):

        self.volunteers = {}
        self.pending = {}


    def register_volunteer(self, information):

        try:


            print("Registering volunteer")

        except Exception as e:
            raise e
