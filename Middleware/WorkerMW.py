# WorkerMW
# Purpose to interact with zookeeper and master to send responses to jobs and receive jobs

import zmq
from Messages import messages_pb2
from Middleware.zookeeper import ZK

class WorkerMW():

    def __init__(self, logger):

        self.logger = logger
        self.port = None
        self.addr = None
        self.poller = None
        self.upcall_obj = None
        self.capacity = None

        self.dealer = None
        self.zkAddr = None
        self.zkPort = None
        self.zk = None

        self.router = None



    def configure(self, args):

        try:

            self.logger.info("WorkerMW::configure - configuring worker")

            self.port = args.port
            self.addr = args.addr
            self.zkPort = args.zkport
            self.zkAddr = args.zkaddr
            self.capacity = args.capacity

            context = zmq.Context()
            self.poller = zmq.Poller()
            self.dealer = context.socket(zmq.DEALER)
            self.router = context.socket(zmq.ROUTER)

            self.poller.register(self.dealer, zmq.POLLIN)
            self.poller.register(self.router, zmq.POLLIN)

            self.router.bind("tcp://*:" + str(self.port))

            self.zk = ZK(self.zkPort, self.zkAddr, self.logger)

            self.logger.info("Worker waiting for main to exist")


            while not self.zk.exists("/main"):
                self.logger.info("Waiting for main to start")

            self.zk.watch_main_change("/main", self)


        except Exception as e:
            raise e

    def connect_main(self, string):

        try:

            self.logger.info("WorkerMW::connect_main - connecting to main at {}".format(string))

            self.dealer.connect("tcp://" + string)

            self.logger.info("WorkerMW::connect_main - connected to main")


        except Exception as e:
            raise e


    def register(self, name):

        try:

            # Zookeeper way of just creating a node
            # self.zk.create(name)

            self.logger.info("WorkerMW::register - Making registration request")

            main_req = messages_pb2.MainReq()
            main_req.msg_type = messages_pb2.TYPE_REGISTER

            reg_info = messages_pb2.RegistrantInfo()
            reg_info.id = name
            reg_info.port = int(self.port)
            reg_info.addr = self.addr
            reg_info.capacity = self.capacity

            reg_req = messages_pb2.RegisterReq()
            reg_req.role = messages_pb2.ROLE_VOLUNTEER
            reg_req.info.CopyFrom(reg_info)

            main_req.register_req.CopyFrom(reg_req)

            buf2send = main_req.SerializeToString()

            self.logger.info("Sending registration request")
            self.dealer.send(buf2send)



        except Exception as e:
            raise e
    def event_loop(self, timeout=None):

        try:

            while True:

                events = dict(self.poller.poll(timeout=timeout))

                if not events:

                    timeout = self.upcall_obj.invoke_operation()

                elif self.dealer in events:
                    self.logger.info("WorkerMW::event_loop - dealer sent a message")
                    timeout = self.handle_response()

                elif self.router in events:
                    self.logger.info("WorkerMW::event_loop - router sent a message")
                    timeout = self.handle_code_or_job()

        except Exception as e:
            raise e


    def send_job_response(self, results, ids):

        try:

            self.logger.info("WorkerMW::send_job_response - sending back response with results below")
            self.logger.info(results)

            job_resp = messages_pb2.JobResp()

            job_resp.status = 1

            for result in results:

                job_res = messages_pb2.JobResult()
                job_res.value = str(result)
                job_res.result = str(results[result])

                job_resp.results.append(job_res)

            main_resp = messages_pb2.MainResp()
            main_resp.msg_type = messages_pb2.TYPE_JOB
            main_resp.job_resp.CopyFrom(job_resp)

            main_resp.msg_type = messages_pb2

            self.logger.info("Sending back")
            self.logger.info(main_resp)

            buf2send = main_resp.SerializeToString()

            self.router.send_multipart(ids + [buf2send])

            self.logger.info("Sent")

        except Exception as e:
            raise e

    def handle_code_or_job(self):

        try:
            self.logger.info("WorkerMW::handle_code_or_job - main sent a job or code to handle")

            recvd = self.router.recv_multipart()
            ids = recvd[:2]
            message = recvd[2]

            main_msg = messages_pb2.MainReq()
            main_msg.ParseFromString(message)

            if main_msg.msg_type == messages_pb2.TYPE_CODE:
                self.upcall_obj.handle_code(main_msg.code_msg)
            elif main_msg.msg_type == messages_pb2.TYPE_JOB:
                self.upcall_obj.handle_job(main_msg.job_msg, ids)


        except Exception as e:
            raise e
    def set_upcall_handle(self, obj):

        try:

            self.upcall_obj = obj

        except Exception as e:
            raise e

    def handle_response(self):
        self.logger.info("WorkerMW::handle_response - main sent a response to handle")
        try:
            timeout = None

            recvd = self.dealer.recv()

            resp = messages_pb2.MainResp()
            resp.ParseFromString(recvd)

            self.logger.info("WorkerMW::handle_response - received response")
            if resp.msg_type == messages_pb2.TYPE_REGISTER:
                self.logger.info("WorkerMW::handle_response - received registration response")
                timeout = self.upcall_obj.register_response(resp.register_resp)

            return timeout

        except Exception as e:
            raise e
