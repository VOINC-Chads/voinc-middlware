# WorkerMW
# Purpose to interact with zookeeper and master to send responses to jobs and receive jobs

import zmq
from Messages import messages_pb2


class WorkerMW():

    def __init__(self, logger):

        self.logger = logger
        self.port = None
        self.addr = None
        self.poller = None
        self.upcall_obj = None

        self.dealer = None
        self.zkAddr = None
        self.zkPort = None


    def configure(self, args):

        try:

            self.port = args.port
            self.addr = args.addr
            self.zkPort = args.zkport
            self.zkAddr = args.zkaddr

            context = zmq.Context()
            self.poller = zmq.Poller()
            self.dealer = context.socket(zmq.DEALER)

            self.poller.register(self.dealer, zmq.POLLIN)


        except Exception as e:
            raise e


    def event_loop(self, timeout=None):

        try:

            while True:

                events = dict(self.poller.poll(timeout=timeout))

                if self.dealer in events:

                    self.handle_response()

        except Exception as e:
            raise e


    def handle_response(self):


        try:

            recvd = self.dealer.recv()


        except Exception as e:
            raise e
