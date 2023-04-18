# Worker
# Works to take in jobs and give back output of the jobs to the Main/Master
import ast
import subprocess

import zmq
import argparse
import logging
from Middleware.WorkerMW import WorkerMW
from enum import Enum

class Worker():

    class State(Enum):
        INITIALIZE = 0,
        CONFIGURE = 1,
        REGISTER = 2,
        PERFORM_JOBS = 3


    def __init__(self, logger):

        self.logger = logger
        self.mw_obj = None
        self.name = None
        self.capacity = None
        self.IP = None
        self.port = None

        self.codeReceived = False


        self.state = self.State.INITIALIZE




    def configure(self, args):

        try:

            self.logger.info("Worker::configure - configuring worker")

            self.name = args.name
            self.capacity = args.capacity

            self.mw_obj = WorkerMW(self.logger)



            self.mw_obj.configure(args)

        except Exception as e:
            raise e

    def register_response(self, resp):

        try:

            self.logger.info("Worker::register_response - got a registration response")
            self.logger.info(resp)

            if resp.status:
                self.logger.info("Worker::register_response - successful registration")
                self.state = self.State.PERFORM_JOBS

                return None

            else:
                raise("Worker with same credentials already exists.")


        except Exception as e:
            raise e

    def handle_register(self, resp):

        try:

            print(resp)

        except Exception as e:
            raise e
    def perform_job(self, job_info):

        try:

            print("Performing job")

        except Exception as e:
            raise e

    def handle_job(self, job_msg, ids):

        try:
            self.logger.info("Worker::handle_job - got the job below")
            self.logger.info(job_msg)

            if not self.codeReceived:
                self.logger.info("Code not received")
            else:
                jobs = job_msg.jobs
                results = {}
                for job in jobs:
                    self.logger.info(job)
                    result = process_function(job)
                    results[job] = result

                self.logger.info("Worker::handle_job - finished all parts of job given")
                self.mw_obj.send_job_response(results, ids)

        except Exception as e:
            raise e

    def handle_code(self, code_msg):

        try:

            self.logger.info("Worker::handle_code - got code below")
            self.logger.info(code_msg)

            reqs = code_msg.requirements
            if len(reqs) > 0:
                reqs = reqs.split("\n")
                self.logger.info(reqs)
                
                for req in reqs:
                    self.logger.info(req)
                    subprocess.check_call(["pip3", "install", f"{req}"])
            process = code_msg.process_code
            execute = code_msg.execute_code

            process = process.replace("\\\\", "\\")
            execute = execute.replace("\\\\", "\\")
            exec(process, globals())
            exec(execute, globals())

            self.codeReceived = True


        except Exception as e:
            raise e

    def invoke_operation(self):


        try:

            self.logger.info("Worker::invoke_operation")

            if (self.state == self.State.REGISTER):
                # send a register msg to discovery service
                self.logger.debug("Worker::invoke_operation - register with the discovery service")
                self.mw_obj.register(self.name)

                return None

        except Exception as e:
            raise e


    def driver(self):

        try:


            # Set the upcall handler from the middleware
            self.mw_obj.set_upcall_handle(self)


            self.state = self.State.REGISTER
            # start the event loop that just continuously runs and checks for requests

            self.mw_obj.event_loop(timeout=0)


        except Exception as e:
            raise e


def parseCmdLineArgs():
    parser = argparse.ArgumentParser(description="Worker Application")

    parser.add_argument("-n", "--name", default="worker", help="Name assigned to master")
    parser.add_argument("-a", "--addr", default="localhost", help="Address process is running on")
    parser.add_argument("-p", "--port", default="8001", help="Port process is running on")
    parser.add_argument("-c", "--capacity", default=2048, help="Storage capacity for job to take place")
    parser.add_argument("-z", "--zkaddr", default="localhost", help="Address zookeeper is running on")
    parser.add_argument("-o", "--zkport", default="2181", help="Port zookeeper is running on")
    parser.add_argument("-l", "--loglevel", type=int, default=logging.INFO,
                        choices=[logging.DEBUG, logging.INFO, logging.WARNING, logging.ERROR, logging.CRITICAL],
                        help="logging level, choices 10,20,30,40,50: default 20=logging.INFO")

    return parser.parse_args()


def main():
    try:
        # obtain a system wide logger and initialize it to debug level to begin with
        logging.debug("Worker - acquire a child logger and then log messages in the child")
        logger = logging.getLogger("Worker")

        # first parse the arguments
        logger.debug("Worker: parse command line arguments")
        args = parseCmdLineArgs()

        # reset the log level to as specified
        logger.debug("Worker: resetting log level to {}".format(args.loglevel))
        logger.setLevel(args.loglevel)
        logger.debug("Worker: effective log level is {}".format(logger.getEffectiveLevel()))

        # Obtain a publisher application
        logger.debug("Worker: obtain the object")
        disc_app = Worker(logger)

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