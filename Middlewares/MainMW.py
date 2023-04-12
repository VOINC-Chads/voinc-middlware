# MainMW
# Purpose is to fulfill requests and message sending.


import zmq
from Messages import messages_pb2
from Middlewares.zookeeper import ZK
import requests
import time

class MainMW():


    def __init__(self, logger):

        self.router = None
        self.logger = None
        self.port = None
        self.addr = None
        self.poller = None
        self.upcall_obj = None

        self.zkAddr = None
        self.zkPort = None
        self.leader = False
        self.zk = None

        self.codeReceived = True

    def create_endpoint(self):

        try:


            self.logger.info("Waiting for code to exist at this endpoint")


            while not self.codeReceived:
                self.logger.info("Waiting on code")

                response = requests.get("http://localhost:3000/getCode")
                self.logger.info(response.json())
                time.sleep(2)



        except Exception as e:
            raise e
    def set_upccall_handle(self, obj):

        try:
            self.upcall_obj = obj
        except Exception as e:
            raise e
    def configure(self, args):

        try:

            self.port = args.port
            self.addr = args.addr
            self.zkPort = args.zkport
            self.zkAddr = args.zkaddr

            self.zk = ZK(self.zkPort, self.zkAddr, self.logger)

            context = zmq.Context()

            self.poller = zmq.Poller()
            self.router = context.socket(zmq.ROUTER)

            self.poller.register(self.router, zmq.POLLIN)

            bind_string = "tcp://*:" + str(self.port)

            self.router.bind(bind_string)

            # Do leader election
            self.leader = self.zk.create(name="/main", eph=True, value=self.addr + ":" + self.port)

            if not self.leader:
                print("Setting subscription sockets to changes")
            else:
                self.create_endpoint()
                print("Setting publisher sockets to report updates")



        except Exception as e:
            raise e

    def event_loop(self, timeout=None):

        try:

            while True:

                events = dict(self.poller.poll(timeout=timeout))

                if self.router in events:

                    self.handle_message()

        except Exception as e:
            raise e


    def handle_message(self):

        try:

            recvd = self.router.recv_multipart()
            id = recvd[0]
            message = recvd[1]

            main_msg = messages_pb2.MainReq()
            main_msg.ParseFromString(message)

            if main_msg.type == messages_pb2.TYPE_REGISTER:
                self.upcall_obj.register_volunteer(main_msg.register_req, id)




        except Exception as e:
            raise e