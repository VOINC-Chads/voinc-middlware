# Worker
# Works to take in jobs and give back output of the jobs to the Main/Master

import zmq
import argparse
import logging
from Middlewares.WorkerMW import WorkerMW

class Worker():

    def __init__(self, logger):

        self.logger = logger
        self.mw_obj = None
        self.name = None



    def configure(self, args):

        try:

            self.name = args.name

            self.mw_obj = WorkerMW(self.logger)
            self.mw_obj.configure(args)

        except Exception as e:
            raise e

    def register(self):

        try:

            self.mw_obj.register(self.name)
            print("Registering")

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



def parseCmdLineArgs():
    parser = argparse.ArgumentParser(description="Main/Master Application")

    parser.add_argument("-n", "--name", default="main", help="Name assigned to master")
    parser.add_argument("-a", "--addr", default="localhost", help="Address process is running on")
    parser.add_argument("-p", "--port", default="5000", help="Port process is running on")
    parser.add_argument("-c", "--capacity", default=2048, help="Storage capacity for job to take place")
    parser.add_argument("-z", "--zkaddr", default="localhost", help="Address zookeeper is running on")
    parser.add_argument("-o", "--zkport", default="2181", help="Port zookeeper is running on")

    parser.add_argument()