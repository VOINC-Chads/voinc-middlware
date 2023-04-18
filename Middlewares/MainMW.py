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
        self.logger = logger
        self.port = None
        self.addr = None
        self.poller = None
        self.upcall_obj = None

        self.codePub = None

        self.zkAddr = None
        self.zkPort = None
        self.leader = False
        self.zk = None

        self.leaderQuorum = None

        self.leaderPort = 7001

        self.dealers = {}

        self.replicaSub = None
        self.leaderPub = None

    def anyof(self, events):

        try:

            for worker in self.dealers:
                socket = self.dealers[worker]
                if socket in events:
                    return True

            return False

        except Exception as e:
            raise e
    def set_dealer(self, addr_port):

        try:

            dealer = self.context.socket(zmq.DEALER)
            conn_string = "tcp://" + addr_port
            dealer.connect(conn_string)

            self.dealers[addr_port] = dealer

            self.poller.register(dealer, zmq.POLLIN)

        except Exception as e:
            raise e

    def set_upcall_handle(self, obj):

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
            self.leaderQuorum = args.leadersize

            self.zk = ZK(self.zkPort, self.zkAddr, self.logger)

            context = zmq.Context()

            self.context = context

            self.poller = zmq.Poller()
            self.router = context.socket(zmq.ROUTER)

            self.poller.register(self.router, zmq.POLLIN)

            self.replicaSub = context.socket(zmq.SUB)
            self.leaderPub = context.socket(zmq.PUB)

            self.poller.register(self.replicaSub, zmq.POLLIN)

            bind_string = "tcp://*:" + str(self.port)

            self.router.bind(bind_string)

            self.zk.create("/replicas/{}".format(args.name), eph=True, value=(self.addr + ":" + str(self.port)))
            while self.zk.get_num_children("/replicas/") < self.leaderQuorum:
                self.logger.info("MainMW::configure - waiting for leader election.")
                time.sleep(0.3)


            # Do leader election
            self.leader = self.zk.create(name="/main", eph=True, value=self.addr + ":" + str(self.port))

            if not self.leader:
                self.logger.info("MainMW::configure - this replica not a leader")
                self.zk.replica_watch_main("/main", value=(self.addr+":"+str(self.port)), obj=self)
                value = self.zk.get_node_data("/main")[0]
                self.logger.info(value)
                if value is not None:
                    decoded = value.decode('utf-8')
                    self.logger.info("MainMW::configure - connecting sub socket {}".format(decoded))
                    self.replicaSub.connect("tcp://" + decoded)
            else:
                self.leaderPub.bind("tcp://*:" + str(self.leaderPort))



        except Exception as e:
            raise e


    def handle_fault(self, leader):

        try:

            self.logger.info("MainMW::handle_fault - fault occurred, so handling now")

            if not leader:
                self.logger.info("MainMW::handle_fault - this is not leader")
                value = self.zk.get_node_data("/main")[0]
                self.logger.info(value)
                if value is not None:
                    decoded = value.decode('utf-8')
                    self.logger.info("Connecting sub {}".format(decoded))
                    self.replicaSub.connect("tcp://" + decoded)

            else:
                self.logger.info("MainMW::handle_fault - this is leader")
                self.leaderPub.bind("tcp://*:" + str(self.leaderPort))


        except Exception as e:
            raise e
    def event_loop(self, timeout=None):

        try:

            self.logger.info("Main::event_loop")

            while True:

                events = dict(self.poller.poll(timeout=timeout))
                self.logger.info(events)
                if self.router in events:

                    self.handle_message()

                elif self.anyof(events):

                    for worker in self.dealers:
                        socket = self.dealers[worker]
                        if socket in events:

                            self.logger.info("Received job response from worker")

                            recvd = socket.recv_multipart()
                            self.logger.info(recvd)
                            # Do some extra consensus work
                            self.router.send_multipart(recvd)

        except Exception as e:
            raise e


    def send_register_response(self, status, id):

        try:

            resp = messages_pb2.MainResp()
            resp.msg_type = messages_pb2.TYPE_REGISTER

            reg_resp = messages_pb2.RegisterResp()
            reg_resp.status = 1 if status else 0

            resp.register_resp.CopyFrom(reg_resp)

            buf2send = resp.SerializeToString()

            self.router.send_multipart([id, buf2send])

        except Exception as e:
            raise e


    def send_to_worker(self, message, id):

        try:

            buf2send = message.SerializeToString()


            for worker in self.dealers:
                socket = self.dealers[worker]
                socket.send_multipart([id, buf2send])



        except Exception as e:
            raise e

    def handle_message(self):

        try:

            self.logger.info("MainMW::handle_message")

            recvd = self.router.recv_multipart()
            id = recvd[0]
            message = recvd[1]

            main_msg = messages_pb2.MainReq()
            main_msg.ParseFromString(message)

            if main_msg.msg_type == messages_pb2.TYPE_REGISTER:
                self.logger.info("Register received")
                self.logger.info(main_msg)
                self.upcall_obj.register_volunteer(main_msg.register_req, id)
            elif main_msg.msg_type == messages_pb2.TYPE_CODE:
                self.logger.info("Code received")
                self.logger.info(main_msg)
                self.send_to_worker(main_msg, id)
            elif main_msg.msg_type == messages_pb2.TYPE_JOB:
                self.logger.info("Job received")
                self.logger.info(main_msg)
                self.send_to_worker(main_msg, id)




        except Exception as e:
            raise e